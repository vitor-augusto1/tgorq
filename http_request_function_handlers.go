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

func (m mainModel) executeRequest() {
	url := m.url.textInput.Value()
	chosenHttpMethod := m.url.chosenMethod
	bodyString := m.request.body.Value()
	headerString := m.request.headers.Value()
	byteBody := bytes.NewBuffer([]byte(bodyString))
	byteHeaders := []byte(headerString)
	if chosenHttpMethod == GET {
		response, err := handleGetMethod(url)
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
	} else if chosenHttpMethod == POST {
		response, err := handlePostMethod(url, byteBody, byteHeaders)
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
	} else if chosenHttpMethod == PUT {
		response, err := handlePutMethod(url, byteBody, byteHeaders)
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
	} else if chosenHttpMethod == DELETE {
		response, err := handleDeleteMethod(url, byteHeaders)
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
	}
	if SaveStateFlag {
		m.storeCurrentState()
	}
}
