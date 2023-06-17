package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/AzureAD/utils"
	"strings"
)

func FetchToken()(string,error) {
	clientid:=utils.ClientID
	tenantid:=utils.TenantID
	clientSecret:=utils.ClientSecret

	data := strings.NewReader(fmt.Sprintf("client_id=%s&scope=https://graph.microsoft.com/.default&client_secret=%s&grant_type=client_credentials", clientid, clientSecret))
	resp, err := http.Post(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantid), "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Printf("error in getting token %s", err)
		return "", err
	}
	
	defer resp.Body.Close()
	var result utils.AuthResponse
	
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Printf("error in parsing response to json format %s", err)
		return "", err
	}
	return result.AccessToken, nil


}