package register

import (
	tea "github.com/charmbracelet/bubbletea"
)

func PingServer(ip string) tea.Msg {

	return nil
}

func (m model) SignUp() tea.Msg {
	login := m.inputs[0].Value()
	password := m.inputs[1].Value()
	m.logger.Info("sign in", "login", login, "password", password)
	response := m.client.Register(login, password)
	// If some error
	if err := response; err != nil {
		m.logger.Error(err)
	}

	return response
}
