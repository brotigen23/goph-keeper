package auth

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/tui/style"
	"github.com/brotigen23/goph-keeper/client/internal/tui/util"
	"github.com/brotigen23/goph-keeper/client/internal/tui/view/auth/login"
	"github.com/brotigen23/goph-keeper/client/internal/tui/view/auth/register"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/tab"
	tea "github.com/charmbracelet/bubbletea"
)

// Message to change page from root model
type LoginSuccessMgs struct{}

// Message to check server connection
type PingServerErr error

const logoStr = "goph-keeper"

type AuthManager struct {
	tabs tab.Tab
}

func NewManager(logger *logger.Logger, client *api.RESTClient) tea.Model {
	log := login.New(logger, client)
	reg := register.New(logger, client)

	ret := AuthManager{
		tabs: *tab.New(
			[]string{"Login", "Register"},
			[]tea.Model{log, reg}),
	}
	return ret
}

func (m AuthManager) Init() tea.Cmd {
	return m.tabs.Init()
}

func (m AuthManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	tabs, cmd := m.tabs.Update(msg)
	cmds = append(cmds, cmd)

	m.tabs = tabs.(tab.Tab)
	return m, tea.Batch(cmds...)
}

func (m AuthManager) View() string {

	var frame string

	logo := util.RenderASCII(logoStr)
	frame += logo

	frame += m.tabs.View()

	frame += style.Gap

	return frame
}
