package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralAxisGrid struct {
	Name string

	AbstractShape

	SpiralAxises []*SpiralAxis
}

func (spiralAxisGrid *SpiralAxisGrid) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	for _, sa := range spiralAxisGrid.SpiralAxises {
		sa.Draw(gongsvgStage, layer, parameter)
	}
}
