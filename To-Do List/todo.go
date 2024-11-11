package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task struct to hold task details
type Task struct {
	Name        string // Name of the task
	Description string // Description of the task
	Done        bool   // Status of the task (completed or not)
}

var tasks []Task

const filename = "To-Do List/saveTask.txt" // Path to the tasks file

func main() {
	loadTasks() // Load tasks from the file at the start of the program
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List:")
		printTasks()

		fmt.Println("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Done")
		fmt.Println("3. Remove Task")
		fmt.Println("4. See All Tasks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addTask(scanner)
		case "2":
			markTaskAsDone(scanner)
		case "3":
			removeTask(scanner) // Remove task by name
		case "4":
			seeAllTasks(scanner) // See all tasks from the file
		case "5":
			saveTasks() // Save tasks before exiting
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// Function to print all tasks
func printTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✓" // Mark as done
		}
		fmt.Printf("%d. [%s] %s: %s\n", i+1, status, task.Name, task.Description)
	}
}

// Function to add a new task
func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter task description: ")
	scanner.Scan()
	description := scanner.Text()

	tasks = append(tasks, Task{Name: name, Description: description, Done: false}) // Add new task
	fmt.Println("Task added.")
	saveTasks() // Save tasks after adding
}

// Function to mark a task as done
func markTaskAsDone(scanner *bufio.Scanner) {
	fmt.Print("Enter task name to mark as done: ")
	scanner.Scan()
	taskName := scanner.Text()

	for i, task := range tasks {
		if strings.EqualFold(task.Name, taskName) { // Case-insensitive comparison
			tasks[i].Done = true
			fmt.Println("Task marked as done.")
			saveTasks() // Save changes to the file
			return
		}
	}
	fmt.Println("Task not found.")
}

// Function to remove a task by name
func removeTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task name to remove: ")
	scanner.Scan()
	taskName := scanner.Text()

	for i, task := range tasks {
		if strings.EqualFold(task.Name, taskName) { // Case-insensitive comparison
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task
			fmt.Println("Task removed.")
			saveTasks() // Save tasks after removal
			return
		}
	}
	fmt.Println("Task not found.")
}

// Function to see all tasks directly from the file
func seeAllTasks(scanner *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks available.")
			return
		}
		fmt.Println("Error loading tasks:", err)
		return
	}
	defer file.Close()

	scannerFile := bufio.NewScanner(file)
	fmt.Println("\nAll Tasks:")
	for scannerFile.Scan() {
		line := scannerFile.Text()
		if strings.HasPrefix(line, "Name: ") {
			fmt.Println(line) // Print the name line
			if scannerFile.Scan() {
				fmt.Println(scannerFile.Text()) // Print the description line
			}
			if scannerFile.Scan() {
				fmt.Println(scannerFile.Text()) // Print the status line
			}
			fmt.Println() // Print an empty line for better readability
		}
	}

	if err := scannerFile.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("Press Enter to return to the main menu...")
	scanner.Scan() // Wait for user to press Enter
}

// Function to save tasks to the file
func saveTasks() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		status := ""
		if task.Done {
			status = "task is done ✓" // Status message for completed tasks
		}
		// Write task details to the file
		_, err := fmt.Fprintf(file, "Name: %s\nDescription: %s\nStatus: %s\n\n", task.Name, task.Description, status)
		if err != nil {
			fmt.Println("Error writing task to file:", err)
		}
	}
}

// Function to load tasks from the file
func loadTasks() {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return // Do nothing if the file does not exist
		}
		fmt.Println("Error loading tasks:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Name: ") {
			name := strings.TrimPrefix(line, "Name: ")
			scanner.Scan()
			description := strings.TrimPrefix(scanner.Text(), "Description: ")
			scanner.Scan()
			statusLine := scanner.Text()
			done := strings.Contains(statusLine, "task is done") // Check if the task is done
			tasks = append(tasks, Task{Name: name, Description: description, Done: done}) // Add loaded task
			scanner.Scan() // Skip the empty line after each task
		}
	}
}