package methods

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"task-api/database"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()


func Post(w http.ResponseWriter, r *http.Request) {
	var taskStruct database.TasksStruct

	fmt.Printf("Type of ResponseWriter: %T\n", w)
	fmt.Printf("%+v\n", r)

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Access the parsed form data
	form := r.PostForm
	fmt.Println("form: ", form)

	// Validate the number of form fields
	if len(form) != 5 {
		http.Error(w, "Form data is incomplete or invalid", http.StatusBadRequest)
		return
	}

	// Decode the form data into the struct
	if err := decoder.Decode(&taskStruct, form); err != nil {
		log.Println("Error decoding form data: ", err)
		http.Error(w, "Failed to decode form data", http.StatusBadRequest)
		return
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
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}

	// Append the new task to the existing tasks
	allTasks := append(existingTasks, newTask)

	// Save the updated tasks to the file
	err = database.SaveTasksToFile(allTasks, "tasks.json")
	if err != nil {
		log.Println("Error saving tasks to file:", err)
		http.Error(w, "Failed to save tasks", http.StatusInternalServerError)
		return
	}

	// Respond with the updated task list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskStruct)
}
