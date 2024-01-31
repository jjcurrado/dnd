package utilities

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func ConnectDB() {

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
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func CreateClient() *openai.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var client = openai.NewClient(os.Getenv("API_KEY"))

	return client
}

func Modifier(score int) int {
	return (score - 10) / 2
}

// TODO: Make this work
// Returns which ability modifier is used in spell attack rolls
func SpellAbility(character Character) string {
	return "Intelligence"
}

// TODO: Make this work
// Returns how many spells and cantrips can be prepared by the character
func PreparedSpells(character Character) (int, int) {

	switch strings.ToLower(character.Class) {

	}

	return 5, 2
}

func PrintCharacter(char Character) {
	subrace := char.Subrace
	if subrace == "" {
		subrace = "None"
	}
	fmt.Printf(
		`-----------------------------------------------------------
		Name: %-40sLevel: %d 
		Race: %-40sClass: %s
		Subrace:%-40sBackground: %s

		Speed: %d
		AC: %d

		STR : %d (%d)
		DEX : %d (%d)
		CON : %d (%d)
		INT : %d (%d)
		WIS : %d (%d)
		CHA : %d (%d) 
-----------------------------------------------------------\n`,
		char.Name, char.Level,
		char.Race, char.Class,
		subrace, char.Background,
		char.Speed,
		char.AC,
		char.Strength, Modifier(char.Strength),
		char.Dexterity, Modifier(char.Dexterity),
		char.Constitution, Modifier(char.Constitution),
		char.Intelligence, Modifier(char.Intelligence),
		char.Wisdom, Modifier(char.Wisdom),
		char.Charisma, Modifier(char.Charisma),
	)
}
func PrintSpell(spell Spell) {
	fmt.Printf(
		`-----------------------------------------------------------
%s
		 
Level: %d

Casting Time: %s
Range: %s
Duration: %s

%v
-----------------------------------------------------------
		`,
		spell.Name,
		spell.Level,
		spell.CastingTime,
		spell.Range,
		spell.Duration,
		spell.Description,
	)
}
