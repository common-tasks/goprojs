package clitaskmanager

import (
	"fmt"
	"os"
	"strings"
)

const taskFile = "tasks.txt"

func Add(task string) {
	fs, _ := os.OpenFile(taskFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeExclusive)
	defer fs.Close()
	_, err := fs.WriteString(task)
	if err != nil {
		fmt.Println("error i writing to the file", err)
	} else {
		fmt.Println("added task ", task)
	}

}
func Remove(task string) {
	tasks, _ := os.ReadFile(taskFile)
	lines := strings.Split(string(tasks), "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == task {
			lines = append(lines[:i], lines[i+1:]...)
			fmt.Println("removed task ", task)
			break
		}
	}
	err:=os.WriteFile(taskFile, []byte(strings.Join(lines, "\n")), 0600)
	if(err!=nil){
		fmt.Printf("error in writing to the file %s",err)
	}

}
func List(){
	tasks,err:=os.ReadFile(taskFile)
	if(err!=nil){
		fmt.Println("error in reading the file ",err)
		return
	}
	fmt.Println(string(tasks))

}
