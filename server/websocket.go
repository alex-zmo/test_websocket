package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Printf("updating: ", t)

			dashboardInfo := GetData()

			jsonString, err := json.Marshal(dashboardInfo)
			if err != nil {
				fmt.Println(err)
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString));
			err != nil {
				fmt.Println(err)
			}
		}
	}


}