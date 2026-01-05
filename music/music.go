package music

func SuggestTones(notes []Note) map[int][]Scale {
	scales := []Scale{}
	for _, n := range notes {
		scales = append(scales, n.FindScales()...)
	}
	scales = RmDupScales(scales)
	result := map[int][]Scale{}
	for i := range scales {
		if scales[i].ScaleType == "major" {
			matchingNoteCount := scales[i].CountMatchingNotes(notes)
			result[matchingNoteCount] = append(result[matchingNoteCount], scales[i])
		}
	}
	return result
}
