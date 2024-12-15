package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"zusk/task-tracker/task"
)

func InProgress(fileName string, id int) error {
	var todos []task.Task

	if id == 0 || id == -1 {
		return fmt.Errorf("provide a valid id")
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return fmt.Errorf("no tasks available, add a task first")
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return err
	}

	for i, task := range todos {
		if task.ID == id {
			task.Start()
			todos[i] = task
		}
	}

	output, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, output, 0644)
}
