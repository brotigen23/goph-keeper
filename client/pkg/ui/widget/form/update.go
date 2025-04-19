package form

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Form[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.changeInput()
		case "enter":
			if m.focus < len(m.inputs)-1 {
				m.changeInput()
			} else {
				cmds = append(cmds, m.EditConfirm)
			}
		}
	}
	cmds = append(cmds, m.updateInputs(msg))
	return m, tea.Batch(cmds...)
}
