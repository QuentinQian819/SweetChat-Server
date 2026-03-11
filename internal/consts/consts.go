package consts

// Message types
const (
	MessageTypeText    = 1
	MessageTypeImage   = 2
	MessageTypeAudio   = 3
	MessageTypeCheckIn = 4 // 报备消息
)

// User/Couple status
const (
	StatusActive   = 1
	StatusInactive = 0
)

// Diary sharing
const (
	DiaryShared   = 1
	DiaryPrivate  = 0
)

// Media types
const (
	MediaTypeImage = "image"
	MediaTypeAudio = "audio"
)

// Cache keys
const (
	CacheKeyOnlineUsers   = "chatbox:online:users"
	CacheKeyUserInfo      = "chatbox:user:info:"
	CacheKeyCoupleInfo    = "chatbox:couple:info:"
	CacheKeyUnreadCount   = "chatbox:unread:"
)

// Redis key patterns
const (
	OnlineUserPattern = "chatbox:online:user:%d"
)
