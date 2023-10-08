// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-server/internal/model"
)

type (
	ITalkRecords interface {
		GetTalkRecord(ctx context.Context, recordId int) (*model.TalkRecordsItem, error)
		HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error)
	}
)

var (
	localTalkRecords ITalkRecords
)

func TalkRecords() ITalkRecords {
	if localTalkRecords == nil {
		panic("implement not found for interface ITalkRecords, forgot register?")
	}
	return localTalkRecords
}

func RegisterTalkRecords(i ITalkRecords) {
	localTalkRecords = i
}
