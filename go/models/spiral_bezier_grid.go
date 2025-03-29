package models

import gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"

type SpiralBezierGrid struct {
	Name string

	Shape

	SpiralBeziers []*SpiralBezier
}

func (spiralBezierGrid *SpiralBezierGrid) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	for _, sc := range spiralBezierGrid.SpiralBeziers {
		sc.Draw(gongsvgStage, layer, p)
	}
}
