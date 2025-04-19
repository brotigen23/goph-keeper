package account

import (
	"github.com/brotigen23/goph-keeper/client/internal/api"
	"github.com/brotigen23/goph-keeper/client/internal/domain"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/form"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab int

type model struct {
	client *api.Client

	table *table.Table[domain.AccountData]
	data  []domain.AccountData

	buffer []domain.AccountData

	accountForm *form.Form[domain.AccountData]

	logger *logger.Logger

	isLoading bool
}

func New(logger *logger.Logger, client *api.Client) tea.Model {

	ret := model{

		table: table.New[domain.AccountData](),

		logger:    logger,
		client:    client,
		isLoading: true,
	}
	return ret
}

func (m model) Init() tea.Cmd {
	return m.FetchData
}
