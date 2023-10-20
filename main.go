package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(
    initialModel(),
    tea.WithANSICompressor(),
    tea.WithAltScreen(),
  )
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ain't no way, boy! %v", err)
		os.Exit(1)
	}
}

func (m mainModel) makeRequest() {
  // Get the user URL
  url := m.url.textInput.Value()
  chosenHttpMethod := m.url.chosenMethod

  if chosenHttpMethod == GET {
    m.handleGetMethod(url)
  } else if chosenHttpMethod == POST {
    bodyString := m.request.body.Value()
    headerString := m.request.headers.Value()
    byteBody := bytes.NewBuffer([]byte(bodyString))
    byteHeaders := []byte(headerString)
    m.handlePostMethod(url, byteBody, byteHeaders)
  } else if chosenHttpMethod == PUT {
    return
  } else if chosenHttpMethod == DELETE {
    return
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
  borderStyle = lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("5")).
                    BorderStyle(lipgloss.RoundedBorder()).
                    Padding(0).Width(160).Height(1)
  responseBorderStyle = lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("10")).
                    BorderStyle(lipgloss.RoundedBorder()).
                    Padding(0).Width(160).Height(1)
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
    case tea.KeyCtrlG.String():
      m.makeRequest()
      m.response.body.GotoTop()
    case tea.KeyCtrlC.String():
      return m, tea.Quit
    // Focus on the URL model
    case tea.KeyCtrlI.String():
      m.focusedModel = FocusMethod
      return m, nil
    case tea.KeyCtrlU.String():
      m.url.textInput.Cursor.SetMode(cursor.CursorBlink)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
      m.focusedModel = FocusUrl
      return m, nil
    case tea.KeyCtrlB.String():
      m.request.body.Cursor.SetMode(cursor.CursorBlink)
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
      m.request.body.Cursor.Focus()
      m.focusedModel = FocusRequestB
      return m, nil
    case tea.KeyCtrlR.String():
      m.request.headers.Cursor.SetMode(cursor.CursorBlink)
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.Focus()
      m.focusedModel = FocusRequestH
      return m, nil
    case tea.KeyCtrlS.String():
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
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
        var currentPage int = m.response.paginator.Page
        switch msg.String() {
        case tea.KeyLeft.String(), tea.KeyRight.String(), "l", "h":
          m.response.paginator, _ = m.response.paginator.Update(msg)
          return m, nil
        // If the user types `ctrl+a` while focus on the response body, the viewport goes to the top
        case tea.KeyCtrlA.String():
          if currentPage == 0 { m.response.body.GotoTop() }
          return m, nil
        // If the user types `ctrl+e` while focus on the response body, the viewport goes to the bottom
        case tea.KeyCtrlE.String():
          if currentPage == 0 { m.response.body.GotoBottom() }
          return m, nil
        default:
          var cmd tea.Cmd
          if currentPage == 0 {
            m.response.body, cmd = m.response.body.Update(msg)
          } else { m.response.headers, cmd = m.response.headers.Update(msg) }
          return m, cmd
        }
      }
    }
  }
  return m, nil
}

func (m mainModel) View() string {
  s := fmt.Sprintf(
    "%s\n\n%s\n\n%s\n\n%s",
    lipgloss.JoinHorizontal(lipgloss.Left, m.url.View()),
    m.request.View(),
    m.response.View(),
    "Press `ctrl+c` or `q` to quit the program...",
  )
  
	return s
}
