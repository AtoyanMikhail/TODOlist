package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

var migrationsUrl = "file:///app/db/migrations"

type Database struct {
	db *sql.DB
}

func New(databaseUrl string) (*Database, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	m, err := migrate.New(migrationsUrl, databaseUrl)
	if err != nil {
        return nil, err
    }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange { 
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
