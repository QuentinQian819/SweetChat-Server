package do

import (
	"time"
)

// Users is the golang structure for table users.
type Users struct {
	Id           uint64      `json:"id" orm:"id,primary"`
	Phone        string      `json:"phone" orm:"phone"`
	Nickname     string      `json:"nickname" orm:"nickname"`
	Avatar       string      `json:"avatar" orm:"avatar"`
	PasswordHash string      `json:"passwordHash" orm:"password_hash"`
	CreatedAt    *time.Time  `json:"createdAt" orm:"created_at"`
	UpdatedAt    *time.Time  `json:"updatedAt" orm:"updated_at"`
}
