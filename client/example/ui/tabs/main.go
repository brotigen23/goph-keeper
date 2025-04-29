package main

import (
	"os"
	"time"

	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/tab"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/table"
	tea "github.com/charmbracelet/bubbletea"
)

type example struct {
	ID           int       `table:"true"`
	StringColumn string    `table:"true"`
	TimeColumn   time.Time `table:"true"`
}

type Root struct {
	Child tea.Model
}

func NewRoot(child tea.Model) *Root {
	return &Root{
		Child: child}
}

func (m Root) Init() tea.Cmd {
	return m.Child.Init()
}

func (m Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	child, cmd := m.Child.Update(msg)
	m.Child = child
	return m, cmd

}

func (m Root) View() string {
	return m.Child.View() + "\n\n[q or ctrl+c to quit]"
}

func main() {
	table1 := table.New[example]()
	table1.Refresh([]example{{StringColumn: "Lol"}, {StringColumn: "Kek"}})
	table2 := table.New[example]()
	table2.Refresh([]example{{StringColumn: "AAAA"}, {StringColumn: "BBBB"}})
	tabs := []table.Model[example]{table1, table2}
	tab := tab.New([]string{"First", "Next"}, tabs)
	root := NewRoot(tab)
	p := tea.NewProgram(root, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}

}
