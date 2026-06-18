package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	HOST = "postgres"
	PORT = "8080"
)

type Database struct {
	Conn *sql.DB
}

func InitDB(user, password, database string) (Database, error) {
	db := Database{}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", HOST, PORT, user, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")

	return db, nil
}
