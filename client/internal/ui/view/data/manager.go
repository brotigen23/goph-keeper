package data

import (
	"github.com/brotigen23/goph-keeper/client/internal/api"
	"github.com/brotigen23/goph-keeper/client/internal/ui/style"
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/data/account"
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/data/text"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/tab"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	user string
	tabs tab.Tab
}

func NewManager(logger *logger.Logger, client *api.Client, user string) tea.Model {
	a := account.New(logger, client)
	c := text.New(logger, client)

	ret := model{
		tabs: *tab.New(
			[]string{"Accounts", "Text"},
			[]tea.Model{a, c}),
		user: user,
	}
	return ret
}

func (m model) Init() tea.Cmd {
	return m.tabs.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	tabs, cmd := m.tabs.Update(msg)
	cmds = append(cmds, cmd)

	m.tabs = tabs.(tab.Tab)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	var frame string

	user := m.user
	frame += "User: " + user

	frame += style.Gap

	frame += m.tabs.View()

	frame += style.Gap

	return frame
}
