package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Hello, world!")
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
