package tts

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"heygem-api/internal/boot"
)

func Train(ctx context.Context, path string) (text string, err error) {
	// 调用api解析语音文本
	resp, err := g.Client().Post(ctx, boot.TTSApi+"/train", "audio=@file:"+path)
	if err != nil {
		return
	}
	text = gjson.New(resp.ReadAllString()).Get("transcribed_text").String()
	return
}

func Invoke(ctx context.Context, path, text, refText, responseType string) (audio []byte, err error) {
	// 调用api解析语音文本
	resp, err := g.Client().Post(ctx, boot.TTSApi+"/tts", fmt.Sprintf("audio=@file:%s&text=%s&ref_text=%s&response_type=%s", path, text, refText, responseType))
	if err != nil {
		return
	}
	audio = resp.ReadAll()
	return
}
