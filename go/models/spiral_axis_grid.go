package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type SpiralLineGrid struct {
	Name string

	Shape

	SpiralLines []*SpiralLine
}

func (spiralLineGrid *SpiralLineGrid) Draw(gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	for _, sa := range spiralLineGrid.SpiralLines {
		sa.Draw(gongsvgStage, layer, parameter)
	}
}
