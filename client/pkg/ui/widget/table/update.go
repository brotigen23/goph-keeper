package table

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Return representation of table
func (m *Model[T]) Update(msg tea.Msg) (*Model[T], tea.Cmd) {
	// TODO: delete if no need
	switch tableMsg := msg.(type) {
	case tea.KeyMsg:
		switch tableMsg.String() {
		}
	}
	t, cmd := m.table.Update(msg)
	m.table = t
	return m, cmd
}
