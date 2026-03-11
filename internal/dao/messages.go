package dao

import (
	"chatBox/internal/dao/daointernal"
)

// MessagesDao is the global DAO instance for messages table.
var MessagesDao = daointernal.NewMessagesDao(nil)
