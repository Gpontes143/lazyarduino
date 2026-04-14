package main

import (
	"encoding/json"
	"os/exec"
)

type Port struct {
	Adress   string `json:"adress"`
	Protocol string `json:"protocol"`
}

type MathchingBoard struct {
	Name string `json:"name"`
	FQBN string `json:"fqbn"`
}

type BoardInfo struct {
	MathchingBoard []MathchingBoard `json:"mathchingBoard"`
	Port           []Port           `json:"port"`
}

// ListBoards executa 'arduino-cli board list --format json'
func ListBoards() ([]BoardInfo, error) {
	// Preparamos o comando externo
	cmd := exec.Command("arduino-cli", "board", "list", "--format", "json")

	// Capturamos a saída (stdout)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Fazemos o Parse do JSON para a nossa slice de structs
	var boards []BoardInfo
	if err := json.Unmarshal(output, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}
