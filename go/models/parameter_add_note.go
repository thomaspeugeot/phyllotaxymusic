package models

import (
	"fmt"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (*Parameter) add_note(measure *m.A_measure,
	circleNote *Circle,
	circleNotes []*Circle,
	i int,
	voice int) (note *m.Note) {

	// Compute duration
	duration := 1
	remainingCircles := circleNotes[i+1:]
	for _, _c := range remainingCircles {
		if _c.isKept {
			break
		}
		duration++
	}

	// Compute m_pitch
	m_pitch := generatePitch(circleNote.Pitch + nbHalfToneAboveC0)

	// Build the note and append it to the measure
	note = buildAndAppendNote(measure, duration, m_pitch, circleNote.Pitch, voice)

	remainder := setNoteType(note, duration)

	if remainder > 0 {
		var tieStart m.Tie
		tieStart.Type = m.Enum_Start_stop_Start
		note.Tie = &tieStart

		// Build the note and append it to the measure
		note = buildAndAppendNote(measure, remainder, m_pitch, circleNote.Pitch, voice)

		var tieStop m.Tie
		tieStop.Type = m.Enum_Start_stop_Stop
		note.Tie = &tieStop

		setNoteType(note, remainder)
	}

	return
}

func setNoteType(note *m.Note, duration int) int {
	var type_ m.Note_type
	note.Type = &type_

	var dot *m.Empty_placement
	var remainder int
	type_.EnclosedText, dot, remainder, _ = generateNoteType(duration)

	if dot != nil {
		note.Dot = append(note.Dot, dot)
	} else {
		// log.Println("empty dot", i)
	}
	return remainder
}

func buildAndAppendNote(
	measure *m.A_measure,
	duration int,
	m_pitch *m.Pitch,
	pitch int,
	voice int,
) *m.Note {
	// Create the note and its associated Group_music_data
	var note m.Note
	var gmd m.Group_music_data
	gmd.Note = &note

	// Append the Group_music_data to measure
	measure.Group_music_data = append(measure.Group_music_data, &gmd)

	note.Duration = fmt.Sprintf("%d", duration)

	// Set the voice
	note.Voice = fmt.Sprintf("%d", voice)

	note.Pitch = m_pitch

	if pitch+nbHalfToneAboveC0 >= 48 {
		note.Staff = 1
	} else {
		note.Staff = 2
	}

	return &note
}
