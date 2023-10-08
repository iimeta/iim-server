package entity

type Contact struct {
	Id        string `bson:"_id,omitempty"`        // 关系ID
	UserId    int    `bson:"user_id,omitempty"`    // 用户id
	FriendId  int    `bson:"friend_id,omitempty"`  // 好友id
	Remark    string `bson:"remark,omitempty"`     // 好友的备注
	Status    int    `bson:"status,omitempty"`     // 好友状态 [0:否;1:是]
	GroupId   string `bson:"group_id,omitempty"`   // 分组id
	CreatedAt int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `bson:"updated_at,omitempty"` // 更新时间
}

type ContactApply struct {
	Id        string `bson:"_id,omitempty"`              // 申请ID
	UserId    int    `bson:"user_id,omitempty"`          // 申请人ID
	Nickname  string `bson:"nickname,omitempty"`         // 申请人昵称
	Avatar    string `bson:"avatar,omitempty,omitempty"` // 申请人头像地址
	FriendId  int    `bson:"friend_id,omitempty"`        // 被申请人
	Remark    string `bson:"remark,omitempty"`           // 申请备注
	CreatedAt int64  `bson:"created_at,omitempty"`       // 申请时间
}
