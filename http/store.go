package store

import (
	"database/sql"
	"github.com/lib/pq"
)

type Store struct {
	connection
	connect(connStr String) (error)
}

func NewStore(&Store) (Store, error) {

}

func (*Store) connect() error{
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disabled"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return err

	err := db.ping()
	return err
}

