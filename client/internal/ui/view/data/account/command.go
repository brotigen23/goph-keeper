package account

import (
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/domain"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) FetchData() tea.Msg {
	response := m.client.GetData("/user/accounts")
	if err := response.Err; err != nil {
		return ErrMsg{Err: err}
	}
	switch response.StatusCode {
	case http.StatusOK:
		var data []domain.AccountData
		err := json.Unmarshal([]byte(response.Body), &data)
		if err != nil {
			return ErrMsg{Err: err}
		}
		return FetchSuccessMsg{
			Accounts: data,
		}
	case http.StatusNoContent:
		return FetchSuccessMsg{
			Accounts: nil,
		}
	}
	return nil
}
