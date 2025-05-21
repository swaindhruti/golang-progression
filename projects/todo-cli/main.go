package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	id        int
	Task      string
	Date      string
	Priority  string
	Completed bool
}

var todosMap = make(map[string][]Todo)

func main() {

	fmt.Println("Welcome to the Go To-Do CLI Application")
	fmt.Println("---------------------------------------")

	var operation string
	var counter int = 0
	fmt.Println("Please enter the operation you want to perfom from the following options:\n1. add\n2. complete\n3.list\n4. quit / exit")

	scanner := bufio.NewScanner(os.Stdin)

	for operation != "exit" {
		fmt.Print("Your choosen operation : ")
		scanner.Scan()

		operation = scanner.Text()
		fmt.Println("You selected:", operation)

		switch operation {
		case "Add", "add":
			fmt.Print("Add task -->")
			scanner.Scan()
			task := scanner.Text()
			date := time.Now().Format("2006-01-02")
			fmt.Print("Set Priority -->")
			scanner.Scan()
			priority := scanner.Text()
			completed := false

			fmt.Println(task)
			fmt.Println(date)
			fmt.Println(priority)
			fmt.Println(completed)

			formattedTodo := formatInputTodo(counter, task, date, priority, completed)
			addTodo(formattedTodo, date)
			counter++

			fmt.Println(todosMap)

		case "Complete", "complete":

			fmt.Print("Date -->")
			scanner.Scan()
			date := scanner.Text()
			fmt.Print("id -->")
			scanner.Scan()
			idText := scanner.Text()
			id, err := strconv.Atoi(idText)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			completeTask(date, id)
			fmt.Println("Completed")
		case "List", "list":
			listAllTodos()
		default:
			fmt.Println("Choose from the provided options")
		}

	}

}

func formatInputTodo(id int, todo string, date string, priority string, completed bool) Todo {
	updatedTodo := Todo{
		id:        id,
		Task:      todo,
		Date:      date,
		Priority:  priority,
		Completed: completed,
	}
	return updatedTodo
}

func addTodo(todo Todo, date string) {
	todosMap[date] = append(todosMap[date], todo)
	fmt.Println("The new todo has been added:", todosMap)
}

func completeTask(date string, id int) {
	for i := range todosMap {
		if i == date {
			for j := range todosMap[i] {
				if todosMap[i][j].id == id {
					todosMap[i][j].Completed = true
					break
				}
			}
		}
	}
}

func listAllTodos() {
	for key, value := range todosMap {
		fmt.Println("Date: ", key)
		for j := range value {
			fmt.Println(value[j])
		}
	}
}
