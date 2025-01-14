package main

import (
	"fmt"
	"net/http"
	"task-api/server"
)

func main() {
	// Start HTTP server
	http.HandleFunc("/health", server.WelcomeHandler)
	http.HandleFunc("/", server.ErrorHandler)
	http.HandleFunc("/tasks", server.TaskHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
