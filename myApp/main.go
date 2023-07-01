package main

import (
	"fmt"
	"myapp/api/handlers"
	"net/http"
)

func main(){
	fmt.Println("starting rest end points")
	http.HandleFunc("/user",handlers.GetAllUsers)
	http.ListenAndServe(":5000",nil)
}