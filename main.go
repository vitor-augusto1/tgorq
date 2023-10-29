package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/spf13/cobra"
)

func main() {
  Execute()
}

// output file vars
var (
  responseBodyOutput = "./response/body.txt"
  responseHeadersOutput = "./response/headers.txt"
)

// Cobra vars
var (
  SaveToFile = false
  SaveState = false
  rootCmd = &cobra.Command{
    Use: "tgorq",
    Short: "Make http requests from the terminal",
    Long: `
    
__/\\\\\\\\\\\\\\\________________________________________________________        
 _\///////\\\/////_________________________________________________________       
  _______\/\\\_________/\\\\\\\\________________________________/\\\\\\\\___      
   _______\/\\\________/\\\////\\\_____/\\\\\_____/\\/\\\\\\\___/\\\////\\\__     
    _______\/\\\_______\//\\\\\\\\\___/\\\///\\\__\/\\\/////\\\_\//\\\\\\\\\__    
     _______\/\\\________\///////\\\__/\\\__\//\\\_\/\\\___\///___\///////\\\__   
      _______\/\\\________/\\_____\\\_\//\\\__/\\\__\/\\\________________\/\\\__  
       _______\/\\\_______\//\\\\\\\\___\///\\\\\/___\/\\\________________\/\\\\_ 
        _______\///_________\////////______\/////_____\///_________________\////__

    A vim-like TUI (Text User Interface) that allows you to make http requests.
    Example: ./tgorq [ -o | --enable-output ] [ -s | --save-state ]
    `,
    Run: func(cmd *cobra.Command, args []string) {
      outputFlagValue, err := cmd.Flags().GetBool("enable-output")
      if err != nil {
        log.Fatal(err)
      }
      saveStateFlagValue, err := cmd.Flags().GetBool("save-state")
      if err != nil {
        log.Fatal(err)
      }
      if saveStateFlagValue {
        SaveState = true
      }
      if outputFlagValue {
        SaveToFile = true
      }
      p := tea.NewProgram(
        initialModel(),
        tea.WithANSICompressor(),
        tea.WithAltScreen(),
      )
      f, err := tea.LogToFile("debug.log", "debug")
      if err != nil {
        log.Fatal(err)
      }
      defer f.Close()
      if _, err := p.Run(); err != nil {
        fmt.Printf("Ain't no way, boy! %v", err)
        os.Exit(1)
      }
    },
  }
)

func init() {
  rootCmd.Flags().BoolP(
    "enable-output", "o",
    false, `Stores the response body and headers in the response directory.`,
  )
  rootCmd.Flags().BoolP(
    "save-state", "s",
    false, `Save the current application state.`,
  )
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}

func (m mainModel) createOutputFile(content string, pathname string) {
  dir := filepath.Dir(pathname)
  if err := os.MkdirAll(dir, os.ModePerm); err != nil {
    log.Println(err)
    return
  }
  file, err := os.Create(pathname)
  if err != nil {
    log.Println("Error creating the output file: ", err)
    return
  }
  defer file.Close()
  
  fmt.Fprintf(file, "%s", content)
}

func (m mainModel) makeRequest() {
  url := m.url.textInput.Value()
  chosenHttpMethod := m.url.chosenMethod
  bodyString := m.request.body.Value()
  headerString := m.request.headers.Value()
  byteBody := bytes.NewBuffer([]byte(bodyString))
  byteHeaders := []byte(headerString)
  if chosenHttpMethod == GET {
    response, err := handleGetMethod(url)
    if err != nil {
      m.response.body.SetContent(err.Error())
      return
    }
    m.response.body.SetContent(response.body)
    m.response.headers.SetContent(response.headers)
    m.rawResponse = response
    if SaveToFile {
      m.createOutputFile(response.body, responseBodyOutput)
      m.createOutputFile(response.headers, responseHeadersOutput)
    }
  } else if chosenHttpMethod == POST {
    response, err := handlePostMethod(url, byteBody, byteHeaders)
    if err != nil {
      m.response.body.SetContent(err.Error())
      return
    }
    m.response.body.SetContent(response.body)
    m.response.headers.SetContent(response.headers)
    m.rawResponse = response
    if SaveToFile {
      m.createOutputFile(response.body, responseBodyOutput)
      m.createOutputFile(response.headers, responseHeadersOutput)
    }
  } else if chosenHttpMethod == PUT {
    response, err := handlePutMethod(url, byteBody, byteHeaders)
    if err != nil {
      m.response.body.SetContent(err.Error())
      return
    }
    m.response.body.SetContent(response.body)
    m.response.headers.SetContent(response.headers)
    m.rawResponse = response
    if SaveToFile {
      m.createOutputFile(response.body, responseBodyOutput)
      m.createOutputFile(response.headers, responseHeadersOutput)
    }
  } else if chosenHttpMethod == DELETE {
    response, err := handleDeleteMethod(url, byteHeaders)
    if err != nil {
      m.response.body.SetContent(err.Error())
      return
    }
    m.response.body.SetContent(response.body)
    m.response.headers.SetContent(response.headers)
    m.rawResponse = response
    if SaveToFile {
      m.createOutputFile(response.body, responseBodyOutput)
      m.createOutputFile(response.headers, responseHeadersOutput)
    }
  }
  if SaveState {
    m.storeCurrentState()
  }
}

// Styles vars
var (
  grey = lipgloss.Color("#6c6c6c")
  activePaginatorStyle = lipgloss.
                        NewStyle().
                        Foreground(lipgloss.Color("#76fd47")).
                        Render("•")
  inactivePaginatorStyle = lipgloss.
                        NewStyle().
                        Foreground(lipgloss.Color("#6c6c6c")).
                        Render("•")
  paginatorStyleInactive = lipgloss.
                                NewStyle().
                                Foreground(grey).
                                Render("-")
  borderStyle = lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("5")).
                    BorderStyle(lipgloss.RoundedBorder()).
                    Padding(0).Width(160).Height(1)
  responseBorderStyle = lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("10")).
                    BorderStyle(lipgloss.RoundedBorder()).
                    Padding(0).Width(160).Height(1)
  inactiveModelStyle = lipgloss.NewStyle().
                    BorderForeground(lipgloss.Color("#6c6c6c")).
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
  rawResponse  *Response
  response     *ResponseModel
  focusedModel FocusedModel

  width        int
  height       int
}

func initialModel() mainModel {
  return mainModel{
    url: InitialUrlModel(),
    request: InitialRequestModel(),
    rawResponse: &Response{},
    response: InitialResponseModel(),
    focusedModel: FocusUrl,
  }
}

func (m mainModel) Init() tea.Cmd {
  if !SaveState {
    return nil
  }
  if (m.stateFileExists()) {
    m.restorePreviousState()
  }
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    log.Println("This is the current window size: ", msg.Width, msg.Height)
    m.width = msg.Width
    m.height = msg.Height
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
      m.response.border = inactiveModelStyle
      m.response.paginator.ActiveDot = inactivePaginatorStyle
      m.url.httpMethodPag.ActiveDot = activePaginatorStyle
      return m, nil
    case tea.KeyCtrlU.String():
      m.url.textInput.Cursor.SetMode(cursor.CursorBlink)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
      m.response.border = inactiveModelStyle
      m.response.paginator.ActiveDot = inactivePaginatorStyle
      m.url.httpMethodPag.ActiveDot = inactivePaginatorStyle
      m.focusedModel = FocusUrl
      return m, nil
    case tea.KeyCtrlB.String():
      m.request.body.Cursor.SetMode(cursor.CursorBlink)
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
      m.response.border = inactiveModelStyle
      m.response.paginator.ActiveDot = inactivePaginatorStyle
      m.url.httpMethodPag.ActiveDot = inactivePaginatorStyle
      m.request.body.Cursor.Focus()
      m.focusedModel = FocusRequestB
      return m, nil
    case tea.KeyCtrlR.String():
      m.request.headers.Cursor.SetMode(cursor.CursorBlink)
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.Focus()
      m.response.border = inactiveModelStyle
      m.response.paginator.ActiveDot = inactivePaginatorStyle
      m.url.httpMethodPag.ActiveDot = inactivePaginatorStyle
      m.focusedModel = FocusRequestH
      return m, nil
    case tea.KeyCtrlS.String():
      m.url.textInput.Cursor.SetMode(cursor.CursorHide)
      m.request.body.Cursor.SetMode(cursor.CursorHide)
      m.request.headers.Cursor.SetMode(cursor.CursorHide)
      m.response.border = responseBorderStyle
      m.response.paginator.ActiveDot = activePaginatorStyle
      m.url.httpMethodPag.ActiveDot = inactivePaginatorStyle
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
    "Press `ctrl+c` to quit the program...",
  )
  lipPlace := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
  return lipPlace
}                                                           
