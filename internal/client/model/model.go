package model

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const minWidth = 105
const minHeigth = 25

type model struct {
	inputs     []textinput.Model
	inputFocus int

	width  int
	height int
}

func initModel() tea.Model {
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
	}
	return ret
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "ctrl+q":
			return m, tea.Quit
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

var gap = "\n\n"

func (m model) View() string {
	if m.width < 105 || m.height < 25 {
		warning := cRed.Render("not enough space")
		warning += gap

		warning += cGreen.Render(strconv.Itoa(minWidth) + " x " + strconv.Itoa(minHeigth))
		warning += "\n"
		if m.width < minWidth {
			warning += cRed.Render(strconv.Itoa(m.width))
		} else {
			warning += cGreen.Render(strconv.Itoa(m.width))
		}
		warning += " x "
		if m.height < minHeigth {
			warning += cRed.Render(strconv.Itoa(m.height))
		} else {
			warning += cGreen.Render(strconv.Itoa(m.height))
		}
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, warning)
	}
	var frame string

	logo := renderLogo()
	frame += logo

	frame += gap

	login := m.inputs[0].View()
	login = Bordered(login, m.inputFocus == 0)
	frame += "login:"
	frame += "\n"
	frame += login
	frame += "\n"

	password := m.inputs[1].View()
	password = Bordered(password, m.inputFocus == 1)
	frame += "password:"
	frame += "\n"
	frame += password
	frame = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, frame)
	return frame
}

func Run() {
	if _, err := tea.NewProgram(initModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
