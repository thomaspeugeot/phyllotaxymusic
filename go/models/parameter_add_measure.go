package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

const nbHalfToneAboveC0 = 24

func (parameter *Parameter) addMeasure(
	part *m.A_part,
	firstVoiceNotes *CircleGrid,
	secondVoiceNotes *CircleGrid,
	isFirstMeasure bool) {
	var measure m.A_measure
	part.Measure = append(part.Measure, &measure)

	if isFirstMeasure {
		measure.Number = "1"
	} else {
		measure.Number = "2"
	}
	measure.Width = "269.19"

	if isFirstMeasure {
		parameter.addAttribute(&measure)
	}

	circleNotes := firstVoiceNotes.Circles
	for i, circleNote := range circleNotes {

		if !circleNote.isKept {
			continue
		}

		parameter.add_note(&measure, circleNote, circleNotes, i)
	}

	// // backup is used for separating voice 1 and 2
	// var backup m.Backup
	// measure.Backup = append(measure.Backup, &backup)

	// backup.Duration = fmt.Sprintf("%d", parameter.NbOfBeatsInTheme)

	// // for the second voice
	// // start with a rest
	// for range parameter.ActualBeatsTemporalShift {
	// 	var noteRest m.Note
	// 	measure.Note = append(measure.Note, &noteRest)

	// 	noteRest.Voice = "2"
	// 	noteRest.Rest = new(m.Rest)
	// 	noteRest.Rest.Measure = m.Enum_Yes_no_Yes
	// 	noteRest.Duration = "1"
	// 	noteRest.Staff = 1

	// 	var noteType m.Note_type
	// 	noteType.EnclosedText = m.Enum_Note_type_value_16th
	// 	noteRest.Type = &noteType
	// }

}
