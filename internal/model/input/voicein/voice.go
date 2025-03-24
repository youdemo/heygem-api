package voicein

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"heygem-api/internal/model"
	"heygem-api/internal/model/entity"
)

type PageInp struct {
	model.PageReq
	Name string `json:"name" dc:"语音名称"`
}

type PageOut struct {
	model.PageRes
	List []*entity.Voice `json:"list" dc:"语音列表"`
}

type SaveInp struct {
	File      *ghttp.UploadFile `json:"file" type:"file" dc:"语音"`
	Name      string            `json:"name"             dc:"名称"`
	AudioText string            `json:"audioText"        dc:"语音文本"`
	Save      bool              `json:"save"             dc:"是否保存"`
}

type SaveOut struct {
	Id        int64  `json:"id"           dc:"id"`
	AudioPath string `json:"audioPath"    dc:"音频路径"`
}
