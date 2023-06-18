package azuread

import (
	"projects/AzureAD/groups"
	"projects/AzureAD/token"
	"projects/AzureAD/users"
)

func fetchToken() (string, error) {
	return token.FetchToken()
}

func fetchUser(token string) (*users.UserResponse, error) {
	return users.FetchUser(token)
}
func fetchGroups(token string) (*groups.GroupResponse, error) {
	return groups.FetchGroups(token)
}
