package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (parameter *Parameter) GenerateSvgShape(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, shape Shape) {

	if shape.GetIsDisplayed() {
		shape.Draw(gongsvgStage, layer, parameter)
	}
}

func (parameter *Parameter) GenerateSvg(gongsvgStage *gongsvg_models.StageStruct) {

	gongsvgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	for _, shape := range parameter.Shapes {
		parameter.GenerateSvgShape(gongsvgStage, layer, shape)
	}

	gongsvgStage.Commit()

}
