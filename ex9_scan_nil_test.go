package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"testing"

	"gopkg.in/guregu/null.v3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "pass"

	createTable  = false
	populateData = false
)

type EntitledUser struct {
	Id    int64       `json:"id"`
	Title null.String `json:"title"`
}

func TestDbUnmarshalling(t *testing.T) {

	db := getDbConnection()
	defer db.Close()

	if createTable {
		createDbTable(db)
	}

	if populateData {
		populateDbTable(db)
	}

	rows, err := db.Query(`SELECT id, title FROM UserTitle`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	fmt.Println("Iterating rows: ")
	for rows.Next() {
		var u EntitledUser
		err = rows.Scan(&u.Id, &u.Title)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", u)
	}
	fmt.Println("done iterating rows: ")
}

func populateDbTable(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO UserTitle(title) VALUES (NULL),('Dr.')`)

	if err != nil {
		panic(err)
	}
}

func createDbTable(db *sql.DB) {
	_, err := db.Exec(
`CREATE TABLE IF NOT EXISTS UserTitle(
ID 				SERIAL PRIMARY KEY     NOT NULL,
TITLE           TEXT)`)

	if err != nil {
		panic(err)
	}
}

func getDbConnection() (db *sql.DB) {
	pattern := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	dataSource := fmt.Sprintf(pattern, host, port, user, password, user)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return
}
