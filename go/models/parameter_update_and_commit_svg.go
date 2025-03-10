package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (parameter *Parameter) UpdateAndCommitSVGStage() {

	parameter.gongsvgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(parameter.gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(parameter.gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	for _, shape := range parameter.Shapes {
		if shape.GetIsDisplayed() {
			shape.Draw(parameter.gongsvgStage, layer, parameter)
		}
	}

	parameter.gongsvgStage.Commit()

}
