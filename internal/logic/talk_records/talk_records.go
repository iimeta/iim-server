package talk_records

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-server/internal/dao"
	"github.com/iimeta/iim-server/internal/model"
	"github.com/iimeta/iim-server/internal/service"
	"github.com/iimeta/iim-server/utility/logger"
)

type sTalkRecords struct{}

func init() {
	service.RegisterTalkRecords(New())
}

func New() service.ITalkRecords {
	return &sTalkRecords{}
}

func (s *sTalkRecords) GetTalkRecord(ctx context.Context, recordId int) (*model.TalkRecordsItem, error) {

	record, user, err := dao.TalkRecords.GetTalkRecord(ctx, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	talkRecordsItem := &model.TalkRecordsItem{
		Id:         record.RecordId,
		Sequence:   record.Sequence,
		MsgId:      record.MsgId,
		TalkType:   record.TalkType,
		MsgType:    record.MsgType,
		UserId:     record.UserId,
		ReceiverId: record.ReceiverId,
		IsRevoke:   record.IsRevoke,
		IsMark:     record.IsMark,
		IsRead:     record.IsRead,
		Content:    record.Content,
		CreatedAt:  gtime.NewFromTimeStamp(record.CreatedAt).String(),
		Extra:      make(map[string]any),
	}

	if user != nil {
		talkRecordsItem.Nickname = user.Nickname
		talkRecordsItem.Avatar = user.Avatar
	}

	if record.Extra != "" {
		talkRecordsItem.Extra, err = gjson.Decode(record.Extra)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	items, err := s.HandleTalkRecords(ctx, []*model.TalkRecordsItem{talkRecordsItem})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return items[0], nil
}

func (s *sTalkRecords) HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error) {

	talkRecordsItems, err := dao.TalkRecords.HandleTalkRecords(ctx, items)
	if err != nil {
		logger.Error(ctx)
		return nil, err
	}

	return talkRecordsItems, nil
}
