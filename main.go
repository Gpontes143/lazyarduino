package main

import (
	"fmt"
	"log"

	"lazyarduino/pkg/commands"
)

func board() string {
	fmt.Println("Buscando placas conectadas...")

	boards, err := commands.ListBoards()
	if err != nil {
		log.Fatalf("Erro ao listar placas: %v", err)
	}

	if len(boards) == 0 {
		fmt.Println("Nenhuma placa encontrada.")
		return ""
	}
	b := boards[0]
	fqbn := ""
	nome := "desconhecido"

	if len(b.MatchingBoards) > 0 {
		nome = b.MatchingBoards[0].Name
		fqbn = b.MatchingBoards[0].FQBN
	}
	fmt.Printf("- Porta: %s | Placa: %s\n", b.Port.Address, nome)
	return fqbn
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

func main() {
	fqbn := board()
	path := "."
	if fqbn == "" {
		fmt.Println("Não foi possivel identificar o fqbn automaticamente!")
	}
	compile(fqbn, path)
}
