package core

import (
	"context"
	"github.com/iimeta/iim-server/utility/logger"
	"github.com/iimeta/iim-server/utility/redis"
)

const (
	RECORD_ID_AUTO_INCREMENT_KEY = "CORE:RECORD_ID_AUTO_INCREMENT"
)

func IncrRecordId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, RECORD_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}
