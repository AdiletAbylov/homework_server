package main

import (
	"fmt"
	"hw_server/handlers"
	"hw_server/model"
	"hw_server/repo"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GameHandler)
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusUnsupportedMediaType, recorder.Code, "Should return 415 because content-type is not set")

	recorder = httptest.NewRecorder()
	req.Header.Add("Content-type", "application/json")
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Should return 400 because Source-Type is not set")

	recorder = httptest.NewRecorder()
	req.Header.Add("Source-type", model.SourceGame)
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Should return 400 because body is empty")

	req, err = createPostRequest("{lala:10}", model.SourceGame)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code, "Should return 500 because json string is not valid")

	amount := 10.15
	transactionID := "sometransctionID12333"
	req, err = createPostRequest(createBodyString(model.WinState, amount, transactionID), model.SourceGame)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()
	repo.InitDB()
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code, "Should return 200")
	user := repo.DefaultUser()

	assert.Equal(t, amount, user.Balance, "User balance should be "+fmt.Sprintf("%f", amount))

	// Request with the same TransactionID shouldn't be processed and user's balance shouldn't change
	req, err = createPostRequest(createBodyString(model.WinState, amount, transactionID), model.SourceGame)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code, "Should return 500 because the same transaction ID was processed before")
	user = repo.DefaultUser()
	assert.Equal(t, amount, user.Balance, "User balance should be "+fmt.Sprintf("%f", amount))

	req, err = createPostRequest(createBodyString(model.WinState, amount, transactionID), model.SourcePayment)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotImplemented, recorder.Code, "Should return 501 because payment transactions are not supported")

	req, err = createPostRequest(createBodyString(model.WinState, amount, transactionID), model.SourceServer)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotImplemented, recorder.Code, "Should return 501 because server transactions are not supported")

}

func createPostRequest(body string, sourceType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Source-type", sourceType)
	return req, err
}

func createBodyString(state string, amount float64, transactionID string) string {
	return "{\"state\": \"" + state + "\", \"amount\": \"" + fmt.Sprintf("%f", amount) + "\", \"transactionId\": \"" + transactionID + "\"}"
}
