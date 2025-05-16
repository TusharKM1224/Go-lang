package routes

import (
	"net/http"

	"github.com/TusharKM1224/handler"
	"github.com/gin-gonic/gin"
)

func WebsocketRoutes(handlerFunc http.HandlerFunc) {
	http.HandleFunc("/ws", handlerFunc)
}

func GinRoutes(r *gin.Engine, Handler handler.NotificationhandlerInterface) {
	r.POST("/notify")
}
