package text

import (
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/domain"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) FetchAccountData() tea.Msg {
	response := m.client.GetAccounts()
	if err := response.Err; err != nil {
		m.logger.Error(err)
		return nil
	}
	m.logger.Info("server response", "status code", response.StatusCode, "body", response.Body)
	switch response.StatusCode {
	case http.StatusOK:
		var accounts []domain.AccountData
		err := json.Unmarshal([]byte(response.Body), &accounts)
		if err != nil {
			m.logger.Error(err)
			return nil
		}
		return FetchSuccessMsg{
			Accounts: accounts,
		}
	case http.StatusNoContent:
		return FetchSuccessMsg{
			Accounts: nil,
		}
	}

	return nil
}
