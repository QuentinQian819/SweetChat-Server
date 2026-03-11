package dao

import (
	"chatBox/internal/dao/daointernal"
)

// CouplesDao is the global DAO instance for couples table.
var CouplesDao = daointernal.NewCouplesDao(nil)
