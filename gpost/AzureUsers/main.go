package main

import (
	"fmt"
	"gpost/AzureUsers/dbops"
)

func main() {
	fmt.Println("common postgreSQL operations")
	filepath := "./utils/metadata.json"
	dbmetadata, err := dbops.GetDBMetaData(filepath)
	if err != nil {
		fmt.Printf("error is %s\n", err)
	}
	fmt.Println(dbmetadata.DBPassword, dbmetadata.DBPort, dbmetadata.DBUserName)
}
