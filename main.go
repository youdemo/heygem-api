package main

import (
	_ "heygem-file/internal/packed"

	_ "heygem-file/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"heygem-file/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
