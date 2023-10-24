package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetMethod(t *testing.T) {
  expectedResponse := "{'data': 'dummy'}"

  nServer := httptest.NewServer(
    http.HandlerFunc(
      func(w http.ResponseWriter, r *http.Request) { 
        fmt.Fprintf(w, expectedResponse) 
      },
    ),
  )
  defer nServer.Close()

  resp, _ := handleGetMethod(nServer.URL)
  if resp.rawResponse != expectedResponse {
    t.Errorf("Expected response to be %s got %s", expectedResponse, resp.rawResponse)
  }
}

