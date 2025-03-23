// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/voicein"
)

type (
	IVoice interface {
		Page(ctx context.Context, inp *voicein.PageInp) (res *voicein.PageOut, err error)
		Find(ctx context.Context, id int64) (res *entity.Voice, err error)
		Save(ctx context.Context, inp *voicein.SaveInp) (res *voicein.SaveOut, err error)
	}
)

var (
	localVoice IVoice
)

func Voice() IVoice {
	if localVoice == nil {
		panic("implement not found for interface IVoice, forgot register?")
	}
	return localVoice
}

func RegisterVoice(i IVoice) {
	localVoice = i
}
