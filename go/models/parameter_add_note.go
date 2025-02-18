package models

import (
	"fmt"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (*Parameter) add_note(measure *m.A_measure,
	circleNote *Circle,
	circleNotes []*Circle,
	i int,
	voice int) {
	var note m.Note

	var group_music_data m.Group_music_data
	group_music_data.Note = &note

	measure.Group_music_data = append(measure.Group_music_data, &group_music_data)

	pitch := generatePitch(circleNote.Pitch + nbHalfToneAboveC0)
	note.Pitch = pitch

	// let musescore computes x and y, this works fine
	// note.Default_x = "400.98"
	// note.Default_y = "-25"

	note.Duration = "1"
	duration := 1
	remainingCircles := circleNotes[i+1:]

	for _, _c := range remainingCircles {
		if _c.isKept {
			break
		}
		duration += 1
	}
	note.Duration = fmt.Sprintf("%d", duration)

	// not need for the moment.
	note.Voice = fmt.Sprintf("%d", voice)

	// Note_type Named source named complex type "note-type"
	// The note-type type indicates the graphic note type. Values range from
	// 1024th to maxima. The size attribute indicates full, cue, grace-cue, or large size.
	// The default is full for regular notes, grace-cue for notes that contain both grace
	// and cue elements, and cue for notes that contain either a cue or a grace element,
	// but not both.
	//
	// we try to avoid it
	var type_ m.Note_type
	note.Type = &type_

	var dot *m.Empty_placement
	type_.EnclosedText, dot, _ = generateNoteType(duration)

	if dot != nil {
		note.Dot = append(note.Dot, dot)
	} else {
		// log.Println("empty dot", i)
	}
	// Stems can be down, up, none, or double. For down and up stems, the
	// position attributes can be used to specify stem length. The relative values specify
	// the end of the stem relative to the program default. Default values specify an
	// absolute end stem position. Negative values of relative-y that would flip a stem
	// instead of shortening it are ignored. A stem element associated with a rest refers
	// to a stemlet.
	//
	// we try to avoid it
	var stem m.Stem
	note.Stem = &stem
	stem.EnclosedText = "down"

	// Beam values include begin, continue, end, forward hook, and backward
	// hook. Up to eight concurrent beams are available to cover up to 1024th notes. Each
	// beam in a note is represented with a separate beam element, starting with the eighth
	// note beam using a number attribute of 1. Note that the beam number does not
	// distinguish sets of beams that overlap, as it does for slur and other elements.
	// Beaming groups are distinguished by being in different voices and/or the presence or
	// absence of grace and cue elements. Beams that have a begin value can also have a fan
	// attribute to indicate accelerandos and ritardandos using fanned beams. The fan
	// attribute may also be used with a continue value if the fanning direction changes on
	// that note. The value is "none" if not specified. The repeater attribute has been
	// deprecated in MusicXML 3.0. Formerly used for tremolos, it needs to be specified
	// with a "yes" value for each beam using it.
	var beam m.Beam
	note.Beam = &beam

	beam.Number = 1
	beam.EnclosedText = "begin"

	if circleNote.Pitch+nbHalfToneAboveC0 >= 48 {
		note.Staff = 1
	} else {
		note.Staff = 2
	}
}
