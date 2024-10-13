package models

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type FrontCurve struct {
	Name string

	Path string
}

type FrontCurveStack struct {
	Name string
	AbstractShape

	FrontCurves []*FrontCurve

	SpiralCircles []*SpiralCircle

	Presentation
}

// Draw implements Shape.
func (frontCurveStack *FrontCurveStack) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	for idx, frontCurve := range frontCurveStack.FrontCurves {

		path := new(gongsvg_models.Path).Stage(gongsvgStage)
		path.Name = "Front Curve Stack " + fmt.Sprintf("%d", idx)
		layer.Paths = append(layer.Paths, path)

		frontCurveStack.Presentation.CopyTo(&path.Presentation)

		// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
		path.Definition = frontCurve.Path
	}

	for _, sc := range frontCurveStack.SpiralCircles {
		sc.Draw(gongsvgStage, layer, p)
	}
}
