package main

import (
	"fmt"

	"crud-backend/db"
	"crud-backend/models"
	"crud-backend/repositories"
)

func main() {
	fmt.Println("Starting worker")

	DB, err := db.InitDb()
	if err != nil {
		fmt.Println("Failed init db:", err)
	}
	db.CreateTable(DB)
	taskRepo := repositories.TaskRepository{DB: DB}

	wikiRepo := repositories.WikipediaArticleRepository{}

	article, err := wikiRepo.GetRandom()
	if err != nil {
		fmt.Println("Failed to get wikipedia article:", err)
	}

	newTask := models.Task{
		Task: fmt.Sprintf("Read %s", article.Url),
	}
	err = taskRepo.Insert(&newTask)
	if err != nil {
		fmt.Println("Failed to insert wikipedia article task:", err)
	}
}
