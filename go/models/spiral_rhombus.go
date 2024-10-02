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

	// not persisted by gong
	X_r0, Y_r0, X_r1, Y_r1, X_r2, Y_r2, X_r3, Y_r3 float64
}

// Draw implements Shape.
func (spiralrhombus *SpiralRhombus) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {
	r := spiralrhombus.Rhombus

	x_r, y_r := spiralrhombus.GenerateCoordinatesFromStruct()
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
}

func (p *Parameter) convertToCircleSpaceCoordsArray(x_s [4]float64, y_s [4]float64) ([4]float64, [4]float64) {
	var x_r [4]float64
	var y_r [4]float64
	for i := range 4 {
		ratio := x_s[i] / p.RotatedAxis.Length
		angle := math.Pi * 2.0 * ratio
		x_r[i] = (p.SpiralInitialRadius + y_s[i]) * math.Cos(angle)
		y_r[i] = (p.SpiralInitialRadius + y_s[i]) * math.Sin(angle)
	}
	return x_r, y_r
}

func (p *Parameter) convertToCircleSpaceCoords(x_s float64, y_s float64) (float64, float64) {
	var x_r float64
	var y_r float64

	ratio := x_s / p.RotatedAxis.Length
	angle := math.Pi * 2.0 * ratio
	x_r = (p.SpiralInitialRadius + y_s) * math.Cos(angle)
	y_r = (p.SpiralInitialRadius + y_s) * math.Sin(angle)

	return x_r, y_r
}

func (sr *SpiralRhombus) GenerateCoordinates(x_r, y_r [4]float64) {
	sr.X_r0 = x_r[0]
	sr.Y_r0 = y_r[0]
	sr.X_r1 = x_r[1]
	sr.Y_r1 = y_r[1]
	sr.X_r2 = x_r[2]
	sr.Y_r2 = y_r[2]
	sr.X_r3 = x_r[3]
	sr.Y_r3 = y_r[3]
}

func (sr *SpiralRhombus) GenerateCoordinatesFromStruct() ([4]float64, [4]float64) {
	x_r := [4]float64{sr.X_r0, sr.X_r1, sr.X_r2, sr.X_r3}
	y_r := [4]float64{sr.Y_r0, sr.Y_r1, sr.Y_r2, sr.Y_r3}
	return x_r, y_r
}
