package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"myapp/api/models"
	"net/http"
	"os"
)

var users = []models.User{
	{
		Name:     "Draupadi",
		LastName: "Murmu",
	},
	{
		Name:     "Ramnath",
		LastName: "Kovind",
	},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/user" {
		fmt.Println("request at ", r.URL.Path)

		writeSomething(w)

		data, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		readSomething(w)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func readSomething(w http.ResponseWriter) {
	f, err := os.Open("/data/hello.txt")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	d, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Printf("content is %s", string(d))
}

func writeSomething(w http.ResponseWriter) {
	content := []byte("hello world")
	f, err := os.OpenFile("/data/hello.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	n, err := f.Write(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Printf("number of bytes written %d\n", n)
}
