package progs

import (
	"fmt"
	"testing"
)

func TestStringProgs(t *testing.T) {
	StringProgs()
}
func TestPalin(t *testing.T) {
	ispalin := Palin("madam")
	fmt.Printf("%v\n", ispalin)
	fmt.Printf("%v\n", Palin("bholab"))
}
func TestAnagram(t *testing.T) {
	fmt.Println(Anagram("kiska","iskak"))
}
