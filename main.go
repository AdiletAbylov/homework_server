package main

import (
	"hw_server/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GameHandler)
	http.ListenAndServe(":5000", nil)
}
