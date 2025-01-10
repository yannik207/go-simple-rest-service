package main

import (
	"fmt"
	"net/http"
	"task-api/server"
)

func main() {

	http.HandleFunc("/hello", server.TasksHandler)
	// http.HandleFunc("/posts/", tasksHandler)

	fmt.Println("Server is running at http://localhost:8080/hello")
	http.ListenAndServe(":8080", nil)
}
