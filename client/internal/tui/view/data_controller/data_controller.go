package datacontroller

import (
	"github.com/brotigen23/goph-keeper/client/internal/app/domain"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/table"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	View = iota
	Create
	Edit
	Delete
)

type DataController[T domain.Model] struct {
	table *table.Model[T]

	form  tea.Model
	state int

	logger *logger.Logger
}

func New[T domain.Model](logger *logger.Logger) tea.Model {
	return DataController[T]{
		table: table.New[T](),
		form:  nil,
		state: View,

		logger: logger,
	}
}

func (m DataController[T]) Init() tea.Cmd {
	return SendMsg(RequestDataMsg[T]{})
}
