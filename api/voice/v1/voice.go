package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/voicein"
)

type PageReq struct {
	g.Meta `path:"/page" tags:"语音管理" method:"get" summary:"语音列表"`
	voicein.PageInp
}
type PageRes struct {
	*voicein.PageOut
}

type FindReq struct {
	g.Meta `path:"/find/:id" tags:"语音管理" method:"get" summary:"查找语音"`
	Id     int64 `json:"id" v:"required"`
}

type FindRes struct {
	Item *entity.Voice
}

type SaveReq struct {
	g.Meta `path:"/save" tags:"语音管理" method:"post" summary:"保存语音"`
	voicein.SaveInp
}

type SaveRes struct {
	*voicein.SaveOut
}

type DelReq struct {
	g.Meta `path:"/del/:id" tags:"语音管理" method:"delete" summary:"删除语音"`
	Id     int64 `json:"id" v:"required"`
}

type DelRes struct{}
