package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/videoin"
)

type PageReq struct {
	g.Meta `path:"/page" tags:"视频管理" method:"get" summary:"视频分页列表"`
	videoin.PageInp
}

type PageRes struct {
	*videoin.PageOut
}

type CountReq struct {
	g.Meta `path:"/count" tags:"视频管理" method:"get" summary:"视频数量"`
	Name   string `p:"name" dc:"名称"`
}

type CountRes struct {
	Count int `json:"count"`
}

type SelectByStatusReq struct {
	g.Meta `path:"/selectByStatus" tags:"视频管理" method:"get" summary:"根据状态查询视频"`
	videoin.SelectByStatusInp
}

type SelectByStatusRes struct {
	List []*entity.Video `json:"list"`
}

type FindReq struct {
	g.Meta `path:"/find" tags:"视频管理" method:"get" summary:"查询视频"`
	Id     int64 `json:"id" v:"required" dc:"id"`
}

type FindRes struct {
	Item *entity.Video `json:"item"`
}

type FindByStatusReq struct {
	g.Meta `path:"/findByStatus" tags:"视频管理" method:"get" summary:"根据状态查询视频"`
	Status string `json:"status" v:"required" dc:"状态"`
}

type FindByStatusRes struct {
	Item *entity.Video `json:"item"`
}

type SaveReq struct {
	g.Meta `path:"/save" tags:"视频管理" method:"post" summary:"保存视频"`
	videoin.SaveInp
}

type SaveRes struct {
	Id int64 `json:"id" dc:"视频ID"`
}

type DelReq struct {
	g.Meta `path:"/del/:id" tags:"视频管理" method:"delete" summary:"删除视频"`
	Id     int64 `json:"id" v:"required" dc:"id"`
}

type DelRes struct{}

type MakeByF2FReq struct {
	g.Meta `path:"/makeByF2F/:id" tags:"视频管理" method:"post" summary:"合成视频"`
	Id     int64 `json:"id" v:"required" dc:"id"`
}

type MakeByF2FRes struct{}
