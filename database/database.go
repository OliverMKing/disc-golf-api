package database

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"discgolfapi.com/m/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST       = "database"
	PORT       = 5432
	DB_RETRIES = 10
)

var ErrMaxDbRetries = fmt.Errorf("Could not connect to database after %d tries", DB_RETRIES)

type Database struct {
	Conn *gorm.DB
}

var Db *Database

func init() {
	db, err := getDatabase(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to init database: %s", err.Error()))
	}

	if err = db.Migrate(); err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to migrate database: %s", err.Error()))
	}

	Db = db
}

func getDatabase(username string, password string, database string) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		HOST, PORT, username, password, database)

	sleepDuration := time.Second
	for i := 0; i < DB_RETRIES; i++ {
		// exponential backoff retry
		time.Sleep(sleepDuration)
		sleepDuration *= 2

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Info().Msg(fmt.Sprintf("Failed to open database connection: %s", err.Error()))
			continue
		}

		db, err := conn.DB()
		if err != nil {
			log.Info().Msg(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Info().Msg(fmt.Sprintf("Failed to ping database: %s", err.Error()))
			continue
		}

		return &Database{Conn: conn}, nil
	}

	return nil, ErrMaxDbRetries

}

func (db *Database) Migrate() error {
	log.Info().Msg(fmt.Sprintf("Migrating %s model", reflect.TypeOf(&models.Disc{}).Name()))
	if err := db.Conn.AutoMigrate(&models.Disc{}); err != nil {
		return err
	}

	return nil
}

func (db *Database) GetDiscs() ([]models.Disc, error) {
	var discs []models.Disc
	if err := db.Conn.Find(&discs).Error; err != nil {
		return nil, err
	}

	return discs, nil
}

func (db *Database) PutDisc(disc *models.Disc) error {
	// check if disc already in database
	var existingDisc models.Disc
	err := db.Conn.First(&existingDisc, "name = ? AND distributor = ?", disc.Name, disc.Distributor).Error

	// create if not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.Conn.Create(&disc).Error; err != nil {
			return err
		}

		return nil
	}

	if err != nil {
		return err
	}

	// update if found
	if err = db.Conn.Model(existingDisc).Updates(disc).Error; err != nil {
		return err
	}

	return nil
}
