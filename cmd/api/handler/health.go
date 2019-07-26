package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) HealthCheckHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"live"}`)
}
