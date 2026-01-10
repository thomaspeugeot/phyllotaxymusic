// generated code - do not edit
package models

import (
	"fmt"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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
	newInstance := *axis
	return &newInstance
}

func (axisgrid *AxisGrid) GongCopy() GongstructIF {
	newInstance := *axisgrid
	return &newInstance
}

func (bezier *Bezier) GongCopy() GongstructIF {
	newInstance := *bezier
	return &newInstance
}

func (beziergrid *BezierGrid) GongCopy() GongstructIF {
	newInstance := *beziergrid
	return &newInstance
}

func (beziergridstack *BezierGridStack) GongCopy() GongstructIF {
	newInstance := *beziergridstack
	return &newInstance
}

func (chapter *Chapter) GongCopy() GongstructIF {
	newInstance := *chapter
	return &newInstance
}

func (circle *Circle) GongCopy() GongstructIF {
	newInstance := *circle
	return &newInstance
}

func (circlegrid *CircleGrid) GongCopy() GongstructIF {
	newInstance := *circlegrid
	return &newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := *content
	return &newInstance
}

func (exporttomusicxml *ExportToMusicxml) GongCopy() GongstructIF {
	newInstance := *exporttomusicxml
	return &newInstance
}

func (frontcurve *FrontCurve) GongCopy() GongstructIF {
	newInstance := *frontcurve
	return &newInstance
}

func (frontcurvestack *FrontCurveStack) GongCopy() GongstructIF {
	newInstance := *frontcurvestack
	return &newInstance
}

func (horizontalaxis *HorizontalAxis) GongCopy() GongstructIF {
	newInstance := *horizontalaxis
	return &newInstance
}

func (key *Key) GongCopy() GongstructIF {
	newInstance := *key
	return &newInstance
}

func (parameter *Parameter) GongCopy() GongstructIF {
	newInstance := *parameter
	return &newInstance
}

func (rhombus *Rhombus) GongCopy() GongstructIF {
	newInstance := *rhombus
	return &newInstance
}

func (rhombusgrid *RhombusGrid) GongCopy() GongstructIF {
	newInstance := *rhombusgrid
	return &newInstance
}

func (shapecategory *ShapeCategory) GongCopy() GongstructIF {
	newInstance := *shapecategory
	return &newInstance
}

func (spiralbezier *SpiralBezier) GongCopy() GongstructIF {
	newInstance := *spiralbezier
	return &newInstance
}

func (spiralbeziergrid *SpiralBezierGrid) GongCopy() GongstructIF {
	newInstance := *spiralbeziergrid
	return &newInstance
}

func (spiralcircle *SpiralCircle) GongCopy() GongstructIF {
	newInstance := *spiralcircle
	return &newInstance
}

func (spiralcirclegrid *SpiralCircleGrid) GongCopy() GongstructIF {
	newInstance := *spiralcirclegrid
	return &newInstance
}

func (spiralline *SpiralLine) GongCopy() GongstructIF {
	newInstance := *spiralline
	return &newInstance
}

func (spirallinegrid *SpiralLineGrid) GongCopy() GongstructIF {
	newInstance := *spirallinegrid
	return &newInstance
}

func (spiralorigin *SpiralOrigin) GongCopy() GongstructIF {
	newInstance := *spiralorigin
	return &newInstance
}

func (spiralrhombus *SpiralRhombus) GongCopy() GongstructIF {
	newInstance := *spiralrhombus
	return &newInstance
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongCopy() GongstructIF {
	newInstance := *spiralrhombusgrid
	return &newInstance
}

func (verticalaxis *VerticalAxis) GongCopy() GongstructIF {
	newInstance := *verticalaxis
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += axis.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := axis.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := axis.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", axis.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for axis := range stage.Axiss_reference {
		if _, ok := stage.Axiss[axis]; !ok {
			axiss_deletedInstances = append(axiss_deletedInstances, axis)
			deletedInstancesStmt += axis.GongMarshallUnstaging(stage)
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
			newInstancesStmt += axisgrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := axisgrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := axisgrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", axisgrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for axisgrid := range stage.AxisGrids_reference {
		if _, ok := stage.AxisGrids[axisgrid]; !ok {
			axisgrids_deletedInstances = append(axisgrids_deletedInstances, axisgrid)
			deletedInstancesStmt += axisgrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += bezier.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := bezier.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := bezier.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", bezier.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for bezier := range stage.Beziers_reference {
		if _, ok := stage.Beziers[bezier]; !ok {
			beziers_deletedInstances = append(beziers_deletedInstances, bezier)
			deletedInstancesStmt += bezier.GongMarshallUnstaging(stage)
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
			newInstancesStmt += beziergrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := beziergrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := beziergrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", beziergrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for beziergrid := range stage.BezierGrids_reference {
		if _, ok := stage.BezierGrids[beziergrid]; !ok {
			beziergrids_deletedInstances = append(beziergrids_deletedInstances, beziergrid)
			deletedInstancesStmt += beziergrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += beziergridstack.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := beziergridstack.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := beziergridstack.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", beziergridstack.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for beziergridstack := range stage.BezierGridStacks_reference {
		if _, ok := stage.BezierGridStacks[beziergridstack]; !ok {
			beziergridstacks_deletedInstances = append(beziergridstacks_deletedInstances, beziergridstack)
			deletedInstancesStmt += beziergridstack.GongMarshallUnstaging(stage)
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
			newInstancesStmt += chapter.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := chapter.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := chapter.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", chapter.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for chapter := range stage.Chapters_reference {
		if _, ok := stage.Chapters[chapter]; !ok {
			chapters_deletedInstances = append(chapters_deletedInstances, chapter)
			deletedInstancesStmt += chapter.GongMarshallUnstaging(stage)
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
			newInstancesStmt += circle.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := circle.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := circle.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", circle.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for circle := range stage.Circles_reference {
		if _, ok := stage.Circles[circle]; !ok {
			circles_deletedInstances = append(circles_deletedInstances, circle)
			deletedInstancesStmt += circle.GongMarshallUnstaging(stage)
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
			newInstancesStmt += circlegrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := circlegrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := circlegrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", circlegrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for circlegrid := range stage.CircleGrids_reference {
		if _, ok := stage.CircleGrids[circlegrid]; !ok {
			circlegrids_deletedInstances = append(circlegrids_deletedInstances, circlegrid)
			deletedInstancesStmt += circlegrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += content.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := content.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", content.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for content := range stage.Contents_reference {
		if _, ok := stage.Contents[content]; !ok {
			contents_deletedInstances = append(contents_deletedInstances, content)
			deletedInstancesStmt += content.GongMarshallUnstaging(stage)
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
			newInstancesStmt += exporttomusicxml.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := exporttomusicxml.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := exporttomusicxml.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", exporttomusicxml.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for exporttomusicxml := range stage.ExportToMusicxmls_reference {
		if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; !ok {
			exporttomusicxmls_deletedInstances = append(exporttomusicxmls_deletedInstances, exporttomusicxml)
			deletedInstancesStmt += exporttomusicxml.GongMarshallUnstaging(stage)
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
			newInstancesStmt += frontcurve.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := frontcurve.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := frontcurve.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", frontcurve.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for frontcurve := range stage.FrontCurves_reference {
		if _, ok := stage.FrontCurves[frontcurve]; !ok {
			frontcurves_deletedInstances = append(frontcurves_deletedInstances, frontcurve)
			deletedInstancesStmt += frontcurve.GongMarshallUnstaging(stage)
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
			newInstancesStmt += frontcurvestack.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := frontcurvestack.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := frontcurvestack.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", frontcurvestack.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for frontcurvestack := range stage.FrontCurveStacks_reference {
		if _, ok := stage.FrontCurveStacks[frontcurvestack]; !ok {
			frontcurvestacks_deletedInstances = append(frontcurvestacks_deletedInstances, frontcurvestack)
			deletedInstancesStmt += frontcurvestack.GongMarshallUnstaging(stage)
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
			newInstancesStmt += horizontalaxis.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := horizontalaxis.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := horizontalaxis.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", horizontalaxis.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for horizontalaxis := range stage.HorizontalAxiss_reference {
		if _, ok := stage.HorizontalAxiss[horizontalaxis]; !ok {
			horizontalaxiss_deletedInstances = append(horizontalaxiss_deletedInstances, horizontalaxis)
			deletedInstancesStmt += horizontalaxis.GongMarshallUnstaging(stage)
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
			newInstancesStmt += key.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := key.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := key.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", key.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for key := range stage.Keys_reference {
		if _, ok := stage.Keys[key]; !ok {
			keys_deletedInstances = append(keys_deletedInstances, key)
			deletedInstancesStmt += key.GongMarshallUnstaging(stage)
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
			newInstancesStmt += parameter.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := parameter.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := parameter.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", parameter.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for parameter := range stage.Parameters_reference {
		if _, ok := stage.Parameters[parameter]; !ok {
			parameters_deletedInstances = append(parameters_deletedInstances, parameter)
			deletedInstancesStmt += parameter.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rhombus.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rhombus.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rhombus.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rhombus.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rhombus := range stage.Rhombuss_reference {
		if _, ok := stage.Rhombuss[rhombus]; !ok {
			rhombuss_deletedInstances = append(rhombuss_deletedInstances, rhombus)
			deletedInstancesStmt += rhombus.GongMarshallUnstaging(stage)
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
			newInstancesStmt += rhombusgrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := rhombusgrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := rhombusgrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", rhombusgrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rhombusgrid := range stage.RhombusGrids_reference {
		if _, ok := stage.RhombusGrids[rhombusgrid]; !ok {
			rhombusgrids_deletedInstances = append(rhombusgrids_deletedInstances, rhombusgrid)
			deletedInstancesStmt += rhombusgrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += shapecategory.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := shapecategory.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := shapecategory.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", shapecategory.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for shapecategory := range stage.ShapeCategorys_reference {
		if _, ok := stage.ShapeCategorys[shapecategory]; !ok {
			shapecategorys_deletedInstances = append(shapecategorys_deletedInstances, shapecategory)
			deletedInstancesStmt += shapecategory.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralbezier.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralbezier.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralbezier.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralbezier.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralbezier := range stage.SpiralBeziers_reference {
		if _, ok := stage.SpiralBeziers[spiralbezier]; !ok {
			spiralbeziers_deletedInstances = append(spiralbeziers_deletedInstances, spiralbezier)
			deletedInstancesStmt += spiralbezier.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralbeziergrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralbeziergrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralbeziergrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralbeziergrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralbeziergrid := range stage.SpiralBezierGrids_reference {
		if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; !ok {
			spiralbeziergrids_deletedInstances = append(spiralbeziergrids_deletedInstances, spiralbeziergrid)
			deletedInstancesStmt += spiralbeziergrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralcircle.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralcircle.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralcircle.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralcircle.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralcircle := range stage.SpiralCircles_reference {
		if _, ok := stage.SpiralCircles[spiralcircle]; !ok {
			spiralcircles_deletedInstances = append(spiralcircles_deletedInstances, spiralcircle)
			deletedInstancesStmt += spiralcircle.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralcirclegrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralcirclegrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralcirclegrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralcirclegrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralcirclegrid := range stage.SpiralCircleGrids_reference {
		if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; !ok {
			spiralcirclegrids_deletedInstances = append(spiralcirclegrids_deletedInstances, spiralcirclegrid)
			deletedInstancesStmt += spiralcirclegrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralline.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralline.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralline.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralline.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralline := range stage.SpiralLines_reference {
		if _, ok := stage.SpiralLines[spiralline]; !ok {
			spirallines_deletedInstances = append(spirallines_deletedInstances, spiralline)
			deletedInstancesStmt += spiralline.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spirallinegrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spirallinegrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spirallinegrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spirallinegrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spirallinegrid := range stage.SpiralLineGrids_reference {
		if _, ok := stage.SpiralLineGrids[spirallinegrid]; !ok {
			spirallinegrids_deletedInstances = append(spirallinegrids_deletedInstances, spirallinegrid)
			deletedInstancesStmt += spirallinegrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralorigin.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralorigin.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralorigin.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralorigin.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralorigin := range stage.SpiralOrigins_reference {
		if _, ok := stage.SpiralOrigins[spiralorigin]; !ok {
			spiralorigins_deletedInstances = append(spiralorigins_deletedInstances, spiralorigin)
			deletedInstancesStmt += spiralorigin.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralrhombus.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralrhombus.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralrhombus.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralrhombus.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralrhombus := range stage.SpiralRhombuss_reference {
		if _, ok := stage.SpiralRhombuss[spiralrhombus]; !ok {
			spiralrhombuss_deletedInstances = append(spiralrhombuss_deletedInstances, spiralrhombus)
			deletedInstancesStmt += spiralrhombus.GongMarshallUnstaging(stage)
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
			newInstancesStmt += spiralrhombusgrid.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := spiralrhombusgrid.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := spiralrhombusgrid.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", spiralrhombusgrid.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for spiralrhombusgrid := range stage.SpiralRhombusGrids_reference {
		if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; !ok {
			spiralrhombusgrids_deletedInstances = append(spiralrhombusgrids_deletedInstances, spiralrhombusgrid)
			deletedInstancesStmt += spiralrhombusgrid.GongMarshallUnstaging(stage)
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
			newInstancesStmt += verticalaxis.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := verticalaxis.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := verticalaxis.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", verticalaxis.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for verticalaxis := range stage.VerticalAxiss_reference {
		if _, ok := stage.VerticalAxiss[verticalaxis]; !ok {
			verticalaxiss_deletedInstances = append(verticalaxiss_deletedInstances, verticalaxis)
			deletedInstancesStmt += verticalaxis.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(verticalaxiss_newInstances)
	lenDeletedInstances += len(verticalaxiss_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += "\n\tstage.Commit()"
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Axiss_reference = make(map[*Axis]*Axis)
	stage.Axiss_referenceOrder = make(map[*Axis]uint) // diff Unstage needs the reference order
	for instance := range stage.Axiss {
		stage.Axiss_reference[instance] = instance.GongCopy().(*Axis)
		stage.Axiss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.AxisGrids_reference = make(map[*AxisGrid]*AxisGrid)
	stage.AxisGrids_referenceOrder = make(map[*AxisGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.AxisGrids {
		stage.AxisGrids_reference[instance] = instance.GongCopy().(*AxisGrid)
		stage.AxisGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Beziers_reference = make(map[*Bezier]*Bezier)
	stage.Beziers_referenceOrder = make(map[*Bezier]uint) // diff Unstage needs the reference order
	for instance := range stage.Beziers {
		stage.Beziers_reference[instance] = instance.GongCopy().(*Bezier)
		stage.Beziers_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.BezierGrids_reference = make(map[*BezierGrid]*BezierGrid)
	stage.BezierGrids_referenceOrder = make(map[*BezierGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.BezierGrids {
		stage.BezierGrids_reference[instance] = instance.GongCopy().(*BezierGrid)
		stage.BezierGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.BezierGridStacks_reference = make(map[*BezierGridStack]*BezierGridStack)
	stage.BezierGridStacks_referenceOrder = make(map[*BezierGridStack]uint) // diff Unstage needs the reference order
	for instance := range stage.BezierGridStacks {
		stage.BezierGridStacks_reference[instance] = instance.GongCopy().(*BezierGridStack)
		stage.BezierGridStacks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint) // diff Unstage needs the reference order
	for instance := range stage.Chapters {
		stage.Chapters_reference[instance] = instance.GongCopy().(*Chapter)
		stage.Chapters_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Circles_reference = make(map[*Circle]*Circle)
	stage.Circles_referenceOrder = make(map[*Circle]uint) // diff Unstage needs the reference order
	for instance := range stage.Circles {
		stage.Circles_reference[instance] = instance.GongCopy().(*Circle)
		stage.Circles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CircleGrids_reference = make(map[*CircleGrid]*CircleGrid)
	stage.CircleGrids_referenceOrder = make(map[*CircleGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.CircleGrids {
		stage.CircleGrids_reference[instance] = instance.GongCopy().(*CircleGrid)
		stage.CircleGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint) // diff Unstage needs the reference order
	for instance := range stage.Contents {
		stage.Contents_reference[instance] = instance.GongCopy().(*Content)
		stage.Contents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ExportToMusicxmls_reference = make(map[*ExportToMusicxml]*ExportToMusicxml)
	stage.ExportToMusicxmls_referenceOrder = make(map[*ExportToMusicxml]uint) // diff Unstage needs the reference order
	for instance := range stage.ExportToMusicxmls {
		stage.ExportToMusicxmls_reference[instance] = instance.GongCopy().(*ExportToMusicxml)
		stage.ExportToMusicxmls_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FrontCurves_reference = make(map[*FrontCurve]*FrontCurve)
	stage.FrontCurves_referenceOrder = make(map[*FrontCurve]uint) // diff Unstage needs the reference order
	for instance := range stage.FrontCurves {
		stage.FrontCurves_reference[instance] = instance.GongCopy().(*FrontCurve)
		stage.FrontCurves_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FrontCurveStacks_reference = make(map[*FrontCurveStack]*FrontCurveStack)
	stage.FrontCurveStacks_referenceOrder = make(map[*FrontCurveStack]uint) // diff Unstage needs the reference order
	for instance := range stage.FrontCurveStacks {
		stage.FrontCurveStacks_reference[instance] = instance.GongCopy().(*FrontCurveStack)
		stage.FrontCurveStacks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.HorizontalAxiss_reference = make(map[*HorizontalAxis]*HorizontalAxis)
	stage.HorizontalAxiss_referenceOrder = make(map[*HorizontalAxis]uint) // diff Unstage needs the reference order
	for instance := range stage.HorizontalAxiss {
		stage.HorizontalAxiss_reference[instance] = instance.GongCopy().(*HorizontalAxis)
		stage.HorizontalAxiss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Keys_reference = make(map[*Key]*Key)
	stage.Keys_referenceOrder = make(map[*Key]uint) // diff Unstage needs the reference order
	for instance := range stage.Keys {
		stage.Keys_reference[instance] = instance.GongCopy().(*Key)
		stage.Keys_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Parameters_reference = make(map[*Parameter]*Parameter)
	stage.Parameters_referenceOrder = make(map[*Parameter]uint) // diff Unstage needs the reference order
	for instance := range stage.Parameters {
		stage.Parameters_reference[instance] = instance.GongCopy().(*Parameter)
		stage.Parameters_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Rhombuss_reference = make(map[*Rhombus]*Rhombus)
	stage.Rhombuss_referenceOrder = make(map[*Rhombus]uint) // diff Unstage needs the reference order
	for instance := range stage.Rhombuss {
		stage.Rhombuss_reference[instance] = instance.GongCopy().(*Rhombus)
		stage.Rhombuss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RhombusGrids_reference = make(map[*RhombusGrid]*RhombusGrid)
	stage.RhombusGrids_referenceOrder = make(map[*RhombusGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.RhombusGrids {
		stage.RhombusGrids_reference[instance] = instance.GongCopy().(*RhombusGrid)
		stage.RhombusGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ShapeCategorys_reference = make(map[*ShapeCategory]*ShapeCategory)
	stage.ShapeCategorys_referenceOrder = make(map[*ShapeCategory]uint) // diff Unstage needs the reference order
	for instance := range stage.ShapeCategorys {
		stage.ShapeCategorys_reference[instance] = instance.GongCopy().(*ShapeCategory)
		stage.ShapeCategorys_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralBeziers_reference = make(map[*SpiralBezier]*SpiralBezier)
	stage.SpiralBeziers_referenceOrder = make(map[*SpiralBezier]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralBeziers {
		stage.SpiralBeziers_reference[instance] = instance.GongCopy().(*SpiralBezier)
		stage.SpiralBeziers_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralBezierGrids_reference = make(map[*SpiralBezierGrid]*SpiralBezierGrid)
	stage.SpiralBezierGrids_referenceOrder = make(map[*SpiralBezierGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralBezierGrids {
		stage.SpiralBezierGrids_reference[instance] = instance.GongCopy().(*SpiralBezierGrid)
		stage.SpiralBezierGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralCircles_reference = make(map[*SpiralCircle]*SpiralCircle)
	stage.SpiralCircles_referenceOrder = make(map[*SpiralCircle]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralCircles {
		stage.SpiralCircles_reference[instance] = instance.GongCopy().(*SpiralCircle)
		stage.SpiralCircles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralCircleGrids_reference = make(map[*SpiralCircleGrid]*SpiralCircleGrid)
	stage.SpiralCircleGrids_referenceOrder = make(map[*SpiralCircleGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralCircleGrids {
		stage.SpiralCircleGrids_reference[instance] = instance.GongCopy().(*SpiralCircleGrid)
		stage.SpiralCircleGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralLines_reference = make(map[*SpiralLine]*SpiralLine)
	stage.SpiralLines_referenceOrder = make(map[*SpiralLine]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralLines {
		stage.SpiralLines_reference[instance] = instance.GongCopy().(*SpiralLine)
		stage.SpiralLines_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralLineGrids_reference = make(map[*SpiralLineGrid]*SpiralLineGrid)
	stage.SpiralLineGrids_referenceOrder = make(map[*SpiralLineGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralLineGrids {
		stage.SpiralLineGrids_reference[instance] = instance.GongCopy().(*SpiralLineGrid)
		stage.SpiralLineGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralOrigins_reference = make(map[*SpiralOrigin]*SpiralOrigin)
	stage.SpiralOrigins_referenceOrder = make(map[*SpiralOrigin]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralOrigins {
		stage.SpiralOrigins_reference[instance] = instance.GongCopy().(*SpiralOrigin)
		stage.SpiralOrigins_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralRhombuss_reference = make(map[*SpiralRhombus]*SpiralRhombus)
	stage.SpiralRhombuss_referenceOrder = make(map[*SpiralRhombus]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralRhombuss {
		stage.SpiralRhombuss_reference[instance] = instance.GongCopy().(*SpiralRhombus)
		stage.SpiralRhombuss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SpiralRhombusGrids_reference = make(map[*SpiralRhombusGrid]*SpiralRhombusGrid)
	stage.SpiralRhombusGrids_referenceOrder = make(map[*SpiralRhombusGrid]uint) // diff Unstage needs the reference order
	for instance := range stage.SpiralRhombusGrids {
		stage.SpiralRhombusGrids_reference[instance] = instance.GongCopy().(*SpiralRhombusGrid)
		stage.SpiralRhombusGrids_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.VerticalAxiss_reference = make(map[*VerticalAxis]*VerticalAxis)
	stage.VerticalAxiss_referenceOrder = make(map[*VerticalAxis]uint) // diff Unstage needs the reference order
	for instance := range stage.VerticalAxiss {
		stage.VerticalAxiss_reference[instance] = instance.GongCopy().(*VerticalAxis)
		stage.VerticalAxiss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (axis *Axis) GongGetOrder(stage *Stage) uint {
	return stage.AxisMap_Staged_Order[axis]
}

func (axis *Axis) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Axiss_referenceOrder[axis]
}

func (axisgrid *AxisGrid) GongGetOrder(stage *Stage) uint {
	return stage.AxisGridMap_Staged_Order[axisgrid]
}

func (axisgrid *AxisGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AxisGrids_referenceOrder[axisgrid]
}

func (bezier *Bezier) GongGetOrder(stage *Stage) uint {
	return stage.BezierMap_Staged_Order[bezier]
}

func (bezier *Bezier) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Beziers_referenceOrder[bezier]
}

func (beziergrid *BezierGrid) GongGetOrder(stage *Stage) uint {
	return stage.BezierGridMap_Staged_Order[beziergrid]
}

func (beziergrid *BezierGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.BezierGrids_referenceOrder[beziergrid]
}

func (beziergridstack *BezierGridStack) GongGetOrder(stage *Stage) uint {
	return stage.BezierGridStackMap_Staged_Order[beziergridstack]
}

func (beziergridstack *BezierGridStack) GongGetReferenceOrder(stage *Stage) uint {
	return stage.BezierGridStacks_referenceOrder[beziergridstack]
}

func (chapter *Chapter) GongGetOrder(stage *Stage) uint {
	return stage.ChapterMap_Staged_Order[chapter]
}

func (chapter *Chapter) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Chapters_referenceOrder[chapter]
}

func (circle *Circle) GongGetOrder(stage *Stage) uint {
	return stage.CircleMap_Staged_Order[circle]
}

func (circle *Circle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Circles_referenceOrder[circle]
}

func (circlegrid *CircleGrid) GongGetOrder(stage *Stage) uint {
	return stage.CircleGridMap_Staged_Order[circlegrid]
}

func (circlegrid *CircleGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CircleGrids_referenceOrder[circlegrid]
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	return stage.ContentMap_Staged_Order[content]
}

func (content *Content) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Contents_referenceOrder[content]
}

func (exporttomusicxml *ExportToMusicxml) GongGetOrder(stage *Stage) uint {
	return stage.ExportToMusicxmlMap_Staged_Order[exporttomusicxml]
}

func (exporttomusicxml *ExportToMusicxml) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ExportToMusicxmls_referenceOrder[exporttomusicxml]
}

func (frontcurve *FrontCurve) GongGetOrder(stage *Stage) uint {
	return stage.FrontCurveMap_Staged_Order[frontcurve]
}

func (frontcurve *FrontCurve) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FrontCurves_referenceOrder[frontcurve]
}

func (frontcurvestack *FrontCurveStack) GongGetOrder(stage *Stage) uint {
	return stage.FrontCurveStackMap_Staged_Order[frontcurvestack]
}

func (frontcurvestack *FrontCurveStack) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FrontCurveStacks_referenceOrder[frontcurvestack]
}

func (horizontalaxis *HorizontalAxis) GongGetOrder(stage *Stage) uint {
	return stage.HorizontalAxisMap_Staged_Order[horizontalaxis]
}

func (horizontalaxis *HorizontalAxis) GongGetReferenceOrder(stage *Stage) uint {
	return stage.HorizontalAxiss_referenceOrder[horizontalaxis]
}

func (key *Key) GongGetOrder(stage *Stage) uint {
	return stage.KeyMap_Staged_Order[key]
}

func (key *Key) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Keys_referenceOrder[key]
}

func (parameter *Parameter) GongGetOrder(stage *Stage) uint {
	return stage.ParameterMap_Staged_Order[parameter]
}

func (parameter *Parameter) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Parameters_referenceOrder[parameter]
}

func (rhombus *Rhombus) GongGetOrder(stage *Stage) uint {
	return stage.RhombusMap_Staged_Order[rhombus]
}

func (rhombus *Rhombus) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Rhombuss_referenceOrder[rhombus]
}

func (rhombusgrid *RhombusGrid) GongGetOrder(stage *Stage) uint {
	return stage.RhombusGridMap_Staged_Order[rhombusgrid]
}

func (rhombusgrid *RhombusGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RhombusGrids_referenceOrder[rhombusgrid]
}

func (shapecategory *ShapeCategory) GongGetOrder(stage *Stage) uint {
	return stage.ShapeCategoryMap_Staged_Order[shapecategory]
}

func (shapecategory *ShapeCategory) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ShapeCategorys_referenceOrder[shapecategory]
}

func (spiralbezier *SpiralBezier) GongGetOrder(stage *Stage) uint {
	return stage.SpiralBezierMap_Staged_Order[spiralbezier]
}

func (spiralbezier *SpiralBezier) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralBeziers_referenceOrder[spiralbezier]
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetOrder(stage *Stage) uint {
	return stage.SpiralBezierGridMap_Staged_Order[spiralbeziergrid]
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralBezierGrids_referenceOrder[spiralbeziergrid]
}

func (spiralcircle *SpiralCircle) GongGetOrder(stage *Stage) uint {
	return stage.SpiralCircleMap_Staged_Order[spiralcircle]
}

func (spiralcircle *SpiralCircle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralCircles_referenceOrder[spiralcircle]
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetOrder(stage *Stage) uint {
	return stage.SpiralCircleGridMap_Staged_Order[spiralcirclegrid]
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralCircleGrids_referenceOrder[spiralcirclegrid]
}

func (spiralline *SpiralLine) GongGetOrder(stage *Stage) uint {
	return stage.SpiralLineMap_Staged_Order[spiralline]
}

func (spiralline *SpiralLine) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralLines_referenceOrder[spiralline]
}

func (spirallinegrid *SpiralLineGrid) GongGetOrder(stage *Stage) uint {
	return stage.SpiralLineGridMap_Staged_Order[spirallinegrid]
}

func (spirallinegrid *SpiralLineGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralLineGrids_referenceOrder[spirallinegrid]
}

func (spiralorigin *SpiralOrigin) GongGetOrder(stage *Stage) uint {
	return stage.SpiralOriginMap_Staged_Order[spiralorigin]
}

func (spiralorigin *SpiralOrigin) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralOrigins_referenceOrder[spiralorigin]
}

func (spiralrhombus *SpiralRhombus) GongGetOrder(stage *Stage) uint {
	return stage.SpiralRhombusMap_Staged_Order[spiralrhombus]
}

func (spiralrhombus *SpiralRhombus) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralRhombuss_referenceOrder[spiralrhombus]
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetOrder(stage *Stage) uint {
	return stage.SpiralRhombusGridMap_Staged_Order[spiralrhombusgrid]
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SpiralRhombusGrids_referenceOrder[spiralrhombusgrid]
}

func (verticalaxis *VerticalAxis) GongGetOrder(stage *Stage) uint {
	return stage.VerticalAxisMap_Staged_Order[verticalaxis]
}

func (verticalaxis *VerticalAxis) GongGetReferenceOrder(stage *Stage) uint {
	return stage.VerticalAxiss_referenceOrder[verticalaxis]
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
	return fmt.Sprintf("__%s__%08d_", axis.GongGetGongstructName(), axis.GongGetReferenceOrder(stage))
}

func (axisgrid *AxisGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axisgrid.GongGetGongstructName(), axisgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axisgrid *AxisGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axisgrid.GongGetGongstructName(), axisgrid.GongGetReferenceOrder(stage))
}

func (bezier *Bezier) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bezier.GongGetGongstructName(), bezier.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bezier *Bezier) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bezier.GongGetGongstructName(), bezier.GongGetReferenceOrder(stage))
}

func (beziergrid *BezierGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergrid.GongGetGongstructName(), beziergrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beziergrid *BezierGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergrid.GongGetGongstructName(), beziergrid.GongGetReferenceOrder(stage))
}

func (beziergridstack *BezierGridStack) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergridstack.GongGetGongstructName(), beziergridstack.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beziergridstack *BezierGridStack) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beziergridstack.GongGetGongstructName(), beziergridstack.GongGetReferenceOrder(stage))
}

func (chapter *Chapter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chapter *Chapter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetReferenceOrder(stage))
}

func (circle *Circle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circle *Circle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circle.GongGetGongstructName(), circle.GongGetReferenceOrder(stage))
}

func (circlegrid *CircleGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegrid.GongGetGongstructName(), circlegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlegrid *CircleGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegrid.GongGetGongstructName(), circlegrid.GongGetReferenceOrder(stage))
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetReferenceOrder(stage))
}

func (exporttomusicxml *ExportToMusicxml) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", exporttomusicxml.GongGetGongstructName(), exporttomusicxml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (exporttomusicxml *ExportToMusicxml) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", exporttomusicxml.GongGetGongstructName(), exporttomusicxml.GongGetReferenceOrder(stage))
}

func (frontcurve *FrontCurve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurve.GongGetGongstructName(), frontcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frontcurve *FrontCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurve.GongGetGongstructName(), frontcurve.GongGetReferenceOrder(stage))
}

func (frontcurvestack *FrontCurveStack) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurvestack.GongGetGongstructName(), frontcurvestack.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frontcurvestack *FrontCurveStack) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frontcurvestack.GongGetGongstructName(), frontcurvestack.GongGetReferenceOrder(stage))
}

func (horizontalaxis *HorizontalAxis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontalaxis.GongGetGongstructName(), horizontalaxis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (horizontalaxis *HorizontalAxis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontalaxis.GongGetGongstructName(), horizontalaxis.GongGetReferenceOrder(stage))
}

func (key *Key) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key *Key) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetReferenceOrder(stage))
}

func (parameter *Parameter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parameter *Parameter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parameter.GongGetGongstructName(), parameter.GongGetReferenceOrder(stage))
}

func (rhombus *Rhombus) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombus.GongGetGongstructName(), rhombus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombus *Rhombus) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombus.GongGetGongstructName(), rhombus.GongGetReferenceOrder(stage))
}

func (rhombusgrid *RhombusGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusgrid.GongGetGongstructName(), rhombusgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusgrid *RhombusGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusgrid.GongGetGongstructName(), rhombusgrid.GongGetReferenceOrder(stage))
}

func (shapecategory *ShapeCategory) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shapecategory.GongGetGongstructName(), shapecategory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shapecategory *ShapeCategory) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shapecategory.GongGetGongstructName(), shapecategory.GongGetReferenceOrder(stage))
}

func (spiralbezier *SpiralBezier) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbezier.GongGetGongstructName(), spiralbezier.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralbezier *SpiralBezier) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbezier.GongGetGongstructName(), spiralbezier.GongGetReferenceOrder(stage))
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbeziergrid.GongGetGongstructName(), spiralbeziergrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralbeziergrid *SpiralBezierGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralbeziergrid.GongGetGongstructName(), spiralbeziergrid.GongGetReferenceOrder(stage))
}

func (spiralcircle *SpiralCircle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcircle.GongGetGongstructName(), spiralcircle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralcircle *SpiralCircle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcircle.GongGetGongstructName(), spiralcircle.GongGetReferenceOrder(stage))
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcirclegrid.GongGetGongstructName(), spiralcirclegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralcirclegrid *SpiralCircleGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralcirclegrid.GongGetGongstructName(), spiralcirclegrid.GongGetReferenceOrder(stage))
}

func (spiralline *SpiralLine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralline.GongGetGongstructName(), spiralline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralline *SpiralLine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralline.GongGetGongstructName(), spiralline.GongGetReferenceOrder(stage))
}

func (spirallinegrid *SpiralLineGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spirallinegrid.GongGetGongstructName(), spirallinegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spirallinegrid *SpiralLineGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spirallinegrid.GongGetGongstructName(), spirallinegrid.GongGetReferenceOrder(stage))
}

func (spiralorigin *SpiralOrigin) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralorigin.GongGetGongstructName(), spiralorigin.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralorigin *SpiralOrigin) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralorigin.GongGetGongstructName(), spiralorigin.GongGetReferenceOrder(stage))
}

func (spiralrhombus *SpiralRhombus) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombus.GongGetGongstructName(), spiralrhombus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralrhombus *SpiralRhombus) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombus.GongGetGongstructName(), spiralrhombus.GongGetReferenceOrder(stage))
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombusgrid.GongGetGongstructName(), spiralrhombusgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spiralrhombusgrid *SpiralRhombusGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spiralrhombusgrid.GongGetGongstructName(), spiralrhombusgrid.GongGetReferenceOrder(stage))
}

func (verticalaxis *VerticalAxis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticalaxis.GongGetGongstructName(), verticalaxis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalaxis *VerticalAxis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticalaxis.GongGetGongstructName(), verticalaxis.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (axis *Axis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Axis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", axis.Name)
	return
}
func (axisgrid *AxisGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxisGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", axisgrid.Name)
	return
}
func (bezier *Bezier) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bezier.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bezier")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bezier.Name)
	return
}
func (beziergrid *BezierGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beziergrid.Name)
	return
}
func (beziergridstack *BezierGridStack) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGridStack")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beziergridstack.Name)
	return
}
func (chapter *Chapter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", chapter.Name)
	return
}
func (circle *Circle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circle.Name)
	return
}
func (circlegrid *CircleGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CircleGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circlegrid.Name)
	return
}
func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
	return
}
func (exporttomusicxml *ExportToMusicxml) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExportToMusicxml")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", exporttomusicxml.Name)
	return
}
func (frontcurve *FrontCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frontcurve.Name)
	return
}
func (frontcurvestack *FrontCurveStack) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurveStack")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frontcurvestack.Name)
	return
}
func (horizontalaxis *HorizontalAxis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "HorizontalAxis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", horizontalaxis.Name)
	return
}
func (key *Key) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", key.Name)
	return
}
func (parameter *Parameter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parameter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Parameter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", parameter.Name)
	return
}
func (rhombus *Rhombus) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rhombus")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rhombus.Name)
	return
}
func (rhombusgrid *RhombusGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rhombusgrid.Name)
	return
}
func (shapecategory *ShapeCategory) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shapecategory.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShapeCategory")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", shapecategory.Name)
	return
}
func (spiralbezier *SpiralBezier) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezier")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralbezier.Name)
	return
}
func (spiralbeziergrid *SpiralBezierGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezierGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralbeziergrid.Name)
	return
}
func (spiralcircle *SpiralCircle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDeclsWithoutNameInit
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralcircle.Name)
	return
}
func (spiralcirclegrid *SpiralCircleGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircleGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralcirclegrid.Name)
	return
}
func (spiralline *SpiralLine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralline.Name)
	return
}
func (spirallinegrid *SpiralLineGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLineGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spirallinegrid.Name)
	return
}
func (spiralorigin *SpiralOrigin) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralOrigin")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralorigin.Name)
	return
}
func (spiralrhombus *SpiralRhombus) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombus")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralrhombus.Name)
	return
}
func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombusGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralrhombusgrid.Name)
	return
}
func (verticalaxis *VerticalAxis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VerticalAxis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", verticalaxis.Name)
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
