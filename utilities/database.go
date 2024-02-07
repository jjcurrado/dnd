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

func InsertSpell(spell Spell) (int64, error) {
	insertStatement := "INSERT INTO spells (name, description, level, school, duration, casting_time, range, components) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	result, err := DBClient.Exec(insertStatement, spell.Name, spell.Description, spell.Level, spell.School, spell.Duration, spell.CastingTime, spell.Range, spell.Components)

	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}

	fmt.Printf("Inserted spell %s!", spell.Name)
	return id, nil
}

func FindSpell(name string) (Spell, error) {
	fmt.Printf("finding spell: %s\n", name)

	row := DBClient.QueryRow(`SELECT name, level, duration, range, description, casting_time, Components, School FROM spells WHERE name=$1;`, name)
	var (
		Name        string
		Level       int
		Duration    string
		Range       string
		Description string
		CastingTime string
		Components  string
		School      string
	)

	err := row.Scan(&Name, &Level, &Duration, &Range, &Description, &CastingTime, &Components, &School)

	if err != nil {
		return Spell{}, err
	}

	fmt.Printf("Found spell %s!\n", Name)
	return Spell{
		Name:        Name,
		Level:       Level,
		Duration:    Duration,
		Range:       Range,
		Description: Description,
		CastingTime: CastingTime,
		Components:  Components,
		School:      School,
	}, nil
}
