package manager

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/service"
	ui "github.com/brotigen23/goph-keeper/client/internal/tui"
	"github.com/brotigen23/goph-keeper/client/internal/tui/style"
	datacontroller "github.com/brotigen23/goph-keeper/client/internal/tui/view/data_controller"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/form"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/tab"
	tea "github.com/charmbracelet/bubbletea"
)

type Manager struct {
	// API
	client *api.RESTClient

	accountsService *service.Accounts
	user            string
	tabs            tab.Tab

	logger *logger.Logger
}

func New(logger *logger.Logger, accountsService *service.Accounts, user string) tea.Model {
	accountTab := datacontroller.New[domain.Account](logger)
	fileTab := datacontroller.New[domain.BinaryData](logger)

	ret := Manager{
		tabs: *tab.New(
			[]string{"Accounts", "Files"},
			[]tea.Model{accountTab, fileTab}),
		user:            user,
		accountsService: accountsService,

		logger: logger,
	}
	return ret
}

func (m Manager) Init() tea.Cmd {
	return m.tabs.Init()
}

func (m Manager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	m.logger.Info("message", "msg", msg)
	switch msg.(type) {
	case datacontroller.RequestDataMsg[domain.Account]:
		accounts, err := m.accountsService.Fetch()
		switch err {
		case nil:
			msg := ui.FetchSuccessMsg[domain.Account]{Data: accounts}
			cmds = append(cmds, func() tea.Msg { return msg })
		default:
			m.logger.Error(err)
		}
	case form.SubmitFormMsg[domain.Account]:
	}
	tabs, cmd := m.tabs.Update(msg)
	cmds = append(cmds, cmd)
	m.tabs = tabs.(tab.Tab)
	return m, tea.Batch(cmds...)
}

func (m Manager) View() string {

	var frame string

	user := m.user
	frame += "User: " + user

	frame += style.Gap

	frame += m.tabs.View()

	frame += style.Gap

	return frame
}
