package table

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func FromExample() table.Styles {
	return table.Styles{
		Header: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false),
		Cell: lipgloss.NewStyle(),
		Selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false),
	}
}

func DefaultStyle() table.Styles {
	return table.Styles{
		Header:   lipgloss.NewStyle().Bold(true).Border(lipgloss.RoundedBorder()),
		Cell:     lipgloss.NewStyle().Border(lipgloss.RoundedBorder()),
		Selected: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")),
	}
}
