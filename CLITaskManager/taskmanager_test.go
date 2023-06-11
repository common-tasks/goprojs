package clitaskmanager

import (
	"testing"
)

func TestTaskManager(t *testing.T) {
	TaskManager("add haggu")
	TaskManager("list")
	TaskManager("remove haggu")
}
