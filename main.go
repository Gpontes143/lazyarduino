package main

import (
	"fmt"
	"log"

	"lazyarduino/pkg/commands"
)

func main() {
	fmt.Println("Buscando placas conectadas...")

	boards, err := commands.ListBoards()
	if err != nil {
		log.Fatalf("Erro ao listar placas: %v", err)
	}

	if len(boards) == 0 {
		fmt.Println("Nenhuma placa encontrada.")
		return
	}

	for _, b := range boards {
		nome := "Desconhecido"
		if len(b.MatchingBoards) > 0 {
			nome = b.MatchingBoards[0].Name
		}
		fmt.Printf("- Porta: %s | Placa: %s\n", b.Port, nome)
	}
}
