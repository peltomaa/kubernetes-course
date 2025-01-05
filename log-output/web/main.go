package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"log-output/utils"
)

func main() {
	fmt.Println("Starting server")

	filePath, err := utils.GetFilePath()
	if err != nil {
		fmt.Println("Error getting file path:", err)
		return
	}
	pongFilePath, err := utils.GetPongFilePath()
	if err != nil {
		fmt.Println("Error getting pong file path:", err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text, err := utils.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading log file:", err)
			text = "CANNOT_READ_LOG_FILE"
		}

		pongText, err := utils.ReadFile(pongFilePath)
		if err != nil {
			fmt.Println("Error reading pong file:", err)
			pongText = "CANNOT_READ_PONG_FILE"
		}

		configText, err := utils.ReadFile(utils.GetConfigFilePath())
		if err != nil {
			fmt.Println("Error reading config value:", err)
			configText = "CANNOT_READ_CONFIG_FILE"
		}

		message := os.Getenv("MESSAGE")

		data := ""
		data += fmt.Sprintf("file content: %s\n", configText)
		data += fmt.Sprintf("MESSAGE=%s\n", message)
		data += text + "\n"
		data += fmt.Sprintf("Ping / Pongs: %s\n", pongText)

		io.WriteString(w, data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server started in port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
