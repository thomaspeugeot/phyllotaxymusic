package models

type Axis struct {
	Name string
	HideableShape

	Angle  float64
	Length float64

	Presentation
}
