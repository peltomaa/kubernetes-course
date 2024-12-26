package main

import (
	"crud-app/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		img, err := utils.GetImage()
		if err != nil {
			println("Failed to get image:", err)
			io.WriteString(w, "Unknown error")
			return
		}
		http.ServeFile(w, r, img)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server started in port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
