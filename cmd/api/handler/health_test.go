package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
)

func TestHealthCheckHandle(t *testing.T) {
	r, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: nil, Cache: nil}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Error("http.StatusAccepted != w.Code", http.StatusOK, w.Code)
	}
}
