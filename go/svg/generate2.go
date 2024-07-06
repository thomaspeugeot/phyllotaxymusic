package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func GenerateSvg2(gongsvgStage *gongsvg_models.StageStruct, parameter *phylotaxymusic_models.Parameter) {

	gongsvgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	if parameter.HorizontalAxis.IsAxisDisplayed {
		drawHorizontalAxis(gongsvgStage, layer, parameter, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis.IsAxisDisplayed {
		drawVerticalAxis(gongsvgStage, layer, parameter, parameter.VerticalAxis)
	}
	if parameter.InitialRhombus.IsDisplayed {
		drawRhombus(gongsvgStage, layer, parameter, parameter.InitialRhombus)
	}

	gongsvgStage.Commit()

}
