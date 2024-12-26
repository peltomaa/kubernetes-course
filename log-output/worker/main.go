package main

import (
	"fmt"
	"log-output/utils"
	"time"

	"github.com/google/uuid"
)

func GetStringOutput(id uuid.UUID) string {
	now := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("%s: %s", now, id)
}

func main() {
	id := uuid.New()
	filePath, err := utils.GetFilePath()
	if err != nil {
		fmt.Println("Error getting file path:", err)
		return
	}
	for {
		text := GetStringOutput(id)
		fmt.Println("write to file", text)
		utils.WriteToFile(filePath, text)
		time.Sleep(time.Second * 5)
	}
}
