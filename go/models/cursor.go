package models

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Cursor struct {
	Name string

	// in degrees
	AngleDegree float64
	Length      float64

	CenterX, CenterY float64

	StartX, EndX    float64
	DurationSeconds float64
	IsMoving        bool

	Presentation
}

func (movingline *Cursor) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {
	themeWidth := parameter.RotatedAxis.Length
	_ = themeWidth

	line := new(gongsvg_models.Line)
	line.Name = movingline.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := DegreesToRadians(movingline.AngleDegree)

	line.X1 = parameter.OriginX + movingline.CenterX
	line.Y1 = parameter.OriginY - movingline.CenterY

	line.X2 = line.X1 + movingline.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - movingline.Length*math.Sin(angleRad)

}
