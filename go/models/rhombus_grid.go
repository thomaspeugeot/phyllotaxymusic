package models

type RhombusGrid struct {
	Name      string
	Reference *Rhombus

	HideableShape

	Rhombuses []*Rhombus
}
