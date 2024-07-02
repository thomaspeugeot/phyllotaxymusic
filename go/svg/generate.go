package svg

import (
	"log"
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func GenerateSvg(gongsvgStage *gongsvg_models.StageStruct, phylotaxymusicStage *phylotaxymusic_models.StageStruct) {

	diagrams := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Diagram](phylotaxymusicStage)

	if len(*diagrams) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	diagram := (*diagrams)["Reference"]
	_ = diagram

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvgStage)
	layer := new(gongsvg_models.Layer).Stage(gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	var presentation gongsvg_models.Presentation
	presentation.Stroke = "black"
	presentation.StrokeOpacity = 0.8
	presentation.StrokeWidth = 2

	// build the first diamond
	base := new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, base)
	base.X1 = diagram.OriginX
	base.Y1 = diagram.OriginY
	base.X2 = base.X1 + diagram.DiamondSide
	base.Y2 = base.Y1
	base.Presentation = presentation

	deltaX := diagram.DiamondSide * math.Cos(math.Pi*diagram.DiamondAngle/180.0)
	deltaY := diagram.DiamondSide * math.Sin(math.Pi*diagram.DiamondAngle/180.0)

	risingSideLeft := new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, risingSideLeft)
	risingSideLeft.X1 = diagram.OriginX
	risingSideLeft.Y1 = diagram.OriginY
	risingSideLeft.X2 = risingSideLeft.X1 + deltaX
	risingSideLeft.Y2 = risingSideLeft.Y1 - deltaY
	risingSideLeft.Presentation = presentation

	// copyLine(gongsvgStage, layer, base, diagram.DiamondSide, 0.0)
	copyLine(gongsvgStage, layer, risingSideLeft, diagram.DiamondSide, 0.0)
	copyLine(gongsvgStage, layer, base, deltaX, -deltaY)

	gongsvgStage.Commit()
}

func copyLine(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Line,
	deltaX, deltaY float64) {

	copy := new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, copy)
	*copy = *origin
	copy.X1 += deltaX
	copy.X2 += deltaX
	copy.Y1 += deltaY
	copy.Y2 += deltaY
}
