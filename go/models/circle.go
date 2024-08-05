package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
)

type Circle struct {
	Name string

	AbstractShape
	CenterX, CenterY float64

	HasBespokeRadius bool
	BespopkeRadius   float64

	Pitch int

	Presentation

	note *gongtone_models.Note
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

	if c.HasBespokeRadius {
		circle.Radius = c.BespopkeRadius
	}

	c.Presentation.CopyTo(&circle.Presentation)
}

func (_c *Circle) move(c *Circle, x, y float64) {
	_c.BespopkeRadius = c.BespopkeRadius
	_c.HasBespokeRadius = c.HasBespokeRadius
	_c.Pitch = c.Pitch
	_c.CenterX += c.CenterX + x
	_c.CenterY += c.CenterY + y
}
