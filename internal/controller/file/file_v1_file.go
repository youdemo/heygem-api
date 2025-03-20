package file

import (
	"context"
	v1 "heygem-file/api/file/v1"
	"heygem-file/internal/service"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	err = service.File().Upload(ctx, &req.UploadInp)
	return
}
