package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/todolist?sslmode=disable")
	if err != nil {
		return nil, err
	}
	log.Print("Connected to db successfully")
	return &Database{db: db}, nil
}

func (db *Database) Close() error {
	return db.db.Close()
}

func (db *Database) GetDb() *sql.DB {
	return db.db
}
