package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	TALK_RECORDS_COLLECTION             = "talk_records"
	TALK_RECORDS_VOTE_COLLECTION        = "talk_records_vote"
	TALK_RECORDS_VOTE_ANSWER_COLLECTION = "talk_records_vote_answer"
)

type TalkRecords struct {
	gmeta.Meta `collection:"talk_records" bson:"-"`
	RecordId   int    `bson:"record_id"`   // 记录ID
	MsgId      string `bson:"msg_id"`      // 消息唯一ID
	Sequence   int64  `bson:"sequence"`    // 消息时序ID
	TalkType   int    `bson:"talk_type"`   // 对话类型[1:私信;2:群聊;]
	MsgType    int    `bson:"msg_type"`    // 消息类型
	UserId     int    `bson:"user_id"`     // 发送者ID[0:系统用户;]
	ReceiverId int    `bson:"receiver_id"` // 接收者ID(用户ID 或 群ID)
	IsRevoke   int    `bson:"is_revoke"`   // 是否撤回消息[0:否;1:是;]
	IsMark     int    `bson:"is_mark"`     // 是否重要消息[0:否;1:是;]
	IsRead     int    `bson:"is_read"`     // 是否已读[0:否;1:是;]
	QuoteId    string `bson:"quote_id"`    // 引用消息ID
	Content    string `bson:"content"`     // 文本消息
	Extra      string `bson:"extra"`       // 扩展信信息
	CreatedAt  int64  `bson:"created_at"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at"`  // 更新时间
}
