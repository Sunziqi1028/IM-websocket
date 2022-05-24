package main

import (
	"IM/service"
	"net/http"
)

func main() {
	go service.MyServer.Start()

	http.HandleFunc("/chat", service.WsHandler)

	http.ListenAndServe(":2022", nil)

}
