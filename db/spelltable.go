package database

import (
	"database/sql"
	"dnd/utilities"
	"fmt"
	"log/slog"
)

type spellTable struct {
	db *sql.DB
}

func newSpellTable(db *sql.DB) *spellTable {
	return &spellTable{db: db}
}

const findSpellQuery = `SELECT id, name, level, duration, range, description, casting_time, Components, School FROM spells WHERE name=$1;`

// Find a spell from the spell table with the given name.
func (s *spellTable) Find(name string) (utilities.Spell, error) {
	row := s.db.QueryRow(findSpellQuery, name)
	slog.Info("searching for spell", "name", name)

	var (
		ID          int64
		Name        string
		Level       int
		Duration    string
		Range       string
		Description string
		CastingTime string
		Components  string
		School      string
	)

	err := row.Scan(&ID, &Name, &Level, &Duration, &Range, &Description, &CastingTime, &Components, &School)

	if err != nil {
		return utilities.Spell{}, err
	}
	slog.Info("Found.")
	return utilities.Spell{
		ID:          ID,
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

const insertSpellQuery = "INSERT INTO spells (name, description, level, school, duration, casting_time, range, components) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

// Insert a spell into the spell table
// returns spell ID on success, or 0 on failure
func (s *spellTable) Insert(spell utilities.Spell) (int64, error) {
	slog.Info("inserting into spell table", "name", spell.Name)

	result, err := s.db.Exec(insertSpellQuery, spell.Name, spell.Description, spell.Level, spell.School, spell.Duration, spell.CastingTime, spell.Range, spell.Components)
	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}
	return id, nil
}
