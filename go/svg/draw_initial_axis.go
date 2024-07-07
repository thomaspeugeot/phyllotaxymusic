package svg

import (
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawInitialAxis(gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	parameter *phylotaxymusic_models.Parameter,
	initialAxis *phylotaxymusic_models.InitialAxis) {

	initialAxisLine := new(gongsvg_models.Line).Stage(gongsvgStage)
	initialAxisLine.Name = "Initial Axis Line"
	layer.Lines = append(layer.Lines, initialAxisLine)

	angleRad := initialAxis.Angle * math.Pi / 180

	initialAxisLine.X1 = parameter.OriginX
	initialAxisLine.Y1 = parameter.OriginY

	initialAxisLine.X2 = parameter.OriginX + initialAxis.Length*math.Cos(angleRad)
	initialAxisLine.Y2 = parameter.OriginY - initialAxis.Length*math.Sin(angleRad)

	initialAxisLine.StrokeWidth = initialAxis.StrokeWidth
	initialAxisLine.StrokeOpacity = 1
	initialAxisLine.Name = "Vertical Axis"
	initialAxisLine.Stroke = gongsvg_models.Black.ToString()
}
