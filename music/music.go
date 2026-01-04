package music

import (
	"errors"
	"slices"
	"strings"
)

func FormatChords(chords []Chord) []string {
	result := []string{}
	for _, chord := range chords {
		result = append(result, chord.Format())
	}
	return result
}

func RmDupChords(chords []Chord) []Chord {
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
	chords = RmDupChords(chords)
	result := map[int][]Chord{}
	for i := range chords {
		matchingNoteCount := chords[i].CountMatchingNotes(notes)
		result[matchingNoteCount] = append(result[matchingNoteCount], chords[i])
	}
	return result
}

func NewNote(s string) (Note, error) {
	//validate input
	if len(s) > 2 || len(s) == 0 || slices.Index(VALID_NOTE_CHARS, s[0]) == -1 || len(s) == 2 && slices.Index(VALID_ACCIDENTAL_CHARS, s[1]) == -1 {
		return Note(""), errors.New("invalid input")
	}

	firstChar := strings.ToLower(s[0:1])

	if len(s) == 1 {
		var note Note = Note(firstChar + firstChar)
		return note, nil
	}

	for i := range ALL_NOTES {
		if s[0] == ALL_NOTES[i][0] {
			if s[1] == '#' {
				return ALL_NOTES[(i+1)%12], nil
			}
			if s[1] == 'b' {
				return ALL_NOTES[(i-1+12)%12], nil
			}
		}
	}

	return Note(""), errors.New("invalid input")
}

func NewNotes(n []string) ([]Note, error) {
	notes := []Note{}
	for _, v := range n {
		newNote, err := NewNote(v)
		if err != nil {
			return []Note{}, errors.New("invalid input")
		}
		notes = append(notes, newNote)
	}
	return notes, nil
}
