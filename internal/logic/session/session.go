package session

import (
	"context"
	"github.com/iimeta/iim-server/internal/service"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

// 获取会话中UserId
func (s *sSession) GetUid(ctx context.Context) int {
	return ctx.Value("uid").(int)
}
