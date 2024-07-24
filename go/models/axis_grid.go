package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type AxisGrid struct {
	Name string

	Reference *Axis

	HideableShape

	Axiss []*Axis
}

func (g *AxisGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {

	for _, c := range g.Axiss {
		c.Draw(gongsvgStage, layer, p)
	}

}
