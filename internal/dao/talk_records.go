package dao

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-server/internal/consts"
	"github.com/iimeta/iim-server/internal/model"
	"github.com/iimeta/iim-server/internal/model/do"
	"github.com/iimeta/iim-server/internal/model/entity"
	"github.com/iimeta/iim-server/utility/db"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

var TalkRecords = NewTalkRecordsDao()

type TalkRecordsDao struct {
	*MongoDB[entity.TalkRecords]
}

func NewTalkRecordsDao(database ...string) *TalkRecordsDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TalkRecordsDao{
		MongoDB: NewMongoDB[entity.TalkRecords](database[0], do.TALK_RECORDS_COLLECTION),
	}
}

// 获取对话消息
func (d *TalkRecordsDao) GetTalkRecord(ctx context.Context, recordId int) (*entity.TalkRecords, *entity.User, error) {

	talkRecords, err := d.FindOne(ctx, bson.M{"record_id": recordId})
	if err != nil {
		return nil, nil, err
	}

	if talkRecords.UserId == 0 {
		return talkRecords, nil, err
	}

	user, err := User.FindUserByUserId(ctx, talkRecords.UserId)
	if err != nil {
		return nil, nil, err
	}

	return talkRecords, user, nil
}

func (d *TalkRecordsDao) HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error) {

	votes := make([]int, 0)
	for _, item := range items {
		switch item.MsgType {
		case consts.ChatMsgTypeVote:
			votes = append(votes, item.Id)
		}
	}

	hashVotes := make(map[int]*entity.TalkRecordsVote)
	if len(votes) > 0 {

		talkRecordsVoteList, err := TalkRecordsVote.Find(ctx, bson.M{"record_id": bson.M{"$in": votes}})
		if err != nil {
			return nil, err
		}

		for _, vote := range talkRecordsVoteList {
			hashVotes[vote.RecordId] = vote
		}
	}

	newItems := make([]*model.TalkRecordsItem, 0, len(items))
	for _, item := range items {

		data := &model.TalkRecordsItem{
			Id:         item.Id,
			MsgId:      item.MsgId,
			Sequence:   item.Sequence,
			TalkType:   item.TalkType,
			MsgType:    item.MsgType,
			UserId:     item.UserId,
			ReceiverId: item.ReceiverId,
			Nickname:   item.Nickname,
			Avatar:     item.Avatar,
			IsRevoke:   item.IsRevoke,
			IsMark:     item.IsMark,
			IsRead:     item.IsRead,
			Content:    item.Content,
			CreatedAt:  item.CreatedAt,
			Extra:      make(map[string]any),
		}

		_ = gjson.Unmarshal(gjson.MustEncode(item.Extra), &data.Extra)

		switch item.MsgType {
		case consts.ChatMsgTypeVote:

			if value, ok := hashVotes[item.Id]; ok {

				options := make(map[string]any)
				opts := make([]any, 0)

				if err := gjson.Unmarshal([]byte(value.AnswerOption), &options); err == nil {
					arr := make([]string, 0, len(options))
					for k := range options {
						arr = append(arr, k)
					}

					sort.Strings(arr)

					for _, v := range arr {
						opts = append(opts, map[string]any{
							"key":   v,
							"value": options[v],
						})
					}
				}

				users := make([]int, 0)
				if uids, err := TalkRecordsVote.GetVoteAnswerUser(ctx, value.Id); err == nil {
					users = uids
				}

				var statistics any

				if res, err := TalkRecordsVote.GetVoteStatistics(ctx, value.Id); err != nil {
					statistics = map[string]any{
						"count":   0,
						"options": map[string]int{},
					}
				} else {
					statistics = res
				}

				data.Extra = map[string]any{
					"detail": map[string]any{
						"id":            value.Id,
						"record_id":     value.RecordId,
						"title":         value.Title,
						"answer_mode":   value.AnswerMode,
						"status":        value.Status,
						"answer_option": opts,
						"answer_num":    value.AnswerNum,
						"answered_num":  value.AnsweredNum,
					},
					"statistics": statistics,
					"vote_users": users, // 已投票成员
				}
			}
		}

		newItems = append(newItems, data)
	}

	return newItems, nil
}

func (d *TalkRecordsDao) FindByRecordId(ctx context.Context, recordId int) (*entity.TalkRecords, error) {
	return d.FindOne(ctx, bson.M{"record_id": recordId})
}
