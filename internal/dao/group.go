package dao

import (
	"context"
	"github.com/iimeta/iim-server/internal/model/do"
	"github.com/iimeta/iim-server/internal/model/entity"
	"github.com/iimeta/iim-server/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var Group = NewGroupDao()

type GroupDao struct {
	*MongoDB[entity.Group]
}

func NewGroupDao(database ...string) *GroupDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupDao{
		MongoDB: NewMongoDB[entity.Group](database[0], do.GROUP_COLLECTION),
	}
}

// 根据groupId查询群信息
func (d *GroupDao) FindGroupByGroupId(ctx context.Context, groupId int) (*entity.Group, error) {

	group, err := d.FindOne(ctx, bson.M{"group_id": groupId})
	if err != nil {
		return nil, err
	}

	return group, nil
}
