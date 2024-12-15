package task

import (
	"fmt"
)

type Task struct {
	ID         int    `json:"id"`
	TaskName   string `json:"task"`
	Completed  bool   `json:"completed"`
	InProgress bool   `json:"inprogress"`
}

func NewTask(id int, taskName string) *Task {
	return &Task{
		ID:        id,
		TaskName:  taskName,
		Completed: false,
	}
}

func (t *Task) Complete() {
	t.Completed = true
	t.InProgress = false
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.InProgress = false
}

func (t *Task) Start() {
	t.InProgress = true
	t.Completed = false
}

func (t *Task) Update(newName string) {
	t.TaskName = newName
}

func (t *Task) StatusEmoji() string {
	if t.Completed {
		return "✅"
	}

	if t.InProgress {
		return "🚧"
	}

	return "❌"
}

func (t *Task) Print() {
	fmt.Printf("%d: %s %s\n", t.ID, t.TaskName, t.StatusEmoji())
}
