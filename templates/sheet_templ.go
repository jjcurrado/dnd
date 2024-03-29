// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	util "dnd/utilities"
	"strconv"
)

func Sheet(c util.Character) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"sheet\" id=\"char-sheet\"><input type=\"hidden\" name=\"id\" value=\"character-uuid-goes-here\"><div class=\"sheet-header block\"><div style=\"margin: 5px 0px 10px 5px\"><b>Level 4</b></div><div class=\"header-row\"><div class=\"header-column\"><!--row--><div><input type=\"text\" name=\"name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(c.Name))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"header-label\">Name</div></div><div><input type=\"text\" name=\"class\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(c.Class))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"header-label\">Class/Subclass</div></div></div><div class=\"header-column\"><!--row--><div><input type=\"text\" name=\"background\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(c.Background))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"header-label\">Background</div></div><div class=\"header-item\"><input type=\"text\" name=\"race\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(c.Race))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"header-label\">Race</div></div></div></div></div><div class=\"main\"><div class=\"stats block\"><b>Strength</b><div><div class=\"score\"><input name=\"strength\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Strength)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><b>Dexterity</b><div><div class=\"score\"><input name=\"dexterity\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Dexterity)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><b>Constitution</b><div><div class=\"score\"><input name=\"constitution\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Constitution)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><b>Wisdom</b><div><div class=\"score\"><input name=\"wisdom\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Wisdom)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><b>Intelligence</b><div><div class=\"score\"><input name=\"intelligence\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Intelligence)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><b>Charisma</b><div><div class=\"score\"><input name=\"charisma\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Charisma)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><div class=\"sheet-center\"><div id=\"top-center\" style=\"\n                    display: flex;\n                    flex-direction: row;\n                    justify-content: space-between;\n                \"><div class=\"block\" style=\"width: 28%\"><b>AC</b><br><input type=\"text\" name=\"ac\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.AC)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"block\" style=\"width: 28%\"><b>Initiative</b><br><input type=\"text\" name=\"initiative\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(util.ModifierString(c.Dexterity)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"block\" style=\"width: 28%\"><b>Speed</b><br><input type=\"text\" name=\"speed\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(c.Speed)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"block\" style=\"display: flex; flex-direction: column; row-gap: 10px\"><b>Attacks and Spellcasting</b> <button class=\"weapon block\"><div>Longsword</div><div>+5</div><div>1d10 + 5</div></button> <button class=\"weapon block\"><div>Longsword</div><div>+5</div><div>1d10 + 5</div></button></div><div style=\"\n                    display: flex;\n                    width: 100%;\n                    flex-direction: row;\n                    justify-content: space-between;\n                \"><div style=\"\n                        display: flex;\n                        flex-direction: column;\n                        row-gap: 8px;\n                        width: 46%;\n                    \"><div style=\"\n                            display: flex;\n                            flex-direction: column;\n                            row-gap: 10px;\n                            width: 100%;\n                            padding-bottom: 10px;\n                        \" class=\"block\"><b>Strength</b> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"athletics\" value=\"false\">Athletics</button></div><div style=\"\n                            display: flex;\n                            flex-direction: column;\n                            row-gap: 10px;\n                            width: 100%;\n                            padding-bottom: 10px;\n                        \" class=\"block\"><b>Dexterity</b> <button class=\"skill\"><input checked type=\"checkbox\" class=\"check\" name=\"acrobatics\" value=\"true\">Acrobatics</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"sleight_of_hand\" value=\"false\">Sleight of Hand</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"stealth\" value=\"false\">Stealth</button></div><div style=\"\n                            display: flex;\n                            flex-direction: column;\n                            row-gap: 10px;\n                            width: 100%;\n                            padding-bottom: 10px;\n                        \" class=\"block\"><b>Charisma</b> <button class=\"skill\"><input checked type=\"checkbox\" class=\"check\" name=\"deception\" value=\"true\">Deception</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"intimidation\" value=\"false\">Intimidation</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"performance\" value=\"false\">Performance</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"persuasion\" value=\"false\">Persuasion</button></div></div><div style=\"\n                        display: flex;\n                        flex-direction: column;\n                        row-gap: 8px;\n                        width: 46%;\n                    \"><div style=\"\n                            display: flex;\n                            flex-direction: column;\n                            row-gap: 10px;\n                            width: 100%;\n                            padding-bottom: 10px;\n                        \" class=\"block\"><b>Wisdom</b> <button class=\"skill\"><input checked type=\"checkbox\" style=\"margin-left: 5px; margin-right: 8px\" name=\"animal_handling\" value=\"true\">Animal Handling</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"insight\" value=\"false\">Insight</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"medicine\" value=\"false\">Medicine</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"perception\" value=\"false\">Perception</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"survival\" value=\"false\">Survival</button></div><div style=\"\n                            display: flex;\n                            flex-direction: column;\n                            row-gap: 10px;\n                            width: 100%;\n                            padding-bottom: 10px;\n                        \" class=\"block\"><b>Intelligence</b> <button class=\"skill\"><input checked type=\"checkbox\" style=\"margin-left: 5px; margin-right: 8px\" name=\"arcana\" value=\"true\">Arcana</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"history\" value=\"false\">History</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"investigation\" value=\"false\">Investigation</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"nature\" value=\"false\">Nature</button> <button class=\"skill\"><input type=\"checkbox\" class=\"check\" name=\"religion\" value=\"false\">Religion</button></div></div></div></div><div class=\"sheet-right\"><button class=\"block\">View Personality Traits</button> <button class=\"block\">View Ideals</button> <button class=\"block\">View Bonds</button> <button class=\"block\">View Flaws</button><div class=\"block\" style=\"\n                    display: flex;\n                    flex-direction: column;\n                    row-gap: 10px;\n                    text-align: center;\n                \"><b>Features and Traits</b> <button class=\"block\" style=\"width: 100%; margin-bottom: 4px\">Lucky</button> <button class=\"block\" style=\"width: 100%\">Halfling Nimbleness</button> <button style=\"margin-top: auto\">+ Add a Feature or Trait</button></div></div></div><input class=\"submit-btn\" style=\"width: 100%\" hx-indicator=\"#char-sheet\" hx-post=\"/upload\" hx-target=\"#char-sheet\" type=\"submit\" value=\"Submit\"></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
