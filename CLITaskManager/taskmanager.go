package clitaskmanager

import (
	"bufio"
	"fmt"
	"os"
	clitaskmanager "projects/CLITaskManager"
	"strings"
)

func TaskManager() {
	fmt.Println("cli task manager")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$>")
		command, _ := reader.ReadString('\n')
		handleCommand(strings.TrimSpace(command))
	}

}
func handleCommand(command string) {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "add":
		{
			clitaskmanager.Add(strings.Join(parts[1:], " "))
		}
	case "remove":
		{
			clitaskmanager.Remove(parts[1])
		}

	case "list":
		{
			clitaskmanager.List()
		}
	default:
		fmt.Println("unknown command ,please user add,remove or list commands")
	}

}
