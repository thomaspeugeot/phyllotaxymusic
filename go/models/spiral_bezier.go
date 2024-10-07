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

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = b.Name
	layer.Lines = append(layer.Lines, line)

	b.Presentation.CopyTo(&line.Presentation)

	line.StrokeOpacity = 0.2

	line.X1 = p.SpiralOriginX + b.StartX
	line.Y1 = p.SpiralOriginY - b.StartY

	line.X2 = p.SpiralOriginX + b.EndX
	line.Y2 = p.SpiralOriginY - b.EndY

	layer.Lines = append(layer.Lines, line)

	lineControl1 := new(gongsvg_models.Line).Stage(gongsvgStage)
	lineControl1.Name = b.Name
	layer.Lines = append(layer.Lines, lineControl1)

	b.Presentation.CopyTo(&lineControl1.Presentation)

	lineControl1.StrokeOpacity = 0.2
	lineControl1.Stroke = gongsvg_models.Green.ToString()

	lineControl1.X1 = p.SpiralOriginX + b.StartX
	lineControl1.Y1 = p.SpiralOriginY - b.StartY

	lineControl1.X2 = p.SpiralOriginX + b.ControlPointStartX
	lineControl1.Y2 = p.SpiralOriginY - b.ControlPointStartY

	layer.Lines = append(layer.Lines, lineControl1)

	lineControl2 := new(gongsvg_models.Line).Stage(gongsvgStage)
	lineControl2.Name = b.Name
	layer.Lines = append(layer.Lines, lineControl2)

	b.Presentation.CopyTo(&lineControl2.Presentation)

	lineControl2.StrokeOpacity = 0.2
	lineControl2.Stroke = gongsvg_models.Blue.ToString()

	lineControl2.X1 = p.SpiralOriginX + b.EndX
	lineControl2.Y1 = p.SpiralOriginY - b.EndY

	lineControl2.X2 = p.SpiralOriginX + b.ControlPointEndX
	lineControl2.Y2 = p.SpiralOriginY - b.ControlPointEndY

	layer.Lines = append(layer.Lines, lineControl2)
}
