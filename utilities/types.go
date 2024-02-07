package utilities

type Character struct {
	Name         string
	Class        string
	Subclass     string
	Level        int
	Race         string
	Subrace      string
	Background   string
	AC           int
	Speed        int
	Strength     int
	Intelligence int
	Constitution int
	Wisdom       int
	Charisma     int
	Dexterity    int
	Appearance   string
	Spellcaster  bool
}

type SpellList struct {
	Spells []string
}

type Options struct {
	Spells     bool `default:"false"`
	Appearance bool `default:"false"`
	Feats      bool `default:"false"`
}

type Spell struct {
	Name        string
	Level       int
	Duration    string
	CastingTime string
	Range       string
	Description string
	Components  string
	School      string
}

type SpellResponse struct {
	Spell
}
