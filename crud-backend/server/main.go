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

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskCtrl.GetTasks(w)
			return
		}
		if r.Method == http.MethodPost {
			taskCtrl.PostTask(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
