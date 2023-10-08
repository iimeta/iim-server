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
	IRobot interface {
		GetRobotByUserId(ctx context.Context, userId int) (*model.Robot, error)
	}
)

var (
	localRobot IRobot
)

func Robot() IRobot {
	if localRobot == nil {
		panic("implement not found for interface IRobot, forgot register?")
	}
	return localRobot
}

func RegisterRobot(i IRobot) {
	localRobot = i
}
