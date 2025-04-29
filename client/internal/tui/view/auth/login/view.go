package login

import (
	"github.com/brotigen23/goph-keeper/client/internal/tui/style"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	var frame string

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
