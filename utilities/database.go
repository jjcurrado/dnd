package utilities

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBClient = ConnectDB()

func ConnectDB() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db
}

func FindSpell(name string) (Spell, error) {
	fmt.Printf("finding spell: %s\n", name)

	row := DBClient.QueryRow(`SELECT name, level, duration, range, description, casting_time FROM spells WHERE name=$1`, name)

	var (
		Name        string
		Level       int
		Duration    string
		Range       string
		Description string
		CastingTime string
	)

	err := row.Scan(&Name, &Level, &Duration, &Range, &Description, &CastingTime)

	if err != nil {
		return Spell{}, err
	}
	return Spell{
		Name:        Name,
		Level:       Level,
		Duration:    Duration,
		Range:       Range,
		Description: Description,
		CastingTime: CastingTime,
	}, nil
}
