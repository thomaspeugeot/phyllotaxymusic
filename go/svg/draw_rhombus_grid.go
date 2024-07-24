package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawRhombusGrid(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *phylotaxymusic_models.Parameter,
	g *phylotaxymusic_models.RhombusGrid,
) {

	for _, r := range g.Rhombuses {
		drawRhombus(gongsvgStage, layer, p, r)
	}
}
