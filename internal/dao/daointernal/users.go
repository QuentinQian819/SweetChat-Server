package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	db   gdb.DB
	table string
}

// UsersColumns holds the column names for table users.
type usersDaoColumns struct {
	Id           string
	Phone        string
	Nickname     string
	Avatar       string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}

var UsersColumns = usersDaoColumns{
	Id:           "id",
	Phone:        "phone",
	Nickname:     "nickname",
	Avatar:       "avatar",
	PasswordHash: "password_hash",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewUsersDao creates a new UsersDao instance.
func NewUsersDao(db gdb.DB) *UsersDao {
	return &UsersDao{
		db:    db,
		table: "users",
	}
}

// Table returns the table name.
func (d *UsersDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *UsersDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *UsersDao) Columns() usersDaoColumns {
	return UsersColumns
}

// Ctx creates a Model instance with context.
func (d *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
