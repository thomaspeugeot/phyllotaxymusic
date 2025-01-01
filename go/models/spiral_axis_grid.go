package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralLineGrid struct {
	Name string

	Shape

	SpiralLines []*SpiralLine
}

func (spiralLineGrid *SpiralLineGrid) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	for _, sa := range spiralLineGrid.SpiralLines {
		sa.Draw(gongsvgStage, layer, parameter)
	}
}
