package svg

import (
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawRhombus(
	gongsvgStage *gongsvg_models.StageStruct,
	axisLayer *gongsvg_models.Layer,
	p *phylotaxymusic_models.Parameter,
	r *phylotaxymusic_models.Rhombus,
) {

	// Calculate half of the side length based on the angle for better accuracy
	halfSideLength := r.SideLength * math.Cos(math.Pi*r.InsideAngle/180.0) // Radians

	// Calculate starting point coordinates based on center, angle, and half side length
	startX := r.CenterX + halfSideLength*math.Cos(math.Pi*r.Angle/180.0)
	startY := r.CenterY + halfSideLength*math.Sin(math.Pi*r.Angle/180.0)

	// Calculate ending point coordinates by shifting based on half side length again
	endX := startX + halfSideLength*math.Cos((math.Pi*r.Angle+90.0)/180.0) // 90 degrees for right side
	endY := startY + halfSideLength*math.Sin((math.Pi*r.Angle+90.0)/180.0)

	line := (&gongsvg_models.Line{
		X1: startX + p.OriginX,
		Y1: p.OriginY - startY,
		X2: endX + p.OriginX,
		Y2: p.OriginY - endY,
		Presentation: gongsvg_models.Presentation{
			Stroke:        gongsvg_models.Black.ToString(),
			StrokeWidth:   r.StrokeWidth,
			StrokeOpacity: 1,
		},
	}).Stage(gongsvgStage)

	axisLayer.Lines = append(axisLayer.Lines, line)

}
