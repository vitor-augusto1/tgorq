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
                    Padding(1).Width(158),
  }
}

func (u Url) Init() tea.Cmd {
  return nil
}

func (u *Url) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  u.url, _ = u.url.Update(msg)
  return u, nil
}

func (u Url) View() string {
  s := u.style.Render(u.url.View())
  return s
}
