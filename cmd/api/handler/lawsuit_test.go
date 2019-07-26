package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
)

func TestLawsuitInformationRequest_Ok(t *testing.T) {
	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = ""

	err := request.Ok()
	if err == nil {
		t.Error("LawsuitNumber is invalid")
	}
}

func TestLawsuitInformationHandlerWithInvalidLawSuitNumber(t *testing.T) {
	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = ""

	payload := new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(request)

	r, err := http.NewRequest("POST", "/lawsuits", payload)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: nil, Cache: nil}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Error("http.StatusBadRequest != w.Code", http.StatusBadRequest, w.Code)
	}
}

func TestLawsuitInformationHandlerWithValidLawSuitNumber(t *testing.T) {
	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = "0710802-55.2018.8.02.0001"

	payload := new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(request)

	r, err := http.NewRequest("POST", "/lawsuits", payload)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: nil, Cache: nil}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusAccepted != w.Code {
		t.Error("http.StatusBadRequest != w.Code", http.StatusBadRequest, w.Code)
	}
}
