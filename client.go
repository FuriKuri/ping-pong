package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultHitChance = "50"
)

// ClientHandler returns the HttpHandler for the client mode.
// Simply returns "HIT" or "MISS" on root path "/".
func ClientHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hitOrMiss(hitChange()))
}

func hitChange() int {
	i, _ := strconv.Atoi(getArgParameter("hit-chance", defaultHitChance))
	return i
}

func getArgParameter(name string, defaultValue string) string {
	argsWithoutProg := os.Args[1:]
	for index, element := range argsWithoutProg {
		if element == "--"+name {
			return argsWithoutProg[index+1]
		}
	}
	return defaultValue
}

func hitOrMiss(hitChange int) string {
	result := rand.Intn(100) + 1
	if result > hitChange {
		return "MISS"
	}
	return "HIT"
}
