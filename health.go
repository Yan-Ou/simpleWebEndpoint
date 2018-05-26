package main

import (
	"github.com/etherlabsio/healthcheck"
	"github.com/etherlabsio/healthcheck/checkers"
	"net/http"
	"time"
)

func health() http.Handler {
	return healthcheck.Handler(

		// WithTimeout allows you to set a max overall timeout.
		healthcheck.WithTimeout(5*time.Second),

		// Checkers fail the status in case of any error.
		healthcheck.WithChecker("heartbeat", checkers.Heartbeat("/")),
		healthcheck.WithObserver(
			"diskspace", checkers.DiskSpace("/var/log", 1),
		),
	)
}
