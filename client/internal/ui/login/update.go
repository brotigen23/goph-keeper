package login

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case PingServerErr:
		m.serverStatus = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
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
	}
	return m, m.updateInputs(msg)
}
