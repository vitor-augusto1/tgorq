package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
  paginatorStyle = lipgloss.
                        NewStyle().
                        Foreground(lipgloss.AdaptiveColor{
                          Light: "235", Dark: "252",
                        }).Render("•")
)

type httpMethod string

const (
  GET = "GET"
  POST = "POST"
  PUT = "PUT"
  DELETE = "DELETE"
)

type Url struct {
   methods []httpMethod
   chosenMethod httpMethod 

   textInput textinput.Model

   borderStyle lipgloss.Style
   httpMethodPag paginator.Model
}

func InitialUrlModel() *Url {
  newTextInput := textinput.New()
  newTextInput.Placeholder = "https://www.example.com/"
  newTextInput.Focus()

  newBorderStyle := lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("36")).
                    BorderStyle(lipgloss.RoundedBorder()).
                    Padding(0).Width(160).Height(1)

  newPaginator.ActiveDot = paginatorStyle
	newPaginator.InactiveDot = paginatorStyle 
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
