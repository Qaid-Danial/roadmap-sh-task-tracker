package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

	ID int `json:"ID"`
	Description string `json:"Description"`
	Status Status `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAT"`

}

func addTask(description string) {

	readTask()

	new_task := tasks{
		ID: (len(taskList) + 1),
		Description: description,
		Status: toDo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	taskList = append(taskList, new_task)

	marshalled, err_marshal := json.MarshalIndent(taskList, "", " ")
	if err_marshal != nil {
		fmt.Printf("Error when marshaling data: %v\n", err_marshal)
	}

	err_write := os.WriteFile("task.json", marshalled, 0644)
	if err_write != nil {
		fmt.Printf("Error when writing data: %v\n", err_write)
	}

}

func readTask() {

	reader, err_read := os.ReadFile("task.json")
	if err_read != nil {
		fmt.Printf("Error when reading data: %v\n", err_read)
	}

	if len(reader) != 0 {
		err_unmarshal := json.Unmarshal(reader, &taskList)
		if err_unmarshal != nil {
			fmt.Printf("Error when unmarshaling data: %v\n", err_unmarshal)
		}
	}

}

func printTasks(statusInt int) {

	readTask()

	for i := range taskList {
		if statusInt == 0 {
			if taskList[i].Status == "to-do" {
				fmt.Printf("\nTask ID: %d | Task Description: %s | Status: %s\n", taskList[i].ID, taskList[i].Description, taskList[i].Status)
				fmt.Printf("Last Modified: %v\n\n",  taskList[i].UpdatedAt.Format(time.ANSIC))
			}
		}

	}
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

			addTask(description)

			fmt.Printf("'%s' is added to the list of tasks\n\n", description)

			// taskList = append(taskList, tasks{
			// 	id: (len(taskList) + 1),
			// 	description: description,
			// 	status: toDo,
			// 	createdAt: time.Now(),
			// 	updatedAt: time.Now(),
			// })

		case "update", "u":
			
			// parameter := strings.ToLower(text_list[1])
			// parameter_list := strings.SplitN(parameter, " ", 2)

			// task_id, _ := strconv.Atoi(parameter_list[0])
			
			// for i, t := range taskList {
			// 	if t.id == task_id {
			// 		taskList[i].description = strings.Trim(parameter_list[1], `"`)
			// 		taskList[i].updatedAt = time.Now()
			// 	}
			// }
			

		case "delete", "d":
			fmt.Printf("\nDeleting Task\n\n")

		case "list", "l":

			status := strings.ToLower(text_list[1])

			switch status{
			case "todo":
				printTasks(0)
			}


		default:
			fmt.Printf("\nInvalid\n\n")

		
		} 
	}

}
