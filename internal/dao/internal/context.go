// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContextDao is the data access object for the table context.
type ContextDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns ContextColumns // columns contains all the column names of Table for convenient usage.
}

// ContextColumns defines and stores column names for the table context.
type ContextColumns struct {
	Id  string // 主键
	Key string // 键
	Val string // 值
}

// contextColumns holds the columns for the table context.
var contextColumns = ContextColumns{
	Id:  "id",
	Key: "key",
	Val: "val",
}

// NewContextDao creates and returns a new DAO object for table data access.
func NewContextDao() *ContextDao {
	return &ContextDao{
		group:   "default",
		table:   "context",
		columns: contextColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContextDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContextDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContextDao) Columns() ContextColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContextDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContextDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ContextDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
