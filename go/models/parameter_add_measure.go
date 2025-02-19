package models

import (
	"fmt"
	"slices"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

const nbHalfToneAboveC0 = 24

func (parameter *Parameter) addMeasure(
	part *m.A_part,
	firstVoiceNotes *CircleGrid,
	secondVoiceNotes *CircleGrid,
	measureNb int) {

	shift := parameter.ActualBeatsTemporalShift
	_ = shift
	pitchDiff := parameter.PitchDifference
	_ = pitchDiff
	firstVoiceCircleNotes := firstVoiceNotes.Circles
	secondVoiceCircleNotes := slices.Clone(secondVoiceNotes.Circles)

	// for i, note := range firstVoiceCircleNotes {
	// 	if note.isKept {
	// 		log.Println("rank", i, "isKept")
	// 	}
	// }

	var measure m.A_measure
	part.Measure = append(part.Measure, &measure)

	measure.Number = fmt.Sprintf("%d", measureNb+1)

	measure.Width = "269.19"

	if measureNb == 0 {
		parameter.addAttribute(&measure)
	}

	if measureNb < 2 {
		for i, circleNote := range firstVoiceCircleNotes {

			if !circleNote.isKept {
				continue
			}
			parameter.add_note(&measure, circleNote, firstVoiceCircleNotes, i, 1)
		}
	}

	// backup is used for separating voice 1 and 2
	var group_music_data m.Group_music_data
	measure.Group_music_data = append(measure.Group_music_data, &group_music_data)

	var backup m.Backup

	group_music_data.Backup = &backup

	backup.Duration = fmt.Sprintf("%d", parameter.NbOfBeatsInTheme)

	// for the second voice
	// start with a rest
	if measureNb == 0 {
		for range shift {
			var group_music_data m.Group_music_data
			measure.Group_music_data = append(measure.Group_music_data, &group_music_data)

			var noteRest m.Note
			group_music_data.Note = &noteRest

			noteRest.Voice = "2"
			noteRest.Rest = new(m.Rest)
			noteRest.Rest.Measure = m.Enum_Yes_no_Yes
			noteRest.Duration = "1"
			noteRest.Staff = 1

			var noteType m.Note_type
			noteType.EnclosedText = m.Enum_Note_type_value_16th
			noteRest.Type = &noteType
		}
	}

	_ = secondVoiceCircleNotes
	// secondVoiceCircleNotes = shiftRight(slices.Clone(secondVoiceCircleNotes), -shift)

	// split second in 2
	firstPart := slices.Clone(secondVoiceNotes.Circles[0 : len(secondVoiceCircleNotes)-shift])
	secondPart := slices.Clone(secondVoiceNotes.Circles[len(secondVoiceCircleNotes)-shift:])

	switch measureNb {
	case 0:
		// uses the note from shift to the end
		secondVoiceCircleNotes = firstPart
		// log.Println(measureNb, len(secondVoiceCircleNotes))
	case 1:
		secondVoiceCircleNotes = slices.Clone(secondPart)
		secondVoiceCircleNotes = append(secondVoiceCircleNotes, slices.Clone(firstPart)...)
		// log.Println(measureNb, len(secondVoiceCircleNotes))
	case 2:
		secondVoiceCircleNotes = slices.Clone(secondPart)
		// log.Println(measureNb, len(secondVoiceCircleNotes))
	}
	for i, circleNote := range secondVoiceCircleNotes {

		if !circleNote.isKept {
			continue
		}

		parameter.add_note(&measure, circleNote, secondVoiceCircleNotes, i, 2)
	}

}
