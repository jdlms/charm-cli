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
			m.chunks = msg.chunks
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
				return m, fetchGitHubRepos(m)
			}
		case "b":
		}
		return m, cmd
	}
	if m.currentView == resultsView {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.chunks[m.reposPage])-1 {
				m.cursor++
			}
		case "enter", " ":
			currentChunk := m.chunks[m.reposPage]
			if m.cursor >= len(currentChunk) {
				break
			}
			repo := currentChunk[m.cursor]

			if _, ok := m.selected[repo.ID]; ok {
				delete(m.selected, repo.ID)
			} else {
				m.selected[repo.ID] = repo
			}
		case "h", "left":
			m.reposPage--
			m.cursor = 0
		case "l", "right":
			m.reposPage++
			m.cursor = 0
		case "b":
			if m.currentView == resultsView {
				m.currentView = inputView
			}
		case "s":
			if m.currentView == resultsView {
				m.currentView = selectedView
			}
		}
	}
	if m.currentView == selectedView {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "b":
			if m.currentView == selectedView {
				m.currentView = resultsView
			}
		}

	}
	return m, nil
}
