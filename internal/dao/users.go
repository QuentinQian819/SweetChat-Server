package dao

import (
	"chatBox/internal/dao/daointernal"
	"github.com/gogf/gf/v2/database/gdb"
)

// UsersDao is the global DAO instance for users table.
var UsersDao = daointernal.NewUsersDao(nil)

// InitDAOs initializes all DAOs with database instance.
func InitDAOs(db gdb.DB) {
	UsersDao = daointernal.NewUsersDao(db)
	CouplesDao = daointernal.NewCouplesDao(db)
	MessagesDao = daointernal.NewMessagesDao(db)
	DiariesDao = daointernal.NewDiariesDao(db)
	DiaryMediaDao = daointernal.NewDiaryMediaDao(db)
}
