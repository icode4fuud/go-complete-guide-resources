package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateEvent_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	svc := &fakeService{} // implement minimal interface
	h := NewEventHandler(svc)
	h.RegisterRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/events", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
