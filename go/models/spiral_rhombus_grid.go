package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type SpiralRhombusGrid struct {
	Name string

	AbstractShape

	RhombusGrid *RhombusGrid
}

func (spiralRhombusGrid *SpiralRhombusGrid) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	for idx, r := range spiralRhombusGrid.RhombusGrid.Rhombuses {
		spiralRhombus := new(SpiralRhombus)

		r2 := r.Copy()

		r2.Stroke = GenerateColor(idx % len(colors))
		r2.StrokeOpacity = 0.5

		spiralRhombus.Rhombus = &r2

		spiralRhombus.Draw(gongsvgStage, layer, p)
	}

}
