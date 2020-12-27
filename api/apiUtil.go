package api

import (
	"log"
	"regexp"
	"strings"
	"encoding/json"
)

// This is the final object to be passed
type Definition struct {
	GuessWords	[]string	// Guessed words returned by api in case of spelling error
	SpellError	bool		// Indicate whether there's a spelling error

	Word		string		// Word
	Pronun		[]string	// Pronunciation
	Func		string		// Function
	Stems		[]string	// Stemed words
	Defs		[]string	// Definitions and uses
}

// Parse Merriam-Webster api result
func ParseMWJson(data []byte) []Definition {
	var entries []Entry
	defs := make([]Definition, 0, 0)

	// Parse json
	err := json.Unmarshal([]byte(data), &entries)
	if err != nil {
		// Spelling error check
		var words []string
		err = json.Unmarshal([]byte(data), &words)
		if err != nil {
			log.Println("Error unmarshaling data:", string(data), "\nError:", err)
			return nil
		}

		defs = append(defs, Definition {
			GuessWords: words,
			SpellError: true,
		})
		return defs
	}

	// Get all entries and definition
	for _, entry := range entries {
		// Get all pronunciation
		pronun := make([]string, 0, 0)
		for _, p := range entry.Hwi.Prs {
			pronun = append(pronun, p.Mw)
		}

		// Get all definitions
		var definitions []string
		for _, d := range entry.Def {
			retrieveDt(d.Sseq, &definitions)
		}

		def := Definition {
			GuessWords: nil,
			SpellError: false,
			Word: sanitizeId(entry.Meta.Id),
			Pronun: pronun,
			Func: entry.Fl,
			Stems: entry.Meta.Stems,
			Defs: definitions,
		}

		// Add definition to list
		defs = append(defs, def)
	}

	return defs
}

// Recursive function to find dt tags
func retrieveDt(elem interface{}, defs *[]string) {
	// Is a map, find dt
	if m, ok := elem.(map[string]interface{}); ok {
		if v, ok := m["dt"]; ok {
			if l, ok := v.([]interface{}); ok {
				if f, ok := l[0].([]interface{}); ok {
					if s, ok := f[1].(string); ok {
						*defs = append(*defs, sanitizeDt(s))
					}
				}
			}
		}
	} else if m, ok := elem.([]interface{}); ok {
		for _, e := range m {
			retrieveDt(e, defs)
		}
	}
}

// Remove {} markers
func sanitizeDt(s string) string {
	reg := regexp.MustCompile("{[^{}]*}")
	markersIdx := reg.FindAllStringIndex(s, -1)

	// No marker
	if len(markersIdx) == 0 {
		return s
	}

	// Have markers
	lastIdx := 0
	result := ""
	for _, idx := range markersIdx {
		marker := s[idx[0]:idx[1]]
		result += s[lastIdx:idx[0]]

		if i := strings.Index(marker, "|"); i != -1 {
			// Link marker
			i += idx[0]
			nextBarIdx := strings.Index(s[i+1:idx[1]], "|")
			if nextBarIdx != -1 {
				result += s[i+1:nextBarIdx+i+1]
			} else {
				result += s[i+1:idx[1]-1]
			}
		}

		lastIdx = idx[1]
	}

	result += s[lastIdx:]

	return result
}

// Remove optional :# trailing
func sanitizeId(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[:idx]
	}
	return s
}