package internal

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
			m.username = msg.username
		}
		return m, nil
	case deleteResultMsg:
		if msg.err != nil {
			m.errorMessage = msg.message
		} else {
			m.successMessage = msg.message
		}

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
			if m.cursor < len(m.chunks[m.reposPage]) {
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
				m.selected[repo.ID-1] = repo
			}
		case "h", "left":
			if m.reposPage > 0 {
				m.reposPage--
				m.cursor = 0
			} else {
				break
			}
		case "l", "right":
			if m.reposPage < len(m.chunks)-1 {
				m.reposPage++
				m.cursor = 0
			} else {
				break
			}
		case "b":
			if m.currentView == resultsView {
				m.currentView = inputView
			}
		case "s":
			if m.currentView == resultsView {
				m.currentView = selectedView
			} else {
				break
			}
		}
	}
	if m.currentView == selectedView {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "d":
			if m.currentView == selectedView {
				if m.tokenInput.Value() != "" {
					m.loading = true
					m.input = m.tokenInput.Value()
					return m, deleteGitHubRepos(m)
				}
			} else {
				break
			}
		case "b":
			if m.currentView == selectedView {
				m.currentView = resultsView
			}
		}

	}
	return m, nil
}
