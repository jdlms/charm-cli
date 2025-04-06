package main

import "fmt"

func (m model) View() string {
	if m.showResults {
		return fmt.Sprintf(
			"Results:\n\n%s\n\nPress b to go back, q to quit.",
			m.results,
		)
	}

	// The regular list view code
	s := "What should we buy at the market?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit, s to send selected items.\n"
	return s
}
