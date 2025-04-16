package main

import (
	"fmt"
)

func (m model) View() string {
	switch m.currentView {
	case inputView:
		return m.inputView()
	case resultsView:
		return m.resultsView()
	case selectedView:
		return m.selectedView()
	default:
		return ""
	}
}

func (m model) inputView() string {
	s := "Enter your Github auth token:\n\n(Press Enter to continue)"

	s += m.tokenInput.View() + "\n\n"

	if m.loading {
		s += "Loading repositories...\n"
	}

	if m.errorMessage != "" {
		s += fmt.Sprintf("Error: %s\n\n", m.errorMessage)
	}

	s += "\nPress Enter to fetch repositories, q to quit.\n"
	return s
}

func (m model) resultsView() string {
	s := "Your current repositories:\n\n"

	if len(m.chunks) == 0 {
		s += "No repositories found.\n"
	} else {
		currentChunk := m.chunks[m.reposPage]
		for i, repo := range currentChunk {
			desc := repo.Description
			if desc == "" {
				desc = "(No description)"
			}

			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.selected[repo.ID-1]; ok {
				checked = "x"
			}

			s += fmt.Sprintf("%s [%s] %d. %s\n", cursor, checked, repo.ID, repo.Name)
			s += fmt.Sprintf("   %s\n", desc)
			s += fmt.Sprintf("   %s\n\n", repo.URL)
		}
	}

	s += "\nPress s to select, l + n to page, b to go back, q to quit.\n"
	return s
}

func (m model) selectedView() string {
	s := "This is the selectedView\n\n"
	if len(m.selected) == 0 {
		s += "No repositories selected.\n"
	} else {
		for i, repo := range m.selected {
			desc := repo.Description
			if desc == "" {
				desc = "(No description)"
			}
			s += fmt.Sprintf("[x] %d. %s\n", i+1, repo.Name)
			s += fmt.Sprintf("   %s\n", desc)
			s += fmt.Sprintf("   %s\n\n", repo.URL)
		}
	}

	if m.errorMessage != "" {
		s += m.errorMessage
		s += "\nPress q to quit.\n"
	}
	if m.successMessage != "" {
		s += m.successMessage
		s += "\nPress q to quit.\n"
	} else {
		s += "\nPress d to delete, b to go back, q to quit.\n"
	}
	return s
}
