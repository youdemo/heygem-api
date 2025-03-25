package tts

import (
	"context"
	"heygem-api/internal/service"

	"heygem-api/api/tts/v1"
)

func (c *ControllerV1) Train(ctx context.Context, req *v1.TrainReq) (res *v1.TrainRes, err error) {
	train, err := service.Tts().Train(ctx, &req.TrainInp)
	if err != nil {
		return nil, err
	}
	res = &v1.TrainRes{TrainOut: train}
	return
}
func (c *ControllerV1) Invoke(ctx context.Context, req *v1.InvokeReq) (res *v1.InvokeRes, err error) {
	_, err = service.Tts().Invoke(ctx, &req.InvokeInp)
	if err != nil {
		return nil, err
	}
	return
}
