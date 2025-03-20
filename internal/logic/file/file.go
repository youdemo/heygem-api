package file

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/os/gfile"
	"heygem-file/internal/model/input/filein"
	"heygem-file/internal/service"
)

type sFile struct {
}

func init() {
	service.RegisterFile(&sFile{})
}

func (s *sFile) Upload(ctx context.Context, inp *filein.UploadInp) (err error) {

	fileName := gfile.Basename(inp.Path)
	if inp.FileType == "video" {
		fileName = "/data/models/heygem/data/voice/data/" + fileName
	} else {
		fileName = "/data/models/heygem/data/face2face/" + fileName
	}
	// 上传文件到远程服务器
	remoteFilePath := fileName
	// 解析文件
	fileContent := gbase64.MustDecodeString(inp.File)
	// 创建文件
	return gfile.PutBytes(remoteFilePath, fileContent)
}
