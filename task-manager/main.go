package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func deleteTasks(id int) {

	tasks := listTasks()
	updatedNewTasks := []Task{}

	for _, t := range tasks {
		if t.ID != id {
			updatedNewTasks = append(updatedNewTasks, t)

		}
	}
	saveTasks(updatedNewTasks)
	fmt.Print("Tasks deleted sucessfully\nUse <list> command to see tasks in line")

}

func addTasks() {
	filename := "tasks.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		saveTasks([]Task{})
	}

	tasks := listTasks()

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: os.Args[2],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func saveTasks(tasks []Task) {
	filename := "tasks.json"

	newData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, newData, 0644)
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func listTasks() []Task {
	filename := "tasks.json"

	data, err := os.ReadFile(filename)
	if err != nil {
		panic("cant read file [from listTasks]")
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)

	if err != nil {
		panic("err in unmarshal [listTasks]")
		return []Task{}
	}

	return tasks

}

func updateTasks(id int, status string) {
	tasks := listTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].UpdatedAt = time.Now()
			tasks[i].Status = status
		}

	}

	saveTasks(tasks)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Print("Error not enough arguments/nUsage : task-cli <command> [Arguments]")
		return
	}

	commands := os.Args[1]

	switch strings.ToLower(commands) {
	case "add":
		addTasks()
	case "list":
		tasks_list := listTasks()
		filter := ""
		if len(os.Args) > 2 {
			filter = strings.ToLower(os.Args[2])
		}

		fmt.Println("ID | Description | Status")
		for _, t := range tasks_list {
			// If no filter we will show all. If filter exists, only show matching status like only "todo" or "in-progress".
			if filter == "" || t.Status == filter {
				fmt.Printf("%d | %s | %s\n", t.ID, t.Description, t.Status)
			}
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing ID or New Description\nUsage: update <id> <description>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newStatus := os.Args[3]
		updateTasks(id, newStatus)

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteTasks(id)
	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		updateTasks(id, "in-progress")
	case "mark-done":
		id, _ := strconv.Atoi(os.Args[2])
		updateTasks(id, "done")

	default:
		fmt.Print("Unkown command usage\n[add/list/update/delete]")
	}

}
