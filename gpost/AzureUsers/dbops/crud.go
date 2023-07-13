package dbops

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDB() error {
	filepath := "../utils/metadata.json"
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	} else {
		fmt.Println("File exists")
	}
	dbMetadata, err := GetDBMetaData(filepath)
	if err != nil {
		return fmt.Errorf("error is %s", err)
	}
	connectionStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbMetadata.DBUserName, dbMetadata.DBName, dbMetadata.DBPassword)
	fmt.Println(connectionStr)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return fmt.Errorf("error in connecting to string %s", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error in db ping %s", err)
	}
	return nil
}
