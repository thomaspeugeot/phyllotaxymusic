package models

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// Rhombus is defined by its center, a side length, a direction (with an angle)
// and an opening angle along this direction
type Rhombus struct {
	Name string

	Shape

	CenterX, CenterY float64

	SideLength float64

	// AngleDegree of one arbitrary Axis of the Rhombus to the horizontal plan
	// in degress.
	AngleDegree float64

	// InsideAngle is the angle inside the rhombus in the edge that
	// crosses the axis of the rhombus
	InsideAngle float64

	Presentation
}

func (r *Rhombus) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	// Convert angle from degrees to radians for computation
	// Calculate half of the inside angles to position the vertices
	// Coordinates of the first vertex (using the angle + half of the inside angle)
	// Coordinates of the second vertex (using the angle - half of the inside angle)
	// Coordinates of the third vertex (opposite of the first vertex)
	// Coordinates of the fourth vertex (opposite of the second vertex)
	x_s, y_s := r.getCoordinates()

	for i := range 4 {
		line := (&gongsvg_models.Line{
			X1: x_s[i%4] + p.OriginX,
			Y1: p.OriginY - y_s[i%4],
			X2: x_s[(i+1)%4] + p.OriginX,
			Y2: p.OriginY - y_s[(i+1)%4],
			Presentation: gongsvg_models.Presentation{
				Stroke:        r.Stroke,
				StrokeWidth:   r.StrokeWidth,
				StrokeOpacity: r.StrokeOpacity,
			},
		}).Stage(gongsvgStage)

		layer.Lines = append(layer.Lines, line)
	}

}

func (r *Rhombus) getCoordinates() ([4]float64, [4]float64) {
	angleRad := r.AngleDegree * math.Pi / 180
	insideAngleRad := r.InsideAngle * math.Pi / 180

	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)

	halfVerticalDiagonal := r.SideLength * sinHalfInsideAngle
	halfHorizontalDiagonal := r.SideLength * cosHalfInsideAngle

	var x_s [4]float64
	var y_s [4]float64

	x_s[0] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad)
	y_s[0] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad)

	x_s[1] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad+math.Pi/2.0)
	y_s[1] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad+math.Pi/2.0)

	x_s[2] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad+math.Pi)
	y_s[2] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad+math.Pi)

	x_s[3] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad+math.Pi*1.5)
	y_s[3] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad+math.Pi*1.5)
	return x_s, y_s
}

func (r Rhombus) Copy() Rhombus {
	return Rhombus{
		Name:         r.Name,
		Shape:        r.Shape, // Assuming is a copyable struct
		CenterX:      r.CenterX,
		CenterY:      r.CenterY,
		SideLength:   r.SideLength,
		AngleDegree:  r.AngleDegree,
		InsideAngle:  r.InsideAngle,
		Presentation: r.Presentation, // Assuming Presentation is a copyable struct
	}
}
