package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (parameter *Parameter) GenerateSvgShape(layer *gongsvg_models.Layer, shape ShapeInterface) {

	if shape.GetIsDisplayed() {
		shape.Draw(parameter.gongsvgStage, layer, parameter)
	}
}

func (parameter *Parameter) UpdateAndCommitSVGStage() {

	parameter.gongsvgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(parameter.gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(parameter.gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	for _, shape := range parameter.Shapes {
		parameter.GenerateSvgShape(layer, shape)
	}

	parameter.gongsvgStage.Commit()

}
