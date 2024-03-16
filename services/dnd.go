package services

import (
	database "dnd/db"
	util "dnd/utilities"
	"fmt"
	"log"
	"sync"
)

type DNDservice struct {
	ai *AIservice
	db *database.DB
}

func NewDNDService(ai *AIservice, db *database.DB) *DNDservice {
	dnd := &DNDservice{}
	dnd.ai = ai
	dnd.db = db
	return dnd
}

func (dnd *DNDservice) CreateCharacter(prompt string, options util.Options) util.Character {
	res := dnd.ai.sendRequest(prompt)
	char := util.Character{}
	err := dnd.ai.getAIResponse(res, &char)
	if err != nil {
		log.Printf("Error received: %s", err.Error())
	}
	// get a spell list
	if char.Spellcaster && options.Spells {
		spells, cantrips := util.PreparedSpells(char)
		prompt := fmt.Sprintf("Create a list of %v spells and %v cantrips for a level %v %v . Only include spells that can be cast by a character of this level and class.", spells, cantrips, char.Level, char.Class)
		res := dnd.ai.sendRequest(prompt)
		l := util.SpellList{}
		err := dnd.ai.getAIResponse(res, &l)
		if err != nil {
			panic(err)
		}

		num_spells := len(l.Spells)
		details := make([]util.Spell, num_spells)
		var wg sync.WaitGroup

		for i := 0; i < num_spells; i++ {
			wg.Add(1)
			go func(SpellName string, index int) {
				defer wg.Done()
				s, err := dnd.db.Spells.Find(SpellName)
				if err == nil && s.Name != "" {
					details[index] = s
				} else {
					prompt := fmt.Sprintf("Give me the details for the spell %v", SpellName)
					res := dnd.ai.sendRequest(prompt)
					s := util.Spell{}
					err := dnd.ai.getAIResponse(res, &s)
					if err != nil {
						panic(err)
					}
					dnd.db.Spells.Insert(s)
					details[index] = s
				}
			}(l.Spells[i], i)
		}
		wg.Wait()
		for _, spell := range details {
			util.PrintSpell(spell)
		}
	}
	return char
}
