package entity

import (
	"time"
)

// Messages is the golang structure for table messages.
type Messages struct {
	Id         uint64      `json:"id" orm:"id,primary"`
	CoupleId   uint64      `json:"coupleId" orm:"couple_id"`
	SenderId   uint64      `json:"senderId" orm:"sender_id"`
	ReceiverId uint64      `json:"receiverId" orm:"receiver_id"`
	MsgType    int8        `json:"msgType" orm:"msg_type"`
	Content    string      `json:"content" orm:"content"`
	IsRead     int8        `json:"isRead" orm:"is_read"`
	CreatedAt  *time.Time  `json:"createdAt" orm:"created_at"`
}

// TableName is the table name for model Messages.
func (m *Messages) TableName() string {
	return "messages"
}
