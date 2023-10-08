package socket

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/iimeta/iim-server/internal/errors"
	"github.com/iimeta/iim-server/utility/logger"
	"github.com/iimeta/iim-server/utility/util"
	"sync/atomic"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

const (
	_MsgEventPing = "ping"
	_MsgEventPong = "pong"
	_MsgEventAck  = "ack"
)

type IClient interface {
	Cid() int64                       // 客户端ID
	Uid() int                         // 客户端关联用户ID
	Close(code int, text string)      // 关闭客户端
	Write(data *ClientResponse) error // 写入数据
	Channel() IChannel                // 获取客户端所属渠道
}

type IStorage interface {
	Bind(ctx context.Context, channel string, cid int64, uid int) error
	UnBind(ctx context.Context, channel string, cid int64) error
}

// WebSocket 客户端连接信息
type Client struct {
	conn     IConn                // 客户端连接
	cid      int64                // 客户端ID/客户端唯一标识
	uid      int                  // 用户ID
	lastTime int64                // 客户端最后心跳时间/心跳检测
	closed   int32                // 客户端是否关闭连接
	channel  IChannel             // 渠道分组
	storage  IStorage             // 缓存服务
	event    IEvent               // 回调方法
	outChan  chan *ClientResponse // 发送通道
}

type ClientOption struct {
	Uid         int         // 用户识别ID
	Channel     IChannel    // 渠道信息
	Storage     IStorage    // 自定义缓存组件, 用于绑定用户与客户端的关系
	IdGenerator IdGenerator // 客户端ID生成器(唯一ID), 默认使用雪花算法
	Buffer      int         // 缓冲区大小根据业务, 自行调整
}

type ClientResponse struct {
	IsAck   bool   `json:"-"`                 // 是否需要 ack 回调
	Sid     string `json:"sid,omitempty"`     // ACK ID
	Event   string `json:"event"`             // 事件名
	Content any    `json:"content,omitempty"` // 事件内容
	Retry   int    `json:"-"`                 // 重试次数(0 默认不重试)
}

// 初始化客户端信息
func NewClient(conn IConn, option *ClientOption, event IEvent) error {

	if option.Buffer <= 0 {
		option.Buffer = 10
	}

	if event == nil {
		panic("event is nil")
	}

	client := &Client{
		conn:     conn,
		uid:      option.Uid,
		lastTime: time.Now().Unix(),
		channel:  option.Channel,
		storage:  option.Storage,
		outChan:  make(chan *ClientResponse, option.Buffer),
		event:    event,
	}

	if option.IdGenerator != nil {
		client.cid = option.IdGenerator.IdGen()
	} else {
		client.cid = defaultIdGenerator.IdGen()
	}

	// 设置客户端连接关闭回调事件
	conn.SetCloseHandler(client.hookClose)

	// 绑定客户端映射关系
	if client.storage != nil {
		ctx := gctx.New()
		err := client.storage.Bind(ctx, client.channel.Name(), client.cid, client.uid)
		if err != nil {
			logger.Error(ctx, "bind client err:", err)
			return err
		}
	}

	// 注册客户端
	client.channel.addClient(client)

	// 触发自定义的 Open 事件
	client.event.Open(client)

	// 注册心跳管理
	health.insert(client)

	return client.init()
}

func (c *Client) Channel() IChannel {
	return c.channel
}

// 获取客户端ID
func (c *Client) Cid() int64 {
	return c.cid
}

// 获取客户端关联的用户ID
func (c *Client) Uid() int {
	return c.uid
}

// 关闭客户端连接
func (c *Client) Close(code int, message string) {
	defer func() {
		if err := c.conn.Close(); err != nil {
			logger.Errorf(gctx.New(), "connection closed failed: %s", err.Error())
		}
	}()

	// 触发客户端关闭回调事件
	if err := c.hookClose(code, message); err != nil {
		logger.Errorf(gctx.New(), "[%s-%d-%d] client close err: %s", c.channel.Name(), c.cid, c.uid, err.Error())
	}
}

func (c *Client) Closed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

// 客户端写入数据
func (c *Client) Write(data *ClientResponse) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf(gctx.New(), "[%s-%d-%d] chan write err: %v", c.channel.Name(), c.cid, c.uid, err)
		}
	}()

	if c.Closed() {
		return errors.New("connection has been closed")
	}

	if data.IsAck {
		data.Sid = util.NewMsgId()
	}

	c.outChan <- data

	return nil
}

// 循环接收客户端推送信息
func (c *Client) loopAccept() {
	defer c.Close(1000, "loop accept closed")

	for {
		data, err := c.conn.Read()
		if err != nil {
			break
		}

		c.lastTime = time.Now().Unix()

		c.handleMessage(data)
	}
}

// 循环推送客户端信息
func (c *Client) loopWrite() {

	ctx := gctx.New()
	timer := time.NewTimer(15 * time.Second)

	defer timer.Stop()

	for {
		timer.Reset(15 * time.Second)

		select {
		case <-timer.C:
			logger.Debugf(ctx, "client empty message cid: %d, uid: %d", c.cid, c.uid)
		case data, ok := <-c.outChan:
			if !ok || c.Closed() {
				return // channel closed
			}

			bt, err := jsoniter.Marshal(data)
			if err != nil {
				logger.Error(ctx, "client json marshal err:", err)
				break
			}

			if err := c.conn.Write(bt); err != nil {
				logger.Errorf(ctx, "[%s-%d-%d] client write err: %v", c.channel.Name(), c.cid, c.uid, err)
				return
			}

			if data.IsAck && data.Retry > 0 {
				data.Retry--

				ackBufferContent := &AckBufferContent{}
				ackBufferContent.cid = c.cid
				ackBufferContent.uid = int64(c.uid)
				ackBufferContent.channel = c.channel.Name()
				ackBufferContent.response = data

				ack.insert(data.Sid, ackBufferContent)
			}
		}
	}
}

// 初始化连接
func (c *Client) init() error {

	// 推送心跳检测配置
	_ = c.Write(&ClientResponse{
		Event: "connect",
		Content: map[string]any{
			"ping_interval": heartbeatInterval,
			"ping_timeout":  heartbeatTimeout,
		}},
	)

	// 启动协程处理推送信息
	_ = grpool.AddWithRecover(gctx.New(), func(ctx context.Context) {
		c.loopWrite()
	}, nil)

	_ = grpool.AddWithRecover(gctx.New(), func(ctx context.Context) {
		c.loopAccept()
	}, nil)

	return nil
}

func (c *Client) hookClose(code int, text string) error {

	if !atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		return nil
	}

	close(c.outChan)

	c.event.Close(c, code, text)

	if c.storage != nil {
		ctx := gctx.New()
		err := c.storage.UnBind(ctx, c.channel.Name(), c.cid)
		if err != nil {
			logger.Error(ctx, "unbind client err:", err)
			return err
		}
	}

	health.delete(c)

	c.channel.delClient(c)

	return nil
}

func (c *Client) handleMessage(data []byte) {

	event, err := c.validate(data)
	if err != nil {
		logger.Errorf(gctx.New(), "validate err:", err)
		return
	}

	switch event {
	case _MsgEventPing:
		_ = c.Write(&ClientResponse{Event: _MsgEventPong})
	case _MsgEventPong:
	case _MsgEventAck:
		ackId := gjson.GetBytes(data, "sid").String()
		if len(ackId) > 0 {
			ack.delete(ackId)
		}
	default: // 触发消息回调
		c.event.Message(c, data)
	}
}

func (c *Client) validate(data []byte) (string, error) {

	if !gjson.ValidBytes(data) {
		return "", errors.New("invalid json")
	}

	event := gjson.GetBytes(data, "event").String()
	if len(event) == 0 {
		return "", errors.New("invalid event")
	}

	return event, nil
}
