package main

import (
	"log"
	"net/http"
	"os"
)

var (
	ProjectName = "unset"
	BuildTime   = "unset"
	Commit      = "unset"
	Release     = "unset"
)

func main() {

	log.Printf("Starting the service ...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := Router(ProjectName, BuildTime, Commit, Release)
	log.Print("The service is up and running.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
