package models

type Note struct {
	Name string

	Frequencies []*Freqency

	Start float64

	Duration float64

	Velocity float64
}

type Freqency struct {
	Name string
}
