package commands

import (
	"encoding/json"
	"os/exec"
)

type BoardListResponse struct {
	DetectedPorts []BoardInfo `json:"detected_ports"`
}

type BoardInfo struct {
	MatchingBoards []MatchingBoard `json:"matching_boards"`
	Port           Port            `json:"port"`
}

type MatchingBoard struct {
	Name string `json:"name"`
	FQBN string `json:"fqbn"`
}

type Port struct {
	Address string `json:"address"`
}

// ListBoards Retorna o json formatado com a Porta e a Placa
func ListBoards() ([]BoardInfo, error) {
	cmd := exec.Command("arduino-cli", "board", "list", "--format", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Criamos a variável do tipo "raiz" (objeto)
	var resp BoardListResponse

	// Fazemos o unmarshal no objeto raiz
	if err := json.Unmarshal(output, &resp); err != nil {
		return nil, err
	}

	return resp.DetectedPorts, nil
}
