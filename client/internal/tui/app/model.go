package app

/*
 * Dont work as it should
 */

import (
	"os"
	"runtime"

	"github.com/brotigen23/goph-keeper/client/internal/core/api/rest"
	"github.com/brotigen23/goph-keeper/client/internal/tui/view/auth"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {

	logger := logger.New().Default()

	logger.Info(
		"OS INFO:",
		"OS:", runtime.GOOS,
		"Arch:", runtime.GOARCH,
		"TERM:", os.Getenv("TERM"))
	client := rest.New("http://localhost:8080")
	// Main
	rootModel := New(logger, client)

	p := tea.NewProgram(rootModel, tea.WithAltScreen())

	_, err := p.Run()

	return err
}

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
	client *rest.Client

	windowSize struct {
		width  int
		height int
	}
}

func New(logger *logger.Logger, client *rest.Client) *model {
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
