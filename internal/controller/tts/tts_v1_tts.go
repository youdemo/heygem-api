package tts

import (
	"context"
	"heygem-api/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"heygem-api/api/tts/v1"
)

func (c *ControllerV1) Train(ctx context.Context, req *v1.TrainReq) (res *v1.TrainRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) Invoke(ctx context.Context, req *v1.InvokeReq) (res *v1.InvokeRes, err error) {
	_, err = service.Tts().Invoke(ctx, &req.InvokeInp)
	if err != nil {
		return nil, err
	}
	return
}
