package main

import (
	"net/http"
)

const (
	defaultMode = "client"
)

func main() {
	http.HandleFunc("/", ClientHandler)
	http.ListenAndServe(":8080", nil)
}

func mode() string {
	return getArgParameter("mode", defaultMode)
}
