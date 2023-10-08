package talk_message

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-server/internal/config"
	"github.com/iimeta/iim-server/internal/consts"
	"github.com/iimeta/iim-server/internal/core"
	"github.com/iimeta/iim-server/internal/dao"
	"github.com/iimeta/iim-server/internal/model"
	"github.com/iimeta/iim-server/internal/model/do"
	"github.com/iimeta/iim-server/internal/service"
	"github.com/iimeta/iim-server/utility/cache"
	"github.com/iimeta/iim-server/utility/logger"
	"github.com/iimeta/iim-server/utility/redis"
	"github.com/iimeta/iim-server/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type sTalkMessage struct {
	unreadStorage  *cache.UnreadStorage
	messageStorage *cache.MessageStorage
	sidStorage     *cache.ServerStorage
	clientStorage  *cache.ClientStorage
}

func init() {
	service.RegisterTalkMessage(New())
}

func New() service.ITalkMessage {
	return &sTalkMessage{
		unreadStorage:  cache.NewUnreadStorage(redis.Client),
		messageStorage: cache.NewMessageStorage(redis.Client),
		sidStorage:     cache.NewSidStorage(redis.Client),
		clientStorage:  cache.NewClientStorage(redis.Client, config.Cfg, cache.NewSidStorage(redis.Client)),
	}
}

// 文本消息
func (s *sTalkMessage) SendText(ctx context.Context, uid int, req *model.TextMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeText,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Content:    util.EscapeHtml(req.Content),
	}

	return s.save(ctx, data)
}

// 代码消息
func (s *sTalkMessage) SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordCode{
			Lang: req.Lang,
			Code: req.Code,
		}),
	}

	return s.save(ctx, data)
}

func (s *sTalkMessage) save(ctx context.Context, data *model.TalkRecord) error {

	if data.RecordId == 0 {
		data.RecordId = core.IncrRecordId(ctx)
	}

	if data.MsgId == "" {
		data.MsgId = util.NewMsgId()
	}

	s.loadReply(ctx, data)

	s.loadSequence(ctx, data)

	id, err := dao.TalkRecords.Insert(ctx, &do.TalkRecords{
		RecordId:   data.RecordId,
		MsgId:      data.MsgId,
		Sequence:   data.Sequence,
		TalkType:   data.TalkType,
		MsgType:    data.MsgType,
		UserId:     data.UserId,
		ReceiverId: data.ReceiverId,
		IsRevoke:   data.IsRevoke,
		IsMark:     data.IsMark,
		IsRead:     data.IsRead,
		QuoteId:    data.QuoteId,
		Content:    data.Content,
		Extra:      data.Extra,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	data.Id = id

	option := make(map[string]string)

	switch data.MsgType {
	case consts.ChatMsgTypeText:
		option["text"] = gstr.SubStr(util.ReplaceImgAll(data.Content), 0, 300)
	default:
		if value, ok := consts.ChatMsgTypeMapping[data.MsgType]; ok {
			option["text"] = value
		} else {
			option["text"] = "[未知消息]"
		}
	}

	s.afterHandle(ctx, data, option)

	return nil
}

func (s *sTalkMessage) loadSequence(ctx context.Context, data *model.TalkRecord) {
	if data.TalkType == consts.ChatGroupMode {
		data.Sequence = dao.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = dao.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (s *sTalkMessage) loadReply(ctx context.Context, data *model.TalkRecord) {

	// 检测是否引用消息
	if data.QuoteId == "" {
		return
	}

	if data.Extra == "" {
		data.Extra = "{}"
	}

	extra := make(map[string]any)

	err := gjson.Unmarshal([]byte(data.Extra), &extra)
	if err != nil {
		logger.Error(ctx, "MessageService Json Decode err:", err)
		return
	}

	record, err := dao.TalkRecords.FindOne(ctx, bson.M{"msg_id": data.QuoteId})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	user, err := dao.User.FindUserByUserId(ctx, record.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	reply := model.TalkRecordReply{
		UserId:   record.UserId,
		Nickname: user.Nickname,
		MsgType:  1,
		Content:  record.Content,
		MsgId:    record.MsgId,
	}

	if record.MsgType != consts.ChatMsgTypeText {
		reply.Content = "[未知消息]"
		if value, ok := consts.ChatMsgTypeMapping[record.MsgType]; ok {
			reply.Content = value
		}
	}

	extra["reply"] = reply

	data.Extra = gjson.MustEncodeString(extra)
}

// 发送消息后置处理
func (s *sTalkMessage) afterHandle(ctx context.Context, record *model.TalkRecord, opt map[string]string) {

	if record.TalkType == consts.ChatPrivateMode {
		s.unreadStorage.Incr(ctx, consts.ChatPrivateMode, record.UserId, record.ReceiverId)
		if record.MsgType == consts.ChatMsgSysText {
			s.unreadStorage.Incr(ctx, 1, record.ReceiverId, record.UserId)
		}
	} else if record.TalkType == consts.ChatGroupMode {
		pipe := redis.Pipeline(ctx)
		for _, uid := range dao.GroupMember.GetMemberIds(ctx, record.ReceiverId) {
			if uid != record.UserId {
				s.unreadStorage.PipeIncr(ctx, pipe, consts.ChatGroupMode, record.ReceiverId, uid)
			}
		}
		if _, err := pipe.Exec(ctx); err != nil {
			logger.Error(ctx, err)
		}
	}

	if err := s.messageStorage.Set(ctx, record.TalkType, record.UserId, record.ReceiverId, &cache.LastCacheMessage{
		Content:  opt["text"],
		Datetime: gtime.Datetime(),
	}); err != nil {
		logger.Error(ctx, err)
	}

	content := gjson.MustEncodeString(map[string]any{
		"event": consts.SubEventImMessage,
		"data": gjson.MustEncodeString(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"talk_type":   record.TalkType,
			"record_id":   record.RecordId,
		}),
	})

	if record.TalkType == consts.ChatPrivateMode {
		sids := s.sidStorage.All(ctx, 1)

		if len(sids) > 3 {

			pipe := redis.Pipeline(ctx)

			for _, sid := range sids {
				for _, uid := range []int{record.UserId, record.ReceiverId} {
					if !s.clientStorage.IsCurrentServerOnline(ctx, sid, consts.ImChannelChat, strconv.Itoa(uid)) {
						continue
					}
					pipe.Publish(ctx, fmt.Sprintf(consts.ImTopicChatPrivate, sid), content)
				}
			}

			if _, err := pipe.Exec(ctx); err != nil {
				logger.Error(ctx, err)
				return
			}
		}
	}

	if _, err := redis.Publish(ctx, consts.ImTopicChat, content); err != nil {
		logger.Error(ctx, "[ALL]消息推送失败 err:", err)
	}
}
