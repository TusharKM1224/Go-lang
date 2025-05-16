package facade

import "github.com/TusharKM1224/service"

type Facade struct {
	serv service.NotificationserviceInterface
}

type NotificationFacadeInterface interface {
}

func NewNotificationFacade(s service.NotificationserviceInterface) NotificationFacadeInterface {
	return &Facade{
		serv: s,
	}
}
