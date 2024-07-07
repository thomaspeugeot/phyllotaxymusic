package models

type CircleGrid struct {
	Name string

	Reference *Circle
	N         int
	M         int

	HideableShape

	Circles []*Circle
}
