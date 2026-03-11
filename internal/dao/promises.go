package dao

import (
	"chatBox/internal/dao/daointernal"
)

// PromisesDao is the global DAO instance for promises table.
var PromisesDao = daointernal.NewPromisesDao(nil)
