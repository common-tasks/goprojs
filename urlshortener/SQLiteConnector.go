package urlshortener

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()")

func SqLiteConnect() {
	fmt.Println("connecting to sqlIte")

}
func RandSeq(n int) string {
	rand.Seed(time.Now().UnixMicro())
	randomsrting := make([]rune, n)
	for i := range randomsrting {
		randomsrting[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomsrting)
}
func Connect() (*sql.DB, error) {
	var sqldb *sql.DB
	sqldb, err := sql.Open("sqlite3", "./urlShortener.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return sqldb, nil

}
func CreateURLMappingTable() bool {
	sqldb, err := Connect()
	if err != nil {
		fmt.Printf("error in connecting to db %s\n", err)
		return false
	}

	creatTableStatement := `
	CREATE TABLE IF NOT EXISTS url_mapping (
		short_url TEXT NOT NULL PRIMARY KEY,
		original_url TEXT NOT NULL
	);
`
	defer sqldb.Close()
	result, err := sqldb.Exec(creatTableStatement)
	if err != nil {
		fmt.Printf("error in execute query %s/n", err)
		return false
	}
	fmt.Println(result)

	return true
}
