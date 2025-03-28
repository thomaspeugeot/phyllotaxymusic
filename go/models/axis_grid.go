package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type AxisGrid struct {
	Name string

	Reference *Axis

	Shape

	Axiss []*Axis
}

func (g *AxisGrid) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {
	for _, c := range g.Axiss {
		c.Draw(gongsvgStage, layer, p)
	}
}
