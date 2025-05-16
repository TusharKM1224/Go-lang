package facade

import (
	"context"

	types "github.com/TusharKM1224/Types"
	"github.com/TusharKM1224/service"
)

type Facade struct {
	serv service.NotificationserviceInterface
}

// CreateRecordFacade implements NotificationFacadeInterface.
func (f *Facade) CreateRecordFacade(ctx context.Context, data *types.Requesttype) {
	transformeddata := types.TranformToDBschema(data)
	f.serv.CreateRecordService(ctx, transformeddata)
}

type NotificationFacadeInterface interface {
	CreateRecordFacade(ctx context.Context, data *types.Requesttype)
}

func NewNotificationFacade(s service.NotificationserviceInterface) NotificationFacadeInterface {
	return &Facade{
		serv: s,
	}
}
