package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	handler := http.HandlerFunc(Hello)
	r := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	handler(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d, Got: %d", http.StatusOK, resp.StatusCode)
	}
	t.Logf("Status code: %d", resp.StatusCode)
	t.Logf("Response: %v", resp.Body)
}
