package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// DiaryMediaDao is the data access object for table diary_media.
type DiaryMediaDao struct {
	db   gdb.DB
	table string
}

// DiaryMediaColumns holds the column names for table diary_media.
type diaryMediaColumns struct {
	Id        string
	DiaryId   string
	MediaUrl  string
	MediaType string
	CreatedAt string
}

var DiaryMediaColumns = diaryMediaColumns{
	Id:        "id",
	DiaryId:   "diary_id",
	MediaUrl:  "media_url",
	MediaType: "media_type",
	CreatedAt: "created_at",
}

// NewDiaryMediaDao creates a new DiaryMediaDao instance.
func NewDiaryMediaDao(db gdb.DB) *DiaryMediaDao {
	return &DiaryMediaDao{
		db:    db,
		table: "diary_media",
	}
}

// Table returns the table name.
func (d *DiaryMediaDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *DiaryMediaDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *DiaryMediaDao) Columns() diaryMediaColumns {
	return DiaryMediaColumns
}

// Ctx creates a Model instance with context.
func (d *DiaryMediaDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
