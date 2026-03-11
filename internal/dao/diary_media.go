package dao

import (
	"chatBox/internal/dao/daointernal"
)

// DiaryMediaDao is the global DAO instance for diary_media table.
var DiaryMediaDao = daointernal.NewDiaryMediaDao(nil)
