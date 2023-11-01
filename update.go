package main

import (
	tea "github.com/charmbracelet/bubbletea"
)


func (m mainModel) updateUrlModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
				m.url.textInput, _ = m.url.textInput.Update(msg)
				return m, nil
}

func (m mainModel) updateMethodModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
				// Update the http method model
				m.url.httpMethodPaginator, _ = m.url.httpMethodPaginator.Update(msg)
				currentPage := m.url.httpMethodPaginator.Page
				m.url.chosenMethod = httpMethod(currentPage)
				return m, nil
}

func (m mainModel) updateRequestBodyModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
				// Now change the focus to the request body textarea
				m.request.body.Focus()
				m.request.body, _ = m.request.body.Update(msg)
				return m, nil
}

func (m mainModel) updateRequestHeaderModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
				// Now change the focus to the request headers textarea
				m.request.headers.Focus()
				m.request.headers, _ = m.request.headers.Update(msg)
				return m, nil
}

func (m mainModel) updateResponseModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
