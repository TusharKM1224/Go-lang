package handler

import "github.com/TusharKM1224/facade"

type Handler struct {
	fac facade.NotificationFacadeInterface
}

type NotificationhandlerInterface interface {
}

func NewNotificationHandler(f facade.NotificationFacadeInterface) NotificationhandlerInterface {
	return &Handler{
		fac: f,
	}
}
