package utils

import (
	"os"
	"path/filepath"
)

func GetFilePath() (string, error) {
	dir := os.Getenv("FILE_DIR")
	if dir == "" {
		dir = "./"
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "pong.txt"), nil
}
