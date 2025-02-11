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

	var measure m.A_measure
	part.Measure = append(part.Measure, &measure)

	measure.Number = fmt.Sprintf("%d", measureNb+1)

	measure.Width = "269.19"

	if measureNb == 0 {
		parameter.addAttribute(&measure)
	}

	circleNotes := firstVoiceNotes.Circles
	if measureNb < 2 {
		for i, circleNote := range circleNotes {

			if !circleNote.isKept {
				continue
			}
			parameter.add_note(&measure, circleNote, circleNotes, i, 1)
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
		for range parameter.ActualBeatsTemporalShift {
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

	switch measureNb {
	case 0:
		circleNotes = slices.Clone(secondVoiceNotes.Circles[parameter.ActualBeatsTemporalShift:])
	case 1:
		circleNotes = slices.Clone(secondVoiceNotes.Circles[:parameter.ActualBeatsTemporalShift])
		circleNotes = append(circleNotes, slices.Clone(secondVoiceNotes.Circles[parameter.ActualBeatsTemporalShift:])...)
	case 2:
		circleNotes = slices.Clone(secondVoiceNotes.Circles[:parameter.ActualBeatsTemporalShift])
	}
	for i, circleNote := range circleNotes {

		if !circleNote.isKept {
			continue
		}

		parameter.add_note(&measure, circleNote, circleNotes, i, 2)
	}

}
