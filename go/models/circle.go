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

	// effect of the note info on the circle
	// will influence how it is drawn and played
	isKept bool

	Presentation

	note *gongtone_models.Note
}

func (circle *Circle) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	svgCircle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, svgCircle)

	svgCircle.CX = p.OriginX + circle.CenterX
	svgCircle.CY = p.OriginY - circle.CenterY
	svgCircle.Radius = p.SideLength / 2.0

	if circle.HasBespokeRadius {
		svgCircle.Radius = circle.BespopkeRadius
	}

	circle.Presentation.CopyTo(&svgCircle.Presentation)

	if !circle.isKept {
		svgCircle.StrokeWidth /= 2.0
	}

	svgText := new(gongsvg_models.Text).Stage(gongsvgStage)
	layer.Texts = append(layer.Texts, svgText)

	svgText.X = p.OriginX + circle.CenterX
	svgText.Y = p.OriginY - circle.CenterY

	svgText.Name = circle.Name
	svgText.Content = circle.Name
	circle.Presentation.CopyTo(&svgText.Presentation)

}

func (_c *Circle) move(c *Circle, x, y float64) {
	_c.BespopkeRadius = c.BespopkeRadius
	_c.HasBespokeRadius = c.HasBespokeRadius
	_c.Pitch = c.Pitch
	_c.CenterX += c.CenterX + x
	_c.CenterY += c.CenterY + y
}
