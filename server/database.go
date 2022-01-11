package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"discgolfapi.com/m/models"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

var Db *Database

func init() {
	db, err := getDatabase(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to init database: %s", err.Error()))
	}
	Db = db
}
func getDatabase(username string, password string, database string) (*Database, error) {
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

func (db *Database) GetDiscs() ([]models.Disc, error) {
	rows, err := db.Conn.Query("SELECT * FROM disc")
	if err != nil {
		return nil, err
	}

	discs := make([]models.Disc, 0)
	for rows.Next() {
		var disc models.Disc

		s := reflect.ValueOf(&disc).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			return nil, err
		}

		discs = append(discs, disc)
	}

	return discs, nil
}
