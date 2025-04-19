package main

import (
	"fmt"
	"os"

	"nuke/internal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(internal.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
