package models

import (
	"github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralRhombus struct {
	Name string

	AbstractShape

	Rhombus *Rhombus
}

// Draw implements Shape.
func (spiralrhombus *SpiralRhombus) Draw(gongsvgStage *models.StageStruct, layer *models.Layer, parameter *Parameter) {
	// panic("unimplemented")
}
