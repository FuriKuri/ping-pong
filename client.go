package main

import (
	"fmt"
	"math/rand"
	"net/http"
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

func hitOrMiss(hitChange int) string {
	result := rand.Intn(100) + 1
	if result > hitChange {
		return "MISS"
	}
	return "HIT"
}
