package worker

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func Connect (pgUser, pgPass, pgBase string ) error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", pgUser, pgPass, pgBase))
	if err != nil{
		fmt.Println(err)
	}
	return nil
}