package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (parameter *Parameter) addAttribute(measure *m.A_measure) {
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

	// For piano, one for the G clef and one for the F clef
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
	fClef.Line = 4
}
