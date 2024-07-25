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
