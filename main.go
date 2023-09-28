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

