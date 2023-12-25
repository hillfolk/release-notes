package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	reqId := "36dea16f-d5b1-4075-8767-1d055758ab0f"
	r = addRequestIdToContext(r, reqId)

	processRequest(w, r)
}

func processRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	logRequest(r)
	return fmt.Fprintf(w, "Hello, world")
}

func logRequest(r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(requestContextKey{})
	if m, ok := v.(requestContextValue); ok {
		log.Printf("Processing request: %s", m.requestId)
	}
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
