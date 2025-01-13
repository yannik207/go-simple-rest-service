package methods

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type TasksStruct struct {
	ID          string    `schema:"ID"`
	Title       string    `schema:"Title"`
	Description string    `schema:"Description"`
	DueDate     time.Time `schema:"DueDate"`
	Completed   bool      `schema:"Completed"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	var taskStruct TasksStruct

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
		statusCode := http.StatusBadRequest
		w.WriteHeader(statusCode)
		http.Error(w, "Form data is incomplete or invalid", statusCode)
		return
	}

	// Decode the form data into the struct
	if err := decoder.Decode(&taskStruct, form); err != nil {
		log.Println("Error decoding form data: ", err)
		http.Error(w, "Failed to decode form data", http.StatusBadRequest)
		return
	}

	log.Println("Decoded task: ", taskStruct)

	// Send the decoded task as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskStruct)
}
