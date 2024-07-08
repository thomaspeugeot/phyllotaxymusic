package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawCircleGrid(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *phylotaxymusic_models.Parameter,
	g *phylotaxymusic_models.CircleGrid,
) {

	for _, c := range g.Circles {
		drawCircle(gongsvgStage, layer, p, c)
	}

}
