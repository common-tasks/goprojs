package main

import (
	"fmt"
	"unicode"
)

func compareStrings(s1, s2 string) bool {
	// if len(s1) != len(s2) {
	// 	return false
	// }

	for i := 0; i < len(s1); i++ {
		if !unicode.IsDigit(rune(s2[i])) && s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(compareStrings("apple", "a2le"))
	fmt.Println(compareStrings("apple", "a2l3")) 
	fmt.Println(compareStrings("apple", "bp2le")) 
}


