package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func generateNoteType(duration int) (v m.Enum_Note_type_value, dot *m.Empty_placement, remainder int, err error) {

	var powerOfTwo int
	powerOfTwo, remainder = highestPowerTwoOfLessOrEqualToDuration(duration)
	switch powerOfTwo {
	case 1:
		v = m.Enum_Note_type_value_16th
	case 2:
		v = m.Enum_Note_type_value_Eighth
	case 4:
		v = m.Enum_Note_type_value_Quarter
	case 8:
		v = m.Enum_Note_type_value_Half
	case 16:
		v = m.Enum_Note_type_value_Whole
	case 32:
		v = m.Enum_Note_type_value_Breve
	case 64:
		v = m.Enum_Note_type_value_Long
	case 128:
		v = m.Enum_Note_type_value_Maxima
	default:
	}

	if remainder == powerOfTwo/2 && remainder != 0 {
		dot = new(m.Empty_placement)
		remainder = 0
	}

	return
}

// highestPowerTwoOfLessOrEqualToDuration returns the greatest power of 2 <= n.
// If n < 1, returns 0.
func highestPowerTwoOfLessOrEqualToDuration(n int) (powOf2 int, remainder int) {
	if n < 1 {
		return 0, 0
	}
	p := 1
	for p <= n {
		p <<= 1
	}
	// When the loop exits, p has overshot by one power of two,
	// so we shift back to get the highest power of 2 <= n.
	powOf2 = p >> 1
	remainder = n - powOf2
	return
}
