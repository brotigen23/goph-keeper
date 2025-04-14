package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch globalMsg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize.height = globalMsg.Height
		m.windowSize.width = globalMsg.Width
	// quit
	case tea.KeyMsg:
		switch globalMsg.String() {
		case "ctrl+c", "esc", "ctrl+q":
			return m, tea.Quit
		}
	}
	switch m.currentPage {
	case loginPage:
		return m.login.Update(msg)
	}
	return m, nil
}
