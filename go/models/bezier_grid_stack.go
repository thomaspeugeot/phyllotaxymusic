package models

import gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

type BezierGridStack struct {
	Name string

	HideableShape

	BezierGrids []*BezierGrid
}

func (s *BezierGridStack) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {
	for _, g := range s.BezierGrids {
		g.Draw(gongsvgStage, layer, p)
	}
}
