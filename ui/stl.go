package ui

import "github.com/charmbracelet/lipgloss"

var (
	Primary   = lipgloss.Color("#C97B68")
	Secondary = lipgloss.Color("#7A8281")
	Accent    = lipgloss.Color("#625954")
	Ornament  = lipgloss.Color("#DBC7BA")
	Yuk       = lipgloss.Color("#B13A2A")

	FooterStyle = lipgloss.NewStyle().
			Foreground(Accent).
			Background(Secondary).
			Italic(true).
			MarginTop(1)
)
