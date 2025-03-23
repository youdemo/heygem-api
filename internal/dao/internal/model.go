// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ModelDao is the data access object for the table model.
type ModelDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns ModelColumns // columns contains all the column names of Table for convenient usage.
}

// ModelColumns defines and stores column names for the table model.
type ModelColumns struct {
	Id        string // 主键
	Name      string // 模型名称
	VideoPath string // 视频地址
	AudioPath string // 音频地址
	VoiceId   string // 音频ID
	CreatedAt string // 创建时间
}

// modelColumns holds the columns for the table model.
var modelColumns = ModelColumns{
	Id:        "id",
	Name:      "name",
	VideoPath: "video_path",
	AudioPath: "audio_path",
	VoiceId:   "voice_id",
	CreatedAt: "created_at",
}

// NewModelDao creates and returns a new DAO object for table data access.
func NewModelDao() *ModelDao {
	return &ModelDao{
		group:   "default",
		table:   "model",
		columns: modelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ModelDao) Columns() ModelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ModelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
