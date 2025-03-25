package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gfile"
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
	InitWeb(ctx)
	_, _ = gcron.AddSingleton(ctx, "*/5 * * * * *", func(ctx context.Context) {
		service.Video().LoopPending(ctx)
	})
}

func InitWeb(ctx context.Context) {
	if gfile.Exists("resource/assets") {
		if gfile.Exists("resource/public/assets") {
			_ = gfile.RemoveAll("resource/public/assets")
		}
		_ = gfile.Move("resource/assets", "resource/public/assets")
		_ = gfile.Move("resource/favicon.ico", "resource/public/")
		_ = gfile.Move("resource/index.html", "resource/public/")
	}

}
