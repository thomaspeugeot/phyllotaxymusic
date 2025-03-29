package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type BezierGrid struct {
	Name string

	Reference *Bezier

	Shape

	Beziers []*Bezier
}

func (g *BezierGrid) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	p *Parameter,

) {
	for _, c := range g.Beziers {
		c.Draw(gongsvgStage, layer, p)
	}
}

func (g *BezierGrid) Move(seed *Bezier, source *BezierGrid, x, y float64) {

	g.Beziers = g.Beziers[:0]

	for _, b := range source.Beziers {
		_b := new(Bezier)
		*_b = *seed
		g.Beziers = append(g.Beziers, _b)

		_b.move(b, x, y)
	}
}
