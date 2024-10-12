package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralLine struct {
	Name string

	AbstractShape

	StartX, EndX, StartY, EndY float64

	Presentation
}

func (sl *SpiralLine) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = sl.Name
	layer.Lines = append(layer.Lines, line)

	line.X1 = parameter.SpiralOriginX + sl.StartX
	line.Y1 = parameter.SpiralOriginY - sl.StartY

	line.X2 = parameter.SpiralOriginX + sl.EndX
	line.Y2 = parameter.SpiralOriginY - sl.EndY

	sl.Presentation.CopyTo(&line.Presentation)

	if parameter.ShowSpiralBezierConstruct {
		circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
		circle.Name = sl.Name

		circle.CX = line.X1
		circle.CY = line.Y1

		sl.Presentation.CopyTo(&circle.Presentation)

		circle.Radius = 15

		layer.Circles = append(layer.Circles, circle)
	}
}
