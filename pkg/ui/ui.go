package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rsteube/vincent"
)

type model struct {
	themes []string
	cursor int
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	return model{
		themes: vincent.Themes(),
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "left", "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "right", "down", "j":
			if m.cursor < len(m.themes)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	t, err := vincent.Load(m.themes[m.cursor])
	if err != nil {
		panic(err.Error()) // TODO handle err
	}
	s := fmt.Sprintf("%v\n\n", t.Name)
	s += t.Render()

	s += "\nPress q to quit.\n"

	return s
}

func Execute() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}