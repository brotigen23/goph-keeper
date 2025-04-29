package main

import (
	"os"
	"time"

	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/table"

	tea "github.com/charmbracelet/bubbletea"
)

type example struct {
	ID           int       `table:"true"`
	StringColumn string    `table:"true"`
	TimeColumn   time.Time `table:"true"`
}

type Root struct {
	data  []example
	table table.Model[example]
}

func NewRoot() *Root {
	data := []example{
		{ID: 1, StringColumn: "yes", TimeColumn: time.Now()},
		{ID: 1, StringColumn: "yes", TimeColumn: time.Now()},
	}
	table := table.New[example]()
	table.Refresh(data)
	return &Root{
		data:  data,
		table: *table,
	}
}

func (m Root) Init() tea.Cmd {
	return nil
}

func (m *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "c":
			m.data = append(m.data, *m.table.GetCurrentItem())
			m.table.Refresh(m.data)
		}
	}

	c, cmd := m.table.Update(msg)
	cmds = append(cmds, cmd)
	m.table = c.(table.Table[example])
	return m, tea.Batch(cmds...)
}

func (m Root) View() string {
	return m.table.View() + "\n\n[q or ctrl+c to quit]"
}

func main() {
	root := NewRoot()

	p := tea.NewProgram(root, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}

}
