package models

import (
	"fmt"
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Bezier struct {
	Name string
	AbstractShape

	StartX, StartY                         float64
	ControlPointStartX, ControlPointStartY float64

	EndX, EndY                         float64
	ControlPointEndX, ControlPointEndY float64

	Presentation
}

// Draw implements Shape.
func (b *Bezier) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	path := new(gongsvg_models.Path).Stage(gongsvgStage)
	path.Name = "Initial Bezier"
	layer.Paths = append(layer.Paths, path)

	b.Presentation.CopyTo(&path.Presentation)

	// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
	path.Definition = fmt.Sprintf("M %0.1f %0.1f C %0.1f %0.1f, %0.1f %0.1f, %0.1f %0.1f",
		p.OriginX+b.StartX, p.OriginY-b.StartY,
		p.OriginX+b.ControlPointStartX, p.OriginY-b.ControlPointStartY,
		p.OriginX+b.ControlPointEndX, p.OriginY-b.ControlPointEndY,
		p.OriginX+b.EndX, p.OriginY-b.EndY,
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

// This function, ComputeYFromX, is designed to compute the Y-coordinate corresponding to
// a given X-coordinate on a cubic Bezier curve. The Bezier curve is defined by a starting point,
// an ending point, and two control points.
func (b *Bezier) ComputeYFromX(x float64) (float64, error) {
	const tolerance = 1e-6
	const maxIterations = 100

	// Function to compute the X value at a given t
	bezierX := func(t float64) float64 {
		return math.Pow(1-t, 3)*b.StartX +
			3*math.Pow(1-t, 2)*t*b.ControlPointStartX +
			3*(1-t)*math.Pow(t, 2)*b.ControlPointEndX +
			math.Pow(t, 3)*b.EndX
	}

	// Derivative of the Bezier X function with respect to t
	bezierXPrime := func(t float64) float64 {
		return -3*math.Pow(1-t, 2)*b.StartX +
			3*(math.Pow(1-t, 2)-2*(1-t)*t)*b.ControlPointStartX +
			3*((1-t)*2*t-math.Pow(t, 2))*b.ControlPointEndX +
			3*math.Pow(t, 2)*b.EndX
	}

	// Newton-Raphson method to find the t for the given x
	t := 0.5 // Initial guess
	for i := 0; i < maxIterations; i++ {
		xAtT := bezierX(t)
		xPrimeAtT := bezierXPrime(t)
		deltaT := (xAtT - x) / xPrimeAtT
		t -= deltaT
		if math.Abs(deltaT) < tolerance {
			break
		}
	}

	// Clamp t to the range [0, 1]
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	// Function to compute the Y value at a given t
	bezierY := func(t float64) float64 {
		return math.Pow(1-t, 3)*b.StartY +
			3*math.Pow(1-t, 2)*t*b.ControlPointStartY +
			3*(1-t)*math.Pow(t, 2)*b.ControlPointEndY +
			math.Pow(t, 3)*b.EndY
	}

	y := bezierY(t)
	return y, nil
}
