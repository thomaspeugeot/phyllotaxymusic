package models

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type SpiralCircle struct {
	Circle

	Path string
}

func (spiralCircle *SpiralCircle) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	svgCircle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, svgCircle)

	svgCircle.CX = p.SpiralOriginX + spiralCircle.CenterX
	svgCircle.CY = p.SpiralOriginY - spiralCircle.CenterY
	svgCircle.Radius = p.SideLength / 2.0

	if spiralCircle.HasBespokeRadius {
		svgCircle.Radius = spiralCircle.BespopkeRadius
	}

	// Now, create an SVG Path representing the same circle
	CX := svgCircle.CX
	CY := svgCircle.CY
	R := svgCircle.Radius

	// Starting point at angle 0 (rightmost point of the circle)
	xStart := CX + R
	yStart := CY

	// Build the path data using two arc commands to form a complete circle
	pathData := fmt.Sprintf(
		"M %f %f "+
			"A %f %f 0 1 0 %f %f "+
			"A %f %f 0 1 0 %f %f "+
			"Z",
		xStart, yStart, // Move to the starting point
		R, R, CX-R, CY, // First arc to the leftmost point
		R, R, xStart, yStart) // Second arc back to the starting point

	spiralCircle.Path = pathData

	spiralCircle.Presentation.CopyTo(&svgCircle.Presentation)

	if spiralCircle.ShowName {
		svgText := new(gongsvg_models.Text).Stage(gongsvgStage)
		layer.Texts = append(layer.Texts, svgText)

		svgText.X = p.SpiralOriginX + spiralCircle.CenterX
		svgText.Y = p.SpiralOriginY - spiralCircle.CenterY

		svgText.Name = spiralCircle.Name
		svgText.Content = spiralCircle.Name
		spiralCircle.Presentation.CopyTo(&svgText.Presentation)
	}

}
