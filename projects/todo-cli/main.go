package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Date      string `json:"date"`
	Priority  string `json:"priority"`
	Completed bool   `json:"completed"`
}

var todosMap = make(map[string][]Todo)

const dataFile = "todos.json"

func main() {
	err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
	}

	counter := getHighestID() + 1

	fmt.Println("Welcome to the Go To-Do CLI Application")
	fmt.Println("---------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)
	operation := ""

	displayMenu()

	for operation != "exit" && operation != "quit" {
		fmt.Print("\nYour operation: ")
		if !scanner.Scan() {
			break
		}

		operation = strings.TrimSpace(scanner.Text())

		switch strings.ToLower(operation) {
		case "add":
			task, priority := promptForTodoDetails(scanner)
			date := time.Now().Format("2006-01-02")

			todo := Todo{
				ID:        counter,
				Task:      task,
				Date:      date,
				Priority:  validatePriority(priority),
				Completed: false,
			}

			addTodo(todo, date)
			counter++

			// Save after changes
			saveTodos()

		case "complete":
			date, id, valid := promptForTaskIdentification(scanner)
			if !valid {
				continue
			}

			err := completeTask(date, id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("✓ Task marked as completed")
			saveTodos()

		case "list":
			listAllTodos()

		case "delete":
			date, id, valid := promptForTaskIdentification(scanner)
			if !valid {
				continue
			}

			err := deleteTodo(date, id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("✓ Task deleted successfully")
			saveTodos()

		case "quit", "exit":
			fmt.Println("Goodbye! Your todos have been saved.")

		case "help":
			displayMenu()

		default:
			fmt.Println("❌ Unknown command. Type 'help' to see available options.")
		}
	}
}

func displayMenu() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  add      - Add a new task")
	fmt.Println("  complete - Mark a task as completed")
	fmt.Println("  list     - List all tasks")
	fmt.Println("  delete   - Delete a task")
	fmt.Println("  help     - Show this menu")
	fmt.Println("  quit     - Exit the application")
}

func promptForTodoDetails(scanner *bufio.Scanner) (string, string) {
	fmt.Print("Task description: ")
	scanner.Scan()
	task := strings.TrimSpace(scanner.Text())

	fmt.Print("Priority (high/medium/low): ")
	scanner.Scan()
	priority := strings.TrimSpace(scanner.Text())

	return task, priority
}

func promptForTaskIdentification(scanner *bufio.Scanner) (string, int, bool) {
	fmt.Print("Date (YYYY-MM-DD): ")
	scanner.Scan()
	date := strings.TrimSpace(scanner.Text())

	fmt.Print("Task ID: ")
	scanner.Scan()
	idText := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idText)
	if err != nil {
		fmt.Println("❌ Invalid ID. Please enter a number.")
		return "", 0, false
	}

	return date, id, true
}

func validatePriority(priority string) string {
	priority = strings.ToLower(strings.TrimSpace(priority))

	switch priority {
	case "high", "medium", "low":
		return priority
	default:
		fmt.Println("! Invalid priority. Using 'medium' as default.")
		return "medium"
	}
}

func getHighestID() int {
	highest := 0
	for _, todos := range todosMap {
		for _, todo := range todos {
			if todo.ID > highest {
				highest = todo.ID
			}
		}
	}
	return highest
}

func addTodo(todo Todo, date string) {
	todosMap[date] = append(todosMap[date], todo)
	fmt.Println("✓ New task added successfully:")
	fmt.Printf("   [%d] %s (Priority: %s)\n", todo.ID, todo.Task, todo.Priority)
}

func completeTask(date string, id int) error {
	todoList, exists := todosMap[date]
	if !exists {
		return fmt.Errorf("no tasks found for date %s", date)
	}

	for i := range todoList {
		if todoList[i].ID == id {
			todoList[i].Completed = true
			todosMap[date] = todoList
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func deleteTodo(date string, id int) error {
	todoList, exists := todosMap[date]
	if !exists {
		return fmt.Errorf("no tasks found for date %s", date)
	}

	for i, todo := range todoList {
		if todo.ID == id {
			todosMap[date] = append(todoList[:i], todoList[i+1:]...)
			if len(todosMap[date]) == 0 {
				delete(todosMap, date)
			}
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func listAllTodos() {
	if len(todosMap) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	dates := make([]string, 0, len(todosMap))
	for date := range todosMap {
		dates = append(dates, date)
	}

	for _, date := range dates {
		todos := todosMap[date]
		fmt.Printf("\n📅 %s (%d tasks):\n", date, len(todos))
		fmt.Println(strings.Repeat("-", 40))

		for _, todo := range todos {
			status := "[ ]"
			if todo.Completed {
				status = "[✓]"
			}

			priorityIcon := "⚪"
			if todo.Priority == "high" {
				priorityIcon = "🔴"
			} else if todo.Priority == "low" {
				priorityIcon = "🔵"
			}

			fmt.Printf("%s %s ID:%d - %s (%s priority)\n",
				status, priorityIcon, todo.ID, todo.Task, todo.Priority)
		}
	}
}

func saveTodos() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(todosMap)
}

func loadTodos() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&todosMap)
}
