package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := tea.NewProgram(
		model("A long time ago in a galaxy far, far away"),
		tea.WithAltScreen(),
	)
	if _, err := app.Run(); err != nil {
		log.Fatalf("failed to run app:\n%v", err)
	}
}

type model string

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return string(m)
}
