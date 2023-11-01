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
      var focusHandler = map[FocusedModel]func(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
        FocusUrlModel: m.updateUrlModel,
        FocusMethodModel: m.updateMethodModel,
        FocusRequestBodyModel: m.updateRequestBodyModel,
        FocusRequestHeaderModel: m.updateRequestHeaderModel,
        FocusResponseModel: m.updateResponseModel,
      }
      return focusHandler[m.focusedModel](msg)
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
