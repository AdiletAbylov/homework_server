package main

import (
	"hw_server/business"
	"hw_server/config"
	"hw_server/handlers"
	"hw_server/helpers"
	"hw_server/repo"
	"net/http"
)

func main() {
	config.ReadConfigs()

	repo.InitDB(config.DBConnectString())
	defer repo.Close()

	http.HandleFunc("/", handlers.BasicHandler)
	http.ListenAndServe(config.ServicePort(), nil)

	// start job
	helpers.DoJobEvery(config.JobTimerDuration(), business.WorkerJob)
}