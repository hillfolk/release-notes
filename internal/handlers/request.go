package handlers

import (
	"context"
	"net/http"
)

type requestContextKey struct {
}

type requestContextValue struct {
	requestId string
}

func addRequestIdToContext(r *http.Request, requestId string) *http.Request {
	c := requestContextValue{
		requestId: requestId,
	}
	currentCtx := r.Context()
	newCtx := context.WithValue(currentCtx, requestContextKey{}, c)
	return r.WithContext(newCtx)
}
