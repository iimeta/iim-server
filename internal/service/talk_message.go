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
	ITalkMessage interface {
		// 文本消息
		SendText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// 代码消息
		SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error
	}
)

var (
	localTalkMessage ITalkMessage
)

func TalkMessage() ITalkMessage {
	if localTalkMessage == nil {
		panic("implement not found for interface ITalkMessage, forgot register?")
	}
	return localTalkMessage
}

func RegisterTalkMessage(i ITalkMessage) {
	localTalkMessage = i
}
