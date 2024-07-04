package svg

import (
	"log"
	"math"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Diamond []*gongsvg_models.Line
type DiamondGrid []*Diamond

func GenerateSvg(gongsvgStage *gongsvg_models.StageStruct, phylotaxymusicStage *phylotaxymusic_models.StageStruct) {

	diagrams := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Diagram](phylotaxymusicStage)

	if len(*diagrams) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	gongsvgStage.Reset()

	diagram := (*diagrams)["Reference"]
	_ = diagram

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(gongsvgStage)
	layer.Display = true
	// svg.Layers = append(svg.Layers, layer)

	rotatedObjectsLayer := (&gongsvg_models.Layer{Name: "Layer 2"}).Stage(gongsvgStage)
	rotatedObjectsLayer.Display = true
	svg.Layers = append(svg.Layers, rotatedObjectsLayer)

	var presentation gongsvg_models.Presentation
	presentation.Stroke = "black"
	presentation.StrokeOpacity = 0.2
	presentation.StrokeWidth = 4

	// build the first diamondRoot
	var diamondRoot Diamond

	diamondBottom := new(gongsvg_models.Line).Stage(gongsvgStage)
	diamondBottom.Name = "B"
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
	diamondLeftSide.Name = "L"
	diamondRoot = append(diamondRoot, diamondLeftSide)
	layer.Lines = append(layer.Lines, diamondLeftSide)
	diamondLeftSide.X1 = diagram.OriginX
	diamondLeftSide.Y1 = diagram.OriginY
	diamondLeftSide.X2 = diamondLeftSide.X1 + deltaX
	diamondLeftSide.Y2 = diamondLeftSide.Y1 - deltaY
	diamondLeftSide.Presentation = presentation

	diamondRightSide := CopyLine(gongsvgStage, layer, diamondLeftSide, diagram.DiamondSideLenght, 0.0)
	diamondRightSide.Name = "R"
	diamondRoot = append(diamondRoot, diamondRightSide)

	diamondTopSide := CopyLine(gongsvgStage, layer, diamondBottom, deltaX, -deltaY)
	diamondTopSide.Name = "T"
	diamondRoot = append(diamondRoot, diamondTopSide)

	circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, circle)
	circle.Presentation = presentation

	circle.CX = diagram.OriginX + float64(diagram.N)*diagram.DiamondSideLenght + float64(diagram.M)*deltaX
	circle.CY = diagram.OriginY - float64(diagram.M)*deltaY
	circle.Radius = diagram.DiamondSideLenght / 2.0

	angleForTheGridEnd := math.Atan2(circle.CY-diagram.OriginY, circle.CX-diagram.OriginX)
	angleForTheGridEnd = 180 * angleForTheGridEnd / math.Pi

	currentDiamond := diamondRoot
	rotatedCurrentDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, diamondRoot, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)

	var rotatedDiamondGrid DiamondGrid
	rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedCurrentDiamond)

	for j := range diagram.M - 1 {
		copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
		rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, copiedDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
		rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)
	}
	for i := range diagram.N - 1 {
		currentDiamond = CopyLineSet(gongsvgStage, layer, diamondRoot, float64(i+1)*diagram.DiamondSideLenght, 0)
		rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, currentDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
		rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)

		for j := range diagram.M - 1 {
			copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
			rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, copiedDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
			rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)
		}
	}

	// rotatedCircle := RotateCircle(gongsvgStage, rotatedObjectsLayer, circle, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
	// _ = rotatedCircle

	//
	// create the circle on each rotated diamond tip in the diamond grid
	generateCircleFromDiamondGrid(gongsvgStage, rotatedObjectsLayer, presentation, diagram, rotatedDiamondGrid)

	gongsvgStage.Commit()
}

func applyRotatedObjectPresentation(lineSet Diamond) {

	for _, _line := range lineSet {
		_line.Stroke = "green"
	}
}

func generateCircleFromDiamondGrid(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	presentation gongsvg_models.Presentation,
	diagram *models.Diagram,
	diamondGrid DiamondGrid) (circles []*gongsvg_models.Circle) {

	for _, diamond := range diamondGrid {

		x, y := computeDiamonCenter(*diamond)
		// if y <= diagram.OriginY {
		circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
		layer.Circles = append(layer.Circles, circle)
		circle.CX = x
		circle.CY = y
		circle.Radius = diagram.DiamondSideLenght / 2.0

		circle.Presentation = presentation
		circles = append(circles, circle)
		// }
	}

	return
}

func computeDiamonCenter(diamond Diamond) (x, y float64) {

	for _, line := range diamond {
		x += line.X1
		x += line.X2
	}
	x /= 8.0

	for _, line := range diamond {
		y += line.Y1
		y += line.Y2
	}
	y /= 8.0

	return
}
