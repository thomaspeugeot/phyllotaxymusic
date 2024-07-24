package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Circle struct {
	Name string

	HideableShape
	CenterX, CenterY float64

	Presentation
}

func (c *Circle) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, circle)

	circle.CX = p.OriginX + c.CenterX
	circle.CY = p.OriginY - c.CenterY
	circle.Radius = p.SideLength / 2.0

	c.Presentation.CopyTo(&circle.Presentation)
}
