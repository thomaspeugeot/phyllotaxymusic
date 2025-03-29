package models

import gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"

type SpiralCircleGrid struct {
	Name string

	Shape

	SpiralRhombusGrid *SpiralRhombusGrid

	SpiralCircles []*SpiralCircle
}

func (spiralCircleGrid *SpiralCircleGrid) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	for _, sc := range spiralCircleGrid.SpiralCircles {
		sc.Draw(gongsvgStage, layer, p)
	}
}
