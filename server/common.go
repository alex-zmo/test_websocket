package server

import (
	"fmt"
	"github.com/gravitational/teleport/teleport_admin_service/model"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", &model.IndexPage{})
}


func info(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
	}
	go Writer(ws)
}

func InitServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/info", info)
	log.Fatal(http.ListenAndServeTLS(":443", "certs/server-cert.pem", "certs/server-key.pem", nil))
}
