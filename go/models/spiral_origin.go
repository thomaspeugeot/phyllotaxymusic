package models

import gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"

type SpiralOrigin struct {
	Name string

	Shape

	Presentation
}

func (spiralOrigin *SpiralOrigin) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {
	{
		line := new(gongsvg_models.Line).Stage(gongsvgStage)
		line.Name = "Spiral Origin"
		layer.Lines = append(layer.Lines, line)

		line.X1 = p.SpiralOriginX + p.OriginCrossWidth/2.0
		line.Y1 = p.SpiralOriginY

		line.X2 = p.SpiralOriginX - p.OriginCrossWidth/2.0
		line.Y2 = p.SpiralOriginY

		p.SpiralOrigin.Presentation.CopyTo(&line.Presentation)
	}
	{
		line := new(gongsvg_models.Line).Stage(gongsvgStage)
		line.Name = "Spiral Origin"
		layer.Lines = append(layer.Lines, line)

		line.X1 = p.SpiralOriginX
		line.Y1 = p.SpiralOriginY + p.OriginCrossWidth/2.0

		line.X2 = p.SpiralOriginX
		line.Y2 = p.SpiralOriginY - p.OriginCrossWidth/2.0

		p.SpiralOrigin.Presentation.CopyTo(&line.Presentation)

	}
}
