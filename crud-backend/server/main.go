package main

import (
	"fmt"
	"net/http"
	"os"

	"crud-backend/controllers"
	"crud-backend/db"
	"crud-backend/repositories"
)

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	fmt.Println("Starting backend server")

	DB, err := db.InitDb()
	if err != nil {
		fmt.Println("Failed init db:", err)
	}

	err = db.CreateTable(DB)
	if err != nil {
		fmt.Println("Failed migrate db:", err)
	}

	taskRepo := repositories.TaskRepository{DB: DB}
	taskCtrl := controllers.TaskController{R: &taskRepo}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR"))
			return
		}

		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
		taskCtrl.GetTasks(w)
	})

	mux.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
		taskCtrl.PostTask(w, r)
	})

	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
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
