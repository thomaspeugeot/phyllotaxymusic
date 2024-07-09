package svg

import (
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawAxis(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *phylotaxymusic_models.Parameter,
	axis *phylotaxymusic_models.Axis) {

	line := new(gongsvg_models.Line).Stage(gongsvgStage)
	line.Name = "Initial Axis Line"
	layer.Lines = append(layer.Lines, line)

	angleRad := axis.Angle * math.Pi / 180

	line.X1 = parameter.OriginX
	line.Y1 = parameter.OriginY

	line.X2 = parameter.OriginX + axis.Length*math.Cos(angleRad)
	line.Y2 = parameter.OriginY - axis.Length*math.Sin(angleRad)

	axis.Presentation.CopyTo(&line.Presentation)

	line.Name = "Vertical Axis"
}
