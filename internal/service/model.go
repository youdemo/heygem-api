// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/modelin"
)

type (
	IModel interface {
		Page(ctx context.Context, inp *modelin.PageInp) (res *modelin.PageOut, err error)
		Count(ctx context.Context, inp *modelin.CountInp) (total int, err error)
		Find(ctx context.Context, id int64) (res *entity.Model, err error)
		// Add 添加模特
		Add(ctx context.Context, inp *modelin.AddInp) (res *modelin.AddOut, err error)
		// Del 删除模特
		Del(ctx context.Context, id int64) (err error)
	}
)

var (
	localModel IModel
)

func Model() IModel {
	if localModel == nil {
		panic("implement not found for interface IModel, forgot register?")
	}
	return localModel
}

func RegisterModel(i IModel) {
	localModel = i
}
