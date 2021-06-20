package server

import (
	"fmt"
	"log"
	"net/http"
	"test_websocket/model"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", &model.DashboardInfo{})
}


func upgrade(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
	}
	go Writer(ws)
}

func InitServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/upgrade", upgrade)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//log.Fatal(http.ListenAndServeTLS(":443", "certs/server-cert.pem", "certs/server-key.pem", nil))
}
