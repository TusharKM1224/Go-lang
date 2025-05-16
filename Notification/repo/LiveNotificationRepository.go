package repo

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

type NotificationRepositoryInterface interface {
}

func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &Repository{
		db: db,
	}
}
