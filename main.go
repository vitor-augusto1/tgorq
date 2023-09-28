package main

type model struct {
	url        textinput.Model
	method     string
  request    Request
  response   Response
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "https://www.example.com/"
	ti.Focus()
  return model{
    url: ti,
    method: "GET",
    request: InitialRequestModel(),
    response: InitalResponseModel(),
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

