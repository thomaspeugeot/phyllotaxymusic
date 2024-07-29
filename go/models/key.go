package models

import (
	"fmt"

	svg "github.com/fullstack-lang/gongsvg/go/models"
)

type Key struct {
	Name string

	AbstractShape

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

	path.Transform = fmt.Sprintf(" translate(%0.1f, %0.1f)", p.OriginX, p.OriginY) + path.Transform
}
