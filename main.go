package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"heygem-api/internal/boot"
	_ "heygem-api/internal/logic"
	_ "heygem-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"heygem-api/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.Init(ctx)
	cmd.Main.Run(ctx)
}
