package dao

import (
	"context"
	"github.com/iimeta/iim-server/internal/consts"
	"github.com/iimeta/iim-server/internal/model/do"
	"github.com/iimeta/iim-server/internal/model/entity"
	"github.com/iimeta/iim-server/utility/db"
	"github.com/iimeta/iim-server/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var GroupMember = NewGroupMemberDao()

type GroupMemberDao struct {
	*MongoDB[entity.GroupMember]
}

func NewGroupMemberDao(database ...string) *GroupMemberDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupMemberDao{
		MongoDB: NewMongoDB[entity.GroupMember](database[0], do.GROUP_MEMBER_COLLECTION),
	}
}

// 获取所有群成员用户ID
func (d *GroupMemberDao) GetMemberIds(ctx context.Context, groupId int) []int {

	groupMemberList, err := d.Find(ctx, bson.M{"group_id": groupId, "is_quit": bson.M{"$ne": consts.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, member := range groupMemberList {
		ids = append(ids, member.UserId)
	}

	return ids
}

// 获取所有群成员ID
func (d *GroupMemberDao) GetUserGroupIds(ctx context.Context, uid int) []int {

	groupMemberList, err := d.Find(ctx, bson.M{"user_id": uid, "is_quit": bson.M{"$ne": consts.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, member := range groupMemberList {
		ids = append(ids, member.GroupId)
	}

	return ids
}
