package clitaskmanager

import (
	"fmt"
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
			Add(strings.Join(parts[1:], " "))
		}
	case "remove":
		{
			Remove(parts[1])
		}

	case "list":
		{
			List()
		}
	default:
		fmt.Println("unknown command ,please user add,remove or list commands")
	}

}
