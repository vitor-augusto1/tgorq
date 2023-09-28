package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type httpMethod string
const (
  GET = "GET"
  POST = "POST"
  PUT = "PUT"
  PATCH = "PATCH"
  DELETE = "DELETE"
)

type MethodModel struct {
  methodType []httpMethod
  chosenMethod httpMethod
  paginator paginator.Model
}

