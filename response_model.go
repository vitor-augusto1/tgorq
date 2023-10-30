package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResponseModel struct {
	body      viewport.Model
	headers   viewport.Model
	paginator paginator.Model

	border lipgloss.Style
}

func InitialResponseModel() *ResponseModel {
	bodyViewPort := viewport.New(160, 9)
	bodyViewPort.SetContent("ResponseModel body")

	headersViewPort := viewport.New(160, 9)
	headersViewPort.SetContent("ResponseModel headers")

	newPaginator := paginator.New()
	newPaginator.Type = paginator.Dots
	newPaginator.SetTotalPages(2)
	newPaginator.ActiveDot = inactivePaginatorStyle
	newPaginator.InactiveDot = paginatorStyleInactive

	return &ResponseModel{
		body:      bodyViewPort,
		headers:   headersViewPort,
		paginator: newPaginator,
		border:    inactiveModelStyle,
	}
}

func (rs ResponseModel) Init() tea.Cmd {
	return nil
}

func (rs ResponseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return rs, nil
}

func (rs ResponseModel) View() string {
	var sBuilder strings.Builder

	responseSlice := []string{rs.body.View(), rs.headers.View()}
	start, end := rs.paginator.GetSliceBounds(len(responseSlice))
	for _, item := range responseSlice[start:end] {
		sBuilder.WriteString(string(item) + "\n")
	}

	sBuilder.WriteString("  " + rs.paginator.View())
	return rs.border.Render(sBuilder.String())
}
