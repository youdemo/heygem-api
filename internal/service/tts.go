// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"heygem-api/internal/model/input/ttsin"
)

type (
	ITts interface {
		Train(ctx context.Context, inp *ttsin.TrainInp) (res *ttsin.TrainOut, err error)
		Invoke(ctx context.Context, inp *ttsin.InvokeInp) (res *ttsin.InvokeOut, err error)
	}
)

var (
	localTts ITts
)

func Tts() ITts {
	if localTts == nil {
		panic("implement not found for interface ITts, forgot register?")
	}
	return localTts
}

func RegisterTts(i ITts) {
	localTts = i
}
