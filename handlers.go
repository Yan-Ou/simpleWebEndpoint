package main

import (
	"github.com/gorilla/mux"
)

func Router(projectName, buildTime, commit, release string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/metadata", metadata(projectName, buildTime, commit, release)).Methods("GET")
	r.Handle("/health", health())

	return r
}
