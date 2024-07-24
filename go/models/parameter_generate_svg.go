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

	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.HorizontalAxis)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.VerticalAxis)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.InitialRhombus)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.InitialCircle)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.InitialRhombusGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.InitialCircleGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.InitialAxis)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.RotatedAxis)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.RotatedRhombus)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.RotatedRhombusGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.RotatedCircleGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.NextRhombus)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.NextCircle)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.GrowingRhombusGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.GrowingCircleGrid)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.GrowingCircleGridLeft)
	parameter.GenerateSvgShape(gongsvgStage, layer, parameter.ConstructionAxis)
	gongsvgStage.Commit()

}
