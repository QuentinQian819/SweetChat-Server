package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// PromisesDao is the data access object for table promises.
type PromisesDao struct {
	db   gdb.DB
	table string
}

// PromisesColumns holds the column names for table promises.
type promisesDaoColumns struct {
	Id          string
	CoupleId    string
	CreatorId   string
	Title       string
	MessageIds  string
	ColorTag    string
	IsCompleted string
	CompletedAt string
	CreatedAt   string
	UpdatedAt   string
}

var PromisesColumns = promisesDaoColumns{
	Id:          "id",
	CoupleId:    "couple_id",
	CreatorId:   "creator_id",
	Title:       "title",
	MessageIds:  "message_ids",
	ColorTag:    "color_tag",
	IsCompleted: "is_completed",
	CompletedAt: "completed_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPromisesDao creates a new PromisesDao instance.
func NewPromisesDao(db gdb.DB) *PromisesDao {
	return &PromisesDao{
		db:    db,
		table: "promises",
	}
}

// Table returns the table name.
func (d *PromisesDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *PromisesDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *PromisesDao) Columns() promisesDaoColumns {
	return PromisesColumns
}

// Ctx creates a Model instance with context.
func (d *PromisesDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
