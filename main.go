package main

import (
	"heygem-file/internal/boot"
	_ "heygem-file/internal/packed"

	_ "heygem-file/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"heygem-file/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.Init(ctx)
	cmd.Main.Run(ctx)
}
