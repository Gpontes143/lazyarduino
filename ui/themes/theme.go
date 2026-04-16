package themes

import "github.com/charmbracelet/lipgloss"

var (
	Red   = lipgloss.Color("#ab1616")
	Black = lipgloss.Color("#121010")
	White = lipgloss.Color("#F2F3F4")

	IconProject = ""
	IconStatus  = " "
	IconBoard   = " "
	IconSerial  = "󱇰"
	BaseStyle   = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(White).
			Padding(0, 1)

	TitleStyle = lipgloss.NewStyle().
			Foreground(White).Bold(true)
)

func GetPanelStyle(focused bool) lipgloss.Style {
	style := BaseStyle // Aqui você já está criando uma cópia automaticamente

	if focused {
		return style.BorderForeground(Red)
	}
	return style.BorderForeground(White)
}
