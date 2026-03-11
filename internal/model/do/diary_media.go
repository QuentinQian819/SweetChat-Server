package do

import (
	"time"
)

// DiaryMedia is the golang structure for table diary_media.
type DiaryMedia struct {
	Id        uint64      `json:"id" orm:"id,primary"`
	DiaryId   uint64      `json:"diaryId" orm:"diary_id"`
	MediaUrl  string      `json:"mediaUrl" orm:"media_url"`
	MediaType string      `json:"mediaType" orm:"media_type"`
	CreatedAt *time.Time  `json:"createdAt" orm:"created_at"`
}
