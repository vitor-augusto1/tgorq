package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
    initialModel(),
    tea.WithAltScreen(),
    tea.WithMouseCellMotion(),
  )
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ain't no way, boy! %v", err)
		os.Exit(1)
	}
}

type FocusedModel int
const (
  FocusUrl FocusedModel = 1 << iota
  FocusMethod
  FocusRequest
  FocusResponse
)

type mainModel struct {
  method     *MethodModel
	url        *Url
  request    Request
  response   Response
  focusedModel FocusedModel
}

func initialModel() mainModel {
  return mainModel{
    method: InitialMethodModel(),
    url: InitialUrlModel(),
    request: InitialRequestModel(),
    response: InitalResponseModel(),
    focusedModel: FocusUrl,
  }
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "q":
      if m.focusedModel == FocusUrl {
        m.url.url, _ = m.url.url.Update(msg)
        return m, nil
      }
      return m, tea.Quit
    case "ctrl+c":
      return m, tea.Quit
}

func (m mainModel) View() string {
  s := fmt.Sprintf(
    "URL: %s\nRequest type: %s\n%s\n%s\n%s",
    m.url.View(), m.method, m.request.View(), m.response.View(),
    "Press `ctrl+c` or `q` to quit the program...",
  )
	return s
}
