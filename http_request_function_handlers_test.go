package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetMethod(t *testing.T) {
  expectedResponse := `{"data": "dummy"}`

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

func TestHandlePostMethod(t *testing.T) {
  expectedResponse := `{"success": "id 101 created"}`
  bodyString := `
  {
    "id": 101,
    "title": "foo",
    "body": "bar",
    "userId": 1
  }
  `
  headerString := `{"Content-type": "application/json; charset=UTF-8"}`
  byteBody := bytes.NewBuffer([]byte(bodyString))
  byteHeaders := []byte(headerString)

  nServer := httptest.NewServer(
    http.HandlerFunc(
      func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
          t.Errorf("Expected a POST request, got %s", r.Method)
          return 
        }
        fmt.Fprintf(w, expectedResponse)
      },
    ),
  )
  defer nServer.Close()

  resp, err := handlePostMethod(nServer.URL, byteBody, byteHeaders)
  if err != nil {
    t.Errorf("POST request failed: %s", err)
    return 
  }
  if resp.rawResponse != expectedResponse {
    t.Errorf("Expected response to be %s got %s", expectedResponse, resp.rawResponse)
  }
}

func TestHandlePutMethod(t *testing.T) {
  expectedResponse := `{"success": "101 updated successfully"}`
  bodyString := `
  {
    "id": 101,
    "title": "bar",
    "body": "foo",
    "userId": 1
  }
  `
  headerString := `{"Content-type": "application/json; charset=UTF-8"}`
  byteBody := bytes.NewBuffer([]byte(bodyString))
  byteHeaders := []byte(headerString)

  nServer := httptest.NewServer(
    http.HandlerFunc(
      func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
          t.Errorf("Expected a PUT request, got %s", r.Method)
          return 
        }
        fmt.Fprintf(w, expectedResponse)
      },
    ),
  )
  defer nServer.Close()

  resp, err := handlePutMethod(nServer.URL, byteBody, byteHeaders)
  if err != nil {
    t.Errorf("PUT request failed: %s", err)
    return 
  }
  if resp.rawResponse != expectedResponse {
    t.Errorf("Expected response to be %s got %s", expectedResponse, resp.rawResponse)
  }
}


