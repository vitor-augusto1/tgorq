package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Request struct {
  body      textarea.Model
  headers   textarea.Model

  border    lipgloss.Style
}

func InitialRequestModel() *Request {
  bodyTextArea := textarea.New()
  bodyTextArea.Placeholder = "{request: body}"

  headersTextArea := textarea.New()
  headersTextArea.Placeholder = "{request: headers}"

  return &Request {
    body: bodyTextArea,
    headers: headersTextArea,

    border: borderStyle,
  }
}

func (rq Request) Init() tea.Cmd {
  	return nil
}

func (rq Request) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return rq, nil
}

func (rq Request) View() string {
  var sBuilder strings.Builder
  
  sBuilder.WriteString(rq.body.View() + "\n\n")
  sBuilder.WriteString(rq.headers.View())
  return rq.border.Render(sBuilder.String())
}
