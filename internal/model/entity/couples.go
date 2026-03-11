package entity

import (
	"time"
)

// Couples is the golang structure for table couples.
type Couples struct {
	Id         uint64      `json:"id" orm:"id,primary"`
	User1Id    uint64      `json:"user1Id" orm:"user1_id"`
	User2Id    uint64      `json:"user2Id" orm:"user2_id"`
	InviteCode string      `json:"inviteCode" orm:"invite_code"`
	Status     int8        `json:"status" orm:"status"`
	CreatedAt  *time.Time  `json:"createdAt" orm:"created_at"`
	UpdatedAt  *time.Time  `json:"updatedAt" orm:"updated_at"`
}

// TableName is the table name for model Couples.
func (m *Couples) TableName() string {
	return "couples"
}
