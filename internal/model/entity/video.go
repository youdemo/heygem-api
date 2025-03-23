// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Video is the golang structure for table video.
type Video struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	FilePath    string      `json:"filePath"    orm:"file_path"    description:"文件地址"`
	Status      string      `json:"status"      orm:"status"       description:"状态"`
	Message     string      `json:"message"     orm:"message"      description:"消息"`
	ModelId     int64       `json:"modelId"     orm:"model_id"     description:"模特ID"`
	AudioPath   string      `json:"audioPath"   orm:"audio_path"   description:"音频地址"`
	Param       string      `json:"param"       orm:"param"        description:"参数"`
	Code        string      `json:"code"        orm:"code"         description:"编码"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Progress    string      `json:"progress"    orm:"progress"     description:"进度"`
	Name        string      `json:"name"        orm:"name"         description:"名称"`
	Duration    int64       `json:"duration"    orm:"duration"     description:"视频时长"`
	TextContent string      `json:"textContent" orm:"text_content" description:"文本内容"`
	VoiceId     int64       `json:"voiceId"     orm:"voice_id"     description:"语音ID"`
}
