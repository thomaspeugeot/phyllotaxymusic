package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type SpiralRhombusGrid struct {
	Name string

	AbstractShape

	RhombusGrid *RhombusGrid
}

func (spiralRhombusGrid *SpiralRhombusGrid) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	for _, r := range spiralRhombusGrid.RhombusGrid.Rhombuses {
		spiralRhombus := new(SpiralRhombus)
		spiralRhombus.Rhombus = r

		spiralRhombus.Draw(gongsvgStage, layer, p)
	}

}
