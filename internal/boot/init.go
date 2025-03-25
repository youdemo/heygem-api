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
	if g.Config().MustGet(ctx, "heygem.web.enabled").Bool() && gfile.Exists("/app/resource/assets") {
		g.Log().Info(ctx, "移动前端文件")
		if gfile.Exists("/app/resource/public/assets") {
			_ = gfile.RemoveAll("/app/resource/public/assets")
		}
		_ = gfile.Move("/app/resource/assets/", "/app/resource/public/")
		_ = gfile.Move("/app/resource/favicon.ico", "/app/resource/public/")
		_ = gfile.Move("/app/resource/index.html", "/app/resource/public/")
	}

}
