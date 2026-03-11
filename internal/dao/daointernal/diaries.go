package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// DiariesDao is the data access object for table diaries.
type DiariesDao struct {
	db   gdb.DB
	table string
}

// DiariesColumns holds the column names for table diaries.
type diariesDaoColumns struct {
	Id        string
	CoupleId  string
	AuthorId  string
	Title     string
	Content   string
	IsShared  string
	Mood      string
	Weather   string
	CreatedAt string
	UpdatedAt string
}

var DiariesColumns = diariesDaoColumns{
	Id:        "id",
	CoupleId:  "couple_id",
	AuthorId:  "author_id",
	Title:     "title",
	Content:   "content",
	IsShared:  "is_shared",
	Mood:      "mood",
	Weather:   "weather",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDiariesDao creates a new DiariesDao instance.
func NewDiariesDao(db gdb.DB) *DiariesDao {
	return &DiariesDao{
		db:    db,
		table: "diaries",
	}
}

// Table returns the table name.
func (d *DiariesDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *DiariesDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *DiariesDao) Columns() diariesDaoColumns {
	return DiariesColumns
}

// Ctx creates a Model instance with context.
func (d *DiariesDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
