package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	req                *http.Request
	reqHeaders         map[string]string
	resp               *http.Response
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
	req, err = http.NewRequest(GET.String(), url, nil)
	if err != nil {
		return nil, err
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range resp.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}
	responseStatusCode = resp.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer resp.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handlePostMethod(url string, body io.Reader, headers []byte) (*Response, error) {
	req, err = http.NewRequest(POST.String(), url, body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &reqHeaders); err != nil {
		return nil, err
	}

	for key, value := range reqHeaders {
		req.Header.Set(key, value)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range resp.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = resp.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer resp.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handlePutMethod(url string, body io.Reader, headers []byte) (*Response, error) {
	req, err = http.NewRequest(PUT.String(), url, body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &reqHeaders); err != nil {
		return nil, err
	}

	for key, value := range reqHeaders {
		req.Header.Set(key, value)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range resp.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = resp.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)

	defer resp.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}

func handleDeleteMethod(url string, headers []byte) (*Response, error) {
	req, err = http.NewRequest(DELETE.String(), url, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(headers, &reqHeaders); err != nil {
		return nil, err
	}

	for key, value := range reqHeaders {
		req.Header.Set(key, value)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseHeaders string
	for k, v := range resp.Header {
		responseHeaders += fmt.Sprintf(
			"%q : %q\n", k, v,
		)
	}

	responseStatusCode = resp.StatusCode
	responseBodyString := fmt.Sprintf(
		"%d\n\n%s",
		responseStatusCode,
		string(responseBody),
	)
	defer resp.Body.Close()

	newResponse := &Response{
		rawResponse: string(responseBody),
		body:        responseBodyString,
		headers:     responseHeaders,
		statusCode:  responseStatusCode,
	}
	return newResponse, nil
}
