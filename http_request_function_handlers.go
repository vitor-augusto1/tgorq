package main

import (
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

type Response struct {
	rawResponse string
	body        string
	headers     string
	statusCode  int
}

func handleGetMethod(url string) (*Response, error) {
	request, err = http.NewRequest(GET.String(), url, nil)
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

func handlePostMethod(url string, body io.Reader, headers []byte) (*Response, error) {
	request, err = http.NewRequest(POST.String(), url, body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &requestHeaders); err != nil {
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

func handlePutMethod(url string, body io.Reader, headers []byte) (*Response, error) {
	request, err = http.NewRequest(PUT.String(), url, body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &requestHeaders); err != nil {
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

func handleDeleteMethod(url string, headers []byte) (*Response, error) {
	request, err = http.NewRequest(DELETE.String(), url, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &requestHeaders); err != nil {
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
