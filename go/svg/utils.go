package svg

import (
	"math"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

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
	originSet Diamond,
	deltaX, deltaY float64) (copy Diamond) {

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
	originSet Diamond,
	deltaX, deltaY, angle float64) (copy Diamond) {

	for _, _line := range originSet {
		copy = append(copy, RotateLine(gongsvgStage, layer, _line, deltaX, deltaY, angle))
	}

	applyRotatedObjectPresentation(copy)

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

func CopyCircle(
	gongsvgStage *gongsvg_models.StageStruct,
	layer *gongsvg_models.Layer,
	origin *gongsvg_models.Circle, x, y float64) (copy *gongsvg_models.Circle) {

	copy = new(gongsvg_models.Circle).Stage(gongsvgStage)
	layer.Circles = append(layer.Circles, copy)
	*copy = *origin

	// Translate the Circle back to its original position
	copy.CX += x
	copy.CY += y

	return
}
