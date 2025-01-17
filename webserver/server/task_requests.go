package server

import (
	"fmt"
	"net/http"
	"task-api/server/methods"
)

func (s *APIServer) TaskHandler(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		return s.Get(w, r)
	// case "POST":
	// 	methods.Post(w, r)
	// case "PUT":
	// 	fmt.Println("PUT")
	// case "DELETE":
	// 	fmt.Println("DELETE")
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) Get(w http.ResponseWriter, r *http.Request) error {
	// loadedTasks, err := database.LoadTasksFromFile("tasks.json")
	// if err != nil {
	// 	fmt.Println("Error loading tasks:", err)
	// 	return
	// }

	statusCode := http.StatusOK //200

	w.WriteHeader(statusCode)

	// Return a response to the user
	// responseMessage := fmt.Sprintf("Status code: %d", statusCode)

	w.Header().Set("Content-Type", "application/json")
	// err = json.NewEncoder(w).Encode(loadedTasks)
	return methods.WriteJSON(w, statusCode, "dope")
}
