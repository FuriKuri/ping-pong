package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	defaultPingEndpoint = "http://ping:3000/"
	defaultPongEndpoint = "http://pong:3000/"
)

// ServerHandler return the HttpFunction for server mode.
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Start game")
	doGame(w)
	log.Println("End game")
}

func doGame(w http.ResponseWriter) {
	for round := 1; round < 100; round++ {
		pingRound := doRequest(pingEndpoint())
		if pingRound == hit {
			logEvent(w, "PING --> HIT")
			pongRound := doRequest(pongEndpoint())
			if pongRound == hit {
				logEvent(w, "PONG --> HIT")
			} else {
				logEvent(w, "PONG --> MISS")
				logEvent(w, "The winner is PING")
				return
			}
		} else {
			logEvent(w, "PING --> MISS")
			logEvent(w, "The winner is PONG")
			return
		}

	}
	logEvent(w, "DRAW")
}

func logEvent(w http.ResponseWriter, msg string) {
	log.Println(msg)
	fmt.Fprintln(w, msg)
}

func pingEndpoint() string {
	return getArgParameter("ping-endpoint", defaultPingEndpoint)
}

func pongEndpoint() string {
	return getArgParameter("pong-endpoint", defaultPongEndpoint)
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
