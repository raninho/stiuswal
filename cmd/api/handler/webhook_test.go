package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
	"github.com/raninho/stiuswal/internal/lawsuit"
	"github.com/raninho/stiuswal/test"
)

func TestWebHookFinishHandler(t *testing.T) {
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

	requestWebhook := lawsuit.ProcessFinishedInput{}
	requestWebhook.LawsuitNumber = response.LawsuitNumber
	requestWebhook.OrderID = response.OrderID

	payload = new(bytes.Buffer)
	_ = json.NewEncoder(payload).Encode(requestWebhook)

	r, err = http.NewRequest("POST", "/webhooks/finish", payload)
	if err != nil {
		t.Fatal(err.Error())
	}

	w = httptest.NewRecorder()

	apiRouter.Router(h).ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Fatal("http.StatusOK != w.Code", http.StatusOK, w.Code)
	}

	ls, err := lawsuit.GetByOrderID(h.DB, requestWebhook.OrderID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ls.Processed != true {
		t.Fatal(errors.New("Webhook is not Ok"))
	}
}

func TestWebHookIdempotencyHandler(t *testing.T) {
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

	r, err = http.NewRequest("GET", "/webhooks/idempotency/"+response.OrderID, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	w = httptest.NewRecorder()

	apiRouter.Router(h).ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Fatal("http.StatusOK != w.Code", http.StatusOK, w.Code)
	}

	ls := lawsuit.New(response.LawsuitNumber)

	r, err = http.NewRequest("GET", "/webhooks/idempotency/"+ls.OrderID, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	w = httptest.NewRecorder()

	apiRouter.Router(h).ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Fatal("http.StatusNotFound != w.Code", http.StatusNotFound, w.Code)
	}

}

