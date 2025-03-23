package video

import (
	"context"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/service"

	"heygem-api/api/video/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {
	page, err := service.Video().Page(ctx, &req.PageInp)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{PageOut: page}
	return
}

func (c *ControllerV1) SelectByStatus(ctx context.Context, req *v1.SelectByStatusReq) (res *v1.SelectByStatusRes, err error) {
	out, err := service.Video().SelectByStatus(ctx, &req.SelectByStatusInp)
	if err != nil {
		return nil, err
	}
	res = &v1.SelectByStatusRes{}
	if out == nil {
		res.List = make([]*entity.Video, 0)
	} else {
		res.List = out
	}
	return
}

func (c *ControllerV1) Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error) {
	total, err := service.Video().Count(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	res = &v1.CountRes{Count: total}
	return
}

func (c *ControllerV1) Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error) {
	item, err := service.Video().Find(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.FindRes{Item: item}
	return
}

func (c *ControllerV1) FindByStatus(ctx context.Context, req *v1.FindByStatusReq) (res *v1.FindByStatusRes, err error) {
	item, err := service.Video().FindByStatus(ctx, req.Status)
	if err != nil {
		return nil, err
	}
	res = &v1.FindByStatusRes{Item: item}
	return
}

func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	id, err := service.Video().Save(ctx, &req.SaveInp)
	if err != nil {
		return nil, err
	}
	res = &v1.SaveRes{Id: id}
	return
}
func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.Video().Del(ctx, req.Id)
	return
}
func (c *ControllerV1) MakeByF2F(ctx context.Context, req *v1.MakeByF2FReq) (res *v1.MakeByF2FRes, err error) {
	err = service.Video().MakeByF2F(ctx, req.Id)
	return
}
