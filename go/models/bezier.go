package models

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Bezier struct {
	Name string
	HideableShape

	StartX, StartY                         float64
	ControlPointStartX, ControlPointStartY float64

	EndX, EndY                         float64
	ControlPointEndX, ControlPointEndY float64

	Presentation
}

// Draw implements Shape.
func (b *Bezier) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, parameter *Parameter) {

	path := new(gongsvg_models.Path).Stage(gongsvgStage)
	path.Name = "Initial Bezier"
	layer.Paths = append(layer.Paths, path)

	b.Presentation.CopyTo(&path.Presentation)

	// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
	path.Definition = fmt.Sprintf("M %0.1f %0.1f C %0.1f %0.1f, %0.1f %0.1f, %0.1f %0.1f",
		parameter.OriginX+b.StartX, parameter.OriginY-b.StartY,
		parameter.OriginX+b.ControlPointStartX, parameter.OriginY-b.ControlPointStartY,
		parameter.OriginX+b.ControlPointEndX, parameter.OriginY-b.ControlPointEndY,
		parameter.OriginX+b.EndX, parameter.OriginY-b.EndY,
	)
}

func (_b *Bezier) move(b *Bezier, x, y float64) {
	_b.StartX = b.StartX + x
	_b.StartY = b.StartY + y
	_b.ControlPointStartX = b.ControlPointStartX + x
	_b.ControlPointStartY = b.ControlPointStartY + y

	_b.EndX = b.EndX + x
	_b.EndY = b.EndY + y
	_b.ControlPointEndX = b.ControlPointEndX + x
	_b.ControlPointEndY = b.ControlPointEndY + y
}
