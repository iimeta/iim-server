package server

import (
	_ "github.com/iimeta/iim-server/internal/core"

	_ "github.com/iimeta/iim-server/internal/packed"

	_ "github.com/iimeta/iim-server/internal/logic"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-server/internal/cmd"
)

// 提供给iim-client引入使用
func init() {
	g.Server().BindHandler("/wss/default.io", cmd.Server)
}
