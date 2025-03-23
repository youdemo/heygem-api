package cmd

import (
	"context"
	"heygem-api/internal/controller/model"
	"heygem-api/internal/controller/tts"
	"heygem-api/internal/controller/video"
	"heygem-api/internal/controller/voice"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/model", func(group *ghttp.RouterGroup) {
					group.Bind(model.NewV1())
				})
				group.Group("/voice", func(group *ghttp.RouterGroup) {
					group.Bind(voice.NewV1())
				})
				group.Group("/tts", func(group *ghttp.RouterGroup) {
					group.Bind(tts.NewV1())
				})
				group.Group("/video", func(group *ghttp.RouterGroup) {
					group.Bind(video.NewV1())
				})
			})
			s.Run()
			return
		},
	}
)
