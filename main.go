package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (t tasks) printTask() {
	fmt.Printf("\nTask ID: %d | Task Description: %s | Status: %s\n", t.id, t.description, t.status)
	fmt.Printf("Last Modified: %v\n\n", t.updatedAt.Format(time.ANSIC))
}

var (
	taskList []tasks
)

func main() {

	// italic := "\033[3m"
	// reset := "\033[0m"

	var command string 

	// var id int

	
	Reader := bufio.NewReader(os.Stdin)

	for {		

		fmt.Print("Task-CLI ")
		// fmt.Printf("\nl2\n")

		// n, err := fmt.Scanf("%s", &command)
		
		// command = strings.ToLower(command)

		text, _ := Reader.ReadString('\n')
		text = strings.TrimSpace(text)

		text_list := strings.SplitN(text, " ", 2)

		command = strings.ToLower(text_list[0])
			
		switch command {
		case "exit", "e":
			fmt.Printf("\nExiting\n\n")
			return

		case "add", "a":

			description := text_list[1]
			description = strings.Trim(description, `"`)

			fmt.Printf("'%s' is added to the list of tasks\n\n", description)

			taskList = append(taskList, tasks{
				id: (len(taskList) + 1),
				description: description,
				status: toDo,
				createdAt: time.Now(),
				updatedAt: time.Now(),
			})

		case "update", "u":
			
			parameter := strings.ToLower(text_list[1])
			parameter_list := strings.SplitN(parameter, " ", 2)

			task_id, _ := strconv.Atoi(parameter_list[0])
			
			for i, t := range taskList {
				if t.id == task_id {
					taskList[i].description = strings.Trim(parameter_list[1], `"`)
					taskList[i].updatedAt = time.Now()
				}
			}
			

		case "delete", "d":
			fmt.Printf("\nDeleting Task\n\n")

		case "list", "l":

			status := strings.ToLower(text_list[1])

			switch status{
			case "todo":
				
				for _, t := range taskList {
					if t.status == "to-do" {
						t.printTask()
					}
				}	
			}


		default:
			fmt.Printf("\nInvalid\n\n")

		
		} 
	}

}