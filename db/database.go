package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// A SQL database and any tables relevant to the application, each with their own
// operations
type DB struct {
	db     *sql.DB
	Spells *spellTable
}

// Instantiates a new SQL database using a .env file then checks the connection
// before returning it
func NewDB() (*DB, error) {
	newDB := &DB{}

	// load in sql db config from env
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open connection and test it
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	spells := newSpellTable(db)
	newDB.db = db
	newDB.Spells = spells

	return newDB, nil
}

// Closes any connections to the SQL database
func (db *DB) Close() {
	db.db.Close()
}
