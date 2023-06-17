package users
import
(
	"net/http"
	"io"
	"encoding/json"
)

type User struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}
type UserResponse struct {
	Users []User `json:"value"`
}

func FetchUser(token string) (*UserResponse, error) {
	req, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/users", nil)
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

	var usersResponse UserResponse
	err = json.Unmarshal(body, &usersResponse)
	if err != nil {
		return nil, err
	}

	return &usersResponse, nil
}