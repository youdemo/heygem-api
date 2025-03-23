// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tts

import (
	"context"

	"heygem-api/api/tts/v1"
)

type ITtsV1 interface {
	Train(ctx context.Context, req *v1.TrainReq) (res *v1.TrainRes, err error)
	Invoke(ctx context.Context, req *v1.InvokeReq) (res *v1.InvokeRes, err error)
}
