package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"tle_fetcher_solution/shared/session"

	"github.com/gorilla/websocket"
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//TODO enter here your deployed solution id
	appid := "417gytn0z6"

	u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s.execute-api.us-east-1.amazonaws.com", appid), Path: "/run"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//TODO enter here your satellites
	req := &session.JoinRequest{
		Satellites: []session.Satellite{
			{Name: "HIBER-3"},
			{Name: "HIBER-4"},
		},
	}
	marshal, err := json.Marshal(req)
	c.WriteMessage(websocket.TextMessage, marshal)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", string(message))
	}
}
