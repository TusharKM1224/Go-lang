package handler

import (
	"net/http"

	types "github.com/TusharKM1224/Types"
	"github.com/TusharKM1224/facade"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	fac facade.NotificationFacadeInterface
}

// CreateRecordhandler implements NotificationhandlerInterface.
func (h *Handler) CreateRecordhandler(c *gin.Context) {
	var data types.Requesttype
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Invalid request format",
			"Err_description": err.Error(),
		})
		return
	}
	h.fac.CreateRecordFacade(c.Request.Context(), &data)

}

type NotificationhandlerInterface interface {
	CreateRecordhandler(c *gin.Context)
}

func NewNotificationHandler(f facade.NotificationFacadeInterface) NotificationhandlerInterface {
	return &Handler{
		fac: f,
	}
}
