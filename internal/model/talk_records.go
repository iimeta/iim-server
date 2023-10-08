package model

type TalkRecord struct {
	Id         string `json:"id"`          // ID
	RecordId   int    `json:"record_id"`   // 记录ID
	MsgId      string `json:"msg_id"`      // 消息唯一ID
	Sequence   int64  `json:"sequence"`    // 消息时序ID
	TalkType   int    `json:"talk_type"`   // 对话类型[1:私信;2:群聊;]
	MsgType    int    `json:"msg_type"`    // 消息类型
	UserId     int    `json:"user_id"`     // 发送者ID[0:系统用户;]
	ReceiverId int    `json:"receiver_id"` // 接收者ID(用户ID 或 群ID)
	IsRevoke   int    `json:"is_revoke"`   // 是否撤回消息[0:否;1:是;]
	IsMark     int    `json:"is_mark"`     // 是否重要消息[0:否;1:是;]
	IsRead     int    `json:"is_read"`     // 是否已读[0:否;1:是;]
	QuoteId    string `json:"quote_id"`    // 引用消息ID
	Content    string `json:"content"`     // 文本消息
	Extra      string `json:"extra"`       // 扩展信信息
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 更新时间
}

type TalkRecordsItem struct {
	Id         int    `json:"id"`
	Sequence   int    `json:"sequence"`
	MsgId      string `json:"msg_id"`
	TalkType   int    `json:"talk_type"`
	MsgType    int    `json:"msg_type"`
	UserId     int    `json:"user_id"`
	ReceiverId int    `json:"receiver_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	IsRevoke   int    `json:"is_revoke"`
	IsMark     int    `json:"is_mark"`
	IsRead     int    `json:"is_read"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Extra      any    `json:"extra"` // 额外参数
}

type TalkRecordReply struct {
	UserId   int    `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	MsgType  int    `json:"msg_type,omitempty"` // 1:文字 2:图片
	Content  string `json:"content,omitempty"`  // 文字或图片连接
	MsgId    string `json:"msg_id,omitempty"`
}

type TalkRecordCode struct {
	Lang string `json:"lang"` // 代码语言
	Code string `json:"code"` // 代码内容
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}
