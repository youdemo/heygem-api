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

type DownloadReq struct {
	g.Meta `path:"/download" tags:"file" method:"post" summary:"下载文件"`
	filein.DownloadInp
}
type DownloadRes struct {
}
