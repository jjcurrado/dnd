package app

import (
	util "dnd/utilities"
	"fmt"
	"sync"
)

type DNDservice struct {
	ai     *AIservice
	spells *spellTable
}

func newDNDService(ai *AIservice, spells *spellTable) *DNDservice {
	dnd := &DNDservice{}
	dnd.ai = ai
	dnd.spells = spells
	return dnd
}

func (dnd *DNDservice) createCharacter(prompt string, options util.Options) util.Character {

	res := dnd.ai.sendRequest(prompt)
	char := util.Character{}
	err := dnd.ai.getAIResponse(res, &char)

	if err != nil {
		panic(err)
	}

	util.PrintCharacter(char)

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

		details := make([]util.Spell, len(l.Spells))
		var wg sync.WaitGroup

		for i := 0; i < len(l.Spells); i++ {
			wg.Add(1)
			go func(SpellName string, index int) {
				defer wg.Done()
				spell, err := dnd.spells.find(SpellName)

				if err == nil && spell.Name != "" {
					details[index] = spell
				} else {
					prompt := fmt.Sprintf("Give me the details for the spell %v", spell)
					res := dnd.ai.sendRequest(prompt)
					s := util.Spell{}

					err := dnd.ai.getAIResponse(res, &s)

					if err != nil {
						panic(err)
					}

					dnd.spells.insert(spell)
					details[index] = spell
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
