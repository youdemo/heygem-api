package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"heygem-api/internal/service"
)

var (
	VideoPath = ""
	AudioPath = ""
	TTSApi    = ""
	F2FApi    = ""
)

func Init(ctx context.Context) {
	VideoPath = g.Config().MustGet(ctx, "heygem.path.video").String()
	AudioPath = g.Config().MustGet(ctx, "heygem.path.audio").String()
	TTSApi = g.Config().MustGet(ctx, "heygem.tts").String()
	F2FApi = g.Config().MustGet(ctx, "heygem.f2f").String()
	_ = grpool.AddWithRecover(gctx.New(), func(ctx context.Context) {
		service.Video().LoopPending(ctx)
	}, nil)
}
