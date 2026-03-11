package main

import (
	_ "chatBox/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"chatBox/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
