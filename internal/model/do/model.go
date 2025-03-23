// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Model is the golang structure of table model for DAO operations like Where/Data.
type Model struct {
	g.Meta    `orm:"table:model, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 模型名称
	VideoPath interface{} // 视频地址
	AudioPath interface{} // 音频地址
	VoiceId   interface{} // 音频ID
	CreatedAt *gtime.Time // 创建时间
}
