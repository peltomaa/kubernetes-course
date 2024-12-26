package main

import (
	"fmt"
	"io"
	"log-output/utils"
	"net/http"
	"os"
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
			fmt.Println("Error reading value:", err)
			io.WriteString(w, "Unknown error")
			return
		}

		pongText, err := utils.ReadFile(pongFilePath)
		if err != nil {
			fmt.Println("Error reading pong value:", err)
			io.WriteString(w, "Unknown error")
			return
		}

		data := fmt.Sprintf("%s\nPing / Pongs: %s\n", text, pongText)
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
