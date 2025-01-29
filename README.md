# Task-Tracker

https://roadmap.sh/projects/task-tracker

## Overview

The **Task Tracker CLI** is a command-line interface application written in Go. It allows users to manage tasks efficiently by adding, updating, deleting, marking tasks as in-progress or done, and listing tasks based on their status. Tasks are stored in memory during execution.

---

## Features

- **Add Tasks**: Add new tasks with a description.
- **Update Tasks**: Modify the description of an existing task.
- **Delete Tasks**: Remove tasks by their unique ID.
- **Change Status**: Mark tasks as `in-progress` or `done`.
- **List Tasks**: View tasks, filtered by status (`todo`, `in-progress`, `done`) or view all tasks.

Each task has the following properties:

- **ID**: A unique identifier.
- **Description**: A short description of the task.
- **Status**: Task state (`todo`, `in-progress`, `done`).
- **CreatedAt**: The timestamp when the task was created.
- **UpdatedAt**: The timestamp when the task was last modified.

---

## Prerequisites

- [Go Programming Language](https://golang.org/) installed (version 1.16 or later).

---

## Installation

1. Clone the repository:

   ```bash
   https://github.com/nabobery/task-tracker.git
   cd task-tracker
   ```

2. Build the CLI:

   ```bash
   # For Linux/Mac
   go build -o task-cli

   # For Windows
   go build -o task-cli.exe main.go
   ```

3. Run the executable:

   ```bash
   ./task-cli
   ```

---

## Usage

### Add a Task

Add a new task with a description.

```bash
task-cli add "Buy groceries"
```

Output:

```bash
Task added successfully (ID: 1)
```

---

### Update a Task

Update a task's description using its ID.

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

Output:

```bash
Task 1 updated successfully
```

---

### Delete a Task

Delete a task by its ID.

```bash
task-cli delete 1
```

Output:

```bash
Task 1 deleted successfully
```

---

### Mark a Task as In Progress

Change the status of a task to `in-progress` using its ID.

```bash
task-cli mark-in-progress 1
```

Output:

```bash
Task 1 marked as in-progress
```

---

### Mark a Task as Done

Change the status of a task to `done` using its ID.

```bash
task-cli mark-done 1
```

Output:

```bash
Task 1 marked as done
```

---

### List All Tasks

List all tasks with their details.

```bash
task-cli list
```

Example Output:

```bash
ID   Status      Created              Description
--   ------      -------              -----------
1    todo       2025-01-20 10:00:00 Buy groceries
```

---

### List Tasks by Status

List tasks filtered by their status: `todo`, `in-progress`, or `done`.

```bash
task-cli list todo
```

Output:

```bash
ID   Status      Created              Description
--   ------      -------              -----------
1    todo       2025-01-20 10:00:00 Buy groceries
```

---

## How It Works

- **Cobra Library**: The CLI is powered by [Cobra](https://github.com/spf13/cobra), a library for building powerful Go CLIs.
- **In-Memory Storage**: Tasks are stored in memory (`tasks` slice) during runtime. To persist data, you can extend this project to store tasks in a file or database.

---

## Extending the Project

1. **Persistent Storage**:
   - Save tasks to a JSON file or a database (e.g., SQLite).
   - Load tasks on application startup and save changes when tasks are modified.

2. **Additional Features**:
   - Add support for due dates or priorities.
   - Implement task search or filtering by keyword.

---

## License

This project is open-source and free to use under the [MIT License](LICENSE). Contributions are welcome!
