package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultHitChance = 50
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hitOrMiss(hitChange()))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func hitChange() int {
	argsWithoutProg := os.Args[1:]
	hitChanceIndex := -1
	for index, element := range argsWithoutProg {
		if element == "--hit-chance" {
			hitChanceIndex = index + 1
		}
	}
	if hitChanceIndex != -1 {
		i, _ := strconv.Atoi(argsWithoutProg[hitChanceIndex])
		return i
	}
	return defaultHitChance
}

func hitOrMiss(hitChange int) string {
	result := rand.Intn(100) + 1
	if result > hitChange {
		return "MISS"
	}
	return "HIT"
}
