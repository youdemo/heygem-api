package file

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"heygem-file/internal/boot"
	"heygem-file/internal/consts"
	"heygem-file/internal/model/input/filein"
	"heygem-file/internal/service"
)

type sFile struct {
}

func init() {
	service.RegisterFile(&sFile{})
}

func (s *sFile) Upload(ctx context.Context, inp *filein.UploadInp) (err error) {
	g.Log().Info(ctx, inp.Path)
	remoteFilePath := ""
	if inp.FileType == consts.AudioType {
		remoteFilePath = boot.AudioPath + inp.Path
	} else {
		remoteFilePath = boot.VideoPath + inp.Path
	}
	// 解析文件
	fileContent := gbase64.MustDecodeString(inp.File)
	// 创建文件
	return gfile.PutBytes(remoteFilePath, fileContent)
}

func (s *sFile) Download(ctx context.Context, inp *filein.DownloadInp) (err error) {
	g.Log().Info(ctx, inp.Path)
	remoteFilePath := ""
	if inp.FileType == consts.AudioType {
		remoteFilePath = boot.AudioPath + inp.Path
	} else {
		remoteFilePath = boot.VideoPath + inp.Path
	}
	// 解析文件
	g.RequestFromCtx(ctx).Response.ServeFileDownload(remoteFilePath)
	// 创建文件
	return
}
