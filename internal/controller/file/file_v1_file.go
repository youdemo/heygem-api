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

func (c *ControllerV1) Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error) {
	err = service.File().Download(ctx, &req.DownloadInp)
	return
}
