// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoDao is the data access object for the table video.
type VideoDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns VideoColumns // columns contains all the column names of Table for convenient usage.
}

// VideoColumns defines and stores column names for the table video.
type VideoColumns struct {
	Id          string // 主键
	FilePath    string // 文件地址
	Status      string // 状态
	Message     string // 消息
	ModelId     string // 模特ID
	AudioPath   string // 音频地址
	Param       string // 参数
	Code        string // 编码
	CreatedAt   string // 创建时间
	Progress    string // 进度
	Name        string // 名称
	Duration    string // 视频时长
	TextContent string // 文本内容
	VoiceId     string // 语音ID
}

// videoColumns holds the columns for the table video.
var videoColumns = VideoColumns{
	Id:          "id",
	FilePath:    "file_path",
	Status:      "status",
	Message:     "message",
	ModelId:     "model_id",
	AudioPath:   "audio_path",
	Param:       "param",
	Code:        "code",
	CreatedAt:   "created_at",
	Progress:    "progress",
	Name:        "name",
	Duration:    "duration",
	TextContent: "text_content",
	VoiceId:     "voice_id",
}

// NewVideoDao creates and returns a new DAO object for table data access.
func NewVideoDao() *VideoDao {
	return &VideoDao{
		group:   "default",
		table:   "video",
		columns: videoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VideoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VideoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VideoDao) Columns() VideoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VideoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VideoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VideoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
