package page

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type contentModel struct {
	spinner  spinner.Model
	quitting bool
	err      error

	logger *log.Logger
}

func NewTContentModel(logger *log.Logger) contentModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return contentModel{
		spinner: s,
		logger:  logger,
	}
}

func (m contentModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m contentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.logger.Debug("TEMPLATE UPDATE", "msg", reflect.TypeOf(msg))
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

func (m contentModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading forever... \n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}
