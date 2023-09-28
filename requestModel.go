package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Request struct {
	reqHeaders map[string]string
	reqBody    map[string]string
}

func (r Request) Init() tea.Cmd {
	return nil
}

func InitialRequestModel() Request {
  return Request{
    reqHeaders: map[string]string{"User-Agent": "Mozilla/5.0"},
    reqBody: map[string]string{"type": "session"},
  }
}

