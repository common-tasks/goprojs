package groups

import (
	"encoding/json"
	"io"
	"net/http"
)

type Group struct {
	ID   string `json:"id"`
	Name string `json:"displayName"`
}
type GroupResponse struct {
	Groups []Group `json:"value"`
}

func FetchGroups(token string) (*GroupResponse, error) {
	req, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/groups", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var grpResponse GroupResponse
	err = json.Unmarshal(body, &grpResponse)
	if err != nil {
		return nil, err
	}
	return &grpResponse, nil
}
