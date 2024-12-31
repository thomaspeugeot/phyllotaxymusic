package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type CircleGrid struct {
	Name string

	Reference *Circle

	Shape

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

func (g *CircleGrid) Move(seed *Circle, source *CircleGrid, x, y float64) {

	g.Circles = g.Circles[:0]

	for _, b := range source.Circles {
		_b := new(Circle)
		*_b = *seed

		g.Circles = append(g.Circles, _b)

		_b.move(b, x, y)
		_b.isKept = b.isKept
	}
}
