package utils

import "os"

func WriteToFile(filePath, text string) error {
	err := os.WriteFile(filePath, []byte(text), 0755)
	if err != nil {
		return err
	}

	return nil
}
