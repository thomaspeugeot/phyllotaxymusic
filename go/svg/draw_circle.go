package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawCircle(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	p *phylotaxymusic_models.Parameter,
	c *phylotaxymusic_models.Circle,
) {
	circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, circle)

	circle.CX = p.OriginX + c.CenterX
	circle.CY = p.OriginY - c.CenterY
	circle.Radius = p.SideLength / 2.0

	c.Presentation.CopyTo(&circle.Presentation)

	circle.Name = "Vertical Axis"
}
