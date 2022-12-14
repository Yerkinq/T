package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	err := tea.NewProgram(&Model{}, tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Model struct {
	count int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			m.count++
		case "left":
			m.count--
		}
	}
	return m, nil
}

func (m *Model) View() string {
	return fmt.Sprintf("count: %d\n\nā Decrease   Increase ā\n\nPress 'ctrl+c' to exit", m.count)
}
