package videoin

import (
	"github.com/gogf/gf/v2/os/gtime"
	"heygem-api/internal/model"
	"heygem-api/internal/model/entity"
)

type PageInp struct {
	model.PageReq
	Name string `json:"name" dc:"名称"`
}

type PageOut struct {
	model.PageRes
	List []*entity.Video `json:"list" dc:"列表"`
}

type SelectByStatusInp struct {
	Status string `json:"status" v:"required" dc:"状态"`
}

type SelectByStatusOut []*entity.Video

type UpdateFields struct {
	ModelId     int64  `json:"modelId"      description:"模特ID"`
	Name        string `json:"name"         description:"名称"`
	TextContent string `json:"textContent"  description:"文本内容"`
	VoiceId     int64  `json:"voiceId"      description:"语音ID"`
	AudioPath   string `json:"audioPath"    description:"音频地址"`
	Message     string `json:"message"      description:"消息"`
	Status      string `json:"status"       description:"状态"`
}

type InsertFields struct {
	UpdateFields
}

type SaveInp struct {
	Id          int64       `json:"id"           description:"主键"`
	FilePath    string      `json:"filePath"     description:"文件地址"`
	Status      string      `json:"status"       description:"状态"`
	Message     string      `json:"message"      description:"消息"`
	ModelId     int64       `json:"modelId"      description:"模特ID"`
	AudioPath   string      `json:"audioPath"    description:"音频地址"`
	Param       string      `json:"param"        description:"参数"`
	Code        string      `json:"code"         description:"编码"`
	CreatedAt   *gtime.Time `json:"createdAt"    description:"创建时间"`
	Progress    int         `json:"progress"     description:"进度"`
	Name        string      `json:"name"         description:"名称"`
	Duration    *gtime.Time `json:"duration"     description:"视频时长"`
	TextContent string      `json:"textContent"  description:"文本内容"`
	VoiceId     int64       `json:"voiceId"      description:"语音ID"`
}

type UpdateInp struct {
	Id       int64  `json:"id"           description:"主键"`
	FilePath string `json:"filePath"     description:"文件地址"`
	Status   string `json:"status"       description:"状态"`
	Message  string `json:"message"      description:"消息"`
}
