package handlers

import (
	"encoding/json"
	"fmt"
	"myapp/api/models"
	"net/http"
)

var users =[]models.User{
	{
		Name: "Draupadi",
		LastName: "Murmu",
	},
	{
		Name: "Ramnath",
		LastName: "Kovind",
	},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	if(r.URL.Path=="/user"){
		fmt.Println("request at ",r.URL.Path)
		data,err:=json.Marshal(users)
		if(err!=nil){
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.Write(data)
	}
}