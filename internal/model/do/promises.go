package do

import (
	"time"
)

// Promises is the golang structure for table promises.
type Promises struct {
	Id          uint64     `json:"id" orm:"id,primary"`
	CoupleId    uint64     `json:"coupleId" orm:"couple_id"`
	CreatorId   uint64     `json:"creatorId" orm:"creator_id"`
	Title       string     `json:"title" orm:"title"`
	MessageIds  string     `json:"messageIds" orm:"message_ids"`
	ColorTag    int        `json:"colorTag" orm:"color_tag"`
	IsCompleted bool       `json:"isCompleted" orm:"is_completed"`
	CompletedAt *time.Time `json:"completedAt,omitempty" orm:"completed_at"`
	CreatedAt   time.Time  `json:"createdAt" orm:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" orm:"updated_at"`
}
