package entity

type Group struct {
	Id        string `bson:"_id,omitempty"`        // ID
	GroupId   int    `bson:"group_id,omitempty"`   // 群聊ID
	Type      int    `bson:"type,omitempty"`       // 群类型[1:普通群;2:企业群;]
	CreatorId int    `bson:"creator_id,omitempty"` // 创建者ID(群主ID)
	GroupName string `bson:"group_name,omitempty"` // 群名称
	Profile   string `bson:"profile,omitempty"`    // 群介绍
	IsDismiss int    `bson:"is_dismiss,omitempty"` // 是否已解散[0:否;1:是;]
	Avatar    string `bson:"avatar,omitempty"`     // 群头像
	MaxNum    int    `bson:"max_num,omitempty"`    // 最大群成员数量
	IsOvert   int    `bson:"is_overt,omitempty"`   // 是否公开可见[0:否;1:是;]
	IsMute    int    `bson:"is_mute,omitempty"`    // 是否全员禁言 [0:否;1:是;], 提示:不包含群主或管理员
	CreatedAt int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `bson:"updated_at,omitempty"` // 更新时间
}

type GroupMember struct {
	Id          string `bson:"_id,omitempty"`           // ID
	GroupId     int    `bson:"group_id,omitempty"`      // 群聊ID
	UserId      int    `bson:"user_id,omitempty"`       // 用户ID
	Leader      int    `bson:"leader,omitempty"`        // 成员属性[0:普通成员;1:管理员;2:群主;]
	UserCard    string `bson:"user_card,omitempty"`     // 群名片
	IsQuit      int    `bson:"is_quit,omitempty"`       // 是否退群[0:否;1:是;]
	IsMute      int    `bson:"is_mute,omitempty"`       // 是否禁言[0:否;1:是;]
	MinRecordId int    `bson:"min_record_id,omitempty"` // 可查看历史记录最小ID
	JoinTime    int64  `bson:"join_time,omitempty"`     // 入群时间
	CreatedAt   int64  `bson:"created_at,omitempty"`    // 创建时间
	UpdatedAt   int64  `bson:"updated_at,omitempty"`    // 更新时间
}
