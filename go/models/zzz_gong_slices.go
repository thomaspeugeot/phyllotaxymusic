// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Axis
	// insertion point per field

	// Compute reverse map for named struct AxisGrid
	// insertion point per field
	stage.AxisGrid_Axiss_reverseMap = make(map[*Axis]*AxisGrid)
	for axisgrid := range stage.AxisGrids {
		_ = axisgrid
		for _, _axis := range axisgrid.Axiss {
			stage.AxisGrid_Axiss_reverseMap[_axis] = axisgrid
		}
	}

	// Compute reverse map for named struct Bezier
	// insertion point per field

	// Compute reverse map for named struct BezierGrid
	// insertion point per field
	stage.BezierGrid_Beziers_reverseMap = make(map[*Bezier]*BezierGrid)
	for beziergrid := range stage.BezierGrids {
		_ = beziergrid
		for _, _bezier := range beziergrid.Beziers {
			stage.BezierGrid_Beziers_reverseMap[_bezier] = beziergrid
		}
	}

	// Compute reverse map for named struct BezierGridStack
	// insertion point per field
	stage.BezierGridStack_BezierGrids_reverseMap = make(map[*BezierGrid]*BezierGridStack)
	for beziergridstack := range stage.BezierGridStacks {
		_ = beziergridstack
		for _, _beziergrid := range beziergridstack.BezierGrids {
			stage.BezierGridStack_BezierGrids_reverseMap[_beziergrid] = beziergridstack
		}
	}

	// Compute reverse map for named struct Chapter
	// insertion point per field

	// Compute reverse map for named struct Circle
	// insertion point per field

	// Compute reverse map for named struct CircleGrid
	// insertion point per field
	stage.CircleGrid_Circles_reverseMap = make(map[*Circle]*CircleGrid)
	for circlegrid := range stage.CircleGrids {
		_ = circlegrid
		for _, _circle := range circlegrid.Circles {
			stage.CircleGrid_Circles_reverseMap[_circle] = circlegrid
		}
	}

	// Compute reverse map for named struct Content
	// insertion point per field
	stage.Content_Chapters_reverseMap = make(map[*Chapter]*Content)
	for content := range stage.Contents {
		_ = content
		for _, _chapter := range content.Chapters {
			stage.Content_Chapters_reverseMap[_chapter] = content
		}
	}

	// Compute reverse map for named struct ExportToMusicxml
	// insertion point per field

	// Compute reverse map for named struct FrontCurve
	// insertion point per field

	// Compute reverse map for named struct FrontCurveStack
	// insertion point per field
	stage.FrontCurveStack_FrontCurves_reverseMap = make(map[*FrontCurve]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		_ = frontcurvestack
		for _, _frontcurve := range frontcurvestack.FrontCurves {
			stage.FrontCurveStack_FrontCurves_reverseMap[_frontcurve] = frontcurvestack
		}
	}
	stage.FrontCurveStack_SpiralCircles_reverseMap = make(map[*SpiralCircle]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		_ = frontcurvestack
		for _, _spiralcircle := range frontcurvestack.SpiralCircles {
			stage.FrontCurveStack_SpiralCircles_reverseMap[_spiralcircle] = frontcurvestack
		}
	}

	// Compute reverse map for named struct HorizontalAxis
	// insertion point per field

	// Compute reverse map for named struct Key
	// insertion point per field

	// Compute reverse map for named struct Parameter
	// insertion point per field

	// Compute reverse map for named struct Rhombus
	// insertion point per field

	// Compute reverse map for named struct RhombusGrid
	// insertion point per field
	stage.RhombusGrid_Rhombuses_reverseMap = make(map[*Rhombus]*RhombusGrid)
	for rhombusgrid := range stage.RhombusGrids {
		_ = rhombusgrid
		for _, _rhombus := range rhombusgrid.Rhombuses {
			stage.RhombusGrid_Rhombuses_reverseMap[_rhombus] = rhombusgrid
		}
	}

	// Compute reverse map for named struct ShapeCategory
	// insertion point per field

	// Compute reverse map for named struct SpiralBezier
	// insertion point per field

	// Compute reverse map for named struct SpiralBezierGrid
	// insertion point per field
	stage.SpiralBezierGrid_SpiralBeziers_reverseMap = make(map[*SpiralBezier]*SpiralBezierGrid)
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		_ = spiralbeziergrid
		for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
			stage.SpiralBezierGrid_SpiralBeziers_reverseMap[_spiralbezier] = spiralbeziergrid
		}
	}

	// Compute reverse map for named struct SpiralCircle
	// insertion point per field

	// Compute reverse map for named struct SpiralCircleGrid
	// insertion point per field
	stage.SpiralCircleGrid_SpiralCircles_reverseMap = make(map[*SpiralCircle]*SpiralCircleGrid)
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		_ = spiralcirclegrid
		for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
			stage.SpiralCircleGrid_SpiralCircles_reverseMap[_spiralcircle] = spiralcirclegrid
		}
	}

	// Compute reverse map for named struct SpiralLine
	// insertion point per field

	// Compute reverse map for named struct SpiralLineGrid
	// insertion point per field
	stage.SpiralLineGrid_SpiralLines_reverseMap = make(map[*SpiralLine]*SpiralLineGrid)
	for spirallinegrid := range stage.SpiralLineGrids {
		_ = spirallinegrid
		for _, _spiralline := range spirallinegrid.SpiralLines {
			stage.SpiralLineGrid_SpiralLines_reverseMap[_spiralline] = spirallinegrid
		}
	}

	// Compute reverse map for named struct SpiralOrigin
	// insertion point per field

	// Compute reverse map for named struct SpiralRhombus
	// insertion point per field

	// Compute reverse map for named struct SpiralRhombusGrid
	// insertion point per field
	stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap = make(map[*SpiralRhombus]*SpiralRhombusGrid)
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		_ = spiralrhombusgrid
		for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
			stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap[_spiralrhombus] = spiralrhombusgrid
		}
	}

	// Compute reverse map for named struct VerticalAxis
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Axiss {
		res = append(res, instance)
	}

	for instance := range stage.AxisGrids {
		res = append(res, instance)
	}

	for instance := range stage.Beziers {
		res = append(res, instance)
	}

	for instance := range stage.BezierGrids {
		res = append(res, instance)
	}

	for instance := range stage.BezierGridStacks {
		res = append(res, instance)
	}

	for instance := range stage.Chapters {
		res = append(res, instance)
	}

	for instance := range stage.Circles {
		res = append(res, instance)
	}

	for instance := range stage.CircleGrids {
		res = append(res, instance)
	}

	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.ExportToMusicxmls {
		res = append(res, instance)
	}

	for instance := range stage.FrontCurves {
		res = append(res, instance)
	}

	for instance := range stage.FrontCurveStacks {
		res = append(res, instance)
	}

	for instance := range stage.HorizontalAxiss {
		res = append(res, instance)
	}

	for instance := range stage.Keys {
		res = append(res, instance)
	}

	for instance := range stage.Parameters {
		res = append(res, instance)
	}

	for instance := range stage.Rhombuss {
		res = append(res, instance)
	}

	for instance := range stage.RhombusGrids {
		res = append(res, instance)
	}

	for instance := range stage.ShapeCategorys {
		res = append(res, instance)
	}

	for instance := range stage.SpiralBeziers {
		res = append(res, instance)
	}

	for instance := range stage.SpiralBezierGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralCircles {
		res = append(res, instance)
	}

	for instance := range stage.SpiralCircleGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralLines {
		res = append(res, instance)
	}

	for instance := range stage.SpiralLineGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralOrigins {
		res = append(res, instance)
	}

	for instance := range stage.SpiralRhombuss {
		res = append(res, instance)
	}

	for instance := range stage.SpiralRhombusGrids {
		res = append(res, instance)
	}

	for instance := range stage.VerticalAxiss {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (axis *Axis) GongCopy() GongstructIF {
	newInstance := new(Axis)
	axis.CopyBasicFields(newInstance)
	return newInstance
}

func (axisgrid *AxisGrid) GongCopy() GongstructIF {
	newInstance := new(AxisGrid)
	axisgrid.CopyBasicFields(newInstance)
	return newInstance
}

func (bezier *Bezier) GongCopy() GongstructIF {
	newInstance := new(Bezier)
	bezier.CopyBasicFields(newInstance)
	return newInstance
}

func (beziergrid *BezierGrid) GongCopy() GongstructIF {
	newInstance := new(BezierGrid)
	beziergrid.CopyBasicFields(newInstance)
	return newInstance
}

func (beziergridstack *BezierGridStack) GongCopy() GongstructIF {
	newInstance := new(BezierGridStack)
	beziergridstack.CopyBasicFields(newInstance)
	return newInstance
}

func (chapter *Chapter) GongCopy() GongstructIF {
	newInstance := new(Chapter)
	chapter.CopyBasicFields(newInstance)
	return newInstance
}

func (circle *Circle) GongCopy() GongstructIF {
	newInstance := new(Circle)
	circle.CopyBasicFields(newInstance)
	return newInstance
}

func (circlegrid *CircleGrid) GongCopy() GongstructIF {
	newInstance := new(CircleGrid)
	circlegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := new(Content)
	content.CopyBasicFields(newInstance)
	return newInstance
}

func (exporttomusicxml *ExportToMusicxml) GongCopy() GongstructIF {
	newInstance := new(ExportToMusicxml)
	exporttomusicxml.CopyBasicFields(newInstance)
	return newInstance
}

func (frontcurve *FrontCurve) GongCopy() GongstructIF {
	newInstance := new(FrontCurve)
	frontcurve.CopyBasicFields(newInstance)
	return newInstance
}

func (frontcurvestack *FrontCurveStack) GongCopy() GongstructIF {
	newInstance := new(FrontCurveStack)
	frontcurvestack.CopyBasicFields(newInstance)
	return newInstance
}

func (horizontalaxis *HorizontalAxis) GongCopy() GongstructIF {
	newInstance := new(HorizontalAxis)
	horizontalaxis.CopyBasicFields(newInstance)
	return newInstance
}

func (key *Key) GongCopy() GongstructIF {
	newInstance := new(Key)
	key.CopyBasicFields(newInstance)
	return newInstance
}

func (parameter *Parameter) GongCopy() GongstructIF {
	newInstance := new(Parameter)
	parameter.CopyBasicFields(newInstance)
	return newInstance
}

func (rhombus *Rhombus) GongCopy() GongstructIF {
	newInstance := new(Rhombus)
	rhombus.CopyBasicFields(newInstance)
	return newInstance
}

func (rhombusgrid *RhombusGrid) GongCopy() GongstructIF {
	newInstance := new(RhombusGrid)
	rhombusgrid.CopyBasicFields(newInstance)
	return newInstance
}

func (shapecategory *ShapeCategory) GongCopy() GongstructIF {
	newInstance := new(ShapeCategory)
	shapecategory.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralbezier *SpiralBezier) GongCopy() GongstructIF {
	newInstance := new(SpiralBezier)
	spiralbezier.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralbeziergrid *SpiralBezierGrid) GongCopy() GongstructIF {
	newInstance := new(SpiralBezierGrid)
	spiralbeziergrid.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralcircle *SpiralCircle) GongCopy() GongstructIF {
	newInstance := new(SpiralCircle)
	spiralcircle.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralcirclegrid *SpiralCircleGrid) GongCopy() GongstructIF {
	newInstance := new(SpiralCircleGrid)
	spiralcirclegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralline *SpiralLine) GongCopy() GongstructIF {
	newInstance := new(SpiralLine)
	spiralline.CopyBasicFields(newInstance)
	return newInstance
}

func (spirallinegrid *SpiralLineGrid) GongCopy() GongstructIF {
	newInstance := new(SpiralLineGrid)
	spirallinegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralorigin *SpiralOrigin) GongCopy() GongstructIF {
	newInstance := new(SpiralOrigin)
	spiralorigin.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralrhombus *SpiralRhombus) GongCopy() GongstructIF {
	newInstance := new(SpiralRhombus)
	spiralrhombus.CopyBasicFields(newInstance)
	return newInstance
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongCopy() GongstructIF {
	newInstance := new(SpiralRhombusGrid)
	spiralrhombusgrid.CopyBasicFields(newInstance)
	return newInstance
}

func (verticalaxis *VerticalAxis) GongCopy() GongstructIF {
	newInstance := new(VerticalAxis)
	verticalaxis.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (axis *Axis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(axis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(axis), uint64(GetOrderPointerGongstruct(stage, axis)))
	return
}

func (axisgrid *AxisGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(axisgrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(axisgrid), uint64(GetOrderPointerGongstruct(stage, axisgrid)))
	return
}

func (bezier *Bezier) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bezier).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bezier), uint64(GetOrderPointerGongstruct(stage, bezier)))
	return
}

func (beziergrid *BezierGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beziergrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(beziergrid), uint64(GetOrderPointerGongstruct(stage, beziergrid)))
	return
}

func (beziergridstack *BezierGridStack) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beziergridstack).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(beziergridstack), uint64(GetOrderPointerGongstruct(stage, beziergridstack)))
	return
}

func (chapter *Chapter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(chapter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(chapter), uint64(GetOrderPointerGongstruct(stage, chapter)))
	return
}

func (circle *Circle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(circle), uint64(GetOrderPointerGongstruct(stage, circle)))
	return
}

func (circlegrid *CircleGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circlegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(circlegrid), uint64(GetOrderPointerGongstruct(stage, circlegrid)))
	return
}

func (content *Content) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(content), uint64(GetOrderPointerGongstruct(stage, content)))
	return
}

func (exporttomusicxml *ExportToMusicxml) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(exporttomusicxml).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(exporttomusicxml), uint64(GetOrderPointerGongstruct(stage, exporttomusicxml)))
	return
}

func (frontcurve *FrontCurve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(frontcurve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(frontcurve), uint64(GetOrderPointerGongstruct(stage, frontcurve)))
	return
}

func (frontcurvestack *FrontCurveStack) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(frontcurvestack).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(frontcurvestack), uint64(GetOrderPointerGongstruct(stage, frontcurvestack)))
	return
}

func (horizontalaxis *HorizontalAxis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(horizontalaxis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(horizontalaxis), uint64(GetOrderPointerGongstruct(stage, horizontalaxis)))
	return
}

func (key *Key) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(key).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(key), uint64(GetOrderPointerGongstruct(stage, key)))
	return
}

func (parameter *Parameter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parameter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(parameter), uint64(GetOrderPointerGongstruct(stage, parameter)))
	return
}

func (rhombus *Rhombus) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombus).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rhombus), uint64(GetOrderPointerGongstruct(stage, rhombus)))
	return
}

func (rhombusgrid *RhombusGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombusgrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rhombusgrid), uint64(GetOrderPointerGongstruct(stage, rhombusgrid)))
	return
}

func (shapecategory *ShapeCategory) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shapecategory).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shapecategory), uint64(GetOrderPointerGongstruct(stage, shapecategory)))
	return
}

func (spiralbezier *SpiralBezier) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralbezier).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralbezier), uint64(GetOrderPointerGongstruct(stage, spiralbezier)))
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralbeziergrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralbeziergrid), uint64(GetOrderPointerGongstruct(stage, spiralbeziergrid)))
	return
}

func (spiralcircle *SpiralCircle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralcircle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralcircle), uint64(GetOrderPointerGongstruct(stage, spiralcircle)))
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralcirclegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralcirclegrid), uint64(GetOrderPointerGongstruct(stage, spiralcirclegrid)))
	return
}

func (spiralline *SpiralLine) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralline).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralline), uint64(GetOrderPointerGongstruct(stage, spiralline)))
	return
}

func (spirallinegrid *SpiralLineGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spirallinegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spirallinegrid), uint64(GetOrderPointerGongstruct(stage, spirallinegrid)))
	return
}

func (spiralorigin *SpiralOrigin) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralorigin).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralorigin), uint64(GetOrderPointerGongstruct(stage, spiralorigin)))
	return
}

func (spiralrhombus *SpiralRhombus) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralrhombus).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralrhombus), uint64(GetOrderPointerGongstruct(stage, spiralrhombus)))
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spiralrhombusgrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spiralrhombusgrid), uint64(GetOrderPointerGongstruct(stage, spiralrhombusgrid)))
	return
}

func (verticalaxis *VerticalAxis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(verticalaxis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(verticalaxis), uint64(GetOrderPointerGongstruct(stage, verticalaxis)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var axiss_newInstances []*Axis
	var axiss_deletedInstances []*Axis

	// parse all staged instances and check if they have a reference
	for axis := range stage.Axiss {
		if ref, ok := stage.Axiss_reference[axis]; !ok {
			axiss_newInstances = append(axiss_newInstances, axis)
			newInstancesSlice = append(newInstancesSlice, axis.GongMarshallIdentifier(stage))
			if stage.Axiss_referenceOrder == nil {
				stage.Axiss_referenceOrder = make(map[*Axis]uint)
			}
			stage.Axiss_referenceOrder[axis] = stage.Axis_stagedOrder[axis]
			newInstancesReverseSlice = append(newInstancesReverseSlice, axis.GongMarshallUnstaging(stage))
			// delete(stage.Axiss_referenceOrder, axis)
			fieldInitializers, pointersInitializations := axis.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Axis_stagedOrder[ref] = stage.Axis_stagedOrder[axis]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := axis.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, axis)
			// delete(stage.Axis_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", axis.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Axiss_reference {
		instance := stage.Axiss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Axiss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			axiss_deletedInstances = append(axiss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(axiss_newInstances)
	lenDeletedInstances += len(axiss_deletedInstances)
	var axisgrids_newInstances []*AxisGrid
	var axisgrids_deletedInstances []*AxisGrid

	// parse all staged instances and check if they have a reference
	for axisgrid := range stage.AxisGrids {
		if ref, ok := stage.AxisGrids_reference[axisgrid]; !ok {
			axisgrids_newInstances = append(axisgrids_newInstances, axisgrid)
			newInstancesSlice = append(newInstancesSlice, axisgrid.GongMarshallIdentifier(stage))
			if stage.AxisGrids_referenceOrder == nil {
				stage.AxisGrids_referenceOrder = make(map[*AxisGrid]uint)
			}
			stage.AxisGrids_referenceOrder[axisgrid] = stage.AxisGrid_stagedOrder[axisgrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, axisgrid.GongMarshallUnstaging(stage))
			// delete(stage.AxisGrids_referenceOrder, axisgrid)
			fieldInitializers, pointersInitializations := axisgrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AxisGrid_stagedOrder[ref] = stage.AxisGrid_stagedOrder[axisgrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := axisgrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, axisgrid)
			// delete(stage.AxisGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", axisgrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.AxisGrids_reference {
		instance := stage.AxisGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AxisGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			axisgrids_deletedInstances = append(axisgrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(axisgrids_newInstances)
	lenDeletedInstances += len(axisgrids_deletedInstances)
	var beziers_newInstances []*Bezier
	var beziers_deletedInstances []*Bezier

	// parse all staged instances and check if they have a reference
	for bezier := range stage.Beziers {
		if ref, ok := stage.Beziers_reference[bezier]; !ok {
			beziers_newInstances = append(beziers_newInstances, bezier)
			newInstancesSlice = append(newInstancesSlice, bezier.GongMarshallIdentifier(stage))
			if stage.Beziers_referenceOrder == nil {
				stage.Beziers_referenceOrder = make(map[*Bezier]uint)
			}
			stage.Beziers_referenceOrder[bezier] = stage.Bezier_stagedOrder[bezier]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bezier.GongMarshallUnstaging(stage))
			// delete(stage.Beziers_referenceOrder, bezier)
			fieldInitializers, pointersInitializations := bezier.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Bezier_stagedOrder[ref] = stage.Bezier_stagedOrder[bezier]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bezier.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bezier)
			// delete(stage.Bezier_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", bezier.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Beziers_reference {
		instance := stage.Beziers_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Beziers[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			beziers_deletedInstances = append(beziers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(beziers_newInstances)
	lenDeletedInstances += len(beziers_deletedInstances)
	var beziergrids_newInstances []*BezierGrid
	var beziergrids_deletedInstances []*BezierGrid

	// parse all staged instances and check if they have a reference
	for beziergrid := range stage.BezierGrids {
		if ref, ok := stage.BezierGrids_reference[beziergrid]; !ok {
			beziergrids_newInstances = append(beziergrids_newInstances, beziergrid)
			newInstancesSlice = append(newInstancesSlice, beziergrid.GongMarshallIdentifier(stage))
			if stage.BezierGrids_referenceOrder == nil {
				stage.BezierGrids_referenceOrder = make(map[*BezierGrid]uint)
			}
			stage.BezierGrids_referenceOrder[beziergrid] = stage.BezierGrid_stagedOrder[beziergrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, beziergrid.GongMarshallUnstaging(stage))
			// delete(stage.BezierGrids_referenceOrder, beziergrid)
			fieldInitializers, pointersInitializations := beziergrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BezierGrid_stagedOrder[ref] = stage.BezierGrid_stagedOrder[beziergrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := beziergrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, beziergrid)
			// delete(stage.BezierGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", beziergrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BezierGrids_reference {
		instance := stage.BezierGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BezierGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			beziergrids_deletedInstances = append(beziergrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(beziergrids_newInstances)
	lenDeletedInstances += len(beziergrids_deletedInstances)
	var beziergridstacks_newInstances []*BezierGridStack
	var beziergridstacks_deletedInstances []*BezierGridStack

	// parse all staged instances and check if they have a reference
	for beziergridstack := range stage.BezierGridStacks {
		if ref, ok := stage.BezierGridStacks_reference[beziergridstack]; !ok {
			beziergridstacks_newInstances = append(beziergridstacks_newInstances, beziergridstack)
			newInstancesSlice = append(newInstancesSlice, beziergridstack.GongMarshallIdentifier(stage))
			if stage.BezierGridStacks_referenceOrder == nil {
				stage.BezierGridStacks_referenceOrder = make(map[*BezierGridStack]uint)
			}
			stage.BezierGridStacks_referenceOrder[beziergridstack] = stage.BezierGridStack_stagedOrder[beziergridstack]
			newInstancesReverseSlice = append(newInstancesReverseSlice, beziergridstack.GongMarshallUnstaging(stage))
			// delete(stage.BezierGridStacks_referenceOrder, beziergridstack)
			fieldInitializers, pointersInitializations := beziergridstack.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BezierGridStack_stagedOrder[ref] = stage.BezierGridStack_stagedOrder[beziergridstack]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := beziergridstack.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, beziergridstack)
			// delete(stage.BezierGridStack_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", beziergridstack.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BezierGridStacks_reference {
		instance := stage.BezierGridStacks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BezierGridStacks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			beziergridstacks_deletedInstances = append(beziergridstacks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(beziergridstacks_newInstances)
	lenDeletedInstances += len(beziergridstacks_deletedInstances)
	var chapters_newInstances []*Chapter
	var chapters_deletedInstances []*Chapter

	// parse all staged instances and check if they have a reference
	for chapter := range stage.Chapters {
		if ref, ok := stage.Chapters_reference[chapter]; !ok {
			chapters_newInstances = append(chapters_newInstances, chapter)
			newInstancesSlice = append(newInstancesSlice, chapter.GongMarshallIdentifier(stage))
			if stage.Chapters_referenceOrder == nil {
				stage.Chapters_referenceOrder = make(map[*Chapter]uint)
			}
			stage.Chapters_referenceOrder[chapter] = stage.Chapter_stagedOrder[chapter]
			newInstancesReverseSlice = append(newInstancesReverseSlice, chapter.GongMarshallUnstaging(stage))
			// delete(stage.Chapters_referenceOrder, chapter)
			fieldInitializers, pointersInitializations := chapter.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Chapter_stagedOrder[ref] = stage.Chapter_stagedOrder[chapter]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := chapter.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, chapter)
			// delete(stage.Chapter_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", chapter.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Chapters_reference {
		instance := stage.Chapters_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Chapters[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			chapters_deletedInstances = append(chapters_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(chapters_newInstances)
	lenDeletedInstances += len(chapters_deletedInstances)
	var circles_newInstances []*Circle
	var circles_deletedInstances []*Circle

	// parse all staged instances and check if they have a reference
	for circle := range stage.Circles {
		if ref, ok := stage.Circles_reference[circle]; !ok {
			circles_newInstances = append(circles_newInstances, circle)
			newInstancesSlice = append(newInstancesSlice, circle.GongMarshallIdentifier(stage))
			if stage.Circles_referenceOrder == nil {
				stage.Circles_referenceOrder = make(map[*Circle]uint)
			}
			stage.Circles_referenceOrder[circle] = stage.Circle_stagedOrder[circle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, circle.GongMarshallUnstaging(stage))
			// delete(stage.Circles_referenceOrder, circle)
			fieldInitializers, pointersInitializations := circle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Circle_stagedOrder[ref] = stage.Circle_stagedOrder[circle]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := circle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, circle)
			// delete(stage.Circle_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", circle.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Circles_reference {
		instance := stage.Circles_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Circles[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			circles_deletedInstances = append(circles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(circles_newInstances)
	lenDeletedInstances += len(circles_deletedInstances)
	var circlegrids_newInstances []*CircleGrid
	var circlegrids_deletedInstances []*CircleGrid

	// parse all staged instances and check if they have a reference
	for circlegrid := range stage.CircleGrids {
		if ref, ok := stage.CircleGrids_reference[circlegrid]; !ok {
			circlegrids_newInstances = append(circlegrids_newInstances, circlegrid)
			newInstancesSlice = append(newInstancesSlice, circlegrid.GongMarshallIdentifier(stage))
			if stage.CircleGrids_referenceOrder == nil {
				stage.CircleGrids_referenceOrder = make(map[*CircleGrid]uint)
			}
			stage.CircleGrids_referenceOrder[circlegrid] = stage.CircleGrid_stagedOrder[circlegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, circlegrid.GongMarshallUnstaging(stage))
			// delete(stage.CircleGrids_referenceOrder, circlegrid)
			fieldInitializers, pointersInitializations := circlegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CircleGrid_stagedOrder[ref] = stage.CircleGrid_stagedOrder[circlegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := circlegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, circlegrid)
			// delete(stage.CircleGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", circlegrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.CircleGrids_reference {
		instance := stage.CircleGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CircleGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			circlegrids_deletedInstances = append(circlegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(circlegrids_newInstances)
	lenDeletedInstances += len(circlegrids_deletedInstances)
	var contents_newInstances []*Content
	var contents_deletedInstances []*Content

	// parse all staged instances and check if they have a reference
	for content := range stage.Contents {
		if ref, ok := stage.Contents_reference[content]; !ok {
			contents_newInstances = append(contents_newInstances, content)
			newInstancesSlice = append(newInstancesSlice, content.GongMarshallIdentifier(stage))
			if stage.Contents_referenceOrder == nil {
				stage.Contents_referenceOrder = make(map[*Content]uint)
			}
			stage.Contents_referenceOrder[content] = stage.Content_stagedOrder[content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, content.GongMarshallUnstaging(stage))
			// delete(stage.Contents_referenceOrder, content)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Content_stagedOrder[ref] = stage.Content_stagedOrder[content]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, content)
			// delete(stage.Content_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", content.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Contents_reference {
		instance := stage.Contents_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Contents[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			contents_deletedInstances = append(contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(contents_newInstances)
	lenDeletedInstances += len(contents_deletedInstances)
	var exporttomusicxmls_newInstances []*ExportToMusicxml
	var exporttomusicxmls_deletedInstances []*ExportToMusicxml

	// parse all staged instances and check if they have a reference
	for exporttomusicxml := range stage.ExportToMusicxmls {
		if ref, ok := stage.ExportToMusicxmls_reference[exporttomusicxml]; !ok {
			exporttomusicxmls_newInstances = append(exporttomusicxmls_newInstances, exporttomusicxml)
			newInstancesSlice = append(newInstancesSlice, exporttomusicxml.GongMarshallIdentifier(stage))
			if stage.ExportToMusicxmls_referenceOrder == nil {
				stage.ExportToMusicxmls_referenceOrder = make(map[*ExportToMusicxml]uint)
			}
			stage.ExportToMusicxmls_referenceOrder[exporttomusicxml] = stage.ExportToMusicxml_stagedOrder[exporttomusicxml]
			newInstancesReverseSlice = append(newInstancesReverseSlice, exporttomusicxml.GongMarshallUnstaging(stage))
			// delete(stage.ExportToMusicxmls_referenceOrder, exporttomusicxml)
			fieldInitializers, pointersInitializations := exporttomusicxml.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExportToMusicxml_stagedOrder[ref] = stage.ExportToMusicxml_stagedOrder[exporttomusicxml]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := exporttomusicxml.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, exporttomusicxml)
			// delete(stage.ExportToMusicxml_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", exporttomusicxml.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ExportToMusicxmls_reference {
		instance := stage.ExportToMusicxmls_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ExportToMusicxmls[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			exporttomusicxmls_deletedInstances = append(exporttomusicxmls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(exporttomusicxmls_newInstances)
	lenDeletedInstances += len(exporttomusicxmls_deletedInstances)
	var frontcurves_newInstances []*FrontCurve
	var frontcurves_deletedInstances []*FrontCurve

	// parse all staged instances and check if they have a reference
	for frontcurve := range stage.FrontCurves {
		if ref, ok := stage.FrontCurves_reference[frontcurve]; !ok {
			frontcurves_newInstances = append(frontcurves_newInstances, frontcurve)
			newInstancesSlice = append(newInstancesSlice, frontcurve.GongMarshallIdentifier(stage))
			if stage.FrontCurves_referenceOrder == nil {
				stage.FrontCurves_referenceOrder = make(map[*FrontCurve]uint)
			}
			stage.FrontCurves_referenceOrder[frontcurve] = stage.FrontCurve_stagedOrder[frontcurve]
			newInstancesReverseSlice = append(newInstancesReverseSlice, frontcurve.GongMarshallUnstaging(stage))
			// delete(stage.FrontCurves_referenceOrder, frontcurve)
			fieldInitializers, pointersInitializations := frontcurve.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FrontCurve_stagedOrder[ref] = stage.FrontCurve_stagedOrder[frontcurve]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := frontcurve.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, frontcurve)
			// delete(stage.FrontCurve_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", frontcurve.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FrontCurves_reference {
		instance := stage.FrontCurves_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FrontCurves[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			frontcurves_deletedInstances = append(frontcurves_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(frontcurves_newInstances)
	lenDeletedInstances += len(frontcurves_deletedInstances)
	var frontcurvestacks_newInstances []*FrontCurveStack
	var frontcurvestacks_deletedInstances []*FrontCurveStack

	// parse all staged instances and check if they have a reference
	for frontcurvestack := range stage.FrontCurveStacks {
		if ref, ok := stage.FrontCurveStacks_reference[frontcurvestack]; !ok {
			frontcurvestacks_newInstances = append(frontcurvestacks_newInstances, frontcurvestack)
			newInstancesSlice = append(newInstancesSlice, frontcurvestack.GongMarshallIdentifier(stage))
			if stage.FrontCurveStacks_referenceOrder == nil {
				stage.FrontCurveStacks_referenceOrder = make(map[*FrontCurveStack]uint)
			}
			stage.FrontCurveStacks_referenceOrder[frontcurvestack] = stage.FrontCurveStack_stagedOrder[frontcurvestack]
			newInstancesReverseSlice = append(newInstancesReverseSlice, frontcurvestack.GongMarshallUnstaging(stage))
			// delete(stage.FrontCurveStacks_referenceOrder, frontcurvestack)
			fieldInitializers, pointersInitializations := frontcurvestack.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FrontCurveStack_stagedOrder[ref] = stage.FrontCurveStack_stagedOrder[frontcurvestack]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := frontcurvestack.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, frontcurvestack)
			// delete(stage.FrontCurveStack_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", frontcurvestack.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FrontCurveStacks_reference {
		instance := stage.FrontCurveStacks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FrontCurveStacks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			frontcurvestacks_deletedInstances = append(frontcurvestacks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(frontcurvestacks_newInstances)
	lenDeletedInstances += len(frontcurvestacks_deletedInstances)
	var horizontalaxiss_newInstances []*HorizontalAxis
	var horizontalaxiss_deletedInstances []*HorizontalAxis

	// parse all staged instances and check if they have a reference
	for horizontalaxis := range stage.HorizontalAxiss {
		if ref, ok := stage.HorizontalAxiss_reference[horizontalaxis]; !ok {
			horizontalaxiss_newInstances = append(horizontalaxiss_newInstances, horizontalaxis)
			newInstancesSlice = append(newInstancesSlice, horizontalaxis.GongMarshallIdentifier(stage))
			if stage.HorizontalAxiss_referenceOrder == nil {
				stage.HorizontalAxiss_referenceOrder = make(map[*HorizontalAxis]uint)
			}
			stage.HorizontalAxiss_referenceOrder[horizontalaxis] = stage.HorizontalAxis_stagedOrder[horizontalaxis]
			newInstancesReverseSlice = append(newInstancesReverseSlice, horizontalaxis.GongMarshallUnstaging(stage))
			// delete(stage.HorizontalAxiss_referenceOrder, horizontalaxis)
			fieldInitializers, pointersInitializations := horizontalaxis.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.HorizontalAxis_stagedOrder[ref] = stage.HorizontalAxis_stagedOrder[horizontalaxis]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := horizontalaxis.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, horizontalaxis)
			// delete(stage.HorizontalAxis_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", horizontalaxis.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.HorizontalAxiss_reference {
		instance := stage.HorizontalAxiss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.HorizontalAxiss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			horizontalaxiss_deletedInstances = append(horizontalaxiss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(horizontalaxiss_newInstances)
	lenDeletedInstances += len(horizontalaxiss_deletedInstances)
	var keys_newInstances []*Key
	var keys_deletedInstances []*Key

	// parse all staged instances and check if they have a reference
	for key := range stage.Keys {
		if ref, ok := stage.Keys_reference[key]; !ok {
			keys_newInstances = append(keys_newInstances, key)
			newInstancesSlice = append(newInstancesSlice, key.GongMarshallIdentifier(stage))
			if stage.Keys_referenceOrder == nil {
				stage.Keys_referenceOrder = make(map[*Key]uint)
			}
			stage.Keys_referenceOrder[key] = stage.Key_stagedOrder[key]
			newInstancesReverseSlice = append(newInstancesReverseSlice, key.GongMarshallUnstaging(stage))
			// delete(stage.Keys_referenceOrder, key)
			fieldInitializers, pointersInitializations := key.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Key_stagedOrder[ref] = stage.Key_stagedOrder[key]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := key.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, key)
			// delete(stage.Key_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", key.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Keys_reference {
		instance := stage.Keys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Keys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			keys_deletedInstances = append(keys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(keys_newInstances)
	lenDeletedInstances += len(keys_deletedInstances)
	var parameters_newInstances []*Parameter
	var parameters_deletedInstances []*Parameter

	// parse all staged instances and check if they have a reference
	for parameter := range stage.Parameters {
		if ref, ok := stage.Parameters_reference[parameter]; !ok {
			parameters_newInstances = append(parameters_newInstances, parameter)
			newInstancesSlice = append(newInstancesSlice, parameter.GongMarshallIdentifier(stage))
			if stage.Parameters_referenceOrder == nil {
				stage.Parameters_referenceOrder = make(map[*Parameter]uint)
			}
			stage.Parameters_referenceOrder[parameter] = stage.Parameter_stagedOrder[parameter]
			newInstancesReverseSlice = append(newInstancesReverseSlice, parameter.GongMarshallUnstaging(stage))
			// delete(stage.Parameters_referenceOrder, parameter)
			fieldInitializers, pointersInitializations := parameter.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Parameter_stagedOrder[ref] = stage.Parameter_stagedOrder[parameter]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := parameter.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, parameter)
			// delete(stage.Parameter_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", parameter.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Parameters_reference {
		instance := stage.Parameters_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Parameters[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parameters_deletedInstances = append(parameters_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parameters_newInstances)
	lenDeletedInstances += len(parameters_deletedInstances)
	var rhombuss_newInstances []*Rhombus
	var rhombuss_deletedInstances []*Rhombus

	// parse all staged instances and check if they have a reference
	for rhombus := range stage.Rhombuss {
		if ref, ok := stage.Rhombuss_reference[rhombus]; !ok {
			rhombuss_newInstances = append(rhombuss_newInstances, rhombus)
			newInstancesSlice = append(newInstancesSlice, rhombus.GongMarshallIdentifier(stage))
			if stage.Rhombuss_referenceOrder == nil {
				stage.Rhombuss_referenceOrder = make(map[*Rhombus]uint)
			}
			stage.Rhombuss_referenceOrder[rhombus] = stage.Rhombus_stagedOrder[rhombus]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rhombus.GongMarshallUnstaging(stage))
			// delete(stage.Rhombuss_referenceOrder, rhombus)
			fieldInitializers, pointersInitializations := rhombus.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Rhombus_stagedOrder[ref] = stage.Rhombus_stagedOrder[rhombus]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rhombus.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rhombus)
			// delete(stage.Rhombus_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rhombus.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Rhombuss_reference {
		instance := stage.Rhombuss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Rhombuss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rhombuss_deletedInstances = append(rhombuss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rhombuss_newInstances)
	lenDeletedInstances += len(rhombuss_deletedInstances)
	var rhombusgrids_newInstances []*RhombusGrid
	var rhombusgrids_deletedInstances []*RhombusGrid

	// parse all staged instances and check if they have a reference
	for rhombusgrid := range stage.RhombusGrids {
		if ref, ok := stage.RhombusGrids_reference[rhombusgrid]; !ok {
			rhombusgrids_newInstances = append(rhombusgrids_newInstances, rhombusgrid)
			newInstancesSlice = append(newInstancesSlice, rhombusgrid.GongMarshallIdentifier(stage))
			if stage.RhombusGrids_referenceOrder == nil {
				stage.RhombusGrids_referenceOrder = make(map[*RhombusGrid]uint)
			}
			stage.RhombusGrids_referenceOrder[rhombusgrid] = stage.RhombusGrid_stagedOrder[rhombusgrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rhombusgrid.GongMarshallUnstaging(stage))
			// delete(stage.RhombusGrids_referenceOrder, rhombusgrid)
			fieldInitializers, pointersInitializations := rhombusgrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RhombusGrid_stagedOrder[ref] = stage.RhombusGrid_stagedOrder[rhombusgrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rhombusgrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rhombusgrid)
			// delete(stage.RhombusGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rhombusgrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.RhombusGrids_reference {
		instance := stage.RhombusGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RhombusGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rhombusgrids_deletedInstances = append(rhombusgrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rhombusgrids_newInstances)
	lenDeletedInstances += len(rhombusgrids_deletedInstances)
	var shapecategorys_newInstances []*ShapeCategory
	var shapecategorys_deletedInstances []*ShapeCategory

	// parse all staged instances and check if they have a reference
	for shapecategory := range stage.ShapeCategorys {
		if ref, ok := stage.ShapeCategorys_reference[shapecategory]; !ok {
			shapecategorys_newInstances = append(shapecategorys_newInstances, shapecategory)
			newInstancesSlice = append(newInstancesSlice, shapecategory.GongMarshallIdentifier(stage))
			if stage.ShapeCategorys_referenceOrder == nil {
				stage.ShapeCategorys_referenceOrder = make(map[*ShapeCategory]uint)
			}
			stage.ShapeCategorys_referenceOrder[shapecategory] = stage.ShapeCategory_stagedOrder[shapecategory]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shapecategory.GongMarshallUnstaging(stage))
			// delete(stage.ShapeCategorys_referenceOrder, shapecategory)
			fieldInitializers, pointersInitializations := shapecategory.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShapeCategory_stagedOrder[ref] = stage.ShapeCategory_stagedOrder[shapecategory]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shapecategory.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shapecategory)
			// delete(stage.ShapeCategory_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", shapecategory.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ShapeCategorys_reference {
		instance := stage.ShapeCategorys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShapeCategorys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shapecategorys_deletedInstances = append(shapecategorys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shapecategorys_newInstances)
	lenDeletedInstances += len(shapecategorys_deletedInstances)
	var spiralbeziers_newInstances []*SpiralBezier
	var spiralbeziers_deletedInstances []*SpiralBezier

	// parse all staged instances and check if they have a reference
	for spiralbezier := range stage.SpiralBeziers {
		if ref, ok := stage.SpiralBeziers_reference[spiralbezier]; !ok {
			spiralbeziers_newInstances = append(spiralbeziers_newInstances, spiralbezier)
			newInstancesSlice = append(newInstancesSlice, spiralbezier.GongMarshallIdentifier(stage))
			if stage.SpiralBeziers_referenceOrder == nil {
				stage.SpiralBeziers_referenceOrder = make(map[*SpiralBezier]uint)
			}
			stage.SpiralBeziers_referenceOrder[spiralbezier] = stage.SpiralBezier_stagedOrder[spiralbezier]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralbezier.GongMarshallUnstaging(stage))
			// delete(stage.SpiralBeziers_referenceOrder, spiralbezier)
			fieldInitializers, pointersInitializations := spiralbezier.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralBezier_stagedOrder[ref] = stage.SpiralBezier_stagedOrder[spiralbezier]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralbezier.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralbezier)
			// delete(stage.SpiralBezier_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralbezier.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralBeziers_reference {
		instance := stage.SpiralBeziers_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralBeziers[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralbeziers_deletedInstances = append(spiralbeziers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralbeziers_newInstances)
	lenDeletedInstances += len(spiralbeziers_deletedInstances)
	var spiralbeziergrids_newInstances []*SpiralBezierGrid
	var spiralbeziergrids_deletedInstances []*SpiralBezierGrid

	// parse all staged instances and check if they have a reference
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		if ref, ok := stage.SpiralBezierGrids_reference[spiralbeziergrid]; !ok {
			spiralbeziergrids_newInstances = append(spiralbeziergrids_newInstances, spiralbeziergrid)
			newInstancesSlice = append(newInstancesSlice, spiralbeziergrid.GongMarshallIdentifier(stage))
			if stage.SpiralBezierGrids_referenceOrder == nil {
				stage.SpiralBezierGrids_referenceOrder = make(map[*SpiralBezierGrid]uint)
			}
			stage.SpiralBezierGrids_referenceOrder[spiralbeziergrid] = stage.SpiralBezierGrid_stagedOrder[spiralbeziergrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralbeziergrid.GongMarshallUnstaging(stage))
			// delete(stage.SpiralBezierGrids_referenceOrder, spiralbeziergrid)
			fieldInitializers, pointersInitializations := spiralbeziergrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralBezierGrid_stagedOrder[ref] = stage.SpiralBezierGrid_stagedOrder[spiralbeziergrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralbeziergrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralbeziergrid)
			// delete(stage.SpiralBezierGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralbeziergrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralBezierGrids_reference {
		instance := stage.SpiralBezierGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralBezierGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralbeziergrids_deletedInstances = append(spiralbeziergrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralbeziergrids_newInstances)
	lenDeletedInstances += len(spiralbeziergrids_deletedInstances)
	var spiralcircles_newInstances []*SpiralCircle
	var spiralcircles_deletedInstances []*SpiralCircle

	// parse all staged instances and check if they have a reference
	for spiralcircle := range stage.SpiralCircles {
		if ref, ok := stage.SpiralCircles_reference[spiralcircle]; !ok {
			spiralcircles_newInstances = append(spiralcircles_newInstances, spiralcircle)
			newInstancesSlice = append(newInstancesSlice, spiralcircle.GongMarshallIdentifier(stage))
			if stage.SpiralCircles_referenceOrder == nil {
				stage.SpiralCircles_referenceOrder = make(map[*SpiralCircle]uint)
			}
			stage.SpiralCircles_referenceOrder[spiralcircle] = stage.SpiralCircle_stagedOrder[spiralcircle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralcircle.GongMarshallUnstaging(stage))
			// delete(stage.SpiralCircles_referenceOrder, spiralcircle)
			fieldInitializers, pointersInitializations := spiralcircle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralCircle_stagedOrder[ref] = stage.SpiralCircle_stagedOrder[spiralcircle]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralcircle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralcircle)
			// delete(stage.SpiralCircle_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralcircle.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralCircles_reference {
		instance := stage.SpiralCircles_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralCircles[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralcircles_deletedInstances = append(spiralcircles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralcircles_newInstances)
	lenDeletedInstances += len(spiralcircles_deletedInstances)
	var spiralcirclegrids_newInstances []*SpiralCircleGrid
	var spiralcirclegrids_deletedInstances []*SpiralCircleGrid

	// parse all staged instances and check if they have a reference
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		if ref, ok := stage.SpiralCircleGrids_reference[spiralcirclegrid]; !ok {
			spiralcirclegrids_newInstances = append(spiralcirclegrids_newInstances, spiralcirclegrid)
			newInstancesSlice = append(newInstancesSlice, spiralcirclegrid.GongMarshallIdentifier(stage))
			if stage.SpiralCircleGrids_referenceOrder == nil {
				stage.SpiralCircleGrids_referenceOrder = make(map[*SpiralCircleGrid]uint)
			}
			stage.SpiralCircleGrids_referenceOrder[spiralcirclegrid] = stage.SpiralCircleGrid_stagedOrder[spiralcirclegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralcirclegrid.GongMarshallUnstaging(stage))
			// delete(stage.SpiralCircleGrids_referenceOrder, spiralcirclegrid)
			fieldInitializers, pointersInitializations := spiralcirclegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralCircleGrid_stagedOrder[ref] = stage.SpiralCircleGrid_stagedOrder[spiralcirclegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralcirclegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralcirclegrid)
			// delete(stage.SpiralCircleGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralcirclegrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralCircleGrids_reference {
		instance := stage.SpiralCircleGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralCircleGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralcirclegrids_deletedInstances = append(spiralcirclegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralcirclegrids_newInstances)
	lenDeletedInstances += len(spiralcirclegrids_deletedInstances)
	var spirallines_newInstances []*SpiralLine
	var spirallines_deletedInstances []*SpiralLine

	// parse all staged instances and check if they have a reference
	for spiralline := range stage.SpiralLines {
		if ref, ok := stage.SpiralLines_reference[spiralline]; !ok {
			spirallines_newInstances = append(spirallines_newInstances, spiralline)
			newInstancesSlice = append(newInstancesSlice, spiralline.GongMarshallIdentifier(stage))
			if stage.SpiralLines_referenceOrder == nil {
				stage.SpiralLines_referenceOrder = make(map[*SpiralLine]uint)
			}
			stage.SpiralLines_referenceOrder[spiralline] = stage.SpiralLine_stagedOrder[spiralline]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralline.GongMarshallUnstaging(stage))
			// delete(stage.SpiralLines_referenceOrder, spiralline)
			fieldInitializers, pointersInitializations := spiralline.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralLine_stagedOrder[ref] = stage.SpiralLine_stagedOrder[spiralline]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralline.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralline)
			// delete(stage.SpiralLine_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralline.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralLines_reference {
		instance := stage.SpiralLines_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralLines[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spirallines_deletedInstances = append(spirallines_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spirallines_newInstances)
	lenDeletedInstances += len(spirallines_deletedInstances)
	var spirallinegrids_newInstances []*SpiralLineGrid
	var spirallinegrids_deletedInstances []*SpiralLineGrid

	// parse all staged instances and check if they have a reference
	for spirallinegrid := range stage.SpiralLineGrids {
		if ref, ok := stage.SpiralLineGrids_reference[spirallinegrid]; !ok {
			spirallinegrids_newInstances = append(spirallinegrids_newInstances, spirallinegrid)
			newInstancesSlice = append(newInstancesSlice, spirallinegrid.GongMarshallIdentifier(stage))
			if stage.SpiralLineGrids_referenceOrder == nil {
				stage.SpiralLineGrids_referenceOrder = make(map[*SpiralLineGrid]uint)
			}
			stage.SpiralLineGrids_referenceOrder[spirallinegrid] = stage.SpiralLineGrid_stagedOrder[spirallinegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spirallinegrid.GongMarshallUnstaging(stage))
			// delete(stage.SpiralLineGrids_referenceOrder, spirallinegrid)
			fieldInitializers, pointersInitializations := spirallinegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralLineGrid_stagedOrder[ref] = stage.SpiralLineGrid_stagedOrder[spirallinegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spirallinegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spirallinegrid)
			// delete(stage.SpiralLineGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spirallinegrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralLineGrids_reference {
		instance := stage.SpiralLineGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralLineGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spirallinegrids_deletedInstances = append(spirallinegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spirallinegrids_newInstances)
	lenDeletedInstances += len(spirallinegrids_deletedInstances)
	var spiralorigins_newInstances []*SpiralOrigin
	var spiralorigins_deletedInstances []*SpiralOrigin

	// parse all staged instances and check if they have a reference
	for spiralorigin := range stage.SpiralOrigins {
		if ref, ok := stage.SpiralOrigins_reference[spiralorigin]; !ok {
			spiralorigins_newInstances = append(spiralorigins_newInstances, spiralorigin)
			newInstancesSlice = append(newInstancesSlice, spiralorigin.GongMarshallIdentifier(stage))
			if stage.SpiralOrigins_referenceOrder == nil {
				stage.SpiralOrigins_referenceOrder = make(map[*SpiralOrigin]uint)
			}
			stage.SpiralOrigins_referenceOrder[spiralorigin] = stage.SpiralOrigin_stagedOrder[spiralorigin]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralorigin.GongMarshallUnstaging(stage))
			// delete(stage.SpiralOrigins_referenceOrder, spiralorigin)
			fieldInitializers, pointersInitializations := spiralorigin.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralOrigin_stagedOrder[ref] = stage.SpiralOrigin_stagedOrder[spiralorigin]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralorigin.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralorigin)
			// delete(stage.SpiralOrigin_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralorigin.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralOrigins_reference {
		instance := stage.SpiralOrigins_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralOrigins[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralorigins_deletedInstances = append(spiralorigins_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralorigins_newInstances)
	lenDeletedInstances += len(spiralorigins_deletedInstances)
	var spiralrhombuss_newInstances []*SpiralRhombus
	var spiralrhombuss_deletedInstances []*SpiralRhombus

	// parse all staged instances and check if they have a reference
	for spiralrhombus := range stage.SpiralRhombuss {
		if ref, ok := stage.SpiralRhombuss_reference[spiralrhombus]; !ok {
			spiralrhombuss_newInstances = append(spiralrhombuss_newInstances, spiralrhombus)
			newInstancesSlice = append(newInstancesSlice, spiralrhombus.GongMarshallIdentifier(stage))
			if stage.SpiralRhombuss_referenceOrder == nil {
				stage.SpiralRhombuss_referenceOrder = make(map[*SpiralRhombus]uint)
			}
			stage.SpiralRhombuss_referenceOrder[spiralrhombus] = stage.SpiralRhombus_stagedOrder[spiralrhombus]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralrhombus.GongMarshallUnstaging(stage))
			// delete(stage.SpiralRhombuss_referenceOrder, spiralrhombus)
			fieldInitializers, pointersInitializations := spiralrhombus.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralRhombus_stagedOrder[ref] = stage.SpiralRhombus_stagedOrder[spiralrhombus]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralrhombus.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralrhombus)
			// delete(stage.SpiralRhombus_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralrhombus.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralRhombuss_reference {
		instance := stage.SpiralRhombuss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralRhombuss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralrhombuss_deletedInstances = append(spiralrhombuss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralrhombuss_newInstances)
	lenDeletedInstances += len(spiralrhombuss_deletedInstances)
	var spiralrhombusgrids_newInstances []*SpiralRhombusGrid
	var spiralrhombusgrids_deletedInstances []*SpiralRhombusGrid

	// parse all staged instances and check if they have a reference
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		if ref, ok := stage.SpiralRhombusGrids_reference[spiralrhombusgrid]; !ok {
			spiralrhombusgrids_newInstances = append(spiralrhombusgrids_newInstances, spiralrhombusgrid)
			newInstancesSlice = append(newInstancesSlice, spiralrhombusgrid.GongMarshallIdentifier(stage))
			if stage.SpiralRhombusGrids_referenceOrder == nil {
				stage.SpiralRhombusGrids_referenceOrder = make(map[*SpiralRhombusGrid]uint)
			}
			stage.SpiralRhombusGrids_referenceOrder[spiralrhombusgrid] = stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spiralrhombusgrid.GongMarshallUnstaging(stage))
			// delete(stage.SpiralRhombusGrids_referenceOrder, spiralrhombusgrid)
			fieldInitializers, pointersInitializations := spiralrhombusgrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SpiralRhombusGrid_stagedOrder[ref] = stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spiralrhombusgrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spiralrhombusgrid)
			// delete(stage.SpiralRhombusGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spiralrhombusgrid.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SpiralRhombusGrids_reference {
		instance := stage.SpiralRhombusGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SpiralRhombusGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spiralrhombusgrids_deletedInstances = append(spiralrhombusgrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spiralrhombusgrids_newInstances)
	lenDeletedInstances += len(spiralrhombusgrids_deletedInstances)
	var verticalaxiss_newInstances []*VerticalAxis
	var verticalaxiss_deletedInstances []*VerticalAxis

	// parse all staged instances and check if they have a reference
	for verticalaxis := range stage.VerticalAxiss {
		if ref, ok := stage.VerticalAxiss_reference[verticalaxis]; !ok {
			verticalaxiss_newInstances = append(verticalaxiss_newInstances, verticalaxis)
			newInstancesSlice = append(newInstancesSlice, verticalaxis.GongMarshallIdentifier(stage))
			if stage.VerticalAxiss_referenceOrder == nil {
				stage.VerticalAxiss_referenceOrder = make(map[*VerticalAxis]uint)
			}
			stage.VerticalAxiss_referenceOrder[verticalaxis] = stage.VerticalAxis_stagedOrder[verticalaxis]
			newInstancesReverseSlice = append(newInstancesReverseSlice, verticalaxis.GongMarshallUnstaging(stage))
			// delete(stage.VerticalAxiss_referenceOrder, verticalaxis)
			fieldInitializers, pointersInitializations := verticalaxis.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.VerticalAxis_stagedOrder[ref] = stage.VerticalAxis_stagedOrder[verticalaxis]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := verticalaxis.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, verticalaxis)
			// delete(stage.VerticalAxis_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", verticalaxis.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.VerticalAxiss_reference {
		instance := stage.VerticalAxiss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.VerticalAxiss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			verticalaxiss_deletedInstances = append(verticalaxiss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(verticalaxiss_newInstances)
	lenDeletedInstances += len(verticalaxiss_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.Axiss_reference = make(map[*Axis]*Axis)
	stage.Axiss_referenceOrder = make(map[*Axis]uint) // diff Unstage needs the reference order
	stage.Axiss_instance = make(map[*Axis]*Axis)
	for instance := range stage.Axiss {
		_copy := instance.GongCopy().(*Axis)
		stage.Axiss_reference[instance] = _copy
		stage.Axiss_instance[_copy] = instance
		stage.Axiss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AxisGrids_reference = make(map[*AxisGrid]*AxisGrid)
	stage.AxisGrids_referenceOrder = make(map[*AxisGrid]uint) // diff Unstage needs the reference order
	stage.AxisGrids_instance = make(map[*AxisGrid]*AxisGrid)
	for instance := range stage.AxisGrids {
		_copy := instance.GongCopy().(*AxisGrid)
		stage.AxisGrids_reference[instance] = _copy
		stage.AxisGrids_instance[_copy] = instance
		stage.AxisGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Beziers_reference = make(map[*Bezier]*Bezier)
	stage.Beziers_referenceOrder = make(map[*Bezier]uint) // diff Unstage needs the reference order
	stage.Beziers_instance = make(map[*Bezier]*Bezier)
	for instance := range stage.Beziers {
		_copy := instance.GongCopy().(*Bezier)
		stage.Beziers_reference[instance] = _copy
		stage.Beziers_instance[_copy] = instance
		stage.Beziers_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BezierGrids_reference = make(map[*BezierGrid]*BezierGrid)
	stage.BezierGrids_referenceOrder = make(map[*BezierGrid]uint) // diff Unstage needs the reference order
	stage.BezierGrids_instance = make(map[*BezierGrid]*BezierGrid)
	for instance := range stage.BezierGrids {
		_copy := instance.GongCopy().(*BezierGrid)
		stage.BezierGrids_reference[instance] = _copy
		stage.BezierGrids_instance[_copy] = instance
		stage.BezierGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BezierGridStacks_reference = make(map[*BezierGridStack]*BezierGridStack)
	stage.BezierGridStacks_referenceOrder = make(map[*BezierGridStack]uint) // diff Unstage needs the reference order
	stage.BezierGridStacks_instance = make(map[*BezierGridStack]*BezierGridStack)
	for instance := range stage.BezierGridStacks {
		_copy := instance.GongCopy().(*BezierGridStack)
		stage.BezierGridStacks_reference[instance] = _copy
		stage.BezierGridStacks_instance[_copy] = instance
		stage.BezierGridStacks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint) // diff Unstage needs the reference order
	stage.Chapters_instance = make(map[*Chapter]*Chapter)
	for instance := range stage.Chapters {
		_copy := instance.GongCopy().(*Chapter)
		stage.Chapters_reference[instance] = _copy
		stage.Chapters_instance[_copy] = instance
		stage.Chapters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Circles_reference = make(map[*Circle]*Circle)
	stage.Circles_referenceOrder = make(map[*Circle]uint) // diff Unstage needs the reference order
	stage.Circles_instance = make(map[*Circle]*Circle)
	for instance := range stage.Circles {
		_copy := instance.GongCopy().(*Circle)
		stage.Circles_reference[instance] = _copy
		stage.Circles_instance[_copy] = instance
		stage.Circles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CircleGrids_reference = make(map[*CircleGrid]*CircleGrid)
	stage.CircleGrids_referenceOrder = make(map[*CircleGrid]uint) // diff Unstage needs the reference order
	stage.CircleGrids_instance = make(map[*CircleGrid]*CircleGrid)
	for instance := range stage.CircleGrids {
		_copy := instance.GongCopy().(*CircleGrid)
		stage.CircleGrids_reference[instance] = _copy
		stage.CircleGrids_instance[_copy] = instance
		stage.CircleGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint) // diff Unstage needs the reference order
	stage.Contents_instance = make(map[*Content]*Content)
	for instance := range stage.Contents {
		_copy := instance.GongCopy().(*Content)
		stage.Contents_reference[instance] = _copy
		stage.Contents_instance[_copy] = instance
		stage.Contents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ExportToMusicxmls_reference = make(map[*ExportToMusicxml]*ExportToMusicxml)
	stage.ExportToMusicxmls_referenceOrder = make(map[*ExportToMusicxml]uint) // diff Unstage needs the reference order
	stage.ExportToMusicxmls_instance = make(map[*ExportToMusicxml]*ExportToMusicxml)
	for instance := range stage.ExportToMusicxmls {
		_copy := instance.GongCopy().(*ExportToMusicxml)
		stage.ExportToMusicxmls_reference[instance] = _copy
		stage.ExportToMusicxmls_instance[_copy] = instance
		stage.ExportToMusicxmls_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FrontCurves_reference = make(map[*FrontCurve]*FrontCurve)
	stage.FrontCurves_referenceOrder = make(map[*FrontCurve]uint) // diff Unstage needs the reference order
	stage.FrontCurves_instance = make(map[*FrontCurve]*FrontCurve)
	for instance := range stage.FrontCurves {
		_copy := instance.GongCopy().(*FrontCurve)
		stage.FrontCurves_reference[instance] = _copy
		stage.FrontCurves_instance[_copy] = instance
		stage.FrontCurves_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FrontCurveStacks_reference = make(map[*FrontCurveStack]*FrontCurveStack)
	stage.FrontCurveStacks_referenceOrder = make(map[*FrontCurveStack]uint) // diff Unstage needs the reference order
	stage.FrontCurveStacks_instance = make(map[*FrontCurveStack]*FrontCurveStack)
	for instance := range stage.FrontCurveStacks {
		_copy := instance.GongCopy().(*FrontCurveStack)
		stage.FrontCurveStacks_reference[instance] = _copy
		stage.FrontCurveStacks_instance[_copy] = instance
		stage.FrontCurveStacks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.HorizontalAxiss_reference = make(map[*HorizontalAxis]*HorizontalAxis)
	stage.HorizontalAxiss_referenceOrder = make(map[*HorizontalAxis]uint) // diff Unstage needs the reference order
	stage.HorizontalAxiss_instance = make(map[*HorizontalAxis]*HorizontalAxis)
	for instance := range stage.HorizontalAxiss {
		_copy := instance.GongCopy().(*HorizontalAxis)
		stage.HorizontalAxiss_reference[instance] = _copy
		stage.HorizontalAxiss_instance[_copy] = instance
		stage.HorizontalAxiss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Keys_reference = make(map[*Key]*Key)
	stage.Keys_referenceOrder = make(map[*Key]uint) // diff Unstage needs the reference order
	stage.Keys_instance = make(map[*Key]*Key)
	for instance := range stage.Keys {
		_copy := instance.GongCopy().(*Key)
		stage.Keys_reference[instance] = _copy
		stage.Keys_instance[_copy] = instance
		stage.Keys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Parameters_reference = make(map[*Parameter]*Parameter)
	stage.Parameters_referenceOrder = make(map[*Parameter]uint) // diff Unstage needs the reference order
	stage.Parameters_instance = make(map[*Parameter]*Parameter)
	for instance := range stage.Parameters {
		_copy := instance.GongCopy().(*Parameter)
		stage.Parameters_reference[instance] = _copy
		stage.Parameters_instance[_copy] = instance
		stage.Parameters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Rhombuss_reference = make(map[*Rhombus]*Rhombus)
	stage.Rhombuss_referenceOrder = make(map[*Rhombus]uint) // diff Unstage needs the reference order
	stage.Rhombuss_instance = make(map[*Rhombus]*Rhombus)
	for instance := range stage.Rhombuss {
		_copy := instance.GongCopy().(*Rhombus)
		stage.Rhombuss_reference[instance] = _copy
		stage.Rhombuss_instance[_copy] = instance
		stage.Rhombuss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RhombusGrids_reference = make(map[*RhombusGrid]*RhombusGrid)
	stage.RhombusGrids_referenceOrder = make(map[*RhombusGrid]uint) // diff Unstage needs the reference order
	stage.RhombusGrids_instance = make(map[*RhombusGrid]*RhombusGrid)
	for instance := range stage.RhombusGrids {
		_copy := instance.GongCopy().(*RhombusGrid)
		stage.RhombusGrids_reference[instance] = _copy
		stage.RhombusGrids_instance[_copy] = instance
		stage.RhombusGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShapeCategorys_reference = make(map[*ShapeCategory]*ShapeCategory)
	stage.ShapeCategorys_referenceOrder = make(map[*ShapeCategory]uint) // diff Unstage needs the reference order
	stage.ShapeCategorys_instance = make(map[*ShapeCategory]*ShapeCategory)
	for instance := range stage.ShapeCategorys {
		_copy := instance.GongCopy().(*ShapeCategory)
		stage.ShapeCategorys_reference[instance] = _copy
		stage.ShapeCategorys_instance[_copy] = instance
		stage.ShapeCategorys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralBeziers_reference = make(map[*SpiralBezier]*SpiralBezier)
	stage.SpiralBeziers_referenceOrder = make(map[*SpiralBezier]uint) // diff Unstage needs the reference order
	stage.SpiralBeziers_instance = make(map[*SpiralBezier]*SpiralBezier)
	for instance := range stage.SpiralBeziers {
		_copy := instance.GongCopy().(*SpiralBezier)
		stage.SpiralBeziers_reference[instance] = _copy
		stage.SpiralBeziers_instance[_copy] = instance
		stage.SpiralBeziers_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralBezierGrids_reference = make(map[*SpiralBezierGrid]*SpiralBezierGrid)
	stage.SpiralBezierGrids_referenceOrder = make(map[*SpiralBezierGrid]uint) // diff Unstage needs the reference order
	stage.SpiralBezierGrids_instance = make(map[*SpiralBezierGrid]*SpiralBezierGrid)
	for instance := range stage.SpiralBezierGrids {
		_copy := instance.GongCopy().(*SpiralBezierGrid)
		stage.SpiralBezierGrids_reference[instance] = _copy
		stage.SpiralBezierGrids_instance[_copy] = instance
		stage.SpiralBezierGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralCircles_reference = make(map[*SpiralCircle]*SpiralCircle)
	stage.SpiralCircles_referenceOrder = make(map[*SpiralCircle]uint) // diff Unstage needs the reference order
	stage.SpiralCircles_instance = make(map[*SpiralCircle]*SpiralCircle)
	for instance := range stage.SpiralCircles {
		_copy := instance.GongCopy().(*SpiralCircle)
		stage.SpiralCircles_reference[instance] = _copy
		stage.SpiralCircles_instance[_copy] = instance
		stage.SpiralCircles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralCircleGrids_reference = make(map[*SpiralCircleGrid]*SpiralCircleGrid)
	stage.SpiralCircleGrids_referenceOrder = make(map[*SpiralCircleGrid]uint) // diff Unstage needs the reference order
	stage.SpiralCircleGrids_instance = make(map[*SpiralCircleGrid]*SpiralCircleGrid)
	for instance := range stage.SpiralCircleGrids {
		_copy := instance.GongCopy().(*SpiralCircleGrid)
		stage.SpiralCircleGrids_reference[instance] = _copy
		stage.SpiralCircleGrids_instance[_copy] = instance
		stage.SpiralCircleGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralLines_reference = make(map[*SpiralLine]*SpiralLine)
	stage.SpiralLines_referenceOrder = make(map[*SpiralLine]uint) // diff Unstage needs the reference order
	stage.SpiralLines_instance = make(map[*SpiralLine]*SpiralLine)
	for instance := range stage.SpiralLines {
		_copy := instance.GongCopy().(*SpiralLine)
		stage.SpiralLines_reference[instance] = _copy
		stage.SpiralLines_instance[_copy] = instance
		stage.SpiralLines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralLineGrids_reference = make(map[*SpiralLineGrid]*SpiralLineGrid)
	stage.SpiralLineGrids_referenceOrder = make(map[*SpiralLineGrid]uint) // diff Unstage needs the reference order
	stage.SpiralLineGrids_instance = make(map[*SpiralLineGrid]*SpiralLineGrid)
	for instance := range stage.SpiralLineGrids {
		_copy := instance.GongCopy().(*SpiralLineGrid)
		stage.SpiralLineGrids_reference[instance] = _copy
		stage.SpiralLineGrids_instance[_copy] = instance
		stage.SpiralLineGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralOrigins_reference = make(map[*SpiralOrigin]*SpiralOrigin)
	stage.SpiralOrigins_referenceOrder = make(map[*SpiralOrigin]uint) // diff Unstage needs the reference order
	stage.SpiralOrigins_instance = make(map[*SpiralOrigin]*SpiralOrigin)
	for instance := range stage.SpiralOrigins {
		_copy := instance.GongCopy().(*SpiralOrigin)
		stage.SpiralOrigins_reference[instance] = _copy
		stage.SpiralOrigins_instance[_copy] = instance
		stage.SpiralOrigins_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralRhombuss_reference = make(map[*SpiralRhombus]*SpiralRhombus)
	stage.SpiralRhombuss_referenceOrder = make(map[*SpiralRhombus]uint) // diff Unstage needs the reference order
	stage.SpiralRhombuss_instance = make(map[*SpiralRhombus]*SpiralRhombus)
	for instance := range stage.SpiralRhombuss {
		_copy := instance.GongCopy().(*SpiralRhombus)
		stage.SpiralRhombuss_reference[instance] = _copy
		stage.SpiralRhombuss_instance[_copy] = instance
		stage.SpiralRhombuss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SpiralRhombusGrids_reference = make(map[*SpiralRhombusGrid]*SpiralRhombusGrid)
	stage.SpiralRhombusGrids_referenceOrder = make(map[*SpiralRhombusGrid]uint) // diff Unstage needs the reference order
	stage.SpiralRhombusGrids_instance = make(map[*SpiralRhombusGrid]*SpiralRhombusGrid)
	for instance := range stage.SpiralRhombusGrids {
		_copy := instance.GongCopy().(*SpiralRhombusGrid)
		stage.SpiralRhombusGrids_reference[instance] = _copy
		stage.SpiralRhombusGrids_instance[_copy] = instance
		stage.SpiralRhombusGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.VerticalAxiss_reference = make(map[*VerticalAxis]*VerticalAxis)
	stage.VerticalAxiss_referenceOrder = make(map[*VerticalAxis]uint) // diff Unstage needs the reference order
	stage.VerticalAxiss_instance = make(map[*VerticalAxis]*VerticalAxis)
	for instance := range stage.VerticalAxiss {
		_copy := instance.GongCopy().(*VerticalAxis)
		stage.VerticalAxiss_reference[instance] = _copy
		stage.VerticalAxiss_instance[_copy] = instance
		stage.VerticalAxiss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Axiss {
		reference := stage.Axiss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AxisGrids {
		reference := stage.AxisGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Beziers {
		reference := stage.Beziers_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BezierGrids {
		reference := stage.BezierGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BezierGridStacks {
		reference := stage.BezierGridStacks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Chapters {
		reference := stage.Chapters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Circles {
		reference := stage.Circles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CircleGrids {
		reference := stage.CircleGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Contents {
		reference := stage.Contents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ExportToMusicxmls {
		reference := stage.ExportToMusicxmls_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FrontCurves {
		reference := stage.FrontCurves_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FrontCurveStacks {
		reference := stage.FrontCurveStacks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.HorizontalAxiss {
		reference := stage.HorizontalAxiss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Keys {
		reference := stage.Keys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Parameters {
		reference := stage.Parameters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rhombuss {
		reference := stage.Rhombuss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RhombusGrids {
		reference := stage.RhombusGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShapeCategorys {
		reference := stage.ShapeCategorys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralBeziers {
		reference := stage.SpiralBeziers_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralBezierGrids {
		reference := stage.SpiralBezierGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralCircles {
		reference := stage.SpiralCircles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralCircleGrids {
		reference := stage.SpiralCircleGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralLines {
		reference := stage.SpiralLines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralLineGrids {
		reference := stage.SpiralLineGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralOrigins {
		reference := stage.SpiralOrigins_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralRhombuss {
		reference := stage.SpiralRhombuss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SpiralRhombusGrids {
		reference := stage.SpiralRhombusGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.VerticalAxiss {
		reference := stage.VerticalAxiss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (axis *Axis) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Axis_stagedOrder[axis]; ok {
		return order
	}
	if order, ok := stage.Axiss_referenceOrder[axis]; ok {
		return order
	} else {
		log.Printf("instance %p of type Axis was not staged and does not have a reference order", axis)
		return 0
	}
}

func (axisgrid *AxisGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AxisGrid_stagedOrder[axisgrid]; ok {
		return order
	}
	if order, ok := stage.AxisGrids_referenceOrder[axisgrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type AxisGrid was not staged and does not have a reference order", axisgrid)
		return 0
	}
}

func (bezier *Bezier) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bezier_stagedOrder[bezier]; ok {
		return order
	}
	if order, ok := stage.Beziers_referenceOrder[bezier]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bezier was not staged and does not have a reference order", bezier)
		return 0
	}
}

func (beziergrid *BezierGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BezierGrid_stagedOrder[beziergrid]; ok {
		return order
	}
	if order, ok := stage.BezierGrids_referenceOrder[beziergrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type BezierGrid was not staged and does not have a reference order", beziergrid)
		return 0
	}
}

func (beziergridstack *BezierGridStack) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BezierGridStack_stagedOrder[beziergridstack]; ok {
		return order
	}
	if order, ok := stage.BezierGridStacks_referenceOrder[beziergridstack]; ok {
		return order
	} else {
		log.Printf("instance %p of type BezierGridStack was not staged and does not have a reference order", beziergridstack)
		return 0
	}
}

func (chapter *Chapter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Chapter_stagedOrder[chapter]; ok {
		return order
	}
	if order, ok := stage.Chapters_referenceOrder[chapter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Chapter was not staged and does not have a reference order", chapter)
		return 0
	}
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Circle_stagedOrder[circle]; ok {
		return order
	}
	if order, ok := stage.Circles_referenceOrder[circle]; ok {
		return order
	} else {
		log.Printf("instance %p of type Circle was not staged and does not have a reference order", circle)
		return 0
	}
}

func (circlegrid *CircleGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CircleGrid_stagedOrder[circlegrid]; ok {
		return order
	}
	if order, ok := stage.CircleGrids_referenceOrder[circlegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type CircleGrid was not staged and does not have a reference order", circlegrid)
		return 0
	}
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Content_stagedOrder[content]; ok {
		return order
	}
	if order, ok := stage.Contents_referenceOrder[content]; ok {
		return order
	} else {
		log.Printf("instance %p of type Content was not staged and does not have a reference order", content)
		return 0
	}
}

func (exporttomusicxml *ExportToMusicxml) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExportToMusicxml_stagedOrder[exporttomusicxml]; ok {
		return order
	}
	if order, ok := stage.ExportToMusicxmls_referenceOrder[exporttomusicxml]; ok {
		return order
	} else {
		log.Printf("instance %p of type ExportToMusicxml was not staged and does not have a reference order", exporttomusicxml)
		return 0
	}
}

func (frontcurve *FrontCurve) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FrontCurve_stagedOrder[frontcurve]; ok {
		return order
	}
	if order, ok := stage.FrontCurves_referenceOrder[frontcurve]; ok {
		return order
	} else {
		log.Printf("instance %p of type FrontCurve was not staged and does not have a reference order", frontcurve)
		return 0
	}
}

func (frontcurvestack *FrontCurveStack) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FrontCurveStack_stagedOrder[frontcurvestack]; ok {
		return order
	}
	if order, ok := stage.FrontCurveStacks_referenceOrder[frontcurvestack]; ok {
		return order
	} else {
		log.Printf("instance %p of type FrontCurveStack was not staged and does not have a reference order", frontcurvestack)
		return 0
	}
}

func (horizontalaxis *HorizontalAxis) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.HorizontalAxis_stagedOrder[horizontalaxis]; ok {
		return order
	}
	if order, ok := stage.HorizontalAxiss_referenceOrder[horizontalaxis]; ok {
		return order
	} else {
		log.Printf("instance %p of type HorizontalAxis was not staged and does not have a reference order", horizontalaxis)
		return 0
	}
}

func (key *Key) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Key_stagedOrder[key]; ok {
		return order
	}
	if order, ok := stage.Keys_referenceOrder[key]; ok {
		return order
	} else {
		log.Printf("instance %p of type Key was not staged and does not have a reference order", key)
		return 0
	}
}

func (parameter *Parameter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Parameter_stagedOrder[parameter]; ok {
		return order
	}
	if order, ok := stage.Parameters_referenceOrder[parameter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Parameter was not staged and does not have a reference order", parameter)
		return 0
	}
}

func (rhombus *Rhombus) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Rhombus_stagedOrder[rhombus]; ok {
		return order
	}
	if order, ok := stage.Rhombuss_referenceOrder[rhombus]; ok {
		return order
	} else {
		log.Printf("instance %p of type Rhombus was not staged and does not have a reference order", rhombus)
		return 0
	}
}

func (rhombusgrid *RhombusGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RhombusGrid_stagedOrder[rhombusgrid]; ok {
		return order
	}
	if order, ok := stage.RhombusGrids_referenceOrder[rhombusgrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type RhombusGrid was not staged and does not have a reference order", rhombusgrid)
		return 0
	}
}

func (shapecategory *ShapeCategory) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShapeCategory_stagedOrder[shapecategory]; ok {
		return order
	}
	if order, ok := stage.ShapeCategorys_referenceOrder[shapecategory]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShapeCategory was not staged and does not have a reference order", shapecategory)
		return 0
	}
}

func (spiralbezier *SpiralBezier) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralBezier_stagedOrder[spiralbezier]; ok {
		return order
	}
	if order, ok := stage.SpiralBeziers_referenceOrder[spiralbezier]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralBezier was not staged and does not have a reference order", spiralbezier)
		return 0
	}
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralBezierGrid_stagedOrder[spiralbeziergrid]; ok {
		return order
	}
	if order, ok := stage.SpiralBezierGrids_referenceOrder[spiralbeziergrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralBezierGrid was not staged and does not have a reference order", spiralbeziergrid)
		return 0
	}
}

func (spiralcircle *SpiralCircle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralCircle_stagedOrder[spiralcircle]; ok {
		return order
	}
	if order, ok := stage.SpiralCircles_referenceOrder[spiralcircle]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralCircle was not staged and does not have a reference order", spiralcircle)
		return 0
	}
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralCircleGrid_stagedOrder[spiralcirclegrid]; ok {
		return order
	}
	if order, ok := stage.SpiralCircleGrids_referenceOrder[spiralcirclegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralCircleGrid was not staged and does not have a reference order", spiralcirclegrid)
		return 0
	}
}

func (spiralline *SpiralLine) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralLine_stagedOrder[spiralline]; ok {
		return order
	}
	if order, ok := stage.SpiralLines_referenceOrder[spiralline]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralLine was not staged and does not have a reference order", spiralline)
		return 0
	}
}

func (spirallinegrid *SpiralLineGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralLineGrid_stagedOrder[spirallinegrid]; ok {
		return order
	}
	if order, ok := stage.SpiralLineGrids_referenceOrder[spirallinegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralLineGrid was not staged and does not have a reference order", spirallinegrid)
		return 0
	}
}

func (spiralorigin *SpiralOrigin) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralOrigin_stagedOrder[spiralorigin]; ok {
		return order
	}
	if order, ok := stage.SpiralOrigins_referenceOrder[spiralorigin]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralOrigin was not staged and does not have a reference order", spiralorigin)
		return 0
	}
}

func (spiralrhombus *SpiralRhombus) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralRhombus_stagedOrder[spiralrhombus]; ok {
		return order
	}
	if order, ok := stage.SpiralRhombuss_referenceOrder[spiralrhombus]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralRhombus was not staged and does not have a reference order", spiralrhombus)
		return 0
	}
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgrid]; ok {
		return order
	}
	if order, ok := stage.SpiralRhombusGrids_referenceOrder[spiralrhombusgrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type SpiralRhombusGrid was not staged and does not have a reference order", spiralrhombusgrid)
		return 0
	}
}

func (verticalaxis *VerticalAxis) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.VerticalAxis_stagedOrder[verticalaxis]; ok {
		return order
	}
	if order, ok := stage.VerticalAxiss_referenceOrder[verticalaxis]; ok {
		return order
	} else {
		log.Printf("instance %p of type VerticalAxis was not staged and does not have a reference order", verticalaxis)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (axis *Axis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axis.GongGetGongstructName(), axis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axis *Axis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axis.GongGetGongstructName(), axis.GongGetOrder(stage))
}

func (axisgrid *AxisGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axisgrid.GongGetGongstructName(), axisgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axisgrid *AxisGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axisgrid.GongGetGongstructName(), axisgrid.GongGetOrder(stage))
}

func (bezier *Bezier) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bezier.GongGetGongstructName(), bezier.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bezier *Bezier) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bezier.GongGetGongstructName(), bezier.GongGetOrder(stage))
}

func (beziergrid *BezierGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergrid.GongGetGongstructName(), beziergrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beziergrid *BezierGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergrid.GongGetGongstructName(), beziergrid.GongGetOrder(stage))
}

func (beziergridstack *BezierGridStack) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergridstack.GongGetGongstructName(), beziergridstack.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beziergridstack *BezierGridStack) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergridstack.GongGetGongstructName(), beziergridstack.GongGetOrder(stage))
}

func (chapter *Chapter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chapter *Chapter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

func (circle *Circle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circle *Circle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

func (circlegrid *CircleGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegrid.GongGetGongstructName(), circlegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlegrid *CircleGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegrid.GongGetGongstructName(), circlegrid.GongGetOrder(stage))
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

func (exporttomusicxml *ExportToMusicxml) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", exporttomusicxml.GongGetGongstructName(), exporttomusicxml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (exporttomusicxml *ExportToMusicxml) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", exporttomusicxml.GongGetGongstructName(), exporttomusicxml.GongGetOrder(stage))
}

func (frontcurve *FrontCurve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurve.GongGetGongstructName(), frontcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frontcurve *FrontCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurve.GongGetGongstructName(), frontcurve.GongGetOrder(stage))
}

func (frontcurvestack *FrontCurveStack) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurvestack.GongGetGongstructName(), frontcurvestack.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frontcurvestack *FrontCurveStack) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurvestack.GongGetGongstructName(), frontcurvestack.GongGetOrder(stage))
}

func (horizontalaxis *HorizontalAxis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontalaxis.GongGetGongstructName(), horizontalaxis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (horizontalaxis *HorizontalAxis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontalaxis.GongGetGongstructName(), horizontalaxis.GongGetOrder(stage))
}

func (key *Key) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key *Key) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetOrder(stage))
}

func (parameter *Parameter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parameter *Parameter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetOrder(stage))
}

func (rhombus *Rhombus) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombus.GongGetGongstructName(), rhombus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombus *Rhombus) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombus.GongGetGongstructName(), rhombus.GongGetOrder(stage))
}

func (rhombusgrid *RhombusGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusgrid.GongGetGongstructName(), rhombusgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusgrid *RhombusGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusgrid.GongGetGongstructName(), rhombusgrid.GongGetOrder(stage))
}

func (shapecategory *ShapeCategory) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shapecategory.GongGetGongstructName(), shapecategory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shapecategory *ShapeCategory) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shapecategory.GongGetGongstructName(), shapecategory.GongGetOrder(stage))
}

func (spiralbezier *SpiralBezier) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbezier.GongGetGongstructName(), spiralbezier.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralbezier *SpiralBezier) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbezier.GongGetGongstructName(), spiralbezier.GongGetOrder(stage))
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbeziergrid.GongGetGongstructName(), spiralbeziergrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralbeziergrid *SpiralBezierGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbeziergrid.GongGetGongstructName(), spiralbeziergrid.GongGetOrder(stage))
}

func (spiralcircle *SpiralCircle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcircle.GongGetGongstructName(), spiralcircle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralcircle *SpiralCircle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcircle.GongGetGongstructName(), spiralcircle.GongGetOrder(stage))
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcirclegrid.GongGetGongstructName(), spiralcirclegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralcirclegrid *SpiralCircleGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcirclegrid.GongGetGongstructName(), spiralcirclegrid.GongGetOrder(stage))
}

func (spiralline *SpiralLine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralline.GongGetGongstructName(), spiralline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralline *SpiralLine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralline.GongGetGongstructName(), spiralline.GongGetOrder(stage))
}

func (spirallinegrid *SpiralLineGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spirallinegrid.GongGetGongstructName(), spirallinegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spirallinegrid *SpiralLineGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spirallinegrid.GongGetGongstructName(), spirallinegrid.GongGetOrder(stage))
}

func (spiralorigin *SpiralOrigin) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralorigin.GongGetGongstructName(), spiralorigin.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralorigin *SpiralOrigin) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralorigin.GongGetGongstructName(), spiralorigin.GongGetOrder(stage))
}

func (spiralrhombus *SpiralRhombus) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombus.GongGetGongstructName(), spiralrhombus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralrhombus *SpiralRhombus) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombus.GongGetGongstructName(), spiralrhombus.GongGetOrder(stage))
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombusgrid.GongGetGongstructName(), spiralrhombusgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralrhombusgrid *SpiralRhombusGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombusgrid.GongGetGongstructName(), spiralrhombusgrid.GongGetOrder(stage))
}

func (verticalaxis *VerticalAxis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticalaxis.GongGetGongstructName(), verticalaxis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalaxis *VerticalAxis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticalaxis.GongGetGongstructName(), verticalaxis.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (axis *Axis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Axis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.Name))
	return
}

func (axisgrid *AxisGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxisGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axisgrid.Name))
	return
}

func (bezier *Bezier) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bezier.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bezier")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.Name))
	return
}

func (beziergrid *BezierGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(beziergrid.Name))
	return
}

func (beziergridstack *BezierGridStack) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGridStack")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(beziergridstack.Name))
	return
}

func (chapter *Chapter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.Name))
	return
}

func (circle *Circle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.Name))
	return
}

func (circlegrid *CircleGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CircleGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circlegrid.Name))
	return
}

func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.Name))
	return
}

func (exporttomusicxml *ExportToMusicxml) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExportToMusicxml")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(exporttomusicxml.Name))
	return
}

func (frontcurve *FrontCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurve.Name))
	return
}

func (frontcurvestack *FrontCurveStack) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurveStack")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.Name))
	return
}

func (horizontalaxis *HorizontalAxis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "HorizontalAxis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.Name))
	return
}

func (key *Key) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Name))
	return
}

func (parameter *Parameter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Parameter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.Name))
	return
}

func (rhombus *Rhombus) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rhombus")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.Name))
	return
}

func (rhombusgrid *RhombusGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombusgrid.Name))
	return
}

func (shapecategory *ShapeCategory) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shapecategory.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShapeCategory")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shapecategory.Name))
	return
}

func (spiralbezier *SpiralBezier) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezier")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.Name))
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezierGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbeziergrid.Name))
	return
}

func (spiralcircle *SpiralCircle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDeclsWithoutNameInit
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Name))
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircleGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcirclegrid.Name))
	return
}

func (spiralline *SpiralLine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.Name))
	return
}

func (spirallinegrid *SpiralLineGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLineGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spirallinegrid.Name))
	return
}

func (spiralorigin *SpiralOrigin) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralOrigin")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.Name))
	return
}

func (spiralrhombus *SpiralRhombus) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombus")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.Name))
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombusGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombusgrid.Name))
	return
}

func (verticalaxis *VerticalAxis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VerticalAxis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.Name))
	return
}

// insertion point for unstaging
func (axis *Axis) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axis.GongGetReferenceIdentifier(stage))
	return
}

func (axisgrid *AxisGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axisgrid.GongGetReferenceIdentifier(stage))
	return
}

func (bezier *Bezier) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bezier.GongGetReferenceIdentifier(stage))
	return
}

func (beziergrid *BezierGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergrid.GongGetReferenceIdentifier(stage))
	return
}

func (beziergridstack *BezierGridStack) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergridstack.GongGetReferenceIdentifier(stage))
	return
}

func (chapter *Chapter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetReferenceIdentifier(stage))
	return
}

func (circle *Circle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetReferenceIdentifier(stage))
	return
}

func (circlegrid *CircleGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegrid.GongGetReferenceIdentifier(stage))
	return
}

func (content *Content) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetReferenceIdentifier(stage))
	return
}

func (exporttomusicxml *ExportToMusicxml) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", exporttomusicxml.GongGetReferenceIdentifier(stage))
	return
}

func (frontcurve *FrontCurve) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurve.GongGetReferenceIdentifier(stage))
	return
}

func (frontcurvestack *FrontCurveStack) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurvestack.GongGetReferenceIdentifier(stage))
	return
}

func (horizontalaxis *HorizontalAxis) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", horizontalaxis.GongGetReferenceIdentifier(stage))
	return
}

func (key *Key) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key.GongGetReferenceIdentifier(stage))
	return
}

func (parameter *Parameter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetReferenceIdentifier(stage))
	return
}

func (rhombus *Rhombus) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombus.GongGetReferenceIdentifier(stage))
	return
}

func (rhombusgrid *RhombusGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusgrid.GongGetReferenceIdentifier(stage))
	return
}

func (shapecategory *ShapeCategory) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shapecategory.GongGetReferenceIdentifier(stage))
	return
}

func (spiralbezier *SpiralBezier) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbezier.GongGetReferenceIdentifier(stage))
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbeziergrid.GongGetReferenceIdentifier(stage))
	return
}

func (spiralcircle *SpiralCircle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcircle.GongGetReferenceIdentifier(stage))
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcirclegrid.GongGetReferenceIdentifier(stage))
	return
}

func (spiralline *SpiralLine) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralline.GongGetReferenceIdentifier(stage))
	return
}

func (spirallinegrid *SpiralLineGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spirallinegrid.GongGetReferenceIdentifier(stage))
	return
}

func (spiralorigin *SpiralOrigin) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralorigin.GongGetReferenceIdentifier(stage))
	return
}

func (spiralrhombus *SpiralRhombus) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombus.GongGetReferenceIdentifier(stage))
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombusgrid.GongGetReferenceIdentifier(stage))
	return
}

func (verticalaxis *VerticalAxis) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticalaxis.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
// end of template
