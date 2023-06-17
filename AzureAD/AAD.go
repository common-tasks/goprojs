package azuread

import (
	"fmt"
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
}
func fetchToken() (string, error) {
	return token.FetchToken()
}

func fetchUser(token string) (*users.UserResponse, error) {
	return users.FetchUser(token)
}
