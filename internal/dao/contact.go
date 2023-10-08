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

var Contact = NewContactDao()

type ContactDao struct {
	*MongoDB[entity.Contact]
}

func NewContactDao(database ...string) *ContactDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ContactDao{
		MongoDB: NewMongoDB[entity.Contact](database[0], do.CONTACT_COLLECTION),
	}
}

func (d *ContactDao) GetContactIds(ctx context.Context, uid int) []int {

	filter := bson.M{
		"user_id": uid,
		"status":  consts.ContactStatusNormal,
	}

	contactList, err := d.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, contact := range contactList {
		ids = append(ids, contact.FriendId)
	}

	return ids
}
