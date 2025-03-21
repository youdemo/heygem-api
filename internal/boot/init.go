package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	VideoPath = ""
	AudioPath = ""
)

func Init(ctx context.Context) {
	VideoPath = g.Config().MustGet(ctx, "path.video").String()
	AudioPath = g.Config().MustGet(ctx, "path.audio").String()
}
