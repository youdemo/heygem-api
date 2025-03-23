package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/modelin"
)

type PageReq struct {
	g.Meta `path:"/page" tags:"模特管理" method:"get" summary:"模特列表"`
	modelin.PageInp
}
type PageRes struct {
	*modelin.PageOut
}

type CountReq struct {
	g.Meta `path:"/count" tags:"模特管理" method:"get" summary:"模特数量"`
	modelin.CountInp
}
type CountRes struct {
	Count int `json:"count"`
}

type FindReq struct {
	g.Meta `path:"/find" tags:"模特管理" method:"get" summary:"查找模特"`
	Id     int64 `json:"id" dc:"模特ID"`
}
type FindRes struct {
	Item *entity.Model `json:"item" dc:"模特"`
}

type AddReq struct {
	g.Meta `path:"/add" tags:"模特管理" method:"post" summary:"添加模特"`
	modelin.AddInp
}
type AddRes struct {
	*modelin.AddOut
}

type DelReq struct {
	g.Meta `path:"/del" tags:"模特管理" method:"delete" summary:"删除"`
	Id     int64 `json:"id" dc:"模特ID"`
}
type DelRes struct{}
