// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Model is the golang structure for table model.
type Model struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	Name      string      `json:"name"      orm:"name"       description:"模型名称"`
	VideoPath string      `json:"videoPath" orm:"video_path" description:"视频地址"`
	AudioPath string      `json:"audioPath" orm:"audio_path" description:"音频地址"`
	VoiceId   int64       `json:"voiceId"   orm:"voice_id"   description:"音频ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
