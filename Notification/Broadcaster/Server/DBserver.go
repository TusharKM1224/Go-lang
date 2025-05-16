package server

import (
	"log"
	"net/http"

	routes "github.com/TusharKM1224/Routes"
	types "github.com/TusharKM1224/Types"
	"github.com/TusharKM1224/facade"
	"github.com/TusharKM1224/handler"
	"github.com/TusharKM1224/repo"
	"github.com/TusharKM1224/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true

	},
}

func Initiateserver(db *gorm.DB) handler.NotificationhandlerInterface {

	notifyrepo := repo.NewNotificationRepository(db)
	notifyserv := service.NewNotificationService(notifyrepo)
	notifacade := facade.NewNotificationFacade(notifyserv)
	notifyhandler := handler.NewNotificationHandler(notifacade)
	go handlemessage()
	routes.WebsocketRoutes(HandleConnections)
	log.Println("Server started on :8080 ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server error: ", err)
	}
	router := gin.Default()
	initiateGinServer(router, notifyhandler)

	return notifyhandler

}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Websocket upgrade error : %v \n", err)
		return
	}
	defer ws.Close()
	types.ClientsMu.Lock()
	types.Clients[ws] = true
	types.ClientsMu.Unlock()
	log.Printf("New Client connected")

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Websocket read error : ", err)
			types.ClientsMu.Lock()
			delete(types.Clients, ws)
			types.ClientsMu.Unlock()
			break
		}
	}
}

func handlemessage() {
	for {
		notif := <-types.Broadcast
		types.ClientsMu.Lock()
		for client := range types.Clients {
			err := client.WriteJSON(notif)
			if err != nil {
				log.Println("Websocket Write error : %v\n", err)
				client.Close()
				delete(types.Clients, client)
			}

		}
		types.ClientsMu.Unlock()
	}
}

func initiateGinServer(r *gin.Engine, h handler.NotificationhandlerInterface) {
	routes.GinRoutes(r, h)
	r.Run(":9090")
}
