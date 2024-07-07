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

	parameters := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Parameter](phylotaxymusicStage)

	if len(*parameters) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	gongsvgStage.Reset()

	parameter := (*parameters)["Reference"]
	_ = parameter

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
	diamondBottom.X1 = parameter.OriginX
	diamondBottom.Y1 = parameter.OriginY
	diamondBottom.X2 = diamondBottom.X1 + parameter.SideLength
	diamondBottom.Y2 = diamondBottom.Y1
	diamondBottom.Presentation = presentation

	deltaX := parameter.SideLength * math.Cos(math.Pi*parameter.InsideAngle/180.0)
	deltaY := parameter.SideLength * math.Sin(math.Pi*parameter.InsideAngle/180.0)

	diamondLeftSide := new(gongsvg_models.Line).Stage(gongsvgStage)
	diamondLeftSide.Name = "L"
	diamondRoot = append(diamondRoot, diamondLeftSide)
	layer.Lines = append(layer.Lines, diamondLeftSide)
	diamondLeftSide.X1 = parameter.OriginX
	diamondLeftSide.Y1 = parameter.OriginY
	diamondLeftSide.X2 = diamondLeftSide.X1 + deltaX
	diamondLeftSide.Y2 = diamondLeftSide.Y1 - deltaY
	diamondLeftSide.Presentation = presentation

	diamondRightSide := CopyLine(gongsvgStage, layer, diamondLeftSide, parameter.SideLength, 0.0)
	diamondRightSide.Name = "R"
	diamondRoot = append(diamondRoot, diamondRightSide)

	diamondTopSide := CopyLine(gongsvgStage, layer, diamondBottom, deltaX, -deltaY)
	diamondTopSide.Name = "T"
	diamondRoot = append(diamondRoot, diamondTopSide)

	circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, circle)
	circle.Presentation = presentation

	circle.CX = parameter.OriginX + float64(parameter.N)*parameter.SideLength + float64(parameter.M)*deltaX
	circle.CY = parameter.OriginY - float64(parameter.M)*deltaY
	circle.Radius = parameter.SideLength / 2.0

	angleForTheGridEndRad := math.Atan2(circle.CY-parameter.OriginY, circle.CX-parameter.OriginX)
	angleForTheGridEndDegree := 180 * angleForTheGridEndRad / math.Pi

	distanceForTheGridEnd := math.Sqrt(
		(circle.CY-parameter.OriginY)*(circle.CY-parameter.OriginY) +
			(circle.CX-parameter.OriginX)*(circle.CX-parameter.OriginX))

	currentDiamond := diamondRoot
	rotatedCurrentDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, diamondRoot, parameter.OriginX, parameter.OriginY, -angleForTheGridEndDegree)

	var rotatedDiamondGrid DiamondGrid
	rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedCurrentDiamond)

	for j := range parameter.M - 1 {
		copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
		rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, copiedDiamond, parameter.OriginX, parameter.OriginY, -angleForTheGridEndDegree)
		rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)
	}
	for i := range parameter.N - 1 {
		currentDiamond = CopyLineSet(gongsvgStage, layer, diamondRoot, float64(i+1)*parameter.SideLength, 0)
		rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, currentDiamond, parameter.OriginX, parameter.OriginY, -angleForTheGridEndDegree)
		rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)

		for j := range parameter.M - 1 {
			copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
			rotatedDiamond := RotateLineSet(gongsvgStage, rotatedObjectsLayer, copiedDiamond, parameter.OriginX, parameter.OriginY, -angleForTheGridEndDegree)
			rotatedDiamondGrid = append(rotatedDiamondGrid, &rotatedDiamond)
		}
	}

	// rotatedCircle := RotateCircle(gongsvgStage, rotatedObjectsLayer, circle, parameter.OriginX, parameter.OriginY, -angleForTheGridEnd)
	// _ = rotatedCircle

	//
	// create the circle on each rotated diamond tip in the diamond grid
	circles := generateCircleFromDiamondGrid(gongsvgStage, rotatedObjectsLayer, presentation, parameter, rotatedDiamondGrid)
	_ = circles
	for _, circle := range circles {
		replicatedCircle := CopyCircle(gongsvgStage, rotatedObjectsLayer, circle, distanceForTheGridEnd, 0)
		replicatedCircle.Presentation.Stroke = "lightblue"
		_ = replicatedCircle
	}
	for _, circle := range circles {
		radians := angleForTheGridEndDegree * math.Pi / 180

		rotatedDeltaX := parameter.SideLength * math.Cos(angleForTheGridEndRad+radians)
		rotatedDeltaY := parameter.SideLength * math.Sin(angleForTheGridEndRad+radians)
		replicatedCircle := CopyCircle(gongsvgStage, rotatedObjectsLayer, circle,
			rotatedDeltaX, rotatedDeltaY)
		replicatedCircle.Presentation.Stroke = "lightgreen"
		_ = replicatedCircle
	}

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
	parameter *models.Parameter,
	diamondGrid DiamondGrid) (circles []*gongsvg_models.Circle) {

	for _, diamond := range diamondGrid {

		x, y := computeDiamonCenter(*diamond)
		// if y <= parameter.OriginY {
		circle := new(gongsvg_models.Circle).Stage(gongsvgStage)
		layer.Circles = append(layer.Circles, circle)
		circle.CX = x
		circle.CY = y
		circle.Radius = parameter.SideLength / 2.0

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
