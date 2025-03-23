package voice

import (
	"context"
	"heygem-api/internal/service"

	"heygem-api/api/voice/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {
	page, err := service.Voice().Page(ctx, &req.PageInp)
	if err != nil {
		return nil, err
	}
	res = &v1.PageRes{PageOut: page}
	return
}

func (c *ControllerV1) Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error) {
	item, err := service.Voice().Find(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.FindRes{Item: item}
	return
}

func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	out, err := service.Voice().Save(ctx, &req.SaveInp)
	if err != nil {
		return nil, err
	}
	res = &v1.SaveRes{SaveOut: out}
	return
}
