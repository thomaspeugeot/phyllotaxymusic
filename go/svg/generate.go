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
	circle.Radius = 20

	angleForTheGridEnd := math.Atan2(circle.CY-diagram.OriginY, circle.CX-diagram.OriginX)
	angleForTheGridEnd = 180 * angleForTheGridEnd / math.Pi

	currentDiamond := diamondRoot
	RotateLineSet(gongsvgStage, layer, diamondRoot, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)

	for j := range diagram.M - 1 {
		copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
		RotateLineSet(gongsvgStage, layer, copiedDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
	}
	for i := range diagram.N - 1 {
		currentDiamond = CopyLineSet(gongsvgStage, layer, diamondRoot, float64(i+1)*diagram.DiamondSideLenght, 0)
		RotateLineSet(gongsvgStage, layer, currentDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)

		for j := range diagram.M - 1 {
			copiedDiamond := CopyLineSet(gongsvgStage, layer, currentDiamond, float64(j+1)*deltaX, -float64(j+1)*deltaY)
			RotateLineSet(gongsvgStage, layer, copiedDiamond, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
		}
	}

	rotatedCircle := RotateCircle(gongsvgStage, layer, circle, diagram.OriginX, diagram.OriginY, -angleForTheGridEnd)
	_ = rotatedCircle

	gongsvgStage.Commit()
}

func CopyLine(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Line,
	deltaX, deltaY float64) (copy *gongsvg_models.Line) {

	copy = new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, copy)
	*copy = *origin
	if deltaY != 0 {
		copy.Name = origin.Name + "v"
	} else {
		copy.Name = origin.Name + "h"
	}
	copy.X1 += deltaX
	copy.X2 += deltaX
	copy.Y1 += deltaY
	copy.Y2 += deltaY

	return
}

func CopyLineSet(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	originSet LineSet,
	deltaX, deltaY float64) (copy LineSet) {

	for _, _line := range originSet {
		copy = append(copy, CopyLine(gongsvgStage, layer, _line, deltaX, deltaY))
	}

	return
}

func RotateLine(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Line, originX, originY, angle float64) (copy *gongsvg_models.Line) {
	// Convert angle from degrees to radians
	radians := angle * math.Pi / 180

	copy = new(gongsvg_models.Line).Stage(gongsvgStage)
	layer.Lines = append(layer.Lines, copy)
	*copy = *origin

	// Translate the line to the origin
	copy.X1 -= originX
	copy.Y1 -= originY
	copy.X2 -= originX
	copy.Y2 -= originY

	// Rotate points around the origin
	tempX1 := copy.X1*math.Cos(radians) - copy.Y1*math.Sin(radians)
	tempY1 := copy.X1*math.Sin(radians) + copy.Y1*math.Cos(radians)
	tempX2 := copy.X2*math.Cos(radians) - copy.Y2*math.Sin(radians)
	tempY2 := copy.X2*math.Sin(radians) + copy.Y2*math.Cos(radians)

	// Translate the line back to its original position
	copy.X1 = tempX1 + originX
	copy.Y1 = tempY1 + originY
	copy.X2 = tempX2 + originX
	copy.Y2 = tempY2 + originY

	return
}

func RotateLineSet(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	originSet LineSet,
	deltaX, deltaY, angle float64) (copy LineSet) {

	for _, _line := range originSet {
		copy = append(copy, RotateLine(gongsvgStage, layer, _line, deltaX, deltaY, angle))
	}

	return
}

func RotateCircle(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Circle, originX, originY, angle float64) (copy *gongsvg_models.Circle) {
	// Convert angle from degrees to radians
	radians := angle * math.Pi / 180

	copy = new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, copy)
	*copy = *origin

	// Translate the Circle to the origin
	copy.CX -= originX
	copy.CY -= originY

	// Rotate points around the origin
	tempCX := copy.CX*math.Cos(radians) - copy.CY*math.Sin(radians)
	tempCY := copy.CX*math.Sin(radians) + copy.CY*math.Cos(radians)

	// Translate the Circle back to its original position
	copy.CX = tempCX + originX
	copy.CY = tempCY + originY

	return
}
