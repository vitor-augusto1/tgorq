package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (m mainModel) handleGetMethod(url string) {
var (
  req *http.Request
  reqHeaders map[string]string
  resp *http.Response
  responseBody []byte
  err error
)

  // Initialize new request
  req, err = http.NewRequest(GET.String(), url, nil)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  // Make the request
  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  // Read the bytes from the response body
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

  // Set response body and headers
  m.response.body.SetContent(string(responseBody))
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer resp.Body.Close()
}


func (m mainModel) handlePostMethod(url string, body io.Reader, headers []byte) {
  var req *http.Request
  var reqHeaders map[string]string
  var res *http.Response
  var responseBody []byte
  var err error

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

  res, err = http.DefaultClient.Do(req)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  responseBody, err = io.ReadAll(res.Body)
  if err != nil {
    m.response.body.SetContent(err.Error())
  }

  var stringToBeStoreInTheResponseHeaderTextArea string
  for k, v := range res.Header {
    stringToBeStoreInTheResponseHeaderTextArea += fmt.Sprintf(
      "%q : %q\n", k, v,
    )
  }

  // Set response body and headers
  m.response.body.SetContent(string(responseBody))
  m.response.headers.SetContent(stringToBeStoreInTheResponseHeaderTextArea)

  defer res.Body.Close()
}
