package music

import (
	"fmt"
	"slices"
	"strings"
)

type Note string

type Chord struct {
	Root   Note
	Symbol string
}

var ALL_NOTES = []Note{"aa", "ab", "bb", "cc", "cd", "dd", "de", "ee", "ff", "fg", "gg", "ga"}

var CHORD_FORMULA_MAP = map[string][]uint8{
	"":     {0, 4, 7},
	"m":    {0, 3, 7},
	"sus2": {0, 2, 7},
	"sus4": {0, 6, 7},
	"7":    {0, 4, 7, 10},
	"m7":   {0, 3, 7, 10},
	"maj7": {0, 4, 7, 11},
	"dim":  {0, 3, 6},
	"dim7": {0, 3, 6, 9},
	"m7b5": {0, 3, 6, 10},
	"aug":  {0, 4, 8},
	"6":    {0, 4, 7, 9},
	"m6":   {0, 3, 7, 9},
}

func (note Note) FindChords() []Chord {
	chords := []Chord{}
	for symbol, formula := range CHORD_FORMULA_MAP {
		for i := range formula {
			distanceToRoot := formula[i]
			root := (slices.Index(ALL_NOTES, note) - int(distanceToRoot)%12 + 12) % 12
			chords = append(chords, Chord{Root: ALL_NOTES[root], Symbol: symbol})
		}
	}
	return chords
}

func (chord Chord) GetNotes() []Note {
	formula := CHORD_FORMULA_MAP[chord.Symbol]
	rootIndex := slices.Index(ALL_NOTES, chord.Root)
	result := []Note{}
	for _, halves := range formula {
		note := ALL_NOTES[(rootIndex+int(halves))%12]
		result = append(result, note)
	}
	return result
}

func (chord Chord) Format() string {
	root := strings.ToUpper(string(chord.Root))
	if root[0] == root[1] {
		return root[0:1] + chord.Symbol
	}
	return fmt.Sprintf(`%v#%v | %vb%v`, root[0:1], chord.Symbol, root[1:2], chord.Symbol)
}

func (chord Chord) CountMatchingNotes(notes []Note) int {
	notesInChord := chord.GetNotes()
	count := 0
	for _, n := range notesInChord {
		if slices.Index(notes, n) > -1 {
			count++
		}
	}
	return count
}

func FormatChords(chords []Chord) []string {
	result := []string{}
	for _, chord := range chords {
		result = append(result, chord.Format())
	}
	return result
}

func RemoveDuplicatedChords(chords []Chord) []Chord {
	uniques := []Chord{}
	for i := 0; i < len(chords); i++ {
		isUnique := true
		for _, unique := range uniques {
			if chords[i].Root == unique.Root && chords[i].Symbol == unique.Symbol {
				isUnique = false
				break
			}
		}
		if isUnique == true {
			uniques = append(uniques, chords[i])
		}
	}
	return uniques
}

func SuggestChords(notes []Note) map[int][]Chord {
	chords := []Chord{}
	for _, n := range notes {
		chords = append(chords, n.FindChords()...)
	}
	chords = RemoveDuplicatedChords(chords)
	result := map[int][]Chord{}
	for i := range chords {
		matchingNoteCount := chords[i].CountMatchingNotes(notes)
		result[matchingNoteCount] = append(result[matchingNoteCount], chords[i])
	}
	return result
}

func NewNote(s string) Note {
	firstChar := strings.ToLower(s[0:1])

	if len(s) == 1 {
		return Note(firstChar + firstChar)
	}

	for i := range ALL_NOTES {
		if s[0] == ALL_NOTES[i][0] {
			if s[1] == '#' {
				return ALL_NOTES[(i+1)%12]
			}
			if s[1] == 'b' {
				return ALL_NOTES[(i-1+12)%12]
			}
		}
	}

	return ""
}

func NewNotes(n []string) []Note {
	notes := []Note{}
	for _, v := range n {
		notes = append(notes, NewNote(v))
	}
	return notes
}
