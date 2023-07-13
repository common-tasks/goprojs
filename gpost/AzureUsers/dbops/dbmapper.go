package dbops

import (
	"encoding/json"
	"io/ioutil"
)

type DBMetadata struct {
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}

func GetDBMetaData(filepath string) (DBMetadata, error) {
	var dbMetadata DBMetadata
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return DBMetadata{}, err
	}

	err = json.Unmarshal(data, &dbMetadata)
	if err != nil {
		return DBMetadata{}, err
	}
	return dbMetadata, nil
}
