package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/os/gtime"
)

var node *snowflake.Node

func init() {

	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

}

func GenerateId() string {
	return node.Generate().String()
}

func NewMsgId() string {
	return fmt.Sprintf("%s_%d", GenerateId(), gtime.Timestamp())
}
