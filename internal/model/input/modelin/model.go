package modelin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"heygem-api/internal/model"
	"heygem-api/internal/model/entity"
)

type PageInp struct {
	model.PageReq
	Name string `json:"name" dc:"模特名称"`
}

type PageOut struct {
	model.PageRes
	List []*entity.Model `json:"list" dc:"模特列表"`
}

type CountInp struct {
	Name string `json:"name" dc:"模特名称"`
}

type AddInp struct {
	File *ghttp.UploadFile `json:"file" type:"file" v:"required" dc:"视频文件"`
	Name string            `json:"name" v:"required" dc:"模特名称"`
}

type AddOut struct {
	Id   int64  `json:"id"`
	Path string `json:"path"`
}
