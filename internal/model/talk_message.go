package model

// 表情消息
type EmoticonMessageReq struct {
	Type       string    `json:"type,omitempty"`
	Receiver   *Receiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
	EmoticonId string    `json:"emoticon_id" v:"required"`
}

// 位置消息
type CardMessageReq struct {
	Type       string    `json:"type,omitempty"`
	UserId     int       `json:"user_id,omitempty" v:"required"`
	Receiver   *Receiver `json:"receiver,omitempty"`
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
}

// 图文消息
type MixedMessageReq struct {
	Type     string          `json:"type,omitempty"`
	Items    []*MixedMessage `json:"items"`
	Receiver *Receiver       `json:"receiver,omitempty"`
	QuoteId  string          `json:"quote_id,omitempty"` // 引用的消息ID
}

// 发送文件消息接口请求参数
type MessageFileReq struct {
	Type       string    `json:"type,omitempty"`
	Receiver   *Receiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
	UploadId   string    `json:"upload_id" v:"required"`
}

// 代码消息
type CodeMessageReq struct {
	Type       string    `json:"type,omitempty"`
	Receiver   *Receiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
	Lang       string    `json:"lang" v:"required"`
	Code       string    `json:"code" v:"required"`
}

// 投票消息接口请求参数
type MessageVoteReq struct {
	Type       string    `json:"type,omitempty"`
	Receiver   *Receiver `json:"receiver,omitempty"` // 消息接收者
	ReceiverId int       `json:"receiver_id" v:"required"`
	Mode       int       `json:"mode" v:"in:0,1"`
	Anonymous  int       `json:"anonymous" v:"in:0,1"`
	Title      string    `json:"title" v:"required"`
	Options    []string  `json:"options"`
}

// 位置消息
type LocationMessageReq struct {
	Type        string    `json:"type,omitempty"`
	Longitude   string    `json:"longitude,omitempty" v:"required"`   // 地理位置 经度
	Latitude    string    `json:"latitude,omitempty" v:"required"`    // 地理位置 纬度
	Description string    `json:"description,omitempty" v:"required"` // 位置描述
	Receiver    *Receiver `json:"receiver,omitempty"`                 // 消息接收者
	TalkType    int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId  int       `json:"receiver_id" v:"required"`
}

// 文本消息
type TextMessageReq struct {
	Type       string    `json:"type,omitempty"` // 消息类型
	Content    string    `json:"content,omitempty" v:"required"`
	Mention    *Mention  `json:"mention,omitempty"`
	QuoteId    string    `json:"quote_id,omitempty"` // 引用的消息ID
	Receiver   *Receiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
	Text       string    `json:"text" v:"required"`
}

// 图片消息
type ImageMessageReq struct {
	Type       string    `json:"type,omitempty"`
	Url        string    `json:"url,omitempty" v:"required"`    // 图片地址
	Width      int       `json:"width,omitempty" v:"required"`  // 图片宽度
	Height     int       `json:"height,omitempty" v:"required"` // 图片高度
	Size       int       `json:"size,omitempty" v:"required"`   // 图片大小
	Receiver   *Receiver `json:"receiver,omitempty"`            // 消息接收者
	QuoteId    string    `json:"quote_id,omitempty"`            // 引用的消息ID
	TalkType   int       `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int       `json:"receiver_id" v:"required"`
}

type Receiver struct {
	TalkType   int `json:"talk_type,omitempty"`   // 对话类型 1:私聊 2:群聊
	ReceiverId int `json:"receiver_id,omitempty"` // 接收者ID, 好友ID或群ID
}

type TextMessage struct {
	AckId   string         `json:"ack_id"`
	Event   string         `json:"event"`
	Content TextMessageReq `json:"content"`
}

type CodeMessage struct {
	AckId   string         `json:"ack_id"`
	Event   string         `json:"event"`
	Content CodeMessageReq `json:"content"`
}

type EmoticonMessage struct {
	MsgId   string             `json:"msg_id"`
	Event   string             `json:"event"`
	Content EmoticonMessageReq `json:"content"`
}

type ImageMessage struct {
	MsgId   string          `json:"msg_id"`
	Event   string          `json:"event"`
	Content ImageMessageReq `json:"content"`
}

type FileMessage struct {
	MsgId   string          `json:"msg_id"`
	Event   string          `json:"event"`
	Content ImageMessageReq `json:"content"`
}

type LocationMessage struct {
	MsgId   string             `json:"msg_id"`
	Event   string             `json:"event"`
	Content LocationMessageReq `json:"content"`
}

type VoteMessage struct {
	MsgId   string         `json:"msg_id"`
	Event   string         `json:"event"`
	Content MessageVoteReq `json:"content"`
}

type KeyboardMessage struct {
	Event   string `json:"event"`
	Content struct {
		SenderID   int `json:"sender_id"`
		ReceiverID int `json:"receiver_id"`
	} `json:"content"`
}

type TalkReadMessage struct {
	Event   string `json:"event"`
	Content struct {
		MsgIds     []int `json:"msg_id"`
		ReceiverId int   `json:"receiver_id"`
	} `json:"content"`
}

type Mention struct {
	All  int   `json:"all,omitempty"`
	Uids []int `json:"uids,omitempty"`
}

type MixedMessage struct {
	Type    int    `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}
