package server

import (
	"github.com/TusharKM1224/facade"
	"github.com/TusharKM1224/handler"
	"github.com/TusharKM1224/repo"
	"github.com/TusharKM1224/service"
	"gorm.io/gorm"
)

func Initiateserver(db *gorm.DB) handler.NotificationhandlerInterface {

	notifyrepo := repo.NewNotificationRepository(db)
	notifyserv := service.NewNotificationService(notifyrepo)
	notifacade := facade.NewNotificationFacade(notifyserv)
	notifyhandler := handler.NewNotificationHandler(notifacade)

	return notifyhandler

}
