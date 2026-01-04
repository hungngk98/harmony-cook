package music

import "slices"

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

func (note Note) Validate() bool {
	if len(note) > 2 ||
		len(note) == 0 ||
		slices.Index(VALID_NOTE_CHARS, note[0]) == -1 ||
		len(note) == 2 && slices.Index(VALID_ACCIDENTAL_CHARS, note[1]) == -1 {
		return false
	}
	return true
}
