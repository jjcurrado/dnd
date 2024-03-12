package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func Modifier(score int) int {
	return (score - 10) / 2
}
func ModifierString(score int) string {
	mod := (score - 10) / 2
	s := strconv.Itoa(mod)
	if mod > 0 {
		s = "+" + s
	} else if mod < 0 {
		s = "-" + s
	}

	return s
}

// Returns which ability modifier is used in spell attack rolls
// TODO: probably should be using some sort of enum instead of raw strings
func SpellAbility(character Character) string {
	switch strings.ToLower(character.Class) {
	case "bard", "sorcerer", "warlock", "paladin":
		return "Charisma"
	case "cleric", "druid", "ranger":
		return "Wisdom"
	case "fighter", "wizard", "rogue":
		return "Intelligence"
	default:
		return "Intelligence"
	}
}

// Returns how many spells and cantrips can be prepared by the character (spells, cantrips)
// Prepare for if-else hell
func PreparedSpells(character Character) (int, int) {

	cantrips := 0
	spells := 0

	level := character.Level

	// some classes have this number of cantrips or a scaled up version (+1 or +2)
	// so we will make that our default
	if level < 4 {
		cantrips = 2
	} else if level < 10 {
		cantrips = 3
	} else {
		cantrips = 4
	}

	switch strings.ToLower(character.Class) {

	case "bard":

		if level < 10 {
			spells = level + 3
		} else if level == 11 || level == 12 {
			spells = 15
		} else if level == 14 {
			spells = 18
		} else if level == 15 || level == 16 {
			spells = 19
		} else if level == 17 {
			spells = 20
		} else {
			spells = 22
		}

		// uses default cantrips

	case "cleric":
		spells = Modifier(character.Wisdom) + character.Level
		cantrips += 1

	case "druid":
		spells = Modifier(character.Wisdom) + character.Level

	// Eldritch Knight
	case "fighter":

		if level > 9 {
			cantrips = 3
		} else {
			cantrips = 2
		}

		if level < 4 {
			spells = 3
		} else if level < 7 {
			spells = 4
		} else if level == 7 {
			spells = 5
		} else if level < 10 {
			spells = 6
		} else if level == 10 {
			spells = 7
		} else if level < 13 {
			spells = 8
		} else if level == 13 {
			spells = 9
		} else if level < 16 {
			spells = 10
		} else if level < 19 {
			spells = 11
		} else if level == 19 {
			spells = 12
		} else {
			spells = 13
		}

	case "paladin":
		spells = Modifier(character.Charisma) + character.Level/2
		cantrips = 0

	case "ranger":
		spells = (level+1)/2 + 1
		cantrips = 0

	// Arcane Trickster
	case "rogue":
		if level > 9 {
			cantrips = 4
		} else {
			cantrips = 3
		}

		if level < 4 {
			spells = 3
		} else if level < 7 {
			spells = 4
		} else if level == 7 {
			spells = 5
		} else if level < 10 {
			spells = 6
		} else if level == 10 {
			spells = 7
		} else if level < 13 {
			spells = 8
		} else if level == 13 {
			spells = 9
		} else if level < 16 {
			spells = 10
		} else if level < 19 {
			spells = 11
		} else if level == 19 {
			spells = 12
		} else {
			spells = 13
		}

	case "sorcerer":

		if level < 11 {
			spells = level + 1
		} else if level < 16 {
			spells = (level+1)/2 + 6
		} else {
			spells = 15
		}

		cantrips += 2

	case "warlock":
		if level < 9 {
			spells = level + 1
		} else {
			spells = (level+1)/2 + 5
		}

		// default cantrips
	case "wizard":
		spells = Modifier(character.Intelligence) + level
		cantrips += 1
	}

	return spells, cantrips
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
-----------------------------------------------------------`,
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
	fmt.Println()
}
func PrintSpell(spell Spell) {
	fmt.Printf(
		`-----------------------------------------------------------
%s

Level: %d

Casting Time: %s
Range: %s
Duration: %s
Components: %s
School: %s

%v
-----------------------------------------------------------`,
		spell.Name,
		spell.Level,
		spell.CastingTime,
		spell.Range,
		spell.Duration,
		spell.Components,
		spell.School,
		spell.Description,
	)
}
