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
	s := fmt.Sprintf("Your current repositories")

	if len(m.repos) == 0 {
		s += "No repositories found.\n"
	} else {
		for i, repo := range m.repos {
			desc := repo.Description
			if desc == "" {
				desc = "(No description)"
			}

			s += fmt.Sprintf("%d. %s\n", i+1, repo.Name)
			s += fmt.Sprintf("   %s\n", desc)
			s += fmt.Sprintf("   %s\n\n", repo.URL)
		}
	}

	s += "\nPress b to go back, q to quit.\n"
	return s
}
