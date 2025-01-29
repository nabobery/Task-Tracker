package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // todo, in-progress, done
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasks []Task

const tasksFile = "tasks.json"

var nextID = 1

// LoadTasks loads tasks from a JSON file
func LoadTasks() error {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		// Create the file if it doesn't exist
		return os.WriteFile(tasksFile, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	// Update nextID
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	nextID = maxID + 1
	return nil
}

// SaveTasks saves tasks to a JSON file
func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func withTaskPersistence(fn func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fn(cmd, args)
		if err := SaveTasks(); err != nil {
			log.Printf("Error saving tasks: %v\n", err)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A simple task manager CLI application",
}

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: withTaskPersistence(func(cmd *cobra.Command, args []string) {
		task := Task{
			ID:          nextID,
			Description: args[0],
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, task)
		nextID++
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)

	}),
}

var updateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "Update a task's description",
	Args:  cobra.ExactArgs(2),
	Run: withTaskPersistence(func(cmd *cobra.Command, args []string) {
		id := 0
		if _, err := fmt.Sscanf(args[0], "%d", &id); err != nil {
			log.Fatalf("Error parsing ID: %v\n", err)
			return
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = args[1]
				tasks[i].UpdatedAt = time.Now()
				fmt.Printf("Task %d updated successfully\n", id)
				return
			}
		}
		fmt.Printf("Task %d not found\n", id)
	}),
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: withTaskPersistence(func(cmd *cobra.Command, args []string) {
		id := 0
		if _, err := fmt.Sscanf(args[0], "%d", &id); err != nil {
			log.Fatalf("Error parsing ID: %v\n", err)
		}
		for i := range tasks {
			if tasks[i].ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				fmt.Printf("Task %d deleted successfully\n", id)
				return
			}
		}
		fmt.Printf("Task %d not found\n", id)
	}),
}

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark a task as in progress",
	Args:  cobra.ExactArgs(1),
	Run: withTaskPersistence(func(cmd *cobra.Command, args []string) {
		updateTaskStatus(args[0], "in-progress")
	}),
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: withTaskPersistence(func(cmd *cobra.Command, args []string) {
		updateTaskStatus(args[0], "done")
	}),
}

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List all tasks or tasks by status",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := ""
		if len(args) > 0 {
			status = args[0]
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		fmt.Println("\nID\tStatus\t\tCreated\t\t\tDescription")
		fmt.Println("--\t------\t\t-------\t\t\t-----------")

		for _, task := range tasks {
			if status == "" || task.Status == status {
				fmt.Printf("%d\t%s\t%s\t%s\n",
					task.ID,
					task.Status,
					task.CreatedAt.Format("2006-01-02 15:04:05"),
					task.Description)
			}
		}
		fmt.Println()
	},
}

func updateTaskStatus(idStr string, status string) {
	id := 0
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		log.Fatalf("Error parsing ID: %v\n", err)
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Task %d marked as %s\n", id, status)
			return
		}
	}
	fmt.Printf("Task %d not found\n", id)
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := LoadTasks(); err != nil {
		log.Fatalf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
