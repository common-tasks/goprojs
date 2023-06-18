package azuread

import (
	"fmt"
	"testing"
)
func TestFetchToken(t *testing.T) {
	_,err:=fetchToken()
	if(err!=nil){
		fmt.Printf("error is %s",err)
	}
	
}

func TestFetchUser(t *testing.T) {
	token,err:=fetchToken()
	if(err!=nil){
		fmt.Printf("error in getting token %s",err)
	}
	resp,err:=fetchUser(token)
	if(err!=nil){
		fmt.Printf("error in getting user %s",err)
	}
	fmt.Printf("total number of users %d ",len(resp.Users))
	for index, usr := range resp.Users {
		fmt.Println(index,usr)
	}
}
func TestFetchGroup(t *testing.T) {
	token,err:=fetchToken()
	if(err!=nil){
		fmt.Printf("error in getting token %s",err)
	}
	resp,err:=fetchGroups(token)
	if(err!=nil){
		fmt.Printf("error in getting user %s",err)
	}
	fmt.Printf("total number of users %d ",len(resp.Groups))
	for index, grp := range resp.Groups {
		fmt.Println(index,grp)
	}
}

