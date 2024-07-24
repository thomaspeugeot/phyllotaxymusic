package models

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// Rhombus is defined by its center, a side length, a direction (with an angle)
// and an opening angle along this direction
type Rhombus struct {
	Name string

	HideableShape

	CenterX, CenterY float64

	SideLength float64

	// Angle of one arbitrary Axis of the Rhombus to the horizontal plan
	// in degress.
	Angle float64

	// InsideAngle is the angle inside the rhombus in the edge that
	// crosses the axis of the rhombus
	InsideAngle float64

	Presentation
}

func (r *Rhombus) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	// Convert angle from degrees to radians for computation
	angleRad := r.Angle * math.Pi / 180
	insideAngleRad := r.InsideAngle * math.Pi / 180

	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)

	// Calculate half of the inside angles to position the vertices
	halfVerticalDiagonal := r.SideLength * sinHalfInsideAngle
	halfHorizontalDiagonal := r.SideLength * cosHalfInsideAngle

	var x_s [4]float64
	var y_s [4]float64

	// Coordinates of the first vertex (using the angle + half of the inside angle)
	x_s[0] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad)
	y_s[0] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad)

	// Coordinates of the second vertex (using the angle - half of the inside angle)
	x_s[1] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad+math.Pi/2.0)
	y_s[1] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad+math.Pi/2.0)

	// Coordinates of the third vertex (opposite of the first vertex)
	x_s[2] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad+math.Pi)
	y_s[2] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad+math.Pi)

	// Coordinates of the fourth vertex (opposite of the second vertex)
	x_s[3] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad+math.Pi*1.5)
	y_s[3] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad+math.Pi*1.5)

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
