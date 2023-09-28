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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.url, cmd = m.url.Update(msg)
	return m, cmd
}

func (m model) View() string {
  s := fmt.Sprintf(
    "URL: %s\nRequest type: %s\n%s\n%s\n%s",
    m.url.View(), m.method, m.request.View(), m.response.View(),
    "Press `ctrl+c` or `q` to quit the program...",
  )
	return s
}
