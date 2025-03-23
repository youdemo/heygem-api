// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Voice is the golang structure for table voice.
type Voice struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	Name      string      `json:"name"      orm:"name"       description:"名称"`
	AudioPath string      `json:"audioPath" orm:"audio_path" description:"语音路径"`
	AudioText string      `json:"audioText" orm:"audio_text" description:"语音文本"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
