package ttsin

import "github.com/gogf/gf/v2/net/ghttp"

type TrainInp struct {
	File *ghttp.UploadFile `json:"file" type:"file" dc:"请上传视频"`
}

type TrainOut struct {
	Text string `json:"text" dc:"文本"`
}

type InvokeInp struct {
	VoiceId      int    `json:"voiceId" dc:"语音ID"`
	Text         string `json:"text" dc:"文本"`
	ResponseType string `json:"responseType" d:"stream" dc:"响应类型"`
}

type InvokeOut struct {
}
