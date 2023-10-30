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

	newServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, expectedResponse)
			},
		),
	)
	defer newServer.Close()

	response, _ := handleGetMethod(newServer.URL)
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
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

	newServer := httptest.NewServer(
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
	defer newServer.Close()

	response, err := handlePostMethod(newServer.URL, byteBody, byteHeaders)
	if err != nil {
		t.Errorf("POST request failed: %s", err)
		return
	}
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
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

	newServer := httptest.NewServer(
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
	defer newServer.Close()

	response, err := handlePutMethod(newServer.URL, byteBody, byteHeaders)
	if err != nil {
		t.Errorf("PUT request failed: %s", err)
		return
	}
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
	}
}

func TestHandleDeleteMethod(t *testing.T) {
	expectedResponse := `{"success": "101 deleted successfully"}`
	headerString := `{"Content-type": "application/json; charset=UTF-8"}`
	byteHeaders := []byte(headerString)

	newServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodDelete {
					t.Errorf("Expected a DELETE request, got %s", r.Method)
					return
				}
				fmt.Fprintf(w, expectedResponse)
			},
		),
	)
	defer newServer.Close()

	response, err := handleDeleteMethod(newServer.URL, byteHeaders)
	if err != nil {
		t.Errorf("Delete request failed: %s", err)
		return
	}
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
	}
}
