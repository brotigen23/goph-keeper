package login

import (
	tea "github.com/charmbracelet/bubbletea"
)

func PingServer(ip string) tea.Msg {

	return nil
}

func (m model) SignIn() tea.Msg {

	login := m.inputs[0].Value()
	password := m.inputs[1].Value()
	m.logger.Info("sign in", "login", login, "password", password)
	response := m.client.Login(login, password)
	// If some error
	if err := response.Err; err != nil {
		m.logger.Error(err)
	}

	if response.StatusCode == 202 {
		return LoginSuccessMsg{
			Username: login,
		}
	}

	return response
}
