package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			m.table.CursorDown()
		case "k":
			m.table.CursorUp()
		case "enter":

		case "c":
			// Create data
		}

	case FetchSuccessMsg:
		if msg.Accounts != nil {
			m.data = msg.Accounts
			m.table.Refresh(m.data)
		}
		m.isLoading = false
	case ErrMsg:
		m.logger.Error(msg.Err)
	}
	return m, tea.Batch(cmds...)
}
