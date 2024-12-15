package cmd

import (
	"encoding/json"
	"os"
	"zusk/task-tracker/task"
)

func Add(fileName, name string) error {

	var todos []task.Task


	data, err := os.ReadFile(fileName)
	if err != nil {
        return err
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &todos); err != nil {
			return err
		}
	}

	newTask := task.NewTask(len(todos) + 1, name)


	todos = append(todos, *newTask)

	output, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, output, 0644)
}