package models

import (
	"fmt"
	"math"

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

	var x_r [4]float64
	var y_r [4]float64
	for i := range 4 {
		angle := math.Pi * 2.0 * x_s[i] / p.RotatedAxis.Length
		x_r[i] = (p.SpiralInitialRadius + y_s[i]) * math.Cos(angle)
		y_r[i] = (p.SpiralInitialRadius + y_s[i]) * math.Sin(angle)
	}
	for i := range 4 {
		line := (&gongsvg_models.Line{
			Name: fmt.Sprintf("%d", i),
			X1:   x_r[i%4] + p.SpiralOriginX,
			Y1:   p.SpiralOriginY - y_r[i%4],
			X2:   x_r[(i+1)%4] + p.SpiralOriginX,
			Y2:   p.SpiralOriginY - y_r[(i+1)%4],
			Presentation: gongsvg_models.Presentation{
				Stroke:        r.Stroke,
				StrokeWidth:   r.StrokeWidth,
				StrokeOpacity: r.StrokeOpacity,
			},
		}).Stage(gongsvgStage)
		layer.Lines = append(layer.Lines, line)
	}

	for i := range 4 {
		line := (&gongsvg_models.Line{
			X1: x_s[i%4] + p.SpiralOriginX,
			Y1: p.SpiralOriginY - y_s[i%4],
			X2: x_s[(i+1)%4] + p.SpiralOriginX,
			Y2: p.SpiralOriginY - y_s[(i+1)%4],
			Presentation: gongsvg_models.Presentation{
				Stroke:        r.Stroke,
				StrokeWidth:   r.StrokeWidth,
				StrokeOpacity: r.StrokeOpacity,
			},
		}).Stage(gongsvgStage)

		layer.Lines = append(layer.Lines, line)
	}
}
