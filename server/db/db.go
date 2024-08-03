package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5433/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// fmt.Println("DB CONNECTION URL : ", connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil

}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
