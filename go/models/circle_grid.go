package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type CircleGrid struct {
	Name string

	Reference *Circle

	HideableShape

	Circles []*Circle
}

func (g *CircleGrid) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {

	for _, c := range g.Circles {
		c.Draw(gongsvgStage, layer, p)
	}

}
