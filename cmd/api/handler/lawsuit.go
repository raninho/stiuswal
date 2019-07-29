package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/raninho/stiuswal/internal/lawsuit"
)

type LawsuitInformationRequest struct {
	LawsuitNumber string `json:"lawsuit-number"`
}

func (l *LawsuitInformationRequest) Ok() error {
	if !lawsuit.IsLawsuitNumber(l.LawsuitNumber) {
		return errors.New("lawsuit-number is invalid")
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

func KeyGetLawsuitInformationHandler(orderID string) string {
	return fmt.Sprintf("get_lawsuit_orderid_%s", orderID)
}

func (h *Handler) GetLawsuitInformationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	if _, err := uuid.Parse(orderID); err != nil {
		Respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	cached, err := h.Cache.Get(KeyGetLawsuitInformationHandler(orderID)).Result()
	if err == nil{
		ls := new(lawsuit.Lawsuit)
		_ = json.Unmarshal([]byte(cached), ls)
		Respond(w, r, http.StatusOK, ls)
		return
	}

	ls, err := lawsuit.GetByOrderID(h.DB, orderID)
	if err != nil {
		Respond(w, r, http.StatusNotFound, err.Error())
		return
	}

	caching, err := json.Marshal(ls)
	if err != nil {
		log.Println("json.Marshal:", err.Error())
	}

	if err := h.Cache.Set(KeyGetLawsuitInformationHandler(orderID), caching, time.Minute*1).Err(); err != nil {
		log.Println("Cache.Set:", err.Error())
	}

	Respond(w, r, http.StatusOK, ls)
}
