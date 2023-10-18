package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type Response struct {
  body      viewport.Model
  headers   viewport.Model
  paginator paginator.Model

  border    lipgloss.Style
}


func InitialResponseModel() *Response {
  bodyViewPort := viewport.New(160, 9)
  bodyViewPort.SetContent("Response body")

  headersViewPort := viewport.New(160, 9)
  headersViewPort.SetContent("Response body")

  newPaginator := paginator.New()
  newPaginator.Type = paginator.Dots
  newPaginator.SetTotalPages(2)
  newPaginator.ActiveDot = paginatorStyle
	newPaginator.InactiveDot = paginatorStyleInactive

  return &Response{
    body: bodyViewPort,
    headers: headersViewPort,
    paginator: newPaginator,
    border: responseBorderStyle,
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
  return rs.border.Render(sBuilder.String())
}
