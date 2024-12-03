package main

import (
	"log"
	"yuksenon/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fragments, err := ui.LoadFragments(
		"frags/ldd.json",   // relative
		"~/frags/ldd.json", // absolute
	)
	if err != nil {
		log.Fatalf("Failed to load fragments: %s", err)
	}

	program := tea.NewProgram(ui.NewModel(fragments), tea.WithAltScreen())
	if err := program.Start(); err != nil {
		log.Fatalf("Failed to start program: %s", err)
	}
}
