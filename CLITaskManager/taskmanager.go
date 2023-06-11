package clitaskmanager

import (
	"fmt"
	"projects/task"
	"strings"
)

func TaskManager(cmd string) {
	fmt.Println("cli task manager")

	handleCommand(strings.TrimSpace(cmd))

}
func handleCommand(command string) {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "add":
		{
			task.AddTask(strings.Join(parts[1:], " "))
		}
	case "remove":
		{
			task.RemoveTask((parts[1]))
		}

	case "list":
		{
			task.ListTask()
		}
	default:
		fmt.Println("unknown command ,please user add,remove or list commands")
	}

}
