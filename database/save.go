package database

import (
	"fmt"
	"os"
	"encoding/json"
)

func saveTasksToFile(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}