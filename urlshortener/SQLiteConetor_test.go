package urlshortener

import (
	"fmt"
	"testing"
)

func TestSqLiteConnect(t *testing.T) {
	_,err:=Connect()
	if(err==nil){
	fmt.Println("connected ")
	}

}
func TestCreateTable(t *testing.T){
	CreateURLMappingTable()
}