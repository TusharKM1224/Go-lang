package service

import "github.com/TusharKM1224/repo"

type Services struct {
	repo repo.NotificationRepositoryInterface
}

type NotificationserviceInterface interface {
}

func NewNotificationService(r repo.NotificationRepositoryInterface) NotificationserviceInterface {
	return &Services{
		repo: r,
	}
}
