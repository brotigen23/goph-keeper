package app

import (
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/auth/login"
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/auth/register"
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/data"
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

	case login.LoginSuccessMsg:
		m.dataPage = data.NewManager(m.logger, m.client, globalMsg.Username)
		m.currentPage = contentPage
		cmd := m.dataPage.Init()
		cmds = append(cmds, cmd)

	case register.SignUpSuccessMsg:
		m.dataPage = data.NewManager(m.logger, m.client, globalMsg.Username)
		m.currentPage = contentPage
		cmd := m.dataPage.Init()
		cmds = append(cmds, cmd)
	}

	// Pages
	switch m.currentPage {
	case loginPage:
		model, cmd := m.authPage.Update(msg)
		m.authPage = model
		cmds = append(cmds, cmd)
	case contentPage:
		model, cmd := m.dataPage.Update(msg)
		m.dataPage = model
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
