package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"pong/db"
	"pong/repositories"
)

func main() {
	pongClickRepo := repositories.PongButtonClick{}
	db, err := db.InitDb()
	if err != nil {
		fmt.Println("Error setting up db:", err)
	}

	count, err := pongClickRepo.Count(db)
	if err != nil {
		fmt.Println("Error reading number:", err)
	}

	fmt.Println("Starting server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		io.WriteString(w, fmt.Sprintf("pong %d", count))

		err := pongClickRepo.Create(db)
		count = count + 1
		if err != nil {
			fmt.Println("Error increasing count:", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server started in port", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
