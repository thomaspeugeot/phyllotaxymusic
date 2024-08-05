package models

// NoteInfo is the info attached to each note in the theme
type NoteInfo struct {
	Name string

	// IsKept is true is the preceeding note is note continuted during
	// the current note
	// For instance, if there are 4 4th in a measure, and is the second
	// node is not kept, the first note will have a duration of 2
	IsKept bool
}
