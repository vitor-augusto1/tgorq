package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type httpMethod string
const (
  GET = "GET"
  POST = "POST"
  PUT = "PUT"
  PATCH = "PATCH"
  DELETE = "DELETE"
)

type MethodModel struct {
  methodType []httpMethod
  chosenMethod httpMethod
  paginator paginator.Model
}

func InitialMethodModel() *MethodModel {
  methodsSlice := []httpMethod{GET, POST, PUT, PATCH, DELETE}
  newPaginator := paginator.New()
  newPaginator.Type = paginator.Dots
  newPaginator.SetTotalPages(len(methodsSlice))
  newPaginator.ActiveDot = lipgloss.NewStyle().
                                    Foreground(lipgloss.AdaptiveColor{
                                               Light: "235", Dark: "252"}).
                                    Render("•")
	newPaginator.InactiveDot = lipgloss.NewStyle().
                                      Foreground(lipgloss.AdaptiveColor{
                                        Light: "250", Dark: "238"}).Render("•")
  return &MethodModel{
    methodType: methodsSlice,
    chosenMethod: GET,
    paginator: newPaginator,
  }
}

func (mm MethodModel) Init() tea.Cmd {
  return nil
}

func (mm MethodModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  mm.paginator, cmd = mm.paginator.Update(msg)
  return mm, cmd
}

