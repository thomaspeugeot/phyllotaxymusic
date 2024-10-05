package models

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Axis struct {
	Name string
	AbstractShape

	// in degrees
	Angle  float64
	Length float64

	CenterX, CenterY float64

	Presentation
}

func (axis *Axis) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = axis.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := DegreesToRadians(axis.Angle)

	line.X1 = parameter.OriginX + axis.CenterX
	line.Y1 = parameter.OriginY - axis.CenterY

	line.X2 = line.X1 + axis.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - axis.Length*math.Sin(angleRad)

	axis.Presentation.CopyTo(&line.Presentation)
}
