package main

import (
	"fmt"
	"net/http"
)

// ServerHandler return the HttpFunction for server mode.
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hitOrMiss(hitChange()")
}
