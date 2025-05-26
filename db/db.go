package db

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "urlshortner.db")
	if err != nil {
		fmt.Println("couldn't initialize database", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Couldn't establish connection to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println("Connection Opened to Database")
	createTable()
}

func createTable() {
	urlTable := `CREATE TABLE IF NOT EXISTS urls (
    id TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL
)`
	_, err := DB.Exec(urlTable)
	if err != nil {
		log.Fatal("could not create table", err)
	}

}
