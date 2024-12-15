# Task Tracker CLI

A simple command-line task management tool written in Go that helps you track your tasks with different states (completed, in progress, not started).

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as:
  - Complete ✅
  - In Progress 🚧
  - Not Started ❌
- List tasks with filters

## Installation

```go
go install
```

## Usage

```shell
task-tracker <command> [options]
```

## Commands

- add: Add a new task

```shell
task-tracker add -n "Task name"
```

- update: Update a task

```shell
task-tracker update -i <id> -n "New task name"
```

- delete: Delete a task

```shell
task-tracker delete -i <id>
```

- list: List all tasks

  - -c: Show only completed tasks
  - -p: Show only in-progress tasks
  - -n: Show only not started tasks

- start: Mark a task as in progress

```shell
task-tracker start -i <id>
```

- complete: Mark a task as completed

```shell
task-tracker complete -i <id>
```

- stop: Mark a task as not started

```shell
task-tracker stop -i <id>
```

- empty: Clear all tasks

- or h: Show help message

## Data Storage

Tasks are stored in a local `tasks.json` file in JSON format.
