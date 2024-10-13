package models

import (
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

	Presentation
}

// Draw implements Shape.
func (frontCurveStack *FrontCurveStack) Draw(gongsvgStage *gongsvg_models.StageStruct, layer *gongsvg_models.Layer, p *Parameter) {

	for _, frontCurve := range frontCurveStack.FrontCurves {

		path := new(gongsvg_models.Path).Stage(gongsvgStage)
		path.Name = "Front Curve Stack"
		layer.Paths = append(layer.Paths, path)

		frontCurveStack.Presentation.CopyTo(&path.Presentation)

		// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
		path.Definition = frontCurve.Path
	}
}
