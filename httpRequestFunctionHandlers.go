package main

import (
	"fmt"
	"io"
	"net/http"
)

func (m mainModel) handleGetMethod(url string) {
  var req *http.Request
  var resp *http.Response
  var responseBody []byte
  var err error

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
