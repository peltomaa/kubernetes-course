package main

import (
	"fmt"
	"net/http"
	"os"

	"crud-backend/controllers"
	"crud-backend/db"
	"crud-backend/nats"
	"crud-backend/repositories"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
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

	nc, err := nats.InitNats()
	if err != nil {
		fmt.Println("Failed init nats:", err)
	}

	taskRepo := repositories.TaskRepository{DB: DB}
	natsRepo := repositories.NatsRepository{Conn: nc}
	taskCtrl := controllers.TaskController{R: &taskRepo, N: &natsRepo}

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

	http.Handle("/", corsMiddleware(mux))

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
