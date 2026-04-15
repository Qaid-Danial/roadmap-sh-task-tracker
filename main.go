package main

import (
	"fmt"
	"strings"
)

func main() {

	var command string

	for {		

		fmt.Print("Task-CLI ")
		n, err := fmt.Scanln(&command)
		
		command = strings.ToLower(command)

		if n != 0 && err == nil {
			
			switch command {
			case "exit", "e":
				fmt.Println("Exiting")
				return

			case "add", "a":
				fmt.Println("Adding Task")

			case "update", "u":
				fmt.Println("Updating Task")

			case "delete", "d":
				fmt.Println("Deleting Task")

			default:
				fmt.Println("Invalid")
			}
		} 
	}

}