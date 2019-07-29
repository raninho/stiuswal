package router

import (
	"github.com/gorilla/mux"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
)

func Router(h *apiHandler.Handler) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/health", h.HealthCheckHandle).Methods("GET")

	router.HandleFunc("/lawsuits", h.LawsuitInformationHandler).Methods("POST")
	router.HandleFunc("/lawsuits/{orderID}", h.GetLawsuitInformationHandler).Methods("GET")

	router.HandleFunc("/webhooks/finish", h.WebHookFinishHandler).Methods("POST")

	return
}
