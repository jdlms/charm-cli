package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case fetchResultMsg:
		m.loading = false
		if msg.err != nil {
			m.errorMessage = msg.err.Error()
		} else {
			m.repos = msg.repos
			m.currentView = resultsView
		}
		return m, nil
	}
	return m, nil
}

func (m model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.currentView == inputView {
		var cmd tea.Cmd
		m.tokenInput, cmd = m.tokenInput.Update(msg)

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.tokenInput.Value() != "" {
				m.loading = true
				m.input = m.tokenInput.Value()
				return m, fetchGitHubRepos(m.input)
			}
		case "b":
		}
		return m, cmd
	}

	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "b":
		if m.currentView == resultsView {
			m.currentView = inputView
		}
	}

	return m, nil
}
