package main

import (
	"fmt"
	"net/http"
	"time"
	"task-api/database"
	"task-api/server"
)

func main() {
	// Create a slice of tasks
	tasks := []database.Task{
		{
			ID:          "1234",
			Title:       "clean up",
			Description: "My girlfriend forces me to clean up my room",
			DueDate:     time.Now(),
			Completed:   false,
		},
	}

	// Save tasks to file
	err := database.SaveTasksToFile(tasks, "tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	// Load tasks from file
	loadedTasks, err := database.LoadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Print loaded tasks
	fmt.Println("Loaded tasks:", loadedTasks)

	// Start HTTP server
	http.HandleFunc("/health", server.WelcomeHandler)
	http.HandleFunc("/", server.ErrorHandler)
	http.HandleFunc("/tasks", server.TaskHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
