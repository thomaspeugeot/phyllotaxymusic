package models

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (parameter *Parameter) GenerateMusicXMLFile() bool {
	log.Println("Export to Music xml requested")

	var scorePartwise m.Score_partwise

	parameter.generateHeaders(&scorePartwise)

	{
		var part m.A_part
		scorePartwise.Part = append(scorePartwise.Part, &part)

		part.Id = "P1"

		{
			parameter.addMeasure(&part)
		}

	}

	// Marshal the books struct back into XML
	output, err := xml.MarshalIndent(scorePartwise, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return true
	}

	// Write the XML to a new file
	filename := generateTimestampedFilename("export", ".musicxml")
	err = os.WriteFile(
		filepath.Join("musicxml", filename),
		[]byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return true
	}
	return false
}

func (parameter *Parameter) addMeasure(part *m.A_part) {
	var measure m.A_measure
	part.Measure = append(part.Measure, &measure)

	measure.Number = "1"
	measure.Width = "269.19"

	var attributes m.Attributes
	measure.Attributes = append(measure.Attributes, &attributes)

	// Musical notation duration is commonly represented as
	// fractions. The divisions element indicates how many divisions per quarter
	// note are used to indicate a note's duration. For example, if duration = 1
	// and divisions = 2, this is an eighth note duration. Duration and divisions
	// are used directly for generating sound output, so they must be chosen to
	// take tuplets into account. Using a divisions element lets us use just one
	// number to represent a duration for each note in the score, while retaining
	// the full power of a fractional representation. If maximum compatibility with
	// Standard MIDI 1.0 files is important, do not have the divisions value exceed
	// 16383.
	attributes.Divisions = "4"

	// The key type represents a key signature. Both traditional and
	// non-traditional key signatures are supported. The optional number attribute refers
	// to staff numbers. If absent, the key signature applies to all staves in the part.
	// Key signatures appear at the start of each system unless the print-object attribute
	// has been set to "no"
	var key m.Key
	attributes.Key = append(attributes.Key, &key)

	// The fifths type represents the number of flats or sharps in a
	// traditional key signature. Negative numbers are used for flats and positive numbers
	// for sharps, reflecting the key's placement within the circle of fifths (hence the
	// type name).
	if parameter.IsMinor {
		key.Fifths = -3 // for C minor
	} else {
		key.Fifths = 0 // for C major
	}

	// Time Named source named complex type "time"
	// Time signatures are represented by the beats element for the numerator
	// and the beat-type element for the denominator. The symbol attribute is used to
	// indicate common and cut time symbols as well as a single number display. Multiple
	// pairs of beat and beat-type elements are used for composite time signatures with
	// multiple denominators, such as 2/4 + 3/8. A composite such as 3+2/8 requires only
	// one beat/beat-type pair. The print-object attribute allows a time signature to be
	// specified but not printed, as is the case for excerpts from the middle of a score.
	// The value is "yes" if not present. The optional number attribute refers to staff
	// numbers within the part. If absent, the time signature applies to all staves in the
	// part.
	var time m.Time
	attributes.Time = append(attributes.Time, &time)
	time.Symbol = m.Enum_Time_symbol_Common

	// The beats element indicates the number of beats, as found in
	// the numerator of a time signature.
	time.Beats = "4"

	// The beat-type element indicates the beat unit, as found in the
	// denominator of a time signature.
	time.Beat_type = "4"

	// For piano, one for the G clef and one for the G clef
	attributes.Staves = 2

	// Clefs are represented by a combination of sign, line, and
	// clef-octave-change elements. The optional number attribute refers to staff numbers
	// within the part. A value of 1 is assumed if not present. Sometimes clefs are added
	// to the staff in non-standard line positions, either to indicate cue passages, or
	// when there are multiple clefs present simultaneously on one staff. In this
	// situation, the additional attribute is set to "yes" and the line value is ignored.
	// The size attribute is used for clefs where the additional attribute is "yes". It is
	// typically used to indicate cue clefs. Sometimes clefs at the start of a measure need
	// to appear after the barline rather than before, as for cues or for use after a
	// repeated section. The after-barline attribute is set to "yes" in this situation. The
	// attribute is ignored for mid-measure clefs. Clefs appear at the start of each system
	// unless the print-object attribute has been set to "no" or the additional attribute
	// has been set to "yes".
	var gClef m.Clef
	attributes.Clef = append(attributes.Clef, &gClef)
	gClef.Number = 1
	gClef.Sign = "G"
	gClef.Line = 2 // G is on the second line of the first stave

	var fClef m.Clef
	attributes.Clef = append(attributes.Clef, &fClef)
	fClef.Number = 2
	fClef.Sign = "F"
	fClef.Line = 4 // F is on the fourth line of the second stave

	firstVoiceNotes := parameter.FirstVoiceNotes
	notes := firstVoiceNotes.Circles
	for _, noteIn := range notes {
		log.Println(noteIn.Pitch)

		if !noteIn.isKept {
			continue
		}

		var note m.Note
		measure.Note = append(measure.Note, &note)

		pitch := generateNote(noteIn.Pitch)
		note.Pitch = pitch

		note.Default_x = "82.98"
		note.Default_y = "-25"

		note.Duration = "1"
		note.Voice = "1"

		var type_ m.Note_type
		note.Type = &type_

		type_.EnclosedText = "eighth"

		var stem m.Stem
		note.Stem = &stem
		stem.EnclosedText = "down"

		var beam m.Beam
		note.Beam = &beam

		beam.Number = 1
		beam.EnclosedText = "begin"

		// parameter.addNote(&measure, m.Enum_Step_A)
		// parameter.addNote(&measure, m.Enum_Step_B)
	}
}

func (parameter *Parameter) addNote(measure *m.A_measure, step m.Enum_Step) {
	var note m.Note
	measure.Note = append(measure.Note, &note)

	note.Default_x = "82.98"
	note.Default_y = "-25"

	var pitch m.Pitch
	note.Pitch = &pitch

	pitch.Step = step
	pitch.Alter = "-1"
	pitch.Octave = 4

	note.Duration = "1"
	note.Voice = "1"

	var type_ m.Note_type
	note.Type = &type_

	type_.EnclosedText = "eighth"

	var stem m.Stem
	note.Stem = &stem
	stem.EnclosedText = "down"

	var beam m.Beam
	note.Beam = &beam

	beam.Number = 1
	beam.EnclosedText = "begin"
}

// generateTimestampedFilename creates a filename using the current timestamp (YYYYMMDD_HHMMSS).
func generateTimestampedFilename(prefix, extension string) string {
	// Format current date and time up to the second
	timestamp := time.Now().Format("20060102_150405")
	// Construct the filename
	return fmt.Sprintf("%s_%s%s", prefix, timestamp, extension)
}
