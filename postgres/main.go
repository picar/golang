package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres database=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	stmt := `create table if not exists users (
		id serial primary key,
		email varchar(64),
		created_at timestamp
	)`

	{
		r, err := db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r)
	}
}
