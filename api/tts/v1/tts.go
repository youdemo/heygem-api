package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"heygem-api/internal/model/input/ttsin"
)

type TrainReq struct {
	g.Meta `path:"/train" tags:"tts" method:"post" summary:"解析语音文本"`
	ttsin.TrainInp
}
type TrainRes struct {
	*ttsin.TrainOut
}

type InvokeReq struct {
	g.Meta `path:"/invoke" tags:"tts" method:"post" summary:"声音克隆"`
	ttsin.InvokeInp
}
type InvokeRes struct {
	*ttsin.InvokeOut
}
