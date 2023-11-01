package main

import (
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)


func (m mainModel) focusOnMethod() (tea.Model, tea.Cmd) {
			m.focusedModel = FocusMethodModel
			m.response.border = StyleInactiveBorder
			m.response.paginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.url.httpMethodPaginator.ActiveDot = StyleActivePageOnPaginator
      return m, nil
}

func (m mainModel) focudOnUrl() (tea.Model, tea.Cmd) {
			m.url.textInput.Cursor.SetMode(cursor.CursorBlink)
			m.request.body.Cursor.SetMode(cursor.CursorHide)
			m.request.headers.Cursor.SetMode(cursor.CursorHide)
			m.response.border = StyleInactiveBorder
			m.response.paginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.url.httpMethodPaginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.focusedModel = FocusUrlModel
			return m, nil
}

func (m mainModel) focusOnRequestBody() (tea.Model, tea.Cmd) {
			m.request.body.Cursor.SetMode(cursor.CursorBlink)
			m.url.textInput.Cursor.SetMode(cursor.CursorHide)
			m.request.headers.Cursor.SetMode(cursor.CursorHide)
			m.response.border = StyleInactiveBorder
			m.response.paginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.url.httpMethodPaginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.request.body.Cursor.Focus()
			m.focusedModel = FocusRequestBodyModel
			return m, nil
}

func (m mainModel) focusOnRequestHeader() (tea.Model, tea.Cmd) {
			m.request.headers.Cursor.SetMode(cursor.CursorBlink)
			m.url.textInput.Cursor.SetMode(cursor.CursorHide)
			m.request.body.Cursor.SetMode(cursor.CursorHide)
			m.request.headers.Cursor.Focus()
			m.response.border = StyleInactiveBorder
			m.response.paginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.url.httpMethodPaginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.focusedModel = FocusRequestHeaderModel
			return m, nil
}

func (m mainModel) focusOnResponse() (tea.Model, tea.Cmd) {
			m.url.textInput.Cursor.SetMode(cursor.CursorHide)
			m.request.body.Cursor.SetMode(cursor.CursorHide)
			m.request.headers.Cursor.SetMode(cursor.CursorHide)
			m.response.border = StyleResponseBorder
			m.response.paginator.ActiveDot = StyleActivePageOnPaginator
			m.url.httpMethodPaginator.ActiveDot = StyleInactivecCurrentPageOnPaginator
			m.focusedModel = FocusResponseModel
			return m, nil
}
