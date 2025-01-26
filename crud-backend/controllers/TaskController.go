package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"crud-backend/models"
	"crud-backend/repositories"
)

type TaskController struct {
	R *repositories.TaskRepository
}

func (c *TaskController) GetTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := c.R.FetchAll()
	if err != nil {
		fmt.Println("Failed to select tasks", err)
		http.Error(w, "Failed to select tasks", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		fmt.Println("Failed to serialize tasks", err)
		http.Error(w, "Failed to serialize tasks", http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		fmt.Println("Failed to decode new task", err)
		http.Error(w, "Failed to decode new task", http.StatusInternalServerError)
		return
	}

	if len(newTask.Task) > 140 {
		fmt.Println("Task exceeds character length of 140")
		http.Error(w, "Task exceeds character length of 140", http.StatusForbidden)
		return
	}

	fmt.Println("Insert new task:", newTask.Task)
	err = c.R.Insert(&newTask)
	if err != nil {
		fmt.Println("Failed to insert new task", err)
		http.Error(w, "Failed to insert new task", http.StatusInternalServerError)
		return
	}

	c.GetTasks(w)
}

func (c *TaskController) PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskId := r.PathValue("id")

	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		fmt.Println("Failed to decode new task", err)
		http.Error(w, "Failed to decode new task", http.StatusInternalServerError)
		return
	}

	if len(newTask.Task) > 140 {
		fmt.Println("Task exceeds character length of 140")
		http.Error(w, "Task exceeds character length of 140", http.StatusForbidden)
		return
	}

	fmt.Println("Update new task id:", taskId, newTask.Task, "done:", newTask.IsDone)
	err = c.R.Update(taskId, &newTask)
	if err != nil {
		fmt.Println("Failed to insert new task", err)
		http.Error(w, "Failed to insert new task", http.StatusInternalServerError)
		return
	}

	c.GetTasks(w)
}
