package page

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type errMsg error

type templateModel struct {
	spinner  spinner.Model
	quitting bool
	err      error

	logger log.Logger
}

func NewTemplateModel(logger *log.Logger) templateModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return templateModel{
		spinner: s,
		logger:  *logger,
	}
}

func (m templateModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m templateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m templateModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading forever... \n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}
