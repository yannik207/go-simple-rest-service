package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/database"
)

func Get(w http.ResponseWriter, r *http.Request) {
	loadedTasks, err := database.LoadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	statusCode := http.StatusOK //200

	w.WriteHeader(statusCode)

	// Return a response to the user
	// responseMessage := fmt.Sprintf("Status code: %d", statusCode)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(loadedTasks)
}
