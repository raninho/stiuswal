package handler

import (
	"encoding/json"
	"net/http"

	"github.com/raninho/stiuswal/internal/lawsuit"
)

func (h *Handler) WebHookFinishHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	request := new(lawsuit.ProcessFinishedInput)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Ok(); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	l, err := lawsuit.ProcessFinished(h.DB, request)
	if err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	Respond(w, r, http.StatusOK, l)
}
