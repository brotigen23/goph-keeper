package form

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Prompt      lipgloss.Style
	PromptFocus lipgloss.Style
	Input       lipgloss.Style
	InputFocus  lipgloss.Style
}

func DefaultStyle() Style {
	return Style{
		Prompt:      lipgloss.NewStyle(),
		PromptFocus: lipgloss.NewStyle(),
		Input:       lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#080808")),
		InputFocus:  lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFA500")),
	}
}
