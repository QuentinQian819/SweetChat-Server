package do

import (
	"time"
)

// Diaries is the golang structure for table diaries.
type Diaries struct {
	Id        uint64      `json:"id" orm:"id,primary"`
	CoupleId  uint64      `json:"coupleId" orm:"couple_id"`
	AuthorId  uint64      `json:"authorId" orm:"author_id"`
	Title     string      `json:"title" orm:"title"`
	Content   string      `json:"content" orm:"content"`
	IsShared  int8        `json:"isShared" orm:"is_shared"`
	Mood      string      `json:"mood" orm:"mood"`
	Weather   string      `json:"weather" orm:"weather"`
	CreatedAt *time.Time  `json:"createdAt" orm:"created_at"`
	UpdatedAt *time.Time  `json:"updatedAt" orm:"updated_at"`
}
