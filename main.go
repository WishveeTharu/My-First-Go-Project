//Create CRUD Functions
//Seperate list to learning and fun
//do

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("--- Welcome to our Todo List App! ---")

	////var T1 = "Finish learning Go Language Basics"
	////var T2 = "Watch a movie"
	////var T3 = "Try to do a creation with Go Language"
	////var T4 = "Make my breakfast by my self"
	////var T5 = "Do a project using Go Language"
	//
	//var taskItems = []string{
	//	"Finish learning Go Language Basics",    // Learning
	//	"Watch a movie",                         // Fun
	//	"Try to do a creation with Go Language", // Learning
	//	"Make my breakfast by my self",          // Fun
	//	"Do a project using Go Language",        // Learning
	//}
	//
	//printTasks(taskItems)
	//fmt.Println()
	//
	//taskItems = addTask(taskItems, "Go for a run")
	//taskItems = addTask(taskItems, "Practice coding in Go")
	//
	//fmt.Println("Updated List")
	//printTasks(taskItems)

	http.ListenAndServe(":8080", nil)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTask = append(taskItems, newTask)
	return updatedTask
}
