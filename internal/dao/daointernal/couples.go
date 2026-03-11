package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// CouplesDao is the data access object for table couples.
type CouplesDao struct {
	db   gdb.DB
	table string
}

// CouplesColumns holds the column names for table couples.
type couplesDaoColumns struct {
	Id         string
	User1Id    string
	User2Id    string
	InviteCode string
	Status     string
	CreatedAt  string
	UpdatedAt  string
}

var CouplesColumns = couplesDaoColumns{
	Id:         "id",
	User1Id:    "user1_id",
	User2Id:    "user2_id",
	InviteCode: "invite_code",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewCouplesDao creates a new CouplesDao instance.
func NewCouplesDao(db gdb.DB) *CouplesDao {
	return &CouplesDao{
		db:    db,
		table: "couples",
	}
}

// Table returns the table name.
func (d *CouplesDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *CouplesDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *CouplesDao) Columns() couplesDaoColumns {
	return CouplesColumns
}

// Ctx creates a Model instance with context.
func (d *CouplesDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
