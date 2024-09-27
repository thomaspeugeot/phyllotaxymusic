package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralRhombus struct {
	Name string

	AbstractShape

	Rhombus *Rhombus
}

// Draw implements Shape.
func (spiralrhombus *SpiralRhombus) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {
	r := spiralrhombus.Rhombus

	x_s, y_s := r.getCoordinates()

	for i := range 4 {
		line := (&gongsvg_models.Line{
			X1: x_s[i%4] + p.SpiralCenterX,
			Y1: p.SpiralCenterY - y_s[i%4],
			X2: x_s[(i+1)%4] + p.SpiralCenterX,
			Y2: p.SpiralCenterY - y_s[(i+1)%4],
			Presentation: gongsvg_models.Presentation{
				Stroke:        r.Stroke,
				StrokeWidth:   r.StrokeWidth,
				StrokeOpacity: r.StrokeOpacity,
			},
		}).Stage(gongsvgStage)

		layer.Lines = append(layer.Lines, line)
	}
}
