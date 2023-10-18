package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type httpMethod int
const (
  GET httpMethod = iota
  POST
  PUT
  DELETE
)

func (hm httpMethod) String() string {
  return []string{"GET", "POST", "PUT", "DELETE"}[hm]
}

type Url struct {
   methods []string
   chosenMethod httpMethod 

   textInput textinput.Model

   borderStyle lipgloss.Style
   httpMethodPag paginator.Model
}

func InitialUrlModel() *Url {
  newTextInput := textinput.New()
  newTextInput.Placeholder = "https://www.example.com/"
  newTextInput.Focus()

  methodsSlice := []string{"GET", "POST", "PUT", "DELETE"}
  newPaginator := paginator.New()
  newPaginator.Type = paginator.Dots
  newPaginator.SetTotalPages(len(methodsSlice))
  newPaginator.ActiveDot = paginatorStyle
	newPaginator.InactiveDot = paginatorStyleInactive

  return &Url {
    methods: methodsSlice,
    chosenMethod: GET,
    textInput: newTextInput,
    borderStyle: borderStyle,
    httpMethodPag: newPaginator,
  }
}

func (u Url) Init() tea.Cmd {
  return nil
}

func (u *Url) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return u, nil
}

func (u Url) View() string {
  var sBuilder strings.Builder
  
  start, end := u.httpMethodPag.GetSliceBounds(len(u.methods))
  for _, method := range u.methods[start:end] {
    sBuilder.WriteString("  " + string(method) + "\n")
  }

  sBuilder.WriteString("  " + u.httpMethodPag.View())

  s := fmt.Sprintf(
    "\n%s\n\n%s\n",
      sBuilder.String(),
      u.textInput.View(),
  )
  return  u.borderStyle.Render(s)
}
