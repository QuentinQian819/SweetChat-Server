package dao

import (
	"chatBox/internal/dao/daointernal"
)

// DiariesDao is the global DAO instance for diaries table.
var DiariesDao = daointernal.NewDiariesDao(nil)
