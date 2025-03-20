package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"heygem-file/internal/model/input/filein"
)

type UploadReq struct {
	g.Meta `path:"/upload" tags:"file" method:"post" summary:"上传文件"`
	filein.UploadInp
}
type UploadRes struct {
}
