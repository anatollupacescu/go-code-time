package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type UserTitle sql.NullString

func (s UserTitle) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON sets *m to a copy of data.
func (m *UserTitle) UnmarshalJSON(data []byte) (e error) {
	if m == nil {
		return errors.New("unmarshalJSON on nil pointer")
	}
	m.String = string(data)
	m.Valid = true
	return
}

type User struct {
	Id    int64     `json:"id"`
	Title UserTitle `json:"title"`
}

func TestMarshaling(t *testing.T) {
	body := []byte(`{ "id": 3, "title": "doc" }`)

	var u User

	if err := json.Unmarshal(body, &u); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", u)
}
