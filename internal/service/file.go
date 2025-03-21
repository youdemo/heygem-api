// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"heygem-file/internal/model/input/filein"
)

type (
	IFile interface {
		Upload(ctx context.Context, inp *filein.UploadInp) (err error)
		Download(ctx context.Context, inp *filein.DownloadInp) (err error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
