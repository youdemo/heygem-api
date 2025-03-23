// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Video is the golang structure of table video for DAO operations like Where/Data.
type Video struct {
	g.Meta      `orm:"table:video, do:true"`
	Id          interface{} // 主键
	FilePath    interface{} // 文件地址
	Status      interface{} // 状态
	Message     interface{} // 消息
	ModelId     interface{} // 模特ID
	AudioPath   interface{} // 音频地址
	Param       interface{} // 参数
	Code        interface{} // 编码
	CreatedAt   *gtime.Time // 创建时间
	Progress    interface{} // 进度
	Name        interface{} // 名称
	Duration    interface{} // 视频时长
	TextContent interface{} // 文本内容
	VoiceId     interface{} // 语音ID
}
