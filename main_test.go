package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	expected := "pong"
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/v1/ping", pongV1)
	r.ServeHTTP(w, req)

	if w.Body.String() != "pong" {
		t.Errorf("expect %v but %s", expected, w.Body.String())
	}
}
