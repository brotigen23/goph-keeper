package root

import (
	"github.com/brotigen23/goph-keeper/client/internal/client"
	"github.com/brotigen23/goph-keeper/client/internal/ui/login"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	tea "github.com/charmbracelet/bubbletea"
)

const minWidth = 105
const minHeigth = 25

const (
	loginPage = iota
	loginContent
)

type model struct {
	currentPage int

	login   tea.Model
	content tea.Model

	logger *logger.Logger
	client *client.Client

	windowSize struct {
		width  int
		height int
	}
}

func New(logger *logger.Logger, client *client.Client) *model {
	return &model{
		currentPage: loginPage,
		login:       login.New(logger, client),

		logger: logger,

		client: client,
	}
}

func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		m.login.Init(),
	}
	return tea.Batch(cmds...)
}
