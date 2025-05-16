package types

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Notification struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" binding:"required"`
	NotificationType string    `gorm:"type:text" json:"notification_type" binding:"required"`
	ReceiverName     string    `gorm:"type:text" json:"receiver_name" binding:"required"`
	ReceiverEmail    string    `gorm:"type:text" json:"receiver_email" binding:"required"`
	ReceiverPhone    string    `gorm:"type:text" json:"receiver_phone" binding:"required"`
	Status           string    `gorm:"type:text" json:"status" binding:"required"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at" binding:"required"`
}

type Notify struct {
	Count            int
	Status           string
	Notificationtype string
	CreatedAt        time.Time
}
type Requesttype struct {
	NType string `binding:"required"`
	Name  string `binding :"required"`
	Email string `binding:"required"`
	Phone string `binding:"required"`
}

var (
	Clients   = make(map[*websocket.Conn]bool)
	ClientsMu sync.Mutex
	Broadcast = make(chan Notify)
)

func TranformToDBschema(req *Requesttype) *Notification {
	return &Notification{
		ID:               uuid.New(),
		NotificationType: req.NType,
		ReceiverName:     req.Name,
		ReceiverEmail:    req.Email,
		ReceiverPhone:    req.Phone,
		Status:           "Unread",
		CreatedAt:        time.Now(),
	}

}
