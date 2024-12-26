package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Todo struct {
	Task string `json:"task"`
}

var todos []Todo

func handleGetTodos(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Failed to serialize todos", http.StatusInternalServerError)
	}
}

func handlePostTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Failed to decode new todo", http.StatusInternalServerError)
	}

	todos = append(todos, newTodo)

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Failed to serialize todos", http.StatusInternalServerError)
	}
}

func main() {
	if todos == nil {
		todos = []Todo{}
	}

	fmt.Println("Starting server")
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleGetTodos(w)
			return
		}
		if r.Method == http.MethodPost {
			handlePostTodos(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
