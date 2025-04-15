package content

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	case FetchSuccessMsg:
		m.isLoading = false
	}
	return m, tea.Batch(cmds...)
}
