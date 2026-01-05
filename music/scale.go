package music

import (
	"fmt"
	"slices"
	"strings"
)

var ALL_SCALES = map[string][]uint8{
	"major": {0, 2, 4, 5, 7, 9, 11},
	"minor": {0, 2, 3, 5, 7, 8, 10},
}

type Scale struct {
	Root      Note
	ScaleType string
}

func (scale Scale) GetNotes() []Note {
	formula := ALL_SCALES[scale.ScaleType]
	rootIndex := slices.Index(ALL_NOTES, scale.Root)
	result := []Note{}
	for _, halves := range formula {
		note := ALL_NOTES[(rootIndex+int(halves))%12]
		result = append(result, note)
	}
	return result
}

func (scale Scale) CountMatchingNotes(notes []Note) int {
	notesInScale := scale.GetNotes()
	count := 0
	for _, n := range notesInScale {
		if slices.Index(notes, n) > -1 {
			count++
		}
	}
	return count
}

func (scale Scale) Format() string {
	root := strings.ToUpper(string(scale.Root))
	if root[0] == root[1] {
		return fmt.Sprintf("%v %v", root[0:1], scale.ScaleType)
	}

	sharps := true
	notes := scale.GetNotes()
	for i := 1; i < len(notes); i++ {
		if notes[i][0] == notes[i-1][0] {
			sharps = false
		}
	}

	if sharps == true {
		return fmt.Sprintf("%v# %v", root[0:1], scale.ScaleType)
	} else {
		return fmt.Sprintf("%vb %v", root[1:], scale.ScaleType)
	}
}

func RmDupScales(scales []Scale) []Scale {
	uniques := []Scale{}
	for i := range scales {
		isUnique := true
		for _, unique := range uniques {
			if scales[i].Root == unique.Root && scales[i].ScaleType == unique.ScaleType {
				isUnique = false
				break
			}
		}
		if isUnique == true {
			uniques = append(uniques, scales[i])
		}
	}
	return uniques
}

func FormatScales(scales []Scale) []string {
	result := []string{}
	for _, scale := range scales {
		result = append(result, scale.Format())
	}
	return result
}
