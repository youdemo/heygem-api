package f2f

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"heygem-api/internal/boot"
	"heygem-api/internal/consts"
	"heygem-api/internal/model"
)

func Submit(ctx context.Context, code string, audioPath string, videoPath string) (param map[string]interface{}, res *gjson.Json, err error) {
	// 调用api解析语音文本
	param = g.Map{
		"code":             code,
		"audio_url":        gstr.Join([]string{consts.CodeData, audioPath}, "/"),
		"video_url":        gstr.Join([]string{consts.CodeData, videoPath}, "/"),
		"chaofen":          0,
		"watermark_switch": 0,
		"pn":               1,
	}
	resp, err := g.Client().ContentJson().Post(ctx, boot.F2FApi+"/submit", param)
	if err != nil {
		return
	}
	res = gjson.New(resp.ReadAllString())
	return
}

func Query(ctx context.Context, code string) (res *model.FaceResponse, err error) {
	// 调用api解析语音文本
	resp, err := g.Client().Get(ctx, boot.F2FApi+"/query?code="+code)
	if err != nil {
		return
	}
	err = gjson.New(resp.ReadAll()).Scan(&res)
	return
}
