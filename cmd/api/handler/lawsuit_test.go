package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
	"github.com/raninho/stiuswal/internal/lawsuit"
	"github.com/raninho/stiuswal/test"
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
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	cache := test.NewRedisTest()
	defer cache.Close()

	queue := test.NewQueueTest()
	defer queue.Close()

	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = "0710802-55.2018.8.02.0001"

	payload := new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(request)

	r, err := http.NewRequest("POST", "/lawsuits", payload)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: db, Cache: cache, Queue: queue}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusAccepted != w.Code {
		t.Fatal("http.StatusBadRequest != w.Code", http.StatusBadRequest, w.Code)
	}

	response := new(apiHandler.LawsuitInformationResponse)
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(&response)
	if err != nil {
		t.Fatal(err.Error())
	}

	if request.LawsuitNumber != response.LawsuitNumber {
		t.Fatal(`request.LawsuitNumber != response.LawsuitNumber`, request.LawsuitNumber, response.LawsuitNumber)
	}

	if _, err := uuid.Parse(response.OrderID); err != nil {
		t.Fatal(`uuid.Parse(response.OrderID)`, response.OrderID)
	}
}

func TestGetLawsuitInformationHandlerWithValidOrderID(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	cache := test.NewRedisTest()
	defer cache.Close()

	queue := test.NewQueueTest()
	defer queue.Close()

	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = "0710802-55.2018.8.02.0001"

	payload := new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(request)

	r, err := http.NewRequest("POST", "/lawsuits", payload)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: db, Cache: cache, Queue: queue}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusAccepted != w.Code {
		t.Fatal("http.StatusAccepted != w.Code", http.StatusBadRequest, w.Code)
	}

	response := new(apiHandler.LawsuitInformationResponse)
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(&response)
	if err != nil {
		t.Fatal(err.Error())
	}

	if request.LawsuitNumber != response.LawsuitNumber {
		t.Fatal(`request.LawsuitNumber != response.LawsuitNumber`, request.LawsuitNumber, response.LawsuitNumber)
	}

	if _, err := uuid.Parse(response.OrderID); err != nil {
		t.Fatal(`uuid.Parse(response.OrderID)`, response.OrderID)
	}

	r, err = http.NewRequest("GET", "/lawsuits/"+response.OrderID, nil)
	if err != nil {
		t.Error(err.Error())
	}

	w = httptest.NewRecorder()
	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Fatal("http.StatusOK != w.Code", http.StatusBadRequest, w.Code)
	}
}

func TestGetLawsuitInformationHandlerWithInvalidOrderID(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	cache := test.NewRedisTest()
	defer cache.Close()

	queue := test.NewQueueTest()
	defer queue.Close()

	request := apiHandler.LawsuitInformationRequest{}
	request.LawsuitNumber = "0710802-55.2018.8.02.0001"

	payload := new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(request)

	r, err := http.NewRequest("POST", "/lawsuits", payload)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: db, Cache: cache, Queue: queue}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusAccepted != w.Code {
		t.Fatal("http.StatusAccepted != w.Code", http.StatusBadRequest, w.Code)
	}

	response := new(apiHandler.LawsuitInformationResponse)
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(&response)
	if err != nil {
		t.Fatal(err.Error())
	}

	if request.LawsuitNumber != response.LawsuitNumber {
		t.Fatal(`request.LawsuitNumber != response.LawsuitNumber`, request.LawsuitNumber, response.LawsuitNumber)
	}

	if _, err := uuid.Parse(response.OrderID); err != nil {
		t.Fatal(`uuid.Parse(response.OrderID)`, response.OrderID)
	}

	r, err = http.NewRequest("GET", "/lawsuits/"+response.OrderID[:len(response.OrderID)-1], nil)
	if err != nil {
		t.Error(err.Error())
	}

	w = httptest.NewRecorder()
	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Fatal("http.StatusBadRequest != w.Code", http.StatusBadRequest, w.Code)
	}
}

func TestGetLawsuitInformationHandlerWithOrderIDNotFound(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	cache := test.NewRedisTest()
	defer cache.Close()

	queue := test.NewQueueTest()
	defer queue.Close()

	orderID := "e9f7f07c-b805-45c6-90bc-21f273b27d63"

	r, err := http.NewRequest("GET", "/lawsuits/"+orderID, nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	h := &apiHandler.Handler{DB: db, Cache: cache, Queue: queue}

	apiRouter.Router(h).ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Fatal("http.StatusNotFound != w.Code", http.StatusBadRequest, w.Code)
	}
}