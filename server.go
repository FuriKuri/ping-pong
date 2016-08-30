package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	defaultPingEndpoint = "http://ping:8080/"
	defaultPongEndpoint = "http://pong:8080/"
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
			fmt.Fprintln(w, "PING --> HIT")
			pongRound := doRequest(pongEndpoint())
			if pongRound == hit {
				fmt.Fprintln(w, "PONG --> HIT")
			} else {
				fmt.Fprintln(w, "PONG --> MISS")
				fmt.Fprintln(w, "The winner is PING")
				return
			}
		} else {
			fmt.Fprintln(w, "PING --> MISS")
			fmt.Fprintln(w, "The winner is PONG")
			return
		}

	}
	fmt.Fprintln(w, "DRAW")
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
