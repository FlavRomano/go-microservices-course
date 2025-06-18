package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Autenticate(t *testing.T) {
	// mocking log
	jsonToReturn := `{
	"error": false,
	"message": "test message"
	}`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email":    "a@gmail.com",
		"password": "thisIsAPassword",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/auth", bytes.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted but got %d", responseRecorder.Code)
	}
}
