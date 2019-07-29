package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

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

func (h *Handler) WebHookIdempotencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	if _, err := uuid.Parse(orderID); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	l, err := lawsuit.GetByOrderID(h.DB, orderID)
	if err != nil {
		Respond(w, r, http.StatusNotFound, err.Error())
		return
	}

	if l.Processed {
		Respond(w, r, http.StatusNotFound, errors.New("LawsuitNumber was processed"))
		return
	}

	Respond(w, r, http.StatusOK, l)
}
