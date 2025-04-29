package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/form"

	tea "github.com/charmbracelet/bubbletea"
)

type BaseExample struct {
	ID         int
	TimeColumn time.Time
}

type Example struct {
	BaseExample
	Some         string `form:"true,50, 30"`
	StringColumn string `form:"true,50,30"`
}

type Root struct {
	Child tea.Model
}

func NewRoot(child tea.Model) *Root {
	return &Root{Child: child}
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
	case form.SubmitFormMsg[Example]:
		fmt.Printf("Form data:\n%v", msg.Data)
	case form.SubmitFormErrorMsg:
		fmt.Printf("Form error field:%s\nError\n%v:\n", msg.Field, msg.Error)
	}

	child, cmd := m.Child.Update(msg)
	m.Child = child
	return m, cmd
}

func (m Root) View() string {
	return m.Child.View() + "\n\n[q or ctrl+c to quit]"
}

func main() {
	form := form.NewWithData(Example{BaseExample: BaseExample{ID: 1213, TimeColumn: time.Now()}, StringColumn: "hidden lol", Some: "Some"})
	root := NewRoot(form)

	p := tea.NewProgram(root, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}

}
