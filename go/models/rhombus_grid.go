package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type RhombusGrid struct {
	Name      string
	Reference *Rhombus

	Shape

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
