package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type SpiralOrigin struct {
	Name string

	AbstractShape

	Presentation
}

var OriginCrossWidth = 20.0

func (spiralOrigin *SpiralOrigin) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {
	{
		line := new(gongsvg_models.Line).Stage(gongsvgStage)
		line.Name = "Spiral Origin"
		layer.Lines = append(layer.Lines, line)

		line.X1 = parameter.SpiralOriginX + OriginCrossWidth/2.0
		line.Y1 = parameter.SpiralOriginY

		line.X2 = parameter.SpiralOriginX - OriginCrossWidth/2.0
		line.Y2 = parameter.SpiralOriginY

		parameter.SpiralOrigin.Presentation.CopyTo(&line.Presentation)
	}
	{
		line := new(gongsvg_models.Line).Stage(gongsvgStage)
		line.Name = "Spiral Origin"
		layer.Lines = append(layer.Lines, line)

		line.X1 = parameter.SpiralOriginX
		line.Y1 = parameter.SpiralOriginY + OriginCrossWidth/2.0

		line.X2 = parameter.SpiralOriginX
		line.Y2 = parameter.SpiralOriginY - OriginCrossWidth/2.0

		parameter.SpiralOrigin.Presentation.CopyTo(&line.Presentation)

	}
}
