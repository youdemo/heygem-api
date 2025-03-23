// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package voice

import (
	"context"

	"heygem-api/api/voice/v1"
)

type IVoiceV1 interface {
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	Find(ctx context.Context, req *v1.FindReq) (res *v1.FindRes, err error)
	Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error)
}
