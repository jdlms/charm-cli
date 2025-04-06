package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input        string
	tokenInput   textinput.Model
	repos        []repository
	loading      bool
	errorMessage string
	currentView  view
}

type repository struct {
	Name        string
	Description string
	URL         string
}

type view int

const (
	inputView view = iota
	resultsView
)

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter your Github auth token"
	ti.EchoMode = textinput.EchoPassword
	ti.Focus()

	return model{
		tokenInput:  ti,
		repos:       []repository{},
		currentView: inputView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
