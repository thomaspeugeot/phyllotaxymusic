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

	if parameter.IsHorizontalAxisDisplayed {
		drawHorizontalAxis(gongsvgStage, layer, parameter)
	}

	gongsvgStage.Commit()

}
