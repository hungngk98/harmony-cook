package music

import (
	"slices"
	"strings"
)

type Note string

var ALL_NOTES = []Note{"aa", "ab", "bb", "cc", "cd", "dd", "de", "ee", "ff", "fg", "gg", "ga"}
var VALID_NOTE_CHARS = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
var VALID_ACCIDENTAL_CHARS = []byte{'b', '#'}

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

func (note Note) FindScales() []Scale {
	scales := []Scale{}
	for scaleType, formula := range ALL_SCALES {
		for i := range formula {
			distanceToRoot := formula[i]
			root := (slices.Index(ALL_NOTES, note) - int(distanceToRoot)%12 + 12) % 12
			scales = append(scales, Scale{Root: ALL_NOTES[root], ScaleType: scaleType})
		}
	}
	return scales
}

func (note Note) Validate() bool {
	if len(note) > 2 ||
		len(note) == 0 ||
		slices.Index(VALID_NOTE_CHARS, note[0]) == -1 ||
		len(note) == 2 && slices.Index(VALID_ACCIDENTAL_CHARS, note[1]) == -1 {
		return false
	}
	return true
}

func NewNote(s string) Note {
	//validate input
	if len(s) > 2 || len(s) == 0 || slices.Index(VALID_NOTE_CHARS, s[0]) == -1 || len(s) == 2 && slices.Index(VALID_ACCIDENTAL_CHARS, s[1]) == -1 {
		return Note("")
	}

	firstChar := strings.ToLower(s[0:1])

	if len(s) == 1 {
		var note Note = Note(firstChar + firstChar)
		return note
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

	return Note("")
}

func NewNotes(n []string) []Note {
	notes := []Note{}
	for _, v := range n {
		newNote := NewNote(v)
		if newNote != "" {
			notes = append(notes, NewNote(v))
		}
	}
	return notes
}
