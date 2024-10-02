package models

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralBezier struct {
	Name string
	AbstractShape

	StartX, StartY                         float64
	ControlPointStartX, ControlPointStartY float64

	EndX, EndY                         float64
	ControlPointEndX, ControlPointEndY float64

	Presentation
}

// Draw implements Shape.
func (b *SpiralBezier) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	path := new(gongsvg_models.Path).Stage(gongsvgStage)
	path.Name = "Initial SpiralBezier"
	layer.Paths = append(layer.Paths, path)

	b.Presentation.CopyTo(&path.Presentation)

	// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
	path.Definition = fmt.Sprintf("M %0.1f %0.1f C %0.1f %0.1f, %0.1f %0.1f, %0.1f %0.1f",
		p.SpiralOriginX+b.StartX, p.SpiralOriginY-b.StartY,
		p.SpiralOriginX+b.ControlPointStartX, p.SpiralOriginY-b.ControlPointStartY,
		p.SpiralOriginX+b.ControlPointEndX, p.SpiralOriginY-b.ControlPointEndY,
		p.SpiralOriginX+b.EndX, p.SpiralOriginY-b.EndY,
	)
}
