package azuread

import (
	"fmt"
	"projects/AzureAD/groups"
	"projects/AzureAD/token"
	"projects/AzureAD/users"
)

func AzureAD() {
	token, err := fetchToken()
	if err != nil {
		fmt.Printf("error in getting token %s", err)
		return
	}
	// fmt.Printf("token is %s", token)

	usrs, err := fetchUser(token)
	if err != nil {
		fmt.Printf("error in getting users %s", err)
	}
	fmt.Printf("users %d", len(usrs.Users))

	grp, err := fetchGroups(token)
	if err != nil {
		fmt.Printf("error in getting groups %s", err)
	}
	fmt.Println("groups ",len(grp.Groups))
}
func fetchToken() (string, error) {
	return token.FetchToken()
}

func fetchUser(token string) (*users.UserResponse, error) {
	return users.FetchUser(token)
}
func fetchGroups(token string) (*groups.GroupResponse, error) {
	return groups.FetchGroups(token)
}
