package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) UpdateAndCommitSVGStage() {

	stager.svgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(stager.svgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(stager.svgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	for _, shape := range stager.parameter.Shapes {
		if shape.GetIsDisplayed() {
			shape.Draw(stager.svgStage, layer, stager.parameter)
		}
	}

	stager.svgStage.Commit()
}
