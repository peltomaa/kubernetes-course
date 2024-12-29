package utils

import (
	"path/filepath"
)

func GetConfigFilePath() string {
	return filepath.Join("/config", "information.txt")
}
