package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"heygem-api/internal/service"
)

var (
	VideoPath = ""
	TTSApi    = ""
	F2FApi    = ""
)

func Init(ctx context.Context) {
	VideoPath = g.Config().MustGet(ctx, "heygem.path.video").String()
	TTSApi = g.Config().MustGet(ctx, "heygem.tts").String()
	F2FApi = g.Config().MustGet(ctx, "heygem.f2f").String()
	_, _ = gcron.AddSingleton(ctx, "*/5 * * * * *", func(ctx context.Context) {
		service.Video().LoopPending(ctx)
	})
}
