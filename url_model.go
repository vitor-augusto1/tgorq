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
	methods      []string
	chosenMethod httpMethod

	textInput textinput.Model

	border   lipgloss.Style
	httpMethodPaginator paginator.Model
}

func InitialUrlModel() *Url {
	newTextInput := textinput.New()
	newTextInput.Placeholder = "https://www.example.com/"
	newTextInput.Focus()

	methodsSlice := []string{"GET", "POST", "PUT", "DELETE"}
	newPaginator := paginator.New()
	newPaginator.Type = paginator.Dots
	newPaginator.SetTotalPages(len(methodsSlice))
	newPaginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
	newPaginator.InactiveDot = StyleInactivePageOnPaginator

	return &Url{
		methods:       methodsSlice,
		chosenMethod:  GET,
		textInput:     newTextInput,
		border:   StyleRequestBorder,
		httpMethodPaginator: newPaginator,
	}
}

func (u Url) Init() tea.Cmd {
	return nil
}

func (u *Url) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return u, nil
}

func (u Url) View() string {
	var stringBuilder strings.Builder

	start, end := u.httpMethodPaginator.GetSliceBounds(len(u.methods))
	for _, method := range u.methods[start:end] {
		stringBuilder.WriteString("  " + string(method) + "\n")
	}

	stringBuilder.WriteString("  " + u.httpMethodPaginator.View())

	s := fmt.Sprintf(
		"\n%s\n\n%s\n",
		stringBuilder.String(),
		u.textInput.View(),
	)
	return u.border.Render(s)
}
