package page

import (
	"github.com/brotigen23/goph-keeper/client/internal/client"
	"github.com/brotigen23/goph-keeper/client/internal/message"
	"github.com/brotigen23/goph-keeper/client/internal/style"
	"github.com/brotigen23/goph-keeper/client/internal/util"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const logoStr = "goph-keeper"

type loginModel struct {
	inputs       []textinput.Model
	inputFocus   int
	serverStatus error
	client       *client.Client

	logger *log.Logger
}

func NewLogin(logger *log.Logger, client *client.Client) tea.Model {
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

	ret := loginModel{
		inputs: inputs,
		logger: logger,
		client: client,
	}
	return ret
}

func (m loginModel) Init() tea.Cmd {
	return func() tea.Msg {
		err := m.client.Ping()
		return message.PingServerErr(err)
	}
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case message.PingServerErr:
		m.serverStatus = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.inputFocus++
			if m.inputFocus > 1 {
				m.inputFocus = 0
			}
			for i := range m.inputs {
				if i == m.inputFocus {
					m.inputs[m.inputFocus].Focus()
					continue
				}
				m.inputs[i].Blur()
			}
		}
	}
	return m, m.updateInputs(msg)
}

func (m loginModel) View() string {

	var frame string

	logo := util.RenderASCII(logoStr)
	frame += logo

	frame += style.Gap

	login := m.inputs[0].View()
	login = style.Bordered(login, m.inputFocus == 0)
	frame += "login:"
	frame += "\n"
	frame += login
	frame += "\n"

	password := m.inputs[1].View()
	password = style.Bordered(password, m.inputFocus == 1)
	frame += "password:"
	frame += "\n"
	frame += password

	frame += style.Gap
	frame += lipgloss.NewStyle().Align(lipgloss.Right, lipgloss.Bottom).Render("Server status:")
	frame += "\n"
	serverStatus := ""
	if m.serverStatus != nil {
		serverStatus += style.ColorRed.Render("Server is not response with error: ")
		serverStatus += style.ColorRed.Render(m.serverStatus.Error())
	} else {
		serverStatus += style.ColorGreen.Render("OK")
	}
	serverStatus = lipgloss.NewStyle().Align(lipgloss.Right, lipgloss.Bottom).Render(serverStatus)
	frame += serverStatus
	return frame
}

func (m *loginModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
