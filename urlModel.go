package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Url struct {
   url textinput.Model
   style lipgloss.Style
}

func InitialUrlModel() *Url {
  ti := textinput.New()
  ti.Placeholder = "https://www.example.com/"
  ti.Focus()
  return &Url {
    url: ti,
    style: lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("36")).
                    BorderStyle(lipgloss.NormalBorder()).
                    Padding(1).Width(80),
  }
}

func (u Url) Init() tea.Cmd {
  return nil
}

