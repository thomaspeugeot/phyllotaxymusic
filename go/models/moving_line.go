package models

import (
	"fmt"
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type MovingLine struct {
	Name string
	AbstractShape

	// in degrees
	AngleDegree float64
	Length      float64

	CenterX, CenterY float64

	StartX, EndX    float64
	DurationSeconds float64
	IsMoving        bool

	Presentation
}

func (movingline *MovingLine) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {
	themeWidth := parameter.RotatedAxis.Length

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = movingline.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := DegreesToRadians(movingline.AngleDegree)

	line.X1 = parameter.OriginX + movingline.CenterX
	line.Y1 = parameter.OriginY - movingline.CenterY

	line.X2 = line.X1 + movingline.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - movingline.Length*math.Sin(angleRad)

	if movingline.IsMoving {
		{
			animateX1 := new(gongsvg_models.Animate).Stage(gongsvgStage)
			animateX1.AttributeName = "x1"
			animateX1.Dur = fmt.Sprintf("%0.3fs", movingline.DurationSeconds)
			animateX1.Values = fmt.Sprintf("%f;%f", line.X1, line.X1+themeWidth)
			animateX1.RepeatCount = `indefinite`

			line.Animates = append(line.Animates, animateX1)
		}
		{
			animateX2 := new(gongsvg_models.Animate).Stage(gongsvgStage)
			animateX2.AttributeName = "x2"
			animateX2.Dur = fmt.Sprintf("%0.3fs", movingline.DurationSeconds)
			animateX2.Values = fmt.Sprintf("%f;%f", line.X1, line.X1+themeWidth)
			animateX2.RepeatCount = `indefinite`

			line.Animates = append(line.Animates, animateX2)
		}

	}

	movingline.Presentation.CopyTo(&line.Presentation)
}
