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
func Anagram(word1 string, word2 string) bool {
    // if lengths are not equal return false
	if len(word1) != len(word2) {
		return false
	}
	//
	charCount := make(map[rune]int)
	for _, c := range word1 {
		charCount[c]++
	}

	for _, character := range word2 {
		count, charExists := charCount[character]
		if !charExists {
			return false
		}
		charCount[character] = count - 1
		if charCount[character] == 0 {
			delete(charCount, character)
		}
	}

	return (len(charCount) == 0)
}
