package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type RhombusGrid struct {
	Name      string
	Reference *Rhombus

	AbstractShape

	Rhombuses []*Rhombus
}

func (g *RhombusGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	for _, r := range g.Rhombuses {
		r.Draw(gongsvgStage, layer, p)
	}

}
