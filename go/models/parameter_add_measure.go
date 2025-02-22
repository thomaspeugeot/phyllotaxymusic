package models

import (
	"fmt"
	"slices"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

const nbHalfToneAboveC0 = 24

// addMeasure
//
// when building second measure, one have, if the second part
// starts with unkept beat, one needs the create a note and tie it to
// the last note of first part of the second voice
//
// Therefore, addMeasure output this note for
// the first measure and input it in the second (done with HI :-))
// measure
func (parameter *Parameter) addMeasure(
	part *m.A_part,
	firstVoiceNotes *CircleGrid,
	secondVoiceNotes *CircleGrid,
	measureNb int,
	lastCircleOfFirstPartSecondVoiceIn *Circle,
	lastNoteOfFirstPartSecondVoiceIn *m.Note,
) (lastCircleFirstPartSecondVoiceOut *Circle,
	lastNoteFirstPartSecondVoiceOut *m.Note) {

	shift := parameter.ActualBeatsTemporalShift
	_ = shift
	pitchDiff := parameter.PitchDifference
	_ = pitchDiff
	firstVoiceCircleNotes := firstVoiceNotes.Circles
	secondVoiceCircleNotes := slices.Clone(secondVoiceNotes.Circles)

	//
	// FIRST VOICE NOTES

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

	if !parameter.SecondVoiceNotes.IsDisplayed {
		return
	}

	//
	// SECOND VOICE NOTES
	//

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

	// split second in 2
	// the first part is the beginning of the theme
	// the second part is the remaining
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

		// if the second part
		// starts with unkept beat, one needs the create a note and tie it to
		// the last note of first part of the second voice
		if i == 0 && measureNb == 1 && !circleNote.isKept {
			var tieStart m.Tie
			tieStart.Type = m.Enum_Start_stop_Start
			lastNoteOfFirstPartSecondVoiceIn.Tie = &tieStart

			// we reproduce the note from the same circle
			note := parameter.add_note(&measure, lastCircleOfFirstPartSecondVoiceIn, secondVoiceCircleNotes, i, 2)

			var tieStop m.Tie
			tieStop.Type = m.Enum_Start_stop_Stop
			note.Tie = &tieStop
		}

		if !circleNote.isKept {
			continue
		}

		note := parameter.add_note(&measure, circleNote, secondVoiceCircleNotes, i, 2)

		lastCircleFirstPartSecondVoiceOut = circleNote
		lastNoteFirstPartSecondVoiceOut = note

	}

	return
}
