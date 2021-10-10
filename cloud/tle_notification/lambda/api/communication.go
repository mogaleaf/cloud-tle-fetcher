package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"tle_manager/tle_notification/lambda/datastore"
	"tle_manager/tle_notification/lambda/event"

	"github.com/gorilla/websocket"
)

type wsManager struct {
	satelliteTleObserver event.SatelliteTleObserver
	datastore            datastore.Datastore
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (wsManager *wsManager) reader(conn *websocket.Conn) {
	// read in a message
	messagetype, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	// print out that message for clarity
	fmt.Println(string(p))
	//for now HIBER-3,HIBER-4
	split := strings.Split(string(p), ",")
	err = wsManager.datastore.SaveSatellites(split)
	if err != nil {
		conn.WriteMessage(messagetype, []byte(err.Error()))
		return
	}
	register, com := wsManager.satelliteTleObserver.Register(split)
	for {
		data := <-com
		err := conn.WriteMessage(messagetype, []byte(data.Tle))
		if err != nil {
			wsManager.satelliteTleObserver.Unregister(register)
			return
		}
	}
}
func (wsManager *wsManager) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	wsManager.reader(ws)
}

func ServeClient(addr string, satelliteTleObserver event.SatelliteTleObserver, datastore datastore.Datastore) {
	manager := &wsManager{
		satelliteTleObserver: satelliteTleObserver,
		datastore:            datastore,
	}
	http.HandleFunc("/listen_satellite", manager.wsEndpoint)
	http.ListenAndServe(addr, nil)
}
