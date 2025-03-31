package model

import (
	"reflect"

	"github.com/brotigen23/goph-keeper/client/internal/page"
	"github.com/brotigen23/goph-keeper/client/internal/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const minWidth = 105
const minHeigth = 25

type MainModel struct {
	pages       []tea.Model
	currentPage int

	logger *log.Logger

	width  int
	height int
}

func NewMainModel(logger *log.Logger) *MainModel {
	pages := []tea.Model{page.NewLogin(logger), page.NewTemplateModel(logger)}
	return &MainModel{
		pages:       pages,
		currentPage: 0,

		logger: logger,
	}
}

func (m MainModel) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	for i := range m.pages {
		cmds = append(cmds, m.pages[i].Init())
	}
	return tea.Batch(cmds...)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.logger.Debug("TEMPLATE UPDATE", "msg", reflect.TypeOf(msg))

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	// quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "ctrl+q":
			return m, tea.Quit
		case "f1":
			if m.currentPage != 0 {
				m.currentPage = 0
				cmd := m.pages[m.currentPage].Init()
				return m, cmd
			}
		case "f2":
			if m.currentPage != 1 {
				m.currentPage = 1
				cmd := m.pages[m.currentPage].Init()
				return m, cmd
			}
		}
	}

	model, cmd := m.pages[m.currentPage].Update(msg)
	m.pages[m.currentPage] = model
	return m, cmd
}

func (m MainModel) View() string {
	warning, ok := util.WarningResolutions(m.width, m.height, minWidth, minHeigth)
	if !ok {
		return warning
	}
	frame := m.pages[m.currentPage].View()
	frame = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, frame)

	return frame
}
