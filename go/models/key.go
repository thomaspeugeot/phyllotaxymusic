package models

import (
	"fmt"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Key struct {
	Name string

	Shape

	Path string

	Presentation
}

// Draw implements Shape.
func (key *Key) Draw(gongsvgStage *svg.StageStruct, layer *svg.Layer, p *Parameter) {

	path := new(svg.Path).Stage(gongsvgStage)
	path.Name = "F Key"
	layer.Paths = append(layer.Paths, path)

	path.Definition = key.Path

	key.Presentation.CopyTo(&path.Presentation)

	path.Transform = fmt.Sprintf("scale(%f, %f) translate(%0.1f, %0.1f) ",
		p.FkeySizeRatio*p.SideLength,
		p.FkeySizeRatio*p.SideLength,
		p.OriginX+p.FkeyOriginRelativeX*p.SideLength,
		p.OriginY-p.FkeyOriginRelativeY*p.SideLength,
	) + path.Transform
}
