package login

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api/api"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	inputs       []textinput.Model
	inputFocus   int
	serverStatus error
	client       *api.RESTClient

	logger *logger.Logger
}

func New(logger *logger.Logger, client *api.RESTClient) tea.Model {
	inputs := make([]textinput.Model, 2)

	loginInput := textinput.New()
	loginInput.Focus()
	loginInput.Prompt = ""
	loginInput.Width = 50
	loginInput.CharLimit = 60
	loginInput.TextStyle = lipgloss.NewStyle().Align(lipgloss.Center)
	inputs[0] = loginInput

	passwordInput := textinput.New()
	passwordInput.Prompt = ""

	passwordInput.Width = 50
	passwordInput.CharLimit = 60
	passwordInput.EchoMode = textinput.EchoPassword
	inputs[1] = passwordInput

	ret := model{
		inputs: inputs,
		logger: logger,
		client: client,
	}
	return ret
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		resp := m.client.Ping()
		return PingServerErr(resp.Err)
	}
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
