package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	choices     []string         // items on the to-do list
	cursor      int              // which item our cursor is pointing at
	selected    map[int]struct{} // which items are selected
	results     string           // the results of "sending"
	showResults bool             // whether to show results or the selection UI
}

func initialModel() model {
	return model{
		choices:     []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected:    make(map[int]struct{}),
		results:     "",
		showResults: false,
	}
}
func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
