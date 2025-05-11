package login

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api/api"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case PingServerErr:
		m.serverStatus = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.changeInput()
		case "enter":
			if m.inputFocus == 0 {
				m.changeInput()
			} else {
				cmds = append(cmds, m.SignIn)
			}
		}
	case api.Response:

	}
	cmds = append(cmds, m.updateInputs(msg))
	return m, tea.Batch(cmds...)
}

func (m *model) changeInput() {
	m.inputFocus++
	if m.inputFocus > 1 {
		m.inputFocus = 0
	}
	for i := range m.inputs {
		if i == m.inputFocus {
			m.inputs[m.inputFocus].Focus()
			continue
		}
		m.inputs[i].Blur()
	}
}
