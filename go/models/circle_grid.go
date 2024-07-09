package models

type CircleGrid struct {
	Name string

	Reference *Circle

	HideableShape

	Circles []*Circle
}
