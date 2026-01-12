package notification

import "time"

type Type string

const (
	TypeEmail Type = "EMAIL"
	TypePush  Type = "PUSH"
)

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Type      string
	Title     string
	Message   string
	IsSent    bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
