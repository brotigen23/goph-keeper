package root

import (
	"github.com/brotigen23/goph-keeper/client/internal/ui/login"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
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

	case login.LoginSuccessMgs:
		m.currentPage = contentPage
		cmd := m.content.Init()
		cmds = append(cmds, cmd)
	}

	// Pages
	switch m.currentPage {
	case loginPage:
		model, cmd := m.login.Update(msg)
		m.login = model
		cmds = append(cmds, cmd)
	case contentPage:
		model, cmd := m.content.Update(msg)
		m.login = model
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
