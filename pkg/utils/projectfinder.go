package utils

import (
	"os"
	"path/filepath"
)

func GetProjectName() string {
	files, err := os.ReadDir(".")
	if err != nil {
		return "Erro ao ler diretório"
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".ino" {
			return file.Name()
		}
	}

	return "Nenhum arquivo .ino encontrado"
}
