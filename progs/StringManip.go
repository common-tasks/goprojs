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
