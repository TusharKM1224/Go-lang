package repo

import (
	"context"
	"time"

	types "github.com/TusharKM1224/Types"
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// CreateRecordRepo implements NotificationRepositoryInterface.
func (r *Repository) CreateRecordRepo(ctx context.Context, data *types.Notification) (int, string, string, time.Time, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		log.Fatal("Error while initiating transaction")

	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var count int64
	if err := tx.Create(data).Count(&count).Error; err != nil {
		log.Fatal("Error while inserting Data in DB")
	}
	if tx.Commit().Error != nil {
		log.Fatal("Error occured while Commiting")
	}

	return int(count), data.Status, data.NotificationType, data.CreatedAt, nil
}

type NotificationRepositoryInterface interface {
	CreateRecordRepo(ctx context.Context, data *types.Notification) (int, string, string, time.Time, error)
}

func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &Repository{
		db: db,
	}
}
