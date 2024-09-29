package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type SpiralCircleGrid struct {
	Name string

	AbstractShape

	SpiralRhombusGrid *SpiralRhombusGrid
}

func (spiralCircleGrid *SpiralCircleGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	for idx, r := range spiralCircleGrid.SpiralRhombusGrid.RhombusGrid.Rhombuses {

		x_s, y_s := r.getCoordinates()
		x_r, y_r := p.convertToCircleSpaceCoords(x_s, y_s)

		sc := new(SpiralCircle)
		sc.Stroke = GenerateColor(idx % len(colors))
		sc.Stroke = gongsvg_models.Black.ToString()
		sc.StrokeOpacity = 0.5
		sc.StrokeWidth = 2

		sc.CenterX = (x_r[0] + x_r[2]) / 2.0
		sc.CenterY = (y_r[0] + y_r[2]) / 2.0
		sc.Draw(gongsvgStage, layer, p)
	}
}
