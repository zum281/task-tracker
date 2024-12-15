package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"zusk/task-tracker/cmd"
)

const ADD = "add"
const UPDATE = "update"
const DELETE = "delete"
const LIST = "list"
const HELP = "h"
const COMPLETE = "complete"
const IN_PROGRESS = "start"
const UNCOMPLETE = "stop"
const EMPTY = "empty"

var COMMANDS = []string{ADD, UPDATE, DELETE, LIST}

func main() {

	command := "h"

	addCmd := flag.NewFlagSet(ADD, flag.ExitOnError)
	nameAddFlag := addCmd.String("n", "", "name of the task")

	updateCmd := flag.NewFlagSet(UPDATE, flag.ExitOnError)
	idUpdateFlag := updateCmd.Int("i", -1, "id of the task")
	nameUpdateFlag := updateCmd.String("n", "", "name of the task")

	deleteCmd := flag.NewFlagSet(DELETE, flag.ExitOnError)
	idDeleteFlag := deleteCmd.Int("i", -1, "id of the task")

	listCmd := flag.NewFlagSet(LIST, flag.ExitOnError)
	doneFlag := listCmd.Bool("c", false, "list all done tasks")
	inProgressFlag := listCmd.Bool("p", false, "list all in progress tasks")
	notStartedFlag := listCmd.Bool("n", false, "list all not started tasks")

	completeCmd := flag.NewFlagSet(COMPLETE, flag.ExitOnError)
	idCompleteFlag := completeCmd.Int("i", -1, "id of the task")

	inProgressCmd := flag.NewFlagSet(IN_PROGRESS, flag.ExitOnError)
	idInProgressFlag := inProgressCmd.Int("i", -1, "id of the task")

	uncompleteCmd := flag.NewFlagSet(UNCOMPLETE, flag.ExitOnError)
	idUncompleteFlag := uncompleteCmd.Int("i", -1, "id of the task")

	emptyCmd := flag.NewFlagSet(EMPTY, flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("command expected, please provide a valid command")
	} else {
		command = os.Args[1]
	}

	f, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	switch command {
	case ADD:
		addCmd.Parse(os.Args[2:])
		err := cmd.Add("tasks.json", *nameAddFlag)
		if err != nil {
			log.Fatal(err)
		}
	case UPDATE:
		updateCmd.Parse(os.Args[2:])
		err := cmd.Update("tasks.json", *idUpdateFlag, *nameUpdateFlag)
		if err != nil {
			log.Fatal(err)
		}
	case DELETE:
		deleteCmd.Parse(os.Args[2:])
		err := cmd.Delete("tasks.json", *idDeleteFlag)
		if err != nil {
			log.Fatal(err)
		}

	case COMPLETE:
		completeCmd.Parse(os.Args[2:])
		err := cmd.Complete("tasks.json", *idCompleteFlag)
		if err != nil {
			log.Fatal(err)
		}
	case IN_PROGRESS:
		inProgressCmd.Parse(os.Args[2:])
		err := cmd.InProgress("tasks.json", *idInProgressFlag)
		if err != nil {
			log.Fatal(err)
		}
	case UNCOMPLETE:
		uncompleteCmd.Parse(os.Args[2:])
		err := cmd.Uncomplete("tasks.json", *idUncompleteFlag)
		if err != nil {
			log.Fatal(err)
		}
	case EMPTY:
		emptyCmd.Parse(os.Args[2:])
		err := os.Truncate("tasks.json", 0)
		if err != nil {
			log.Fatal(err)
		}
	case LIST:
		listCmd.Parse(os.Args[2:])
		err := cmd.List("tasks.json", *doneFlag, *inProgressFlag, *notStartedFlag)
		if err != nil {
			log.Fatal(err)
		}
	case HELP:
		fmt.Println("Usage: task-tracker <command> [options]")
		fmt.Println("\na -n <name> : add a task")
		fmt.Println("u -i <id> -n <name> : update a task")
		fmt.Println("rm -i <id> : delete a task")
		fmt.Println("list : list all tasks")
		fmt.Println("list -(c|p|n) : list all (completed|in progress|not started) tasks")
		fmt.Println("h : help. show this message")
		fmt.Println()
	default:
		fmt.Println(fmt.Errorf("invalid command, use 'h' for help"))
	}
}
