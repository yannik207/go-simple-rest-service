package server

import (
	"fmt"
	"log"
	"net/http"
	"task-api/database"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (s *APIServer) Post(w http.ResponseWriter, r *http.Request) error {
	var taskStruct database.TasksStruct

	fmt.Printf("Type of ResponseWriter: %T\n", w)
	fmt.Printf("%+v\n", r)

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("Failed to parse form data %d\n", http.StatusBadRequest)
	}

	// Access the parsed form data
	form := r.PostForm
	fmt.Println("form: ", form)

	// Validate the number of form fields
	if len(form) != 5 {
		return fmt.Errorf("Form data is incomplete or invalid %d\n", http.StatusBadRequest)
	}

	// Decode the form data into the struct
	if err := decoder.Decode(&taskStruct, form); err != nil {
		log.Println("Error decoding form data: ", err)
		return fmt.Errorf("Failed to decode form data %d\n", http.StatusBadRequest)
	}

	log.Println("Decoded task: ", taskStruct)

	// Create a new task
	newTask := database.TasksStruct{
		ID:          form.Get("ID"),
		Title:       form.Get("Title"),
		Description: form.Get("Description"),
		DueDate:     form.Get("DueDate"),
		Completed:   form.Get("Completed"),
	}

	// Load existing tasks from the file
	existingTasks, err := database.LoadTasksFromFile("tasks.json")
	if err != nil {
		log.Println("Error loading tasks from file:", err)
		return fmt.Errorf("Failed to load tasks %d\n", http.StatusInternalServerError)
	}

	// Append the new task to the existing tasks
	allTasks := append(existingTasks, newTask)

	// Save the updated tasks to the file
	err = database.SaveTasksToFile(allTasks, "tasks.json")
	if err != nil {
		log.Println("Error saving tasks to file:", err)
		return fmt.Errorf("Failed to save tasks %d\n", http.StatusInternalServerError)
	}

	return WriteJSON(w, http.StatusAccepted, taskStruct)
}
