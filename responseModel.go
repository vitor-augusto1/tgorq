package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Response struct {
	resHeaders map[string]string
	resBody    string
}

func InitalResponseModel() Response {
  return Response{
    resHeaders: map[string]string{"Access-Control-Allow-Origin": "*"},
    resBody: "{}",
  }
}

func (r Response) Init() tea.Cmd {
  return nil
}

func (r Response) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return r, nil
}

