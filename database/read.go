package database

import (
	"fmt"
	"encoding/json"
	"os"
)

func loadTasksFromFile(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}