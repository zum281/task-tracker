package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"zusk/task-tracker/task"
)

func List(fileName string, done bool, pending bool, notStarted bool) error {

	var todos []task.Task

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

	if done {
		fmt.Println("Completed tasks:")
		for _, task := range todos {
			if task.Completed {
				task.Print()
			}
		}
		return nil
	}

	if pending {
		fmt.Println("In progress tasks:")
		for _, task := range todos {
			if !task.Completed {
				task.Print()
			}
		}
		return nil
	}

	if notStarted {
		fmt.Println("Not started tasks:")
		for _, task := range todos {
			if !task.Completed {
				task.Print()
			}
		}
		return nil
	}

	fmt.Println("All tasks:")
	for _, task := range todos {
		task.Print()
	}

	return nil

}
