package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

// generateNote takes
//
// nbHalfTone: the number of half tone of the note above C0
//
// Pitch is represented as a combination of the step of the diatonic
// scale, the chromatic alteration, and the octave.
func generateNote(nbHalfTone int) *m.Pitch {
	pitch := &m.Pitch{}

	// Which octave are we in?
	octave := nbHalfTone / 12
	// Where are we in the 12-tone chromatic cycle?
	noteIndex := nbHalfTone % 12
	// If nbHalfTone can be negative, you may want to handle negative mod properly.

	// Map our noteIndex (0–11) to diatonic step + optional sharp
	switch noteIndex {
	case 0:
		pitch.Step = m.Enum_Step_C
		pitch.Alter = "" // Natural
	case 1:
		pitch.Step = m.Enum_Step_C
		pitch.Alter = "1" // C#
	case 2:
		pitch.Step = m.Enum_Step_D
		pitch.Alter = ""
	case 3:
		pitch.Step = m.Enum_Step_D
		pitch.Alter = "1" // D#
	case 4:
		pitch.Step = m.Enum_Step_E
		pitch.Alter = ""
	case 5:
		pitch.Step = m.Enum_Step_F
		pitch.Alter = ""
	case 6:
		pitch.Step = m.Enum_Step_F
		pitch.Alter = "1" // F#
	case 7:
		pitch.Step = m.Enum_Step_G
		pitch.Alter = ""
	case 8:
		pitch.Step = m.Enum_Step_G
		pitch.Alter = "1" // G#
	case 9:
		pitch.Step = m.Enum_Step_A
		pitch.Alter = ""
	case 10:
		pitch.Step = m.Enum_Step_A
		pitch.Alter = "1" // A#
	case 11:
		pitch.Step = m.Enum_Step_B
		pitch.Alter = ""
	}

	pitch.Octave = octave
	return pitch
}
