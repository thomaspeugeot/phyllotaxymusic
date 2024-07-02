package svg

import (
	"log"
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type LineSet []*gongsvg_models.Line

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
	presentation.StrokeOpacity = 0.2
	presentation.StrokeWidth = 4

	// build the first diamondRoot
	var diamondRoot LineSet

	diamondBottom := new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, diamondBottom)
	diamondRoot = append(diamondRoot, diamondBottom)
	diamondBottom.X1 = diagram.OriginX
	diamondBottom.Y1 = diagram.OriginY
	diamondBottom.X2 = diamondBottom.X1 + diagram.DiamondSideLenght
	diamondBottom.Y2 = diamondBottom.Y1
	diamondBottom.Presentation = presentation

	deltaX := diagram.DiamondSideLenght * math.Cos(math.Pi*diagram.DiamondAngle/180.0)
	deltaY := diagram.DiamondSideLenght * math.Sin(math.Pi*diagram.DiamondAngle/180.0)

	diamondLeftSide := new(gongsvg_models.Line).Stage(gongsvgStage)
	diamondRoot = append(diamondRoot, diamondLeftSide)
	layer.Lines = append(layer.Lines, diamondLeftSide)
	diamondLeftSide.X1 = diagram.OriginX
	diamondLeftSide.Y1 = diagram.OriginY
	diamondLeftSide.X2 = diamondLeftSide.X1 + deltaX
	diamondLeftSide.Y2 = diamondLeftSide.Y1 - deltaY
	diamondLeftSide.Presentation = presentation

	diamondRightSide := copyLine(gongsvgStage, layer, diamondLeftSide, diagram.DiamondSideLenght, 0.0)
	diamondRoot = append(diamondRoot, diamondRightSide)

	diamondTopSide := copyLine(gongsvgStage, layer, diamondBottom, deltaX, -deltaY)
	diamondRoot = append(diamondRoot, diamondTopSide)

	currentDiamond := diamondRoot
	for j := range diagram.M {
		copyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
	}
	for i := range diagram.N {
		currentDiamond = copyLineSet(gongsvgStage, layer, diamondRoot, float64(i+1)*diagram.DiamondSideLenght, 0)
		for j := range diagram.M {
			copyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
		}

	}

	gongsvgStage.Commit()
}

func copyLine(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Line,
	deltaX, deltaY float64) (copy *gongsvg_models.Line) {

	copy = new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, copy)
	*copy = *origin
	copy.X1 += deltaX
	copy.X2 += deltaX
	copy.Y1 += deltaY
	copy.Y2 += deltaY

	return
}

func copyLineSet(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	originSet LineSet,
	deltaX, deltaY float64) (copy LineSet) {

	for _, _line := range originSet {
		copy = append(copy, copyLine(gongsvgStage, layer, _line, deltaX, deltaY))
	}

	return
}
