package progs

import (
	"fmt"
)

func StringProgs() {
	fmt.Println(revString("Hola"))
}
func revString(s string) string {
	r := []rune(s)
	start := 0
	end := len(s) - 1
	for start < end {
		temp := r[start]
		r[start] = r[end]
		r[end] = temp
		start++
		end--

	}

	return string(r)
}
func Palin(s string) bool {
	ispal := true
	run := []rune(s)

	endIndex := len(s) - 1

	for startIndex := 0; startIndex < endIndex; startIndex++ {
		if run[startIndex] != run[endIndex] {
			ispal = false
			break
		}
		endIndex--
	}
	return ispal
}
