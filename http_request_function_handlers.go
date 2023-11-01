package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	request                *http.Request
	requestHeaders         map[string]string
	response               *http.Response
	responseBody       []byte
	responseStatusCode int
	err                error
)


type RequestStruct struct {
  url           string
  chosenMethod  httpMethod
  byteRequestBody      *bytes.Buffer
  byteRequestHeader    []byte
}

type Response struct {
	rawResponse string
	body        string
	headers     string
	statusCode  int
}

func handleGetMethod(requestS RequestStruct) (*Response, error) {
	request, err = http.NewRequest(GET.String(), requestS.url, nil)
	if err != nil {
		return nil, err
	}

	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range response.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}
	responseStatusCode = response.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer response.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handlePostMethod(requestS RequestStruct) (*Response, error) {
	request, err = http.NewRequest(POST.String(), requestS.url, requestS.byteRequestBody)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(requestS.byteRequestHeader, &requestHeaders); err != nil {
		return nil, err
	}

	for key, value := range requestHeaders {
		request.Header.Set(key, value)
	}

	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range response.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = response.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer response.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handlePutMethod(requestS RequestStruct) (*Response, error) {
	request, err = http.NewRequest(PUT.String(), requestS.url, requestS.byteRequestBody)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(requestS.byteRequestHeader, &requestHeaders); err != nil {
		return nil, err
	}

	for key, value := range requestHeaders {
		request.Header.Set(key, value)
	}

	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range response.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = response.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)

	defer response.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handleDeleteMethod(requestS RequestStruct) (*Response, error) {
	request, err = http.NewRequest(DELETE.String(), requestS.url, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(requestS.byteRequestHeader, &requestHeaders); err != nil {
		return nil, err
	}

	for key, value := range requestHeaders {
		request.Header.Set(key, value)
	}

	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range response.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = response.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer response.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func (m mainModel) returnRequestStruct() RequestStruct {
	url := m.url.textInput.Value()
	chosenHttpMethod := m.url.chosenMethod
	bodyString := m.request.body.Value()
	headerString := m.request.headers.Value()
	byteBody := bytes.NewBuffer([]byte(bodyString))
	byteHeaders := []byte(headerString)
  newRequestStruct := RequestStruct{
    url: url,
    chosenMethod: chosenHttpMethod,
    byteRequestBody: byteBody,
    byteRequestHeader: byteHeaders,

  }
  return newRequestStruct
}

func (m mainModel) setContentsOnRequestResponse() {
  newRequestStruct := m.returnRequestStruct()
  chosenHttpMethod := newRequestStruct.chosenMethod
  var requestHandler = map[httpMethod]func(RequestStruct) (*Response, error) {
    GET: handleGetMethod,
    POST: handlePostMethod,
    PUT: handlePutMethod,
    DELETE: handleDeleteMethod,
  }
  response, err := requestHandler[chosenHttpMethod](newRequestStruct)
  if err != nil {
    m.response.body.SetContent(err.Error())
    return
  }
  m.response.body.SetContent(response.body)
  m.response.headers.SetContent(response.headers)
  m.rawResponse = response
  if SaveToFileFlag {
    m.saveResponseOutputToFile(response.body, responseBodyOutputPath)
    m.saveResponseOutputToFile(response.headers, responseHeadersOutputPath)
  }
  if SaveStateFlag {
    m.storeCurrentState()
  }
}

func (m mainModel) executeRequest() {
  m.setContentsOnRequestResponse()
}
