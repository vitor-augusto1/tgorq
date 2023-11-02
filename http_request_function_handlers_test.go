package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockReturnRequestStruct(url string, method httpMethod) RequestStruct {
  bodyString := `{"foo": "bar"}`
  headerString := `{"Content-type": "application/json; charset=UTF-8"}`
	byteBody := bytes.NewBuffer([]byte(bodyString))
	byteHeaders := []byte(headerString)
  newRequestStruct := RequestStruct {
    url: url,
    chosenMethod: method,
    byteRequestBody: byteBody,
    byteRequestHeader: byteHeaders,

  }
  return newRequestStruct
}

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
  newMockedRequestStruct := mockReturnRequestStruct(newServer.URL, GET)
	response, _ := handleRequest(newMockedRequestStruct)
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
	}
}

func TestHandlePostMethod(t *testing.T) {
	expectedResponse := `{"success": "id 101 created"}`
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
  newMockedRequestStruct := mockReturnRequestStruct(newServer.URL, POST)
	response, err := handleRequest(newMockedRequestStruct)
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
  newMockedRequestStruct := mockReturnRequestStruct(newServer.URL, PUT)
	response, err := handleRequest(newMockedRequestStruct)
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
  newMockedRequestStruct := mockReturnRequestStruct(newServer.URL, DELETE)
	response, err := handleRequest(newMockedRequestStruct)
	if err != nil {
		t.Errorf("Delete request failed: %s", err)
		return
	}
	if response.rawResponse != expectedResponse {
		t.Errorf("Expected response to be %s got %s", expectedResponse, response.rawResponse)
	}
}
