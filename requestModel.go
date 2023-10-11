package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Request struct {
  body      textarea.Model
  headers   textarea.Model
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
  return sBuilder.String()
}
