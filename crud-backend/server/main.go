package main

import (
	"fmt"
	"net/http"
	"os"

	"crud-backend/controllers"
	"crud-backend/db"
	"crud-backend/repositories"
)

func main() {
	fmt.Println("Starting backend server")

	DB, err := db.InitDb()
	if err != nil {
		fmt.Println("Failed init db:", err)
	}
	db.CreateTable(DB)
	taskRepo := repositories.TaskRepository{DB: DB}
	taskCtrl := controllers.TaskController{R: &taskRepo}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR"))
			return
		}

		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		taskCtrl.GetTasks(w)
	})

	mux.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		taskCtrl.PostTask(w, r)
	})

	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		taskCtrl.PutTask(w, r)
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
