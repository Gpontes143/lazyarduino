package themes

import "github.com/charmbracelet/lipgloss"

var (
	// cores

	Red   = lipgloss.Color("#ab1616")
	Black = lipgloss.Color("#121010")
	White = lipgloss.Color("#F2F3F4")

	// estilo da caixa
	BaseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(White).
			Padding(0, 1)

	TitleStyle = lipgloss.NewStyle().
			Foreground(White).Bold(true)
)
