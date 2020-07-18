package main

import (
	"hw_server/business"
	"hw_server/handlers"
	"hw_server/helpers"
	"hw_server/repo"
	"net/http"
)

func main() {
	repo.InitDB()
	http.HandleFunc("/", handlers.GameHandler)
	http.ListenAndServe(":5000", nil)

	defer repo.Close()
	helpers.DoJobEvery(10*60, business.WorkerJob)
}
