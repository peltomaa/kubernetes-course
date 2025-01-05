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
	pongClickRepo := repositories.PongClickRepository{}
	db, err := db.InitDb()
	if err != nil {
		fmt.Println("Error setting up db:", err)
	}

	count, err := pongClickRepo.Count(db)
	if err != nil {
		fmt.Println("Error reading number:", err)
	}

	fmt.Println("Starting server")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world, this is pong!")
	})

	mux.HandleFunc("GET /pingpong", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("pong %d", count))

		err := pongClickRepo.Create(db)
		count = count + 1
		if err != nil {
			fmt.Println("Error increasing count:", err)
		}
	})

	http.Handle("/", mux)

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
