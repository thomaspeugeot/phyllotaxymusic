package models

import (
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

	SpeedX float64

	Presentation
}

func (movingline *MovingLine) Draw(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *Parameter,
) {

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = movingline.Name
	layer.Lines = append(layer.Lines, line)

	angleRad := DegreesToRadians(movingline.AngleDegree)

	line.X1 = parameter.OriginX + movingline.CenterX
	line.Y1 = parameter.OriginY - movingline.CenterY

	line.X2 = line.X1 + movingline.Length*math.Cos(angleRad)
	line.Y2 = line.Y1 - movingline.Length*math.Sin(angleRad)

	// animateX1 := new(gongsvg_models.Animate).Stage(gongsvgStage)
	// animateX1.AttributeName = "x1"
	// animateX1.Dur = "2s"
	// animateX1.Values = fmt.Sprintf("%d;%d", int(line.X1), int(line.X1+300))
	// animateX1.RepeatCount = `indefinite`

	// line.Animates = append(line.Animates, animateX1)

	movingline.Presentation.CopyTo(&line.Presentation)
}
