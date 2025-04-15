package content

import (
	"fmt"

	"github.com/brotigen23/goph-keeper/client/internal/api"
	"github.com/brotigen23/goph-keeper/client/internal/domain"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	tea "github.com/charmbracelet/bubbletea"
)

func StringAccounts(a []domain.AccountData) string {
	var ret string
	for _, v := range a {
		ret += fmt.Sprintf(`
		id: %v, 
		user id: %v,
		metadata id: %v,
		login: %s,
		password: %s,
		created at: %s,
		updated at: %s \n`,
			v.ID,
			v.UserID,
			v.MetadataID,
			v.Login,
			v.Password,
			v.CreatedAt,
			v.UpdatedAt)
	}
	return ret
}

type model struct {
	client *api.Client
	table  []domain.AccountData

	logger *logger.Logger

	isLoading bool
}

func New(logger *logger.Logger, client *api.Client) tea.Model {

	ret := model{
		logger:    logger,
		client:    client,
		isLoading: true,
	}
	return ret
}

func (m model) Init() tea.Cmd {
	return m.FetchAccountData
}
