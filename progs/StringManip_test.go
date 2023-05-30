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
func TestCountCharOccurence(t *testing.T) {
	fmt.Println(CountCharOccurence("hellohallo",'l'))
	fmt.Println(CountCharOccurence01("hellohallo",'l'))
}
func TestFirstNonRepeat(t *testing.T) {
	res,err:=FirstNonRepeatingChar("aaaa")
	if(err==nil){
	fmt.Printf("%v\n",string(res))
	}else{
		fmt.Printf("%s",err)
	}
}
