package server

import (
	"database/sql"
	"fmt"
)

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func GetDatabase(username string, password string, database string) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db := &Database{Conn: conn}
	err = db.Conn.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
