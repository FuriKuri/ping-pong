package main

import (
	"net/http"
)

const (
	serverMode  = "server"
	clientMode  = "client"
	defaultMode = clientMode
)

func main() {
	if mode() == serverMode {
		http.HandleFunc("/", ServerHandler)
	} else {
		http.HandleFunc("/", ClientHandler)
	}
	http.ListenAndServe(":8080", nil)
}

func mode() string {
	return getArgParameter("mode", defaultMode)
}
