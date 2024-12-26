package utils

import (
	"os"
	"path/filepath"
)

func GetPongFilePath() (string, error) {
	dir := os.Getenv("FILE_DIR")
	if dir == "" {
		dir = "../pong/"
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "pong.txt"), nil
}
