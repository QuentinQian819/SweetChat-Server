package daointernal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

// MessagesDao is the data access object for table messages.
type MessagesDao struct {
	db   gdb.DB
	table string
}

// MessagesColumns holds the column names for table messages.
type messagesDaoColumns struct {
	Id         string
	CoupleId   string
	SenderId   string
	ReceiverId string
	MsgType    string
	Content    string
	IsRead     string
	CreatedAt  string
}

var MessagesColumns = messagesDaoColumns{
	Id:         "id",
	CoupleId:   "couple_id",
	SenderId:   "sender_id",
	ReceiverId: "receiver_id",
	MsgType:    "msg_type",
	Content:    "content",
	IsRead:     "is_read",
	CreatedAt:  "created_at",
}

// NewMessagesDao creates a new MessagesDao instance.
func NewMessagesDao(db gdb.DB) *MessagesDao {
	return &MessagesDao{
		db:    db,
		table: "messages",
	}
}

// Table returns the table name.
func (d *MessagesDao) Table() string {
	return d.table
}

// DB returns the database instance.
func (d *MessagesDao) DB() gdb.DB {
	return d.db
}

// Columns returns the column names.
func (d *MessagesDao) Columns() messagesDaoColumns {
	return MessagesColumns
}

// Ctx creates a Model instance with context.
func (d *MessagesDao) Ctx(ctx context.Context) *gdb.Model {
	return d.db.Model(d.table).Ctx(ctx)
}
