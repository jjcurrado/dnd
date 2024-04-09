package database

import (
	"database/sql"
	"dnd/utilities"
	"fmt"
)

/*
type table interface {
	find(string) (any, error)
	insert(any) (int64, error)
}
*/

type characterTable struct {
	db *sql.DB
}

func newCharacterTable(db *sql.DB) *characterTable {
	return &characterTable{db: db}
}

func (s *characterTable) Find(name string) (utilities.Spell, error) {
	row := s.db.QueryRow(`SELECT name, level, duration, range, description, casting_time, Components, School FROM spells WHERE name=$1;`, name)

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
		return utilities.Spell{}, err
	}

	return utilities.Spell{
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

func (s *characterTable) Insert(spell utilities.Spell) (int64, error) {
	insertStatement := "INSERT INTO spells (name, description, level, school, duration, casting_time, range, components) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	result, err := s.db.Exec(insertStatement, spell.Name, spell.Description, spell.Level, spell.School, spell.Duration, spell.CastingTime, spell.Range, spell.Components)
	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error inserting spell: %v", err)
	}
	return id, nil
}
