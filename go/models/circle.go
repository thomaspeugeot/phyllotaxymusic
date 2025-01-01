package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
)

type Circle struct {
	Name string

	Shape
	CenterX, CenterY float64

	HasBespokeRadius bool
	BespopkeRadius   float64

	Presentation

	// following field are used if the Circle is a note
	isANote bool

	Pitch int

	// effect of the note info on the circle
	// will influence how it is drawn and played
	isKept bool

	note *gongtone_models.Note

	ShowName bool

	Impl   CircleUpdater
	BeatNb int // rank within the theme
}

type CircleUpdater interface {
	Updated()
}

func NewCircleUpdaterImpl(parameter *Parameter, beatNb int) *CircleUpdaterImpl {
	return &CircleUpdaterImpl{
		parameter: parameter,
		beatNb:    beatNb,
	}
}

type CircleUpdaterImpl struct {
	beatNb    int
	parameter *Parameter
}

func (circleUpdaterImpl *CircleUpdaterImpl) Updated() {

	circleUpdaterImpl.parameter.ToggleNotePlayed(circleUpdaterImpl.beatNb)
	circleUpdaterImpl.parameter.UpdatePhyllotaxyStage()
	circleUpdaterImpl.parameter.UpdateAndCommitCursorStage()
	circleUpdaterImpl.parameter.UpdateAndCommitSVGStage()
	circleUpdaterImpl.parameter.UpdateAndCommitToneStage()
	circleUpdaterImpl.parameter.UpdateAndCommitTreeStage()
	circleUpdaterImpl.parameter.CommitPhyllotaxymusicStage()
}

// RectUpdated implements models.RectImplInterface.
func (circle *Circle) RectUpdated(updatedRect *gongsvg_models.Rect) {

	log.Println("circle updated")

	circle.Impl.Updated()
}

func (circle *Circle) Draw(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *Parameter,
) {

	// add a reactive gongsvg rect
	if circle.isANote {
		svgRect := new(gongsvg_models.Rect).Stage(gongsvgStage)
		layer.Rects = append(layer.Rects, svgRect)

		rectWidth := 20.0

		svgRect.X = p.OriginX + circle.CenterX - rectWidth/2.0
		svgRect.Y = p.OriginY - circle.CenterY - rectWidth/2.0
		svgRect.Width = rectWidth
		svgRect.Height = rectWidth
		svgRect.Name = circle.Name
		circle.Presentation.CopyTo(&svgRect.Presentation)

		if !circle.isKept {
			svgRect.StrokeWidth /= 2.0
		} else {

		}

		// put the callback
		svgRect.Impl = circle
		circle.Impl = NewCircleUpdaterImpl(p, circle.BeatNb)
	} else {
		svgCircle := new(gongsvg_models.Circle).Stage(gongsvgStage)
		layer.Circles = append(layer.Circles, svgCircle)

		svgCircle.CX = p.OriginX + circle.CenterX
		svgCircle.CY = p.OriginY - circle.CenterY
		svgCircle.Radius = p.SideLength / 2.0

		if circle.HasBespokeRadius {
			svgCircle.Radius = circle.BespopkeRadius
		}

		circle.Presentation.CopyTo(&svgCircle.Presentation)
	}

	if circle.ShowName {
		svgText := new(gongsvg_models.Text).Stage(gongsvgStage)
		layer.Texts = append(layer.Texts, svgText)

		svgText.X = p.OriginX + circle.CenterX
		svgText.Y = p.OriginY - circle.CenterY

		svgText.Name = circle.Name
		svgText.Content = circle.Name
		circle.Presentation.CopyTo(&svgText.Presentation)
	}
}

func (_c *Circle) move(c *Circle, x, y float64) {
	_c.BespopkeRadius = c.BespopkeRadius
	_c.HasBespokeRadius = c.HasBespokeRadius
	_c.Pitch = c.Pitch
	_c.CenterX += c.CenterX + x
	_c.CenterY += c.CenterY + y
}
