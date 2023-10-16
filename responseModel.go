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
  bodyTextArea.MaxHeight = 0
  bodyTextArea.CharLimit = 0
  bodyTextArea.SetWidth(150)
  bodyTextArea.SetHeight(11)

  headersTextArea := textarea.New()
  headersTextArea.Placeholder = "Response Header"
  headersTextArea.MaxHeight = 0
  headersTextArea.CharLimit = 0
  headersTextArea.SetWidth(150)
  headersTextArea.SetHeight(11)

  newPaginator := paginator.New()
  newPaginator.Type = paginator.Dots
  newPaginator.SetTotalPages(2)
  newPaginator.ActiveDot = paginatorStyle
	newPaginator.InactiveDot = paginatorStyleInactive

  return &Response{
    body: bodyTextArea,
    headers: headersTextArea,
    paginator: newPaginator,
  }
}

func (rs Response) Init() tea.Cmd {
  return nil
}

func (rs Response) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return rs, nil
}

func (rs Response) View() string {
  var sBuilder strings.Builder
  
  responseSlice := []string{rs.body.View(), rs.headers.View()}
  start, end := rs.paginator.GetSliceBounds(len(responseSlice))
  for _, item := range responseSlice[start:end] {
    sBuilder.WriteString(string(item) + "\n")
  }

  sBuilder.WriteString("  " + rs.paginator.View())

  return sBuilder.String()
}
