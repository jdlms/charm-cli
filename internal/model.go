package internal

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input          string
	tokenInput     textinput.Model
	username       string
	cursor         int
	reposPage      int
	selected       map[int]repository
	chunks         map[int][]repository
	loading        bool
	errorMessage   string
	successMessage string
	currentView    view
}

type repository struct {
	ID          int
	Name        string
	Description string
	URL         string
}

type view int

const (
	inputView view = iota
	resultsView
	selectedView
)

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter your Github auth token"
	ti.EchoMode = textinput.EchoPassword
	ti.Focus()

	return model{
		tokenInput:  ti,
		selected:    make(map[int]repository),
		chunks:      make(map[int][]repository),
		currentView: inputView,
		cursor:      0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
