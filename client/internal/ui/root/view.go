package root

import (
	"github.com/brotigen23/goph-keeper/client/internal/ui/util"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	warning, ok := util.WarningResolutions(m.windowSize.width, m.windowSize.height, minWidth, minHeigth)
	if !ok {
		return warning
	}
	var frame string
	switch m.currentPage {
	case loginPage:
		frame = m.login.View()
	}

	frame = lipgloss.Place(m.windowSize.width, m.windowSize.height, lipgloss.Center, lipgloss.Center, frame)

	return frame
}
