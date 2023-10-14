package main

import (
	"io"
	"net/http"
)

func (m mainModel) handleGetMethod(url string) {
  res, err := http.Get(url)
  if err != nil {
    m.response.body.SetValue(err.Error())
  }
  // Receive the bytes from the request
  body, err := io.ReadAll(res.Body)
  if err != nil {
    m.response.body.SetValue(err.Error())
  }

  defer res.Body.Close()

  // Convert the bytes to string
  toStringBody := string(body)
  m.response.body.SetValue(toStringBody)
  m.response.headers.SetValue(string(body)
}
