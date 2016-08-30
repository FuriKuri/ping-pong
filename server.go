package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	pingEndpoint = "http://ping:8080/"
	pongEndpoint = "http://pong:8080/"
)

// ServerHandler return the HttpFunction for server mode.
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Start game")
	fmt.Fprintf(w, doRequest(pingEndpoint))
	fmt.Fprintf(w, doRequest(pongEndpoint))
}

func doRequest(endpoint string) string {
	resp, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	s := string(body[:])
	return s
}
