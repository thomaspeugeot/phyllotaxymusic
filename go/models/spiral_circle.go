package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type SpiralCircle struct {
	Circle
}

func (spiralCircle *SpiralCircle) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
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
