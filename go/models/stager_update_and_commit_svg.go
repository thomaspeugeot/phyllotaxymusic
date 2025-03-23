package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (parameter *Parameter) UpdateAndCommitSVGStage() {

	parameter.svgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(parameter.svgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(parameter.svgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	for _, shape := range parameter.Shapes {
		if shape.GetIsDisplayed() {
			shape.Draw(parameter.svgStage, layer, parameter)
		}
	}

	parameter.svgStage.Commit()

}
