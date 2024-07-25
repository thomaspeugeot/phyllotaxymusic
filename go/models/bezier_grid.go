package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type BezierGrid struct {
	Name string

	Reference *Bezier

	HideableShape

	Beziers []*Bezier
}

func (g *BezierGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {
	for _, c := range g.Beziers {
		c.Draw(gongsvgStage, layer, p)
	}
}
