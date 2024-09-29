package models

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralAxis struct {
	Name string
	AbstractShape

	Angle  float64
	Length float64

	CenterX, CenterY float64

	Presentation
}

func (axis *SpiralAxis) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = axis.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := axis.Angle * math.Pi / 180

	line.X1 = parameter.SpiralOriginX + axis.CenterX
	line.Y1 = parameter.SpiralOriginY - axis.CenterY

	line.X2 = line.X1 + axis.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - axis.Length*math.Sin(angleRad)

	axis.Presentation.CopyTo(&line.Presentation)
}
