package main

import (
	"fmt"
	"strings"
	"time"
)

type Status string

const (
	toDo Status = "to-do"
	inProgress Status = "in-progress"
	done Status = "done"
)

type tasks struct {

	id int
	description string 
	status Status
	createdAt time.Time
	updatedAt time.Time

}

func main() {

	// var command string 
	// var description string
	// var id int

	var command1, command2 string

	fmt.Print("> ")

	command1 = strings.ToLower(command1)
	command2 = strings.ToLower(command2)

	fmt.Scanf("%s", &command1)

	if command1 == "pipi" {
		fmt.Scanf("%q", &command2)
	} else if command1 == "popo" {
		fmt.Scanf("%s", &command2)
	}


	fmt.Println(command1)
	fmt.Println(command2)
	


	// for {		

	// 	fmt.Print("Task-CLI ")
	// 	n, err := fmt.Scanln(&command)
		
	// 	command = strings.ToLower(command)

	// 	if n != 0 && err == nil {
			
	// 		switch command {
	// 		case "exit", "e":
	// 			fmt.Printf("\nExiting\n\n")
	// 			return

	// 		case "add", "a":
	// 			fmt.Printf("\nAdding Task\n\n")

	// 		case "update", "u":
	// 			fmt.Printf("\nUpdating Task\n\n")
				

	// 		case "delete", "d":
	// 			fmt.Printf("\nDeleting Task\n\n")


	// 		default:
	// 			fmt.Printf("\nInvalid\n\n")

	// 		}
	// 	} 
	// }

}