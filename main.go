//Finish learning Go Language Basics
//Watch a movie
//Try to do a creation with Go Language
//Make my breakfast by my self
//Scroll Facebook
//Search new functionalities from internet

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Task struct {
	Description string
	Priority    string
	Category    string
}

var tasks []Task // Slice to store tasks

func addTask() {
	if len(tasks) >= 10 {
		fmt.Println("Maximum limit is 10. You have no spaces.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Get task description
	fmt.Println("Enter the task description:")
	if scanner.Scan() {
		description := scanner.Text()

		// Get priority level
		var priority string
		for {
			fmt.Println("Enter the priority (High, Medium, Low):")
			if scanner.Scan() {
				priority = strings.Title(scanner.Text())
				if priority == "High" || priority == "Medium" || priority == "Low" {
					break
				}
				fmt.Println("Invalid input. Please enter High, Medium, or Low.")
			}
		}

		// Get category
		var category string
		for {
			fmt.Println("Enter the category (Learning, Fun):")
			if scanner.Scan() {
				category = strings.Title(scanner.Text())
				if category == "Learning" || category == "Fun" {
					break
				}
				fmt.Println("Invalid input. Please enter Learning or Fun.")
			}
		}

		// Add task to the list
		tasks = append(tasks, Task{
			Description: description,
			Priority:    priority,
			Category:    category,
		})
		fmt.Println("Task added successfully!")
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("There are no any items to show.")
		return
	}

	fmt.Println("List of my Todos:")
	for i, task := range tasks {
		fmt.Printf("%d. [%s | %s] %s\n", i+1, task.Priority, task.Category, task.Description)
	}
}

func updateTask() {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks to update.")
		return
	}

	var taskNumber int
	fmt.Println("Enter the number of the task to update:")
	fmt.Scanln(&taskNumber)

	if taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("There is no task in that position.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Update task description
	fmt.Println("Enter updated task description:")
	if scanner.Scan() {
		tasks[taskNumber-1].Description = scanner.Text()
	}

	// Update priority
	for {
		fmt.Println("Enter updated priority (High, Medium, Low):")
		if scanner.Scan() {
			priority := strings.Title(scanner.Text())
			if priority == "High" || priority == "Medium" || priority == "Low" {
				tasks[taskNumber-1].Priority = priority
				break
			}
			fmt.Println("Invalid input. Please enter High, Medium, or Low.")
		}
	}

	// Update category
	for {
		fmt.Println("Enter updated category (Learning, Fun):")
		if scanner.Scan() {
			category := strings.Title(scanner.Text())
			if category == "Learning" || category == "Fun" {
				tasks[taskNumber-1].Category = category
				break
			}
			fmt.Println("Invalid input. Please enter Learning or Fun.")
		}
	}

	fmt.Println("Task updated successfully!")
}

func deleteTask() {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks to delete.")
		return
	}

	var taskNumber int
	fmt.Println("Enter the number of the task to delete:")
	fmt.Scanln(&taskNumber)

	if taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("There is no task in that position.")
		return
	}

	fmt.Printf("Are you sure you want to delete this item? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		confirmation := strings.ToLower(scanner.Text())
		if confirmation == "yes" {
			// Remove the task
			tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
			fmt.Println("Task deleted successfully.")
		} else {
			fmt.Println("Task not deleted.")
		}
	}
}

func priorityList() {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks to show.")
		return
	}

	// Sort tasks based on priority: High -> Medium -> Low
	sort.SliceStable(tasks, func(i, j int) bool {
		priorityOrder := map[string]int{"High": 1, "Medium": 2, "Low": 3}
		return priorityOrder[tasks[i].Priority] < priorityOrder[tasks[j].Priority]
	})

	fmt.Println("Tasks sorted by priority:")
	for i, task := range tasks {
		fmt.Printf("%d. [%s] %s (%s)\n", i+1, task.Priority, task.Description, task.Category)
	}
}

func categoryBased() {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks to show.")
		return
	}

	// Separate tasks by category
	var learningTasks, funTasks []Task
	for _, task := range tasks {
		if task.Category == "Learning" {
			learningTasks = append(learningTasks, task)
		} else if task.Category == "Fun" {
			funTasks = append(funTasks, task)
		}
	}

	fmt.Println("Tasks categorized by Learning:")
	for i, task := range learningTasks {
		fmt.Printf("%d. [%s] %s\n", i+1, task.Priority, task.Description)
	}

	fmt.Println("\nTasks categorized by Fun:")
	for i, task := range funTasks {
		fmt.Printf("%d. [%s] %s\n", i+1, task.Priority, task.Description)
	}
}

func main() {
	fmt.Println("--- Welcome to our Todo List App! ---")

	for {
		fmt.Println("\nMenu of Functions")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View all Tasks")
		fmt.Println("3. Update a Task")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Priority List")
		fmt.Println("6. Category Based")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			updateTask()
		case 4:
			deleteTask()
		case 5:
			priorityList()
		case 6:
			categoryBased()
		case 7:
			fmt.Println("Exiting Todo List App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
