package http

import (
	"net/http"

	"release-notes/internal/handlers"
)

func Run(port string) error {
	if len(port) == 0 {
		port = ":8080"
	}

	mux := http.NewServeMux()
	setupHandler(mux)
	return http.ListenAndServe(port, mux)
}

func setupHandler(mux *http.ServeMux) {

	mux.HandleFunc("/healthz", handlers.HealthCheck)
	mux.HandleFunc("/api", handlers.Hello)

}
