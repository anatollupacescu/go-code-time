package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/guregu/null.v3"
)

type UserSql struct {
	Id    int64       `json:"id"`
	Title null.String `json:"title"`
}

func TestUnmarshalling(t *testing.T) {
	body := []byte(`{ "id": 3 }`)

	var u UserSql

	if err := json.Unmarshal(body, &u); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", u)
}
