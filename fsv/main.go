package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/datatug/firestore-viewer/fsviewer"
	"log"
)

func main() {
	app, err := fsviewer.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
