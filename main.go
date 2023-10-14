package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

var (
  orange = lipgloss.Color("#fd6a47")
  paginatorStyle = lipgloss.
                        NewStyle().
                        Foreground(lipgloss.AdaptiveColor{
                          Light: "235", Dark: "252",
                        }).Render("•")
  paginatorStyleInactive = lipgloss.
                                NewStyle().
                                Foreground(orange).
                                Render("-")
)

type FocusedModel int
const (
  FocusUrl FocusedModel = 1 << iota
  FocusMethod
  FocusRequestB
  FocusRequestH
  FocusResponse
)

type mainModel struct {
	url          *Url
  request      *Request
  response     *Response
  focusedModel FocusedModel
}

func initialModel() mainModel {
  return mainModel{
    url: InitialUrlModel(),
    request: InitialRequestModel(),
    response: InitialResponseModel(),
    focusedModel: FocusUrl,
  }
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    // Handling the msg being sent
    switch msg.String() {
    // Quit the program if its not focus on the URL model
    case "q":
      if m.focusedModel == FocusUrl {
        m.url.textInput, _ = m.url.textInput.Update(msg)
        return m, nil
      }
      return m, tea.Quit
    case tea.KeyCtrlC.String():
      return m, tea.Quit
    // Focus on the URL model
    case tea.KeyCtrlI.String():
      m.focusedModel = FocusMethod
      return m, nil
    case tea.KeyCtrlU.String():
      m.focusedModel = FocusUrl
      return m, nil
    case tea.KeyCtrlB.String():
      m.focusedModel = FocusRequestB
      return m, nil
    case tea.KeyCtrlR.String():
      m.focusedModel = FocusRequestH
      return m, nil
    case tea.KeyCtrlS.String():
      m.focusedModel = FocusResponse
      return m, nil
    default:
      // Handling each focused model
      switch m.focusedModel {
      // Updating the URL
      case FocusUrl:
        m.url.textInput, _ = m.url.textInput.Update(msg)
        return m, nil
      case FocusMethod:
        // Update the http method model
        m.url.httpMethodPag, _ = m.url.httpMethodPag.Update(msg)
        currentPage := m.url.httpMethodPag.Page
        m.url.chosenMethod = httpMethod(currentPage)
        return m, nil
      case FocusRequestB:
        // Now change the focus to the request body textarea
        m.request.body.Focus()
        m.request.body, _ = m.request.body.Update(msg)
        return m, nil
      case FocusRequestH:
        // Now change the focus to the request headers textarea
        m.request.headers.Focus()
        m.request.headers, _ = m.request.headers.Update(msg)
        return m, nil
      case FocusResponse:
        m.response.paginator, _ = m.response.paginator.Update(msg)
        return m, nil
      }
      return m, nil
    }
  }
  return m, nil
}

func (m mainModel) View() string {
  s := fmt.Sprintf(
    "%s\n\n%s\n\n%s",
    lipgloss.JoinHorizontal(lipgloss.Left, m.url.View()),
    m.request.View(),
    "Press `ctrl+c` or `q` to quit the program...",
  )
  
	return s
}
