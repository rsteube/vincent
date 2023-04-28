package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rsteube/vincent"
)

type model struct {
	schemes []string
	cursor  int
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	return model{
		schemes: vincent.Schemes(),
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "left", "h":
			m.cursor -= 10
			if m.cursor < 0 {
				m.cursor = 0
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.schemes)-1 {
				m.cursor++
			}

		case "right", "l":
			m.cursor += 10
			if max := len(m.schemes) - 1; m.cursor > max {
				m.cursor = max
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	t, err := vincent.Load(m.schemes[m.cursor])
	if err != nil {
		return err.Error()
	}
	s := fmt.Sprintf("%v\n\n", t.Name)
	s += t.Render()

	s += "\nPress arrows to navigate, q to quit.\n"

	return s
}

func Execute() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
