package types

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generator_v4();primaryKey"`
	NotificationType string    `gorm:"type:text"`
	Reciever_Name    string    `gorm :"type:text"`
	Reciever_Email   string    `gorm :"type:text"`
	Reciever_Phone   string    `gorm :"type:text"`
	Status           string    `gorm :"type:text"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
}

type Notify struct {
	Count            int
	Status           string
	Notificationtype string
	CreatedAt        time.Time
}
