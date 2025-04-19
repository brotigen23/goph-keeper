package app

import (
	"github.com/brotigen23/goph-keeper/client/internal/api"
	"github.com/brotigen23/goph-keeper/client/internal/ui/view/auth"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	tea "github.com/charmbracelet/bubbletea"
)

const minWidth = 105
const minHeigth = 25

const (
	loginPage = iota
	contentPage
)

type model struct {
	currentPage int

	authPage tea.Model
	dataPage tea.Model

	logger *logger.Logger
	client *api.Client

	windowSize struct {
		width  int
		height int
	}
}

func New(logger *logger.Logger, client *api.Client) *model {
	return &model{
		currentPage: loginPage,

		authPage: auth.NewManager(logger, client),

		logger: logger,

		client: client,
	}
}

func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		m.authPage.Init(),
	}
	return tea.Batch(cmds...)
}
