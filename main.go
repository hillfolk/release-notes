package main

import (
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthz", checkHandler)
}

func main() {
	listenAddr := os.Getenv("LISTEN_PORT")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	mux := http.NewServeMux()
	setupHandlers(mux)
	http.ListenAndServe(listenAddr, mux)
}
