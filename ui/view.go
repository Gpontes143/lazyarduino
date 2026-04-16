package ui

import (
	"lazyarduino/ui/themes"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	// Prevenção para quando o terminal ainda não reportou o tamanho
	if m.Width == 0 {
		return "A carregar interface..."
	}

	// 1. Definição das Proporções
	// Coluna esquerda fixa em 35 caracteres, direita ocupa o resto
	leftWidth := 35
	rightWidth := m.Width - leftWidth - 4 // -4 compensa bordas e espaçamentos

	// 2. CONSTRUÇÃO DA COLUNA ESQUERDA
	// Painel [1] Status
	statusStyle := themes.GetPanelStyle(m.Focused == 1)
	projectLine := themes.TitleStyle.Render(themes.IconProject + " " + m.ProjectName)
	statusContent := m.Spinner.View() + " " + themes.IconStatus + " " + m.StatusMsg
	statusBody := projectLine + "\n" + statusContent
	statusBox := statusStyle.Width(leftWidth).Render(
		themes.TitleStyle.Render("[1] Status") + "\n\n" + statusBody,
	)

	// Painel [2] Placas
	boardStyle := themes.GetPanelStyle(m.Focused == 2)
	// Ajustamos o tamanho interno da lista para caber no box
	m.List.SetSize(leftWidth-2, 10)
	boardBox := boardStyle.Width(leftWidth).Render(m.List.View())

	// Painel [3] Recursos (Apenas placeholder por enquanto)
	recursosStyle := themes.GetPanelStyle(m.Focused == 3)
	recursosBox := recursosStyle.Width(leftWidth).Render(
		themes.TitleStyle.Render("[3] Recursos") + "\n" + "Flash: --\nRAM:   --",
	)

	// Junta a coluna esquerda na vertical
	leftCol := lipgloss.JoinVertical(lipgloss.Left, statusBox, boardBox, recursosBox)

	// 3. CONSTRUÇÃO DA COLUNA DIREITA

	// Painel [0] Serial Monitor
	serialStyle := themes.GetPanelStyle(m.Focused == 0)
	// Altura dinâmica baseada na janela, subtraindo o espaço do log
	serialHeight := m.Height - 12
	m.Serial.Width = rightWidth - 2
	m.Serial.Height = serialHeight

	serialBox := serialStyle.Width(rightWidth).Render(
		themes.TitleStyle.Render("[0] Serial Monitor") + "\n" + m.Serial.View(),
	)

	// Painel de Log (Fixo na parte de baixo da direita)
	logBox := themes.BaseStyle.Width(rightWidth).Height(3).Render(
		"Command Log: Pronto",
	)

	rightCol := lipgloss.JoinVertical(lipgloss.Left, serialBox, logBox)

	// 4. JUNÇÃO FINAL
	return lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)
}
