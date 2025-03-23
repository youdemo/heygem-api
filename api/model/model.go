// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"context"

	"heygem-api/api/model/v1"
)

type IModelV1 interface {
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error)
	Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
