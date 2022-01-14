package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"discgolfapi.com/m/models"

	"github.com/lib/pq"
)

const (
	HOST            = "database"
	PORT            = 5432
	DB_PING_RETRIES = 10
)

const ERR_INT_RESP = -1

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

	sleepDuration := time.Second
	for i := 0; i < DB_PING_RETRIES; i++ {
		time.Sleep(sleepDuration)

		err = db.Conn.Ping()
		if err == nil {
			break
		}

		// exponential backoff retry
		sleepDuration *= 2
	}
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

func (db *Database) PutDisc(disc models.Disc) error {
	// use a transaction to rollback if something fails
	transaction, err := db.Conn.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()

	// add measurements
	maxWeightId, err := putMeasurement(transaction, disc.MaxWeight)
	if err != nil {
		return err
	}

	diameterId, err := putMeasurement(transaction, disc.Diameter)
	if err != nil {
		return err
	}

	heightId, err := putMeasurement(transaction, disc.Diameter)
	if err != nil {
		return err
	}

	rimDepthId, err := putMeasurement(transaction, disc.RimDepth)
	if err != nil {
		return err
	}

	discSql := `
	INSERT INTO disc (name, distributor, max_weight_id, diameter_id, height_id, rim_depth_id, speed, glide, turn, fade, stability, primary_use, plastic_grades, link)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
	`
	_, err = transaction.Exec(discSql, disc.Name, disc.Distributor, maxWeightId, diameterId, heightId, rimDepthId, disc.Speed, disc.Glide, disc.Turn, disc.Fade, disc.Stability, disc.PrimaryUse, pq.Array(disc.PlasticGrades), disc.Link)
	if err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func putMeasurement(transaction *sql.Tx, measurement models.Measurement) (int64, error) {
	insertId := ERR_INT_RESP
	err := transaction.QueryRow(
		"INSERT INTO measurement (value, unit) VALUES($1, $2) RETURNING id;",
		measurement.Value, measurement.Unit).Scan(&insertId)
	if err != nil {
		return int64(insertId), err
	}

	return int64(insertId), nil
}
