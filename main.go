package main

import (
	"fmt"
	"log"
	"os"

	"lazyarduino/pkg/commands"
	"lazyarduino/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func board() (string, string) {
	fmt.Println("Buscando placas conectadas...")

	boards, err := commands.ListBoards()
	if err != nil {
		log.Fatalf("Erro ao listar placas: %v", err)
	}

	if len(boards) == 0 {
		fmt.Println("Nenhuma placa encontrada.")
		return "", ""
	}
	b := boards[0]
	fqbn := ""
	nome := "desconhecido"
	port := ""
	port = b.Port.Address
	if len(b.MatchingBoards) > 0 {
		nome = b.MatchingBoards[0].Name
		fqbn = b.MatchingBoards[0].FQBN
	}
	fmt.Printf("- Porta: %s | Placa: %s\n", port, nome)
	return fqbn, port
}

func compile(path string, fqbn string) {
	fmt.Printf("Compilando...")
	out, err := commands.Compile(fqbn, path)
	if err != nil {
		fmt.Printf("Erro de compilação: \n%s", out)
		return
	}
	fmt.Printf("Compilado com sucesso")
}

func upload(port string, fqbn string, path string) {
	fmt.Println("Upload....")
	out, err := commands.Upload(port, fqbn, path)
	if err != nil {
		fmt.Printf("Erro ao inicia upload: \n%s", out)
		return
	}
	fmt.Println("Processo terminado com sucesso")
}

func main() {
	tela := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := tela.Run(); err != nil {
		fmt.Printf("Erro ao iniciar lazyarduino: %v", err)
		os.Exit(1)
	}
	// fqbn, port := board()
	// path := "."
	// if fqbn == "" || port == "" {
	// 	fmt.Println("Não foi possivel identificar o fqbn automaticamente ou a placa!")
	// }
	//
	// compile(fqbn, path)
	// upload(port, fqbn, path)
}
