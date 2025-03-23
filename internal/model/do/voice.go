// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Voice is the golang structure of table voice for DAO operations like Where/Data.
type Voice struct {
	g.Meta    `orm:"table:voice, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 名称
	AudioPath interface{} // 语音路径
	AudioText interface{} // 语音文本
	CreatedAt *gtime.Time // 创建时间
}
