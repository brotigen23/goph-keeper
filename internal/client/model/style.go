package model

import "github.com/charmbracelet/lipgloss"

// TODO:
/*
	[colors.primary]
	# hard contrast background = = '#1d2021'
	background = '#282828'
	# soft contrast background = = '#32302f'
	foreground = '#ebdbb2'

	# Normal colors
	[colors.normal]
	black   = '#282828'
	red     = '#cc241d'
	green   = '#98971a'
	yellow  = '#d79921'
	blue    = '#458588'
	magenta = '#b16286'
	cyan    = '#689d6a'
	white   = '#a89984'

	# Bright colors
	[colors.bright]
	black   = '#928374'
	red     = '#fb4934'
	green   = '#b8bb26'
	yellow  = '#fabd2f'
	blue    = '#83a598'
	magenta = '#d3869b'
	cyan    = '#8ec07c'
	white   = '#ebdbb2'
*/
var (
	cBackground = lipgloss.NewStyle().Foreground(lipgloss.Color("#1d2021"))

	cRed   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	cGreen = lipgloss.NewStyle().Foreground(lipgloss.Color("#008000"))
)

var (
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("100"))
)

func Centered(width, height int, str string) string {
	centered := lipgloss.NewStyle().Width(width).Height(height).Align(lipgloss.Center)

	return centered.Render(str)
}

func Bordered(str string, active bool) string {
	if active {
		border := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFA500"))
		return border.Render(str)
	}
	border := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#080808"))
	return border.Render(str)
}
