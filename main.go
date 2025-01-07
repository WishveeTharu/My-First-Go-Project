////Seperate list to learning and fun

////Finish learning Go Language Basics
////Watch a movie
////Try to do a creation with Go Language
////Make my breakfast by my self
////Do a project using Go Language
////Scroll Facebook
////Search new functionalities from internet
////Try out those new functionalities
////Walk
////Sleep at 12PM

package main

import (
	"bufio"
	"fmt"
	"os"
)

var count int

func addTask(taskItems []string) {
	if count >= 10 {
		fmt.Println("Maximum limit is 10. You have no spaces.")
		return
	}

	fmt.Println("Enter new task:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		taskItems[count] = scanner.Text()
		count++
		fmt.Println("Task added successfully.")
	}
}

func printTasks(taskItems []string) {
	if count == 0 {
		fmt.Println("There are no any items to show.")
		return
	}

	fmt.Println("List of my Todos:")
	for i := 0; i < count; i++ {
		fmt.Printf("%d: %s\n", i+1, taskItems[i])
	}
}

func updateTask(taskItems []string) {
	if count == 0 {
		fmt.Println("There are no tasks to update.")
		return
	}

	fmt.Println("Enter the number of the task to update:")
	var taskNumber int
	fmt.Scanln(&taskNumber)

	if taskNumber < 1 || taskNumber > count {
		fmt.Println("There is no task in that position.")
		return
	}

	fmt.Println("Enter updated task:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		taskItems[taskNumber-1] = scanner.Text()
		fmt.Println("Task updated successfully.")
	}
}

func deleteTask(taskItems []string) {
	if count == 0 {
		fmt.Println("There are no tasks to delete.")
		return
	}

	fmt.Println("Enter the number of the task to delete:")
	var taskNumber int
	fmt.Scanln(&taskNumber)

	if taskNumber < 1 || taskNumber > count {
		fmt.Println("There is no task in that position.")
		return
	}

	fmt.Printf("Are you sure you want to delete this item? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		confirmation := scanner.Text()
		if confirmation == "yes" {
			// Shift tasks to remove the deleted one
			for i := taskNumber - 1; i < count-1; i++ {
				taskItems[i] = taskItems[i+1]
			}
			taskItems[count-1] = "" // Clear the last task
			count--
			fmt.Println("Task deleted successfully.")
		} else {
			fmt.Println("Task not deleted.")
		}
	}
}

func main() {
	var taskItems [10]string // Define an array with a fixed size of 10

	fmt.Println("--- Welcome to our Todo List App! ---")
	fmt.Println("** First you have to add your To do List (Only 10 will be accepted) **")

	scanner := bufio.NewScanner(os.Stdin)

	for index := 0; index < len(taskItems); index++ {
		fmt.Printf("Enter task %d: ", index+1)
		if scanner.Scan() {
			task := scanner.Text()
			taskItems[index] = task
			count++
		}
	}

	for {
		fmt.Println()
		fmt.Println("Menu of Functions")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View all Tasks")
		fmt.Println("3. Update a Task")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(taskItems[:])
		case 2:
			printTasks(taskItems[:])
		case 3:
			updateTask(taskItems[:])
		case 4:
			deleteTask(taskItems[:])
		case 5:
			fmt.Println("Exiting Todo List App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
