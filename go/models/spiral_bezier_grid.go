package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type SpiralBezierGrid struct {
	Name string

	AbstractShape

	SpiralBeziers []*SpiralBezier
}

func (spiralBezierGrid *SpiralBezierGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	for _, sc := range spiralBezierGrid.SpiralBeziers {
		sc.Draw(gongsvgStage, layer, p)
	}
}