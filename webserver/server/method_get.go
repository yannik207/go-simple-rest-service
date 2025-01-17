package server

import (
	"fmt"
	"net/http"
	"task-api/database"
)

func (s *APIServer) Get(w http.ResponseWriter, r *http.Request) error {
	loadedTasks, err := database.LoadTasksFromFile("tasks.json")
	if err != nil {
		return fmt.Errorf("Error loading tasks: %s", err)
	}

	statusCode := http.StatusOK //200

	// Return a response to the user
	// responseMessage := fmt.Sprintf("Status code: %d", statusCode)

	// err = json.NewEncoder(w).Encode(loadedTasks)
	return WriteJSON(w, statusCode, loadedTasks)
}
