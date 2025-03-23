// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package video

import (
	"context"

	"heygem-api/api/video/v1"
)

type IVideoV1 interface {
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error)
	SelectByStatus(ctx context.Context, req *v1.SelectByStatusReq) (res *v1.SelectByStatusRes, err error)
	Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error)
	FindByStatus(ctx context.Context, req *v1.FindByStatusReq) (res *v1.FindByStatusRes, err error)
	Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	MakeByF2F(ctx context.Context, req *v1.MakeByF2FReq) (res *v1.MakeByF2FRes, err error)
}
