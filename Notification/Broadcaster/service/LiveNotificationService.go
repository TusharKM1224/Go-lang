package service

import (
	"context"

	types "github.com/TusharKM1224/Types"
	"github.com/TusharKM1224/repo"
)

type Services struct {
	repo repo.NotificationRepositoryInterface
}

// CreateRecordService implements NotificationserviceInterface.
func (s *Services) CreateRecordService(ctx context.Context, data *types.Notification) {

	count, status, Ntypes, created_at, _ := s.repo.CreateRecordRepo(ctx, data)
	notifyNewData(&types.Notify{
		Count:            count,
		Status:           status,
		Notificationtype: Ntypes,
		CreatedAt:        created_at,
	})

}

func notifyNewData(notification *types.Notify) {
	types.Broadcast <- *notification

}

type NotificationserviceInterface interface {
	CreateRecordService(ctx context.Context, data *types.Notification)
}

func NewNotificationService(r repo.NotificationRepositoryInterface) NotificationserviceInterface {
	return &Services{
		repo: r,
	}
}
