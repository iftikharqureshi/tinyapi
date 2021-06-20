package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(`unexpected status code: got "%v" want "%v"`, status, http.StatusOK)
	}

	expected := `Hello from root`
	if rr.Body.String() != expected {
		t.Errorf(`unexpected response body: got "%v" want "%v"`,
			rr.Body.String(), expected)
	}
}
