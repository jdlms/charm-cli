package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			// Add a new case for sending selected items
		case "s":
			// Create a slice to hold selected items
			selectedItems := []string{}

			// Collect all selected items
			for i := range m.choices {
				if _, ok := m.selected[i]; ok {
					selectedItems = append(selectedItems, m.choices[i])
				}
			}

			// Join the selected items into a string
			m.results = "You selected: " + strings.Join(selectedItems, ", ")
			m.showResults = true

			return m, nil

		// Add a case to go back to the list view
		case "b":
			if m.showResults {
				m.showResults = false
				return m, nil
			}
		}
	}
	return m, nil
}
