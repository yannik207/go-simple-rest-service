package database

import (
	"encoding/json"
	"os"
)

type TasksStruct struct {
	ID          string    `schema:"ID"`
	Title       string    `schema:"Title"`
	Description string    `schema:"Description"`
	DueDate     string    `schema:"DueDate"`
	Completed   string     `schema:"Completed"`
}

func SaveTasksToFile(tasks []TasksStruct, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
