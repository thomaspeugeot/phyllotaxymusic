package models

type Circle struct {
	Name string

	HideableShape
	CenterX, CenterY float64

	Presentation
}
