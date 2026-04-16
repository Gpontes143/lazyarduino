package main

import (
	"fmt"
	"os"

	"lazyarduino/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tela := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := tela.Run(); err != nil {
		fmt.Printf("Erro ao iniciar lazyarduino: %v", err)
		os.Exit(1)
	}
}
