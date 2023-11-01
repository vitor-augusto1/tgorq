package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

// Styles vars
var (
	color                = termenv.EnvColorProfile().Color
	help                 = termenv.Style{}.Foreground(color("241")).Styled
	grey                 = lipgloss.Color("#6c6c6c")
	StyleActivePageOnPaginator = lipgloss.
				NewStyle().
				Foreground(lipgloss.Color("#76fd47")).
				Render("•")
	StyleInactivecCurrentPageOnPaginator = lipgloss.
				NewStyle().
				Foreground(lipgloss.Color("#6c6c6c")).
				Render("•")
	StyleInactivePageOnPaginator = lipgloss.
				NewStyle().
				Foreground(grey).
				Render("-")
	StyleRequestBorder = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("5")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(0).Width(160).Height(1)
	StyleResponseBorder = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("10")).
				BorderStyle(lipgloss.RoundedBorder()).
				Padding(0).Width(160).Height(1)
	StyleInactiveBorder = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("#6c6c6c")).
				BorderStyle(lipgloss.RoundedBorder()).
				Padding(0).Width(160).Height(1)
)

// output file vars
var (
	responseBodyOutputPath    = "./response/body.txt"
	responseHeadersOutputPath = "./response/headers.txt"
)

// Cobra vars
var (
	SaveToFileFlag = false
	SaveStateFlag  = false
	rootCmd    = &cobra.Command{
		Use:   "tgorq",
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
				SaveStateFlag = true
			}
			if outputFlagValue {
				SaveToFileFlag = true
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

type FocusedModel int

const (
	FocusUrlModel FocusedModel = 1 << iota
	FocusMethodModel
	FocusRequestBodyModel
	FocusRequestHeaderModel
	FocusResponseModel
)

type mainModel struct {
	url          *Url
	request      *Request
	rawResponse  *Response
	response     *ResponseModel
	focusedModel FocusedModel

	width  int
	height int
}

func initialModel() mainModel {
	return mainModel{
		url:          InitialUrlModel(),
		request:      InitialRequestModel(),
		rawResponse:  &Response{},
		response:     InitialResponseModel(),
		focusedModel: FocusUrlModel,
	}
}

func (m mainModel) Init() tea.Cmd {
	if !SaveStateFlag {
		return nil
	}
	if m.stateFileExists() {
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
			m.executeRequest()
			m.response.body.GotoTop()
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		// Focus on the URL model
		case tea.KeyCtrlI.String():
      return m.focusOnMethod()
		case tea.KeyCtrlU.String():
      return m.focudOnUrl()
		case tea.KeyCtrlB.String():
      return m.focusOnRequestBody()
		case tea.KeyCtrlR.String():
      return m.focusOnRequestHeader()
		case tea.KeyCtrlS.String():
      return m.focusOnResponse()
		default:
			// Handling each focused model
			switch m.focusedModel {
			// Updating the URL
			case FocusUrlModel:
				m.url.textInput, _ = m.url.textInput.Update(msg)
				return m, nil
			case FocusMethodModel:
				// Update the http method model
				m.url.httpMethodPaginator, _ = m.url.httpMethodPaginator.Update(msg)
				currentPage := m.url.httpMethodPaginator.Page
				m.url.chosenMethod = httpMethod(currentPage)
				return m, nil
			case FocusRequestBodyModel:
				// Now change the focus to the request body textarea
				m.request.body.Focus()
				m.request.body, _ = m.request.body.Update(msg)
				return m, nil
			case FocusRequestHeaderModel:
				// Now change the focus to the request headers textarea
				m.request.headers.Focus()
				m.request.headers, _ = m.request.headers.Update(msg)
				return m, nil
			case FocusResponseModel:
				var currentPage int = m.response.paginator.Page
				switch msg.String() {
				case tea.KeyLeft.String(), tea.KeyRight.String(), "l", "h":
					m.response.paginator, _ = m.response.paginator.Update(msg)
					return m, nil
				// If the user types `ctrl+a` while focus on the response body, the viewport goes to the top
				case tea.KeyCtrlA.String():
					if currentPage == 0 {
						m.response.body.GotoTop()
					}
					return m, nil
				// If the user types `ctrl+e` while focus on the response body, the viewport goes to the bottom
				case tea.KeyCtrlE.String():
					if currentPage == 0 {
						m.response.body.GotoBottom()
					}
					return m, nil
				default:
					var cmd tea.Cmd
					if currentPage == 0 {
						m.response.body, cmd = m.response.body.Update(msg)
					} else {
						m.response.headers, cmd = m.response.headers.Update(msg)
					}
					return m, cmd
				}
			}
		}
	}
	return m, nil
}

func (m mainModel) View() string {
	s := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n\n\n%s",
		lipgloss.JoinHorizontal(lipgloss.Left, m.url.View()),
		m.request.View(),
		m.response.View(),
		help("ctrl+c: Quit • ctrl+u: URL • ctrl+i: Method • ctrl+b: Request Body • ctrl+r: Request Header • ctrl+s: Response • ctrl+g: Make request"),
	)
	lipPlace := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
	return lipPlace
}
