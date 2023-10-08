package entity

type TalkRecords struct {
	Id         string `bson:"_id"`         // ID
	RecordId   int    `bson:"record_id"`   // 记录ID
	MsgId      string `bson:"msg_id"`      // 消息唯一ID
	Sequence   int    `bson:"sequence"`    // 消息时序ID
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

type TalkRecordsVote struct {
	Id           string `bson:"_id"`           // 投票ID
	RecordId     int    `bson:"record_id"`     // 消息记录ID
	UserId       int    `bson:"user_id"`       // 用户ID
	Title        string `bson:"title"`         // 投票标题
	AnswerMode   int    `bson:"answer_mode"`   // 答题模式[0:单选;1:多选;]
	AnswerOption string `bson:"answer_option"` // 答题选项
	AnswerNum    int    `bson:"answer_num"`    // 应答人数
	AnsweredNum  int    `bson:"answered_num"`  // 已答人数
	IsAnonymous  int    `bson:"is_anonymous"`  // 匿名投票[0:否;1:是;]
	Status       int    `bson:"status"`        // 投票状态[0:投票中;1:已完成;]
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 更新时间
}

type TalkRecordsVoteAnswer struct {
	Id        string `bson:"_id"`        // 答题ID
	VoteId    string `bson:"vote_id"`    // 投票ID
	UserId    int    `bson:"user_id"`    // 用户ID
	Option    string `bson:"option"`     // 投票选项[A、B、C 、D、E、F]
	CreatedAt int64  `bson:"created_at"` // 答题时间
}
