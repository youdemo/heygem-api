package model

import (
	"context"
	"heygem-api/internal/service"

	"heygem-api/api/model/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {
	page, err := service.Model().Page(ctx, &req.PageInp)
	if err != nil {
		return nil, err
	}
	res = &v1.PageRes{PageOut: page}
	return
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	add, err := service.Model().Add(ctx, &req.AddInp)
	if err != nil {
		return nil, err
	}
	res = &v1.AddRes{AddOut: add}
	return
}
func (c *ControllerV1) Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error) {
	total, err := service.Model().Count(ctx, &req.CountInp)
	res = &v1.CountRes{Count: total}
	return
}
func (c *ControllerV1) Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error) {
	model, err := service.Model().Find(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.FindRes{Item: model}
	return
}
func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.Model().Del(ctx, req.Id)
	return
}
