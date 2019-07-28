package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/raninho/stiuswal/internal/lawsuit"
)

type LawsuitInformationRequest struct {
	LawsuitNumber string `json:"lawsuit-number"`
}

func (l *LawsuitInformationRequest) Ok() error {
	if l.LawsuitNumber == "" {
		return errors.New("lawsuit-number is blank")
	}

	return nil
}

type LawsuitInformationResponse struct {
	OrderID       string `json:"order-id"`
	LawsuitNumber string `json:"lawsuit-number"`
}

func (h *Handler) LawsuitInformationHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	request := new(LawsuitInformationRequest)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Ok(); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ls, err := lawsuit.Process(h.DB, h.Queue, request.LawsuitNumber)
	if err != nil {
		Respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	response := new(LawsuitInformationResponse)
	response.OrderID = ls.OrderID
	response.LawsuitNumber = ls.LawsuitNumber

	Respond(w, r, http.StatusAccepted, response)
}
