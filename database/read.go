package database

import (
	"encoding/json"
	"os"
)

func LoadTasksFromFile(filename string) ([]TasksStruct, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks []TasksStruct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}
