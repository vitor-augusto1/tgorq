package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)


type Response struct {
  body      textarea.Model
  headers   textarea.Model
  paginator paginator.Model
}


func InitialResponseModel() *Response {
  bodyTextArea := textarea.New()
  bodyTextArea.Placeholder = "Reponse Body"
  bodyTextArea.SetValue("")

  headersTextArea := textarea.New()
  headersTextArea.Placeholder = "Response Header"
  headersTextArea.SetValue("")

}
