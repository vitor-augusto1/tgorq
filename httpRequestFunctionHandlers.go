package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
  req *http.Request
  reqHeaders map[string]string
  resp *http.Response
  responseBody []byte
  responseStatusCode int
  err error
)

func (m mainModel) handleGetMethod(url string) {
  req, err = http.NewRequest(GET.String(), url, nil)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  responseBody, err = io.ReadAll(resp.Body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  var stringToBeStoreInTheResponseHeaderTextArea string
  for k, v := range resp.Header {
    stringToBeStoreInTheResponseHeaderTextArea += fmt.Sprintf(
      "%q : %q\n", k, v,
    )
  }

  responseStatusCode = resp.StatusCode
  responseBodyString := fmt.Sprintf(
    "%d\n\n%s",
    responseStatusCode,
    string(responseBody),
  )
  m.response.body.SetContent(responseBodyString)
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer resp.Body.Close()
}


func (m mainModel) handlePostMethod(url string, body io.Reader, headers []byte) {
  req, err = http.NewRequest(POST.String(), url, body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  if err = json.Unmarshal(headers, &reqHeaders); err != nil {
    m.response.body.SetContent(err.Error())
  }

  for key, value := range reqHeaders {
    req.Header.Set(key, value)
  }

  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  responseBody, err = io.ReadAll(resp.Body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  var stringToBeStoreInTheResponseHeaderTextArea string
  for k, v := range resp.Header {
    stringToBeStoreInTheResponseHeaderTextArea += fmt.Sprintf(
      "%q : %q\n", k, v,
    )
  }

  responseStatusCode = resp.StatusCode
  responseBodyString := fmt.Sprintf(
    "%d\n\n%s",
    responseStatusCode,
    string(responseBody),
  )
  m.response.body.SetContent(responseBodyString)
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer resp.Body.Close()
}

func  (m mainModel) handlePutMethod(url string, body io.Reader, headers []byte) {
  var req *http.Request
  var reqHeaders map[string]string
  var resp *http.Response
  var responseBody []byte
  var err error

  req, err = http.NewRequest(PUT.String(), url, body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  if err = json.Unmarshal(headers, &reqHeaders); err != nil {
    m.response.body.SetContent(err.Error())
  }

  for key, value := range reqHeaders {
    req.Header.Set(key, value)
  }

  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  responseBody, err = io.ReadAll(resp.Body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  var stringToBeStoreInTheResponseHeaderTextArea string
  for k, v := range resp.Header {
    stringToBeStoreInTheResponseHeaderTextArea += fmt.Sprintf(
      "%q : %q\n", k, v,
    )
  }

  responseStatusCode = resp.StatusCode
  responseBodyString := fmt.Sprintf(
    "%d\n\n%s",
    responseStatusCode,
    string(responseBody),
  )
  m.response.body.SetContent(responseBodyString)
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer resp.Body.Close()
}

func (m mainModel) handleDeleteMethod(url string, headers []byte) {
  req, err = http.NewRequest(DELETE.String(), url, nil)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  if err = json.Unmarshal(headers, &reqHeaders); err != nil {
    m.response.body.SetContent(err.Error())
  }

  for key, value := range reqHeaders {
    req.Header.Set(key, value)
  }

  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  responseBody, err = io.ReadAll(resp.Body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  var stringToBeStoreInTheResponseHeaderTextArea string
  for k, v := range resp.Header {
    stringToBeStoreInTheResponseHeaderTextArea += fmt.Sprintf(
      "%q : %q\n", k, v,
    )
  }


  responseStatusCode = resp.StatusCode
  responseBodyString := fmt.Sprintf(
    "%d\n\n%s",
    responseStatusCode,
    string(responseBody),
  )
  m.response.body.SetContent(responseBodyString)
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer resp.Body.Close()
}
