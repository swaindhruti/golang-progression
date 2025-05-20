package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Task      string
	Date      time.Time
	Priority  string
	Completed bool
}

var todosMap = make(map[time.Time]Todo)

func main() {

	fmt.Println("Welcome to the Go To-Do CLI Application")
	fmt.Println("---------------------------------------")

	var operation string
	fmt.Println("Please enter the operation you want to perfom from the following options:\n1. Add\n2. Complete\n3. Quit")

	fmt.Print("Your choosen operation : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation = scanner.Text()
	fmt.Println("You selected:", operation)

	switch operation {
	case "Add":
		fmt.Print("Add task -->")
		scanner.Scan()
		task := scanner.Text()
		date := time.Now()
		fmt.Print("Set Priority -->")
		scanner.Scan()
		priority := scanner.Text()
		completed := false

		fmt.Println(task)
		fmt.Println(date)
		fmt.Println(priority)
		fmt.Println(completed)

	case "Complete":
		fmt.Println("Completed")
	case "Quit":
		fmt.Println("Completed")
	default:
		fmt.Println("Choose from the provided options")
	}

}

// func formatInputTodo(todo string) {
// }

// func addTodo(todo Todo) {
// 	for index := range todosMap {
// 		if time.Now() == index {
// 			todosMap[index] = todo
// 		}
// 	}

// 	fmt.Println("The new todo has been added: ", todo)
// }

// func completeTodo(todo Todo) {
// }
