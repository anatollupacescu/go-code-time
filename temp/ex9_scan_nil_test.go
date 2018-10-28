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
)

type EntitledUser struct {
	Id    int64       `json:"id"`
	Title null.String `json:"title"`
}

func TestDbUnmarshalling(t *testing.T) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, user)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS UserTitle(
   ID SERIAL PRIMARY KEY     NOT NULL,
   TITLE           TEXT    )`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO UserTitle(title) VALUES(NULL)`)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT id, title FROM UserTitle`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var u EntitledUser
		err = rows.Scan(&u.Id, &u.Title)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", u)
	}

}
