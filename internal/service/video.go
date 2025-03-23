// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/videoin"
)

type (
	IVideo interface {
		Page(ctx context.Context, inp *videoin.PageInp) (res *videoin.PageOut, err error)
		SelectByStatus(ctx context.Context, inp *videoin.SelectByStatusInp) (res []*entity.Video, err error)
		Count(ctx context.Context, name string) (total int, err error)
		Find(ctx context.Context, id int64) (res *entity.Video, err error)
		FindByStatus(ctx context.Context, status string) (res *entity.Video, err error)
		Save(ctx context.Context, inp *videoin.SaveInp) (id int64, err error)
		Del(ctx context.Context, id int64) (err error)
		MakeByF2F(ctx context.Context, id int64) (err error)
		// LoopPending 处理未合成的视频
		LoopPending(ctx context.Context)
	}
)

var (
	localVideo IVideo
)

func Video() IVideo {
	if localVideo == nil {
		panic("implement not found for interface IVideo, forgot register?")
	}
	return localVideo
}

func RegisterVideo(i IVideo) {
	localVideo = i
}
