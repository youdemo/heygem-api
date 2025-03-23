// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VoiceDao is the data access object for the table voice.
type VoiceDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns VoiceColumns // columns contains all the column names of Table for convenient usage.
}

// VoiceColumns defines and stores column names for the table voice.
type VoiceColumns struct {
	Id        string // 主键
	Name      string // 名称
	AudioPath string // 语音路径
	AudioText string // 语音文本
	CreatedAt string // 创建时间
}

// voiceColumns holds the columns for the table voice.
var voiceColumns = VoiceColumns{
	Id:        "id",
	Name:      "name",
	AudioPath: "audio_path",
	AudioText: "audio_text",
	CreatedAt: "created_at",
}

// NewVoiceDao creates and returns a new DAO object for table data access.
func NewVoiceDao() *VoiceDao {
	return &VoiceDao{
		group:   "default",
		table:   "voice",
		columns: voiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VoiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VoiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VoiceDao) Columns() VoiceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VoiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VoiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VoiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
