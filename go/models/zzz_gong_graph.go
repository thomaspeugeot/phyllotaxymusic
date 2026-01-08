// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Axis:
		ok = stage.IsStagedAxis(target)

	case *AxisGrid:
		ok = stage.IsStagedAxisGrid(target)

	case *Bezier:
		ok = stage.IsStagedBezier(target)

	case *BezierGrid:
		ok = stage.IsStagedBezierGrid(target)

	case *BezierGridStack:
		ok = stage.IsStagedBezierGridStack(target)

	case *Chapter:
		ok = stage.IsStagedChapter(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *CircleGrid:
		ok = stage.IsStagedCircleGrid(target)

	case *Content:
		ok = stage.IsStagedContent(target)

	case *ExportToMusicxml:
		ok = stage.IsStagedExportToMusicxml(target)

	case *FrontCurve:
		ok = stage.IsStagedFrontCurve(target)

	case *FrontCurveStack:
		ok = stage.IsStagedFrontCurveStack(target)

	case *HorizontalAxis:
		ok = stage.IsStagedHorizontalAxis(target)

	case *Key:
		ok = stage.IsStagedKey(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *Rhombus:
		ok = stage.IsStagedRhombus(target)

	case *RhombusGrid:
		ok = stage.IsStagedRhombusGrid(target)

	case *ShapeCategory:
		ok = stage.IsStagedShapeCategory(target)

	case *SpiralBezier:
		ok = stage.IsStagedSpiralBezier(target)

	case *SpiralBezierGrid:
		ok = stage.IsStagedSpiralBezierGrid(target)

	case *SpiralCircle:
		ok = stage.IsStagedSpiralCircle(target)

	case *SpiralCircleGrid:
		ok = stage.IsStagedSpiralCircleGrid(target)

	case *SpiralLine:
		ok = stage.IsStagedSpiralLine(target)

	case *SpiralLineGrid:
		ok = stage.IsStagedSpiralLineGrid(target)

	case *SpiralOrigin:
		ok = stage.IsStagedSpiralOrigin(target)

	case *SpiralRhombus:
		ok = stage.IsStagedSpiralRhombus(target)

	case *SpiralRhombusGrid:
		ok = stage.IsStagedSpiralRhombusGrid(target)

	case *VerticalAxis:
		ok = stage.IsStagedVerticalAxis(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Axis:
		ok = stage.IsStagedAxis(target)

	case *AxisGrid:
		ok = stage.IsStagedAxisGrid(target)

	case *Bezier:
		ok = stage.IsStagedBezier(target)

	case *BezierGrid:
		ok = stage.IsStagedBezierGrid(target)

	case *BezierGridStack:
		ok = stage.IsStagedBezierGridStack(target)

	case *Chapter:
		ok = stage.IsStagedChapter(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *CircleGrid:
		ok = stage.IsStagedCircleGrid(target)

	case *Content:
		ok = stage.IsStagedContent(target)

	case *ExportToMusicxml:
		ok = stage.IsStagedExportToMusicxml(target)

	case *FrontCurve:
		ok = stage.IsStagedFrontCurve(target)

	case *FrontCurveStack:
		ok = stage.IsStagedFrontCurveStack(target)

	case *HorizontalAxis:
		ok = stage.IsStagedHorizontalAxis(target)

	case *Key:
		ok = stage.IsStagedKey(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *Rhombus:
		ok = stage.IsStagedRhombus(target)

	case *RhombusGrid:
		ok = stage.IsStagedRhombusGrid(target)

	case *ShapeCategory:
		ok = stage.IsStagedShapeCategory(target)

	case *SpiralBezier:
		ok = stage.IsStagedSpiralBezier(target)

	case *SpiralBezierGrid:
		ok = stage.IsStagedSpiralBezierGrid(target)

	case *SpiralCircle:
		ok = stage.IsStagedSpiralCircle(target)

	case *SpiralCircleGrid:
		ok = stage.IsStagedSpiralCircleGrid(target)

	case *SpiralLine:
		ok = stage.IsStagedSpiralLine(target)

	case *SpiralLineGrid:
		ok = stage.IsStagedSpiralLineGrid(target)

	case *SpiralOrigin:
		ok = stage.IsStagedSpiralOrigin(target)

	case *SpiralRhombus:
		ok = stage.IsStagedSpiralRhombus(target)

	case *SpiralRhombusGrid:
		ok = stage.IsStagedSpiralRhombusGrid(target)

	case *VerticalAxis:
		ok = stage.IsStagedVerticalAxis(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAxis(axis *Axis) (ok bool) {

	_, ok = stage.Axiss[axis]

	return
}

func (stage *Stage) IsStagedAxisGrid(axisgrid *AxisGrid) (ok bool) {

	_, ok = stage.AxisGrids[axisgrid]

	return
}

func (stage *Stage) IsStagedBezier(bezier *Bezier) (ok bool) {

	_, ok = stage.Beziers[bezier]

	return
}

func (stage *Stage) IsStagedBezierGrid(beziergrid *BezierGrid) (ok bool) {

	_, ok = stage.BezierGrids[beziergrid]

	return
}

func (stage *Stage) IsStagedBezierGridStack(beziergridstack *BezierGridStack) (ok bool) {

	_, ok = stage.BezierGridStacks[beziergridstack]

	return
}

func (stage *Stage) IsStagedChapter(chapter *Chapter) (ok bool) {

	_, ok = stage.Chapters[chapter]

	return
}

func (stage *Stage) IsStagedCircle(circle *Circle) (ok bool) {

	_, ok = stage.Circles[circle]

	return
}

func (stage *Stage) IsStagedCircleGrid(circlegrid *CircleGrid) (ok bool) {

	_, ok = stage.CircleGrids[circlegrid]

	return
}

func (stage *Stage) IsStagedContent(content *Content) (ok bool) {

	_, ok = stage.Contents[content]

	return
}

func (stage *Stage) IsStagedExportToMusicxml(exporttomusicxml *ExportToMusicxml) (ok bool) {

	_, ok = stage.ExportToMusicxmls[exporttomusicxml]

	return
}

func (stage *Stage) IsStagedFrontCurve(frontcurve *FrontCurve) (ok bool) {

	_, ok = stage.FrontCurves[frontcurve]

	return
}

func (stage *Stage) IsStagedFrontCurveStack(frontcurvestack *FrontCurveStack) (ok bool) {

	_, ok = stage.FrontCurveStacks[frontcurvestack]

	return
}

func (stage *Stage) IsStagedHorizontalAxis(horizontalaxis *HorizontalAxis) (ok bool) {

	_, ok = stage.HorizontalAxiss[horizontalaxis]

	return
}

func (stage *Stage) IsStagedKey(key *Key) (ok bool) {

	_, ok = stage.Keys[key]

	return
}

func (stage *Stage) IsStagedParameter(parameter *Parameter) (ok bool) {

	_, ok = stage.Parameters[parameter]

	return
}

func (stage *Stage) IsStagedRhombus(rhombus *Rhombus) (ok bool) {

	_, ok = stage.Rhombuss[rhombus]

	return
}

func (stage *Stage) IsStagedRhombusGrid(rhombusgrid *RhombusGrid) (ok bool) {

	_, ok = stage.RhombusGrids[rhombusgrid]

	return
}

func (stage *Stage) IsStagedShapeCategory(shapecategory *ShapeCategory) (ok bool) {

	_, ok = stage.ShapeCategorys[shapecategory]

	return
}

func (stage *Stage) IsStagedSpiralBezier(spiralbezier *SpiralBezier) (ok bool) {

	_, ok = stage.SpiralBeziers[spiralbezier]

	return
}

func (stage *Stage) IsStagedSpiralBezierGrid(spiralbeziergrid *SpiralBezierGrid) (ok bool) {

	_, ok = stage.SpiralBezierGrids[spiralbeziergrid]

	return
}

func (stage *Stage) IsStagedSpiralCircle(spiralcircle *SpiralCircle) (ok bool) {

	_, ok = stage.SpiralCircles[spiralcircle]

	return
}

func (stage *Stage) IsStagedSpiralCircleGrid(spiralcirclegrid *SpiralCircleGrid) (ok bool) {

	_, ok = stage.SpiralCircleGrids[spiralcirclegrid]

	return
}

func (stage *Stage) IsStagedSpiralLine(spiralline *SpiralLine) (ok bool) {

	_, ok = stage.SpiralLines[spiralline]

	return
}

func (stage *Stage) IsStagedSpiralLineGrid(spirallinegrid *SpiralLineGrid) (ok bool) {

	_, ok = stage.SpiralLineGrids[spirallinegrid]

	return
}

func (stage *Stage) IsStagedSpiralOrigin(spiralorigin *SpiralOrigin) (ok bool) {

	_, ok = stage.SpiralOrigins[spiralorigin]

	return
}

func (stage *Stage) IsStagedSpiralRhombus(spiralrhombus *SpiralRhombus) (ok bool) {

	_, ok = stage.SpiralRhombuss[spiralrhombus]

	return
}

func (stage *Stage) IsStagedSpiralRhombusGrid(spiralrhombusgrid *SpiralRhombusGrid) (ok bool) {

	_, ok = stage.SpiralRhombusGrids[spiralrhombusgrid]

	return
}

func (stage *Stage) IsStagedVerticalAxis(verticalaxis *VerticalAxis) (ok bool) {

	_, ok = stage.VerticalAxiss[verticalaxis]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Axis:
		stage.StageBranchAxis(target)

	case *AxisGrid:
		stage.StageBranchAxisGrid(target)

	case *Bezier:
		stage.StageBranchBezier(target)

	case *BezierGrid:
		stage.StageBranchBezierGrid(target)

	case *BezierGridStack:
		stage.StageBranchBezierGridStack(target)

	case *Chapter:
		stage.StageBranchChapter(target)

	case *Circle:
		stage.StageBranchCircle(target)

	case *CircleGrid:
		stage.StageBranchCircleGrid(target)

	case *Content:
		stage.StageBranchContent(target)

	case *ExportToMusicxml:
		stage.StageBranchExportToMusicxml(target)

	case *FrontCurve:
		stage.StageBranchFrontCurve(target)

	case *FrontCurveStack:
		stage.StageBranchFrontCurveStack(target)

	case *HorizontalAxis:
		stage.StageBranchHorizontalAxis(target)

	case *Key:
		stage.StageBranchKey(target)

	case *Parameter:
		stage.StageBranchParameter(target)

	case *Rhombus:
		stage.StageBranchRhombus(target)

	case *RhombusGrid:
		stage.StageBranchRhombusGrid(target)

	case *ShapeCategory:
		stage.StageBranchShapeCategory(target)

	case *SpiralBezier:
		stage.StageBranchSpiralBezier(target)

	case *SpiralBezierGrid:
		stage.StageBranchSpiralBezierGrid(target)

	case *SpiralCircle:
		stage.StageBranchSpiralCircle(target)

	case *SpiralCircleGrid:
		stage.StageBranchSpiralCircleGrid(target)

	case *SpiralLine:
		stage.StageBranchSpiralLine(target)

	case *SpiralLineGrid:
		stage.StageBranchSpiralLineGrid(target)

	case *SpiralOrigin:
		stage.StageBranchSpiralOrigin(target)

	case *SpiralRhombus:
		stage.StageBranchSpiralRhombus(target)

	case *SpiralRhombusGrid:
		stage.StageBranchSpiralRhombusGrid(target)

	case *VerticalAxis:
		stage.StageBranchVerticalAxis(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAxis(axis *Axis) {

	// check if instance is already staged
	if IsStaged(stage, axis) {
		return
	}

	axis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if axis.ShapeCategory != nil {
		StageBranch(stage, axis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchAxisGrid(axisgrid *AxisGrid) {

	// check if instance is already staged
	if IsStaged(stage, axisgrid) {
		return
	}

	axisgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if axisgrid.Reference != nil {
		StageBranch(stage, axisgrid.Reference)
	}
	if axisgrid.ShapeCategory != nil {
		StageBranch(stage, axisgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _axis := range axisgrid.Axiss {
		StageBranch(stage, _axis)
	}

}

func (stage *Stage) StageBranchBezier(bezier *Bezier) {

	// check if instance is already staged
	if IsStaged(stage, bezier) {
		return
	}

	bezier.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bezier.ShapeCategory != nil {
		StageBranch(stage, bezier.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBezierGrid(beziergrid *BezierGrid) {

	// check if instance is already staged
	if IsStaged(stage, beziergrid) {
		return
	}

	beziergrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziergrid.Reference != nil {
		StageBranch(stage, beziergrid.Reference)
	}
	if beziergrid.ShapeCategory != nil {
		StageBranch(stage, beziergrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bezier := range beziergrid.Beziers {
		StageBranch(stage, _bezier)
	}

}

func (stage *Stage) StageBranchBezierGridStack(beziergridstack *BezierGridStack) {

	// check if instance is already staged
	if IsStaged(stage, beziergridstack) {
		return
	}

	beziergridstack.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziergridstack.ShapeCategory != nil {
		StageBranch(stage, beziergridstack.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziergrid := range beziergridstack.BezierGrids {
		StageBranch(stage, _beziergrid)
	}

}

func (stage *Stage) StageBranchChapter(chapter *Chapter) {

	// check if instance is already staged
	if IsStaged(stage, chapter) {
		return
	}

	chapter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if IsStaged(stage, circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circle.ShapeCategory != nil {
		StageBranch(stage, circle.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCircleGrid(circlegrid *CircleGrid) {

	// check if instance is already staged
	if IsStaged(stage, circlegrid) {
		return
	}

	circlegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circlegrid.Reference != nil {
		StageBranch(stage, circlegrid.Reference)
	}
	if circlegrid.ShapeCategory != nil {
		StageBranch(stage, circlegrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _circle := range circlegrid.Circles {
		StageBranch(stage, _circle)
	}

}

func (stage *Stage) StageBranchContent(content *Content) {

	// check if instance is already staged
	if IsStaged(stage, content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		StageBranch(stage, _chapter)
	}

}

func (stage *Stage) StageBranchExportToMusicxml(exporttomusicxml *ExportToMusicxml) {

	// check if instance is already staged
	if IsStaged(stage, exporttomusicxml) {
		return
	}

	exporttomusicxml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if exporttomusicxml.Parameter != nil {
		StageBranch(stage, exporttomusicxml.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFrontCurve(frontcurve *FrontCurve) {

	// check if instance is already staged
	if IsStaged(stage, frontcurve) {
		return
	}

	frontcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFrontCurveStack(frontcurvestack *FrontCurveStack) {

	// check if instance is already staged
	if IsStaged(stage, frontcurvestack) {
		return
	}

	frontcurvestack.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frontcurvestack.ShapeCategory != nil {
		StageBranch(stage, frontcurvestack.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frontcurve := range frontcurvestack.FrontCurves {
		StageBranch(stage, _frontcurve)
	}
	for _, _spiralcircle := range frontcurvestack.SpiralCircles {
		StageBranch(stage, _spiralcircle)
	}

}

func (stage *Stage) StageBranchHorizontalAxis(horizontalaxis *HorizontalAxis) {

	// check if instance is already staged
	if IsStaged(stage, horizontalaxis) {
		return
	}

	horizontalaxis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if horizontalaxis.ShapeCategory != nil {
		StageBranch(stage, horizontalaxis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchKey(key *Key) {

	// check if instance is already staged
	if IsStaged(stage, key) {
		return
	}

	key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.ShapeCategory != nil {
		StageBranch(stage, key.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParameter(parameter *Parameter) {

	// check if instance is already staged
	if IsStaged(stage, parameter) {
		return
	}

	parameter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parameter.InitialRhombus != nil {
		StageBranch(stage, parameter.InitialRhombus)
	}
	if parameter.InitialCircle != nil {
		StageBranch(stage, parameter.InitialCircle)
	}
	if parameter.InitialRhombusGrid != nil {
		StageBranch(stage, parameter.InitialRhombusGrid)
	}
	if parameter.InitialCircleGrid != nil {
		StageBranch(stage, parameter.InitialCircleGrid)
	}
	if parameter.InitialAxis != nil {
		StageBranch(stage, parameter.InitialAxis)
	}
	if parameter.RotatedAxis != nil {
		StageBranch(stage, parameter.RotatedAxis)
	}
	if parameter.RotatedRhombus != nil {
		StageBranch(stage, parameter.RotatedRhombus)
	}
	if parameter.RotatedRhombusGrid != nil {
		StageBranch(stage, parameter.RotatedRhombusGrid)
	}
	if parameter.RotatedCircleGrid != nil {
		StageBranch(stage, parameter.RotatedCircleGrid)
	}
	if parameter.NextRhombus != nil {
		StageBranch(stage, parameter.NextRhombus)
	}
	if parameter.NextCircle != nil {
		StageBranch(stage, parameter.NextCircle)
	}
	if parameter.GrowingRhombusGridSeed != nil {
		StageBranch(stage, parameter.GrowingRhombusGridSeed)
	}
	if parameter.GrowingRhombusGrid != nil {
		StageBranch(stage, parameter.GrowingRhombusGrid)
	}
	if parameter.GrowingCircleGridSeed != nil {
		StageBranch(stage, parameter.GrowingCircleGridSeed)
	}
	if parameter.GrowingCircleGrid != nil {
		StageBranch(stage, parameter.GrowingCircleGrid)
	}
	if parameter.GrowingCircleGridLeftSeed != nil {
		StageBranch(stage, parameter.GrowingCircleGridLeftSeed)
	}
	if parameter.GrowingCircleGridLeft != nil {
		StageBranch(stage, parameter.GrowingCircleGridLeft)
	}
	if parameter.ConstructionAxis != nil {
		StageBranch(stage, parameter.ConstructionAxis)
	}
	if parameter.ConstructionAxisGrid != nil {
		StageBranch(stage, parameter.ConstructionAxisGrid)
	}
	if parameter.ConstructionCircle != nil {
		StageBranch(stage, parameter.ConstructionCircle)
	}
	if parameter.ConstructionCircleGrid != nil {
		StageBranch(stage, parameter.ConstructionCircleGrid)
	}
	if parameter.GrowthCurveSeed != nil {
		StageBranch(stage, parameter.GrowthCurveSeed)
	}
	if parameter.GrowthCurve != nil {
		StageBranch(stage, parameter.GrowthCurve)
	}
	if parameter.GrowthCurveShiftedRightSeed != nil {
		StageBranch(stage, parameter.GrowthCurveShiftedRightSeed)
	}
	if parameter.GrowthCurveShiftedRight != nil {
		StageBranch(stage, parameter.GrowthCurveShiftedRight)
	}
	if parameter.GrowthCurveNextSeed != nil {
		StageBranch(stage, parameter.GrowthCurveNextSeed)
	}
	if parameter.GrowthCurveNext != nil {
		StageBranch(stage, parameter.GrowthCurveNext)
	}
	if parameter.GrowthCurveNextShiftedRightSeed != nil {
		StageBranch(stage, parameter.GrowthCurveNextShiftedRightSeed)
	}
	if parameter.GrowthCurveNextShiftedRight != nil {
		StageBranch(stage, parameter.GrowthCurveNextShiftedRight)
	}
	if parameter.GrowthCurveStack != nil {
		StageBranch(stage, parameter.GrowthCurveStack)
	}
	if parameter.SpiralRhombusGridSeed != nil {
		StageBranch(stage, parameter.SpiralRhombusGridSeed)
	}
	if parameter.SpiralRhombusGrid != nil {
		StageBranch(stage, parameter.SpiralRhombusGrid)
	}
	if parameter.SpiralCircleSeed != nil {
		StageBranch(stage, parameter.SpiralCircleSeed)
	}
	if parameter.SpiralCircleGrid != nil {
		StageBranch(stage, parameter.SpiralCircleGrid)
	}
	if parameter.SpiralCircleFullGrid != nil {
		StageBranch(stage, parameter.SpiralCircleFullGrid)
	}
	if parameter.SpiralConstructionOuterLineSeed != nil {
		StageBranch(stage, parameter.SpiralConstructionOuterLineSeed)
	}
	if parameter.SpiralConstructionInnerLineSeed != nil {
		StageBranch(stage, parameter.SpiralConstructionInnerLineSeed)
	}
	if parameter.SpiralConstructionOuterLineGrid != nil {
		StageBranch(stage, parameter.SpiralConstructionOuterLineGrid)
	}
	if parameter.SpiralConstructionInnerLineGrid != nil {
		StageBranch(stage, parameter.SpiralConstructionInnerLineGrid)
	}
	if parameter.SpiralConstructionCircleGrid != nil {
		StageBranch(stage, parameter.SpiralConstructionCircleGrid)
	}
	if parameter.SpiralConstructionOuterLineFullGrid != nil {
		StageBranch(stage, parameter.SpiralConstructionOuterLineFullGrid)
	}
	if parameter.SpiralBezierSeed != nil {
		StageBranch(stage, parameter.SpiralBezierSeed)
	}
	if parameter.SpiralBezierGrid != nil {
		StageBranch(stage, parameter.SpiralBezierGrid)
	}
	if parameter.SpiralBezierFullGrid != nil {
		StageBranch(stage, parameter.SpiralBezierFullGrid)
	}
	if parameter.FrontCurveStack != nil {
		StageBranch(stage, parameter.FrontCurveStack)
	}
	if parameter.Fkey != nil {
		StageBranch(stage, parameter.Fkey)
	}
	if parameter.PitchLines != nil {
		StageBranch(stage, parameter.PitchLines)
	}
	if parameter.BeatLines != nil {
		StageBranch(stage, parameter.BeatLines)
	}
	if parameter.FirstVoice != nil {
		StageBranch(stage, parameter.FirstVoice)
	}
	if parameter.FirstVoiceShiftedRigth != nil {
		StageBranch(stage, parameter.FirstVoiceShiftedRigth)
	}
	if parameter.SecondVoice != nil {
		StageBranch(stage, parameter.SecondVoice)
	}
	if parameter.SecondVoiceShiftedRight != nil {
		StageBranch(stage, parameter.SecondVoiceShiftedRight)
	}
	if parameter.FirstVoiceNotes != nil {
		StageBranch(stage, parameter.FirstVoiceNotes)
	}
	if parameter.FirstVoiceNotesShiftedRight != nil {
		StageBranch(stage, parameter.FirstVoiceNotesShiftedRight)
	}
	if parameter.SecondVoiceNotes != nil {
		StageBranch(stage, parameter.SecondVoiceNotes)
	}
	if parameter.SecondVoiceNotesShiftedRight != nil {
		StageBranch(stage, parameter.SecondVoiceNotesShiftedRight)
	}
	if parameter.HorizontalAxis != nil {
		StageBranch(stage, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis != nil {
		StageBranch(stage, parameter.VerticalAxis)
	}
	if parameter.SpiralOrigin != nil {
		StageBranch(stage, parameter.SpiralOrigin)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRhombus(rhombus *Rhombus) {

	// check if instance is already staged
	if IsStaged(stage, rhombus) {
		return
	}

	rhombus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rhombus.ShapeCategory != nil {
		StageBranch(stage, rhombus.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRhombusGrid(rhombusgrid *RhombusGrid) {

	// check if instance is already staged
	if IsStaged(stage, rhombusgrid) {
		return
	}

	rhombusgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rhombusgrid.Reference != nil {
		StageBranch(stage, rhombusgrid.Reference)
	}
	if rhombusgrid.ShapeCategory != nil {
		StageBranch(stage, rhombusgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgrid.Rhombuses {
		StageBranch(stage, _rhombus)
	}

}

func (stage *Stage) StageBranchShapeCategory(shapecategory *ShapeCategory) {

	// check if instance is already staged
	if IsStaged(stage, shapecategory) {
		return
	}

	shapecategory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralBezier(spiralbezier *SpiralBezier) {

	// check if instance is already staged
	if IsStaged(stage, spiralbezier) {
		return
	}

	spiralbezier.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralbezier.ShapeCategory != nil {
		StageBranch(stage, spiralbezier.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralBezierGrid(spiralbeziergrid *SpiralBezierGrid) {

	// check if instance is already staged
	if IsStaged(stage, spiralbeziergrid) {
		return
	}

	spiralbeziergrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralbeziergrid.ShapeCategory != nil {
		StageBranch(stage, spiralbeziergrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
		StageBranch(stage, _spiralbezier)
	}

}

func (stage *Stage) StageBranchSpiralCircle(spiralcircle *SpiralCircle) {

	// check if instance is already staged
	if IsStaged(stage, spiralcircle) {
		return
	}

	spiralcircle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralcircle.ShapeCategory != nil {
		StageBranch(stage, spiralcircle.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralCircleGrid(spiralcirclegrid *SpiralCircleGrid) {

	// check if instance is already staged
	if IsStaged(stage, spiralcirclegrid) {
		return
	}

	spiralcirclegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralcirclegrid.ShapeCategory != nil {
		StageBranch(stage, spiralcirclegrid.ShapeCategory)
	}
	if spiralcirclegrid.SpiralRhombusGrid != nil {
		StageBranch(stage, spiralcirclegrid.SpiralRhombusGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
		StageBranch(stage, _spiralcircle)
	}

}

func (stage *Stage) StageBranchSpiralLine(spiralline *SpiralLine) {

	// check if instance is already staged
	if IsStaged(stage, spiralline) {
		return
	}

	spiralline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralline.ShapeCategory != nil {
		StageBranch(stage, spiralline.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralLineGrid(spirallinegrid *SpiralLineGrid) {

	// check if instance is already staged
	if IsStaged(stage, spirallinegrid) {
		return
	}

	spirallinegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spirallinegrid.ShapeCategory != nil {
		StageBranch(stage, spirallinegrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralline := range spirallinegrid.SpiralLines {
		StageBranch(stage, _spiralline)
	}

}

func (stage *Stage) StageBranchSpiralOrigin(spiralorigin *SpiralOrigin) {

	// check if instance is already staged
	if IsStaged(stage, spiralorigin) {
		return
	}

	spiralorigin.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralorigin.ShapeCategory != nil {
		StageBranch(stage, spiralorigin.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralRhombus(spiralrhombus *SpiralRhombus) {

	// check if instance is already staged
	if IsStaged(stage, spiralrhombus) {
		return
	}

	spiralrhombus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombus.ShapeCategory != nil {
		StageBranch(stage, spiralrhombus.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSpiralRhombusGrid(spiralrhombusgrid *SpiralRhombusGrid) {

	// check if instance is already staged
	if IsStaged(stage, spiralrhombusgrid) {
		return
	}

	spiralrhombusgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombusgrid.ShapeCategory != nil {
		StageBranch(stage, spiralrhombusgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
		StageBranch(stage, _spiralrhombus)
	}

}

func (stage *Stage) StageBranchVerticalAxis(verticalaxis *VerticalAxis) {

	// check if instance is already staged
	if IsStaged(stage, verticalaxis) {
		return
	}

	verticalaxis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalaxis.ShapeCategory != nil {
		StageBranch(stage, verticalaxis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Axis:
		toT := CopyBranchAxis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AxisGrid:
		toT := CopyBranchAxisGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bezier:
		toT := CopyBranchBezier(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BezierGrid:
		toT := CopyBranchBezierGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BezierGridStack:
		toT := CopyBranchBezierGridStack(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Chapter:
		toT := CopyBranchChapter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Circle:
		toT := CopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGrid:
		toT := CopyBranchCircleGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Content:
		toT := CopyBranchContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExportToMusicxml:
		toT := CopyBranchExportToMusicxml(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FrontCurve:
		toT := CopyBranchFrontCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FrontCurveStack:
		toT := CopyBranchFrontCurveStack(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HorizontalAxis:
		toT := CopyBranchHorizontalAxis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key:
		toT := CopyBranchKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Parameter:
		toT := CopyBranchParameter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rhombus:
		toT := CopyBranchRhombus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusGrid:
		toT := CopyBranchRhombusGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShapeCategory:
		toT := CopyBranchShapeCategory(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralBezier:
		toT := CopyBranchSpiralBezier(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralBezierGrid:
		toT := CopyBranchSpiralBezierGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralCircle:
		toT := CopyBranchSpiralCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralCircleGrid:
		toT := CopyBranchSpiralCircleGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralLine:
		toT := CopyBranchSpiralLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralLineGrid:
		toT := CopyBranchSpiralLineGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralOrigin:
		toT := CopyBranchSpiralOrigin(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralRhombus:
		toT := CopyBranchSpiralRhombus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpiralRhombusGrid:
		toT := CopyBranchSpiralRhombusGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalAxis:
		toT := CopyBranchVerticalAxis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAxis(mapOrigCopy map[any]any, axisFrom *Axis) (axisTo *Axis) {

	// axisFrom has already been copied
	if _axisTo, ok := mapOrigCopy[axisFrom]; ok {
		axisTo = _axisTo.(*Axis)
		return
	}

	axisTo = new(Axis)
	mapOrigCopy[axisFrom] = axisTo
	axisFrom.CopyBasicFields(axisTo)

	//insertion point for the staging of instances referenced by pointers
	if axisFrom.ShapeCategory != nil {
		axisTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, axisFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAxisGrid(mapOrigCopy map[any]any, axisgridFrom *AxisGrid) (axisgridTo *AxisGrid) {

	// axisgridFrom has already been copied
	if _axisgridTo, ok := mapOrigCopy[axisgridFrom]; ok {
		axisgridTo = _axisgridTo.(*AxisGrid)
		return
	}

	axisgridTo = new(AxisGrid)
	mapOrigCopy[axisgridFrom] = axisgridTo
	axisgridFrom.CopyBasicFields(axisgridTo)

	//insertion point for the staging of instances referenced by pointers
	if axisgridFrom.Reference != nil {
		axisgridTo.Reference = CopyBranchAxis(mapOrigCopy, axisgridFrom.Reference)
	}
	if axisgridFrom.ShapeCategory != nil {
		axisgridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, axisgridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _axis := range axisgridFrom.Axiss {
		axisgridTo.Axiss = append(axisgridTo.Axiss, CopyBranchAxis(mapOrigCopy, _axis))
	}

	return
}

func CopyBranchBezier(mapOrigCopy map[any]any, bezierFrom *Bezier) (bezierTo *Bezier) {

	// bezierFrom has already been copied
	if _bezierTo, ok := mapOrigCopy[bezierFrom]; ok {
		bezierTo = _bezierTo.(*Bezier)
		return
	}

	bezierTo = new(Bezier)
	mapOrigCopy[bezierFrom] = bezierTo
	bezierFrom.CopyBasicFields(bezierTo)

	//insertion point for the staging of instances referenced by pointers
	if bezierFrom.ShapeCategory != nil {
		bezierTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, bezierFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBezierGrid(mapOrigCopy map[any]any, beziergridFrom *BezierGrid) (beziergridTo *BezierGrid) {

	// beziergridFrom has already been copied
	if _beziergridTo, ok := mapOrigCopy[beziergridFrom]; ok {
		beziergridTo = _beziergridTo.(*BezierGrid)
		return
	}

	beziergridTo = new(BezierGrid)
	mapOrigCopy[beziergridFrom] = beziergridTo
	beziergridFrom.CopyBasicFields(beziergridTo)

	//insertion point for the staging of instances referenced by pointers
	if beziergridFrom.Reference != nil {
		beziergridTo.Reference = CopyBranchBezier(mapOrigCopy, beziergridFrom.Reference)
	}
	if beziergridFrom.ShapeCategory != nil {
		beziergridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, beziergridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bezier := range beziergridFrom.Beziers {
		beziergridTo.Beziers = append(beziergridTo.Beziers, CopyBranchBezier(mapOrigCopy, _bezier))
	}

	return
}

func CopyBranchBezierGridStack(mapOrigCopy map[any]any, beziergridstackFrom *BezierGridStack) (beziergridstackTo *BezierGridStack) {

	// beziergridstackFrom has already been copied
	if _beziergridstackTo, ok := mapOrigCopy[beziergridstackFrom]; ok {
		beziergridstackTo = _beziergridstackTo.(*BezierGridStack)
		return
	}

	beziergridstackTo = new(BezierGridStack)
	mapOrigCopy[beziergridstackFrom] = beziergridstackTo
	beziergridstackFrom.CopyBasicFields(beziergridstackTo)

	//insertion point for the staging of instances referenced by pointers
	if beziergridstackFrom.ShapeCategory != nil {
		beziergridstackTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, beziergridstackFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziergrid := range beziergridstackFrom.BezierGrids {
		beziergridstackTo.BezierGrids = append(beziergridstackTo.BezierGrids, CopyBranchBezierGrid(mapOrigCopy, _beziergrid))
	}

	return
}

func CopyBranchChapter(mapOrigCopy map[any]any, chapterFrom *Chapter) (chapterTo *Chapter) {

	// chapterFrom has already been copied
	if _chapterTo, ok := mapOrigCopy[chapterFrom]; ok {
		chapterTo = _chapterTo.(*Chapter)
		return
	}

	chapterTo = new(Chapter)
	mapOrigCopy[chapterFrom] = chapterTo
	chapterFrom.CopyBasicFields(chapterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCircle(mapOrigCopy map[any]any, circleFrom *Circle) (circleTo *Circle) {

	// circleFrom has already been copied
	if _circleTo, ok := mapOrigCopy[circleFrom]; ok {
		circleTo = _circleTo.(*Circle)
		return
	}

	circleTo = new(Circle)
	mapOrigCopy[circleFrom] = circleTo
	circleFrom.CopyBasicFields(circleTo)

	//insertion point for the staging of instances referenced by pointers
	if circleFrom.ShapeCategory != nil {
		circleTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, circleFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCircleGrid(mapOrigCopy map[any]any, circlegridFrom *CircleGrid) (circlegridTo *CircleGrid) {

	// circlegridFrom has already been copied
	if _circlegridTo, ok := mapOrigCopy[circlegridFrom]; ok {
		circlegridTo = _circlegridTo.(*CircleGrid)
		return
	}

	circlegridTo = new(CircleGrid)
	mapOrigCopy[circlegridFrom] = circlegridTo
	circlegridFrom.CopyBasicFields(circlegridTo)

	//insertion point for the staging of instances referenced by pointers
	if circlegridFrom.Reference != nil {
		circlegridTo.Reference = CopyBranchCircle(mapOrigCopy, circlegridFrom.Reference)
	}
	if circlegridFrom.ShapeCategory != nil {
		circlegridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, circlegridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _circle := range circlegridFrom.Circles {
		circlegridTo.Circles = append(circlegridTo.Circles, CopyBranchCircle(mapOrigCopy, _circle))
	}

	return
}

func CopyBranchContent(mapOrigCopy map[any]any, contentFrom *Content) (contentTo *Content) {

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*Content)
		return
	}

	contentTo = new(Content)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.CopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range contentFrom.Chapters {
		contentTo.Chapters = append(contentTo.Chapters, CopyBranchChapter(mapOrigCopy, _chapter))
	}

	return
}

func CopyBranchExportToMusicxml(mapOrigCopy map[any]any, exporttomusicxmlFrom *ExportToMusicxml) (exporttomusicxmlTo *ExportToMusicxml) {

	// exporttomusicxmlFrom has already been copied
	if _exporttomusicxmlTo, ok := mapOrigCopy[exporttomusicxmlFrom]; ok {
		exporttomusicxmlTo = _exporttomusicxmlTo.(*ExportToMusicxml)
		return
	}

	exporttomusicxmlTo = new(ExportToMusicxml)
	mapOrigCopy[exporttomusicxmlFrom] = exporttomusicxmlTo
	exporttomusicxmlFrom.CopyBasicFields(exporttomusicxmlTo)

	//insertion point for the staging of instances referenced by pointers
	if exporttomusicxmlFrom.Parameter != nil {
		exporttomusicxmlTo.Parameter = CopyBranchParameter(mapOrigCopy, exporttomusicxmlFrom.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFrontCurve(mapOrigCopy map[any]any, frontcurveFrom *FrontCurve) (frontcurveTo *FrontCurve) {

	// frontcurveFrom has already been copied
	if _frontcurveTo, ok := mapOrigCopy[frontcurveFrom]; ok {
		frontcurveTo = _frontcurveTo.(*FrontCurve)
		return
	}

	frontcurveTo = new(FrontCurve)
	mapOrigCopy[frontcurveFrom] = frontcurveTo
	frontcurveFrom.CopyBasicFields(frontcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFrontCurveStack(mapOrigCopy map[any]any, frontcurvestackFrom *FrontCurveStack) (frontcurvestackTo *FrontCurveStack) {

	// frontcurvestackFrom has already been copied
	if _frontcurvestackTo, ok := mapOrigCopy[frontcurvestackFrom]; ok {
		frontcurvestackTo = _frontcurvestackTo.(*FrontCurveStack)
		return
	}

	frontcurvestackTo = new(FrontCurveStack)
	mapOrigCopy[frontcurvestackFrom] = frontcurvestackTo
	frontcurvestackFrom.CopyBasicFields(frontcurvestackTo)

	//insertion point for the staging of instances referenced by pointers
	if frontcurvestackFrom.ShapeCategory != nil {
		frontcurvestackTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, frontcurvestackFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frontcurve := range frontcurvestackFrom.FrontCurves {
		frontcurvestackTo.FrontCurves = append(frontcurvestackTo.FrontCurves, CopyBranchFrontCurve(mapOrigCopy, _frontcurve))
	}
	for _, _spiralcircle := range frontcurvestackFrom.SpiralCircles {
		frontcurvestackTo.SpiralCircles = append(frontcurvestackTo.SpiralCircles, CopyBranchSpiralCircle(mapOrigCopy, _spiralcircle))
	}

	return
}

func CopyBranchHorizontalAxis(mapOrigCopy map[any]any, horizontalaxisFrom *HorizontalAxis) (horizontalaxisTo *HorizontalAxis) {

	// horizontalaxisFrom has already been copied
	if _horizontalaxisTo, ok := mapOrigCopy[horizontalaxisFrom]; ok {
		horizontalaxisTo = _horizontalaxisTo.(*HorizontalAxis)
		return
	}

	horizontalaxisTo = new(HorizontalAxis)
	mapOrigCopy[horizontalaxisFrom] = horizontalaxisTo
	horizontalaxisFrom.CopyBasicFields(horizontalaxisTo)

	//insertion point for the staging of instances referenced by pointers
	if horizontalaxisFrom.ShapeCategory != nil {
		horizontalaxisTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, horizontalaxisFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKey(mapOrigCopy map[any]any, keyFrom *Key) (keyTo *Key) {

	// keyFrom has already been copied
	if _keyTo, ok := mapOrigCopy[keyFrom]; ok {
		keyTo = _keyTo.(*Key)
		return
	}

	keyTo = new(Key)
	mapOrigCopy[keyFrom] = keyTo
	keyFrom.CopyBasicFields(keyTo)

	//insertion point for the staging of instances referenced by pointers
	if keyFrom.ShapeCategory != nil {
		keyTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, keyFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParameter(mapOrigCopy map[any]any, parameterFrom *Parameter) (parameterTo *Parameter) {

	// parameterFrom has already been copied
	if _parameterTo, ok := mapOrigCopy[parameterFrom]; ok {
		parameterTo = _parameterTo.(*Parameter)
		return
	}

	parameterTo = new(Parameter)
	mapOrigCopy[parameterFrom] = parameterTo
	parameterFrom.CopyBasicFields(parameterTo)

	//insertion point for the staging of instances referenced by pointers
	if parameterFrom.InitialRhombus != nil {
		parameterTo.InitialRhombus = CopyBranchRhombus(mapOrigCopy, parameterFrom.InitialRhombus)
	}
	if parameterFrom.InitialCircle != nil {
		parameterTo.InitialCircle = CopyBranchCircle(mapOrigCopy, parameterFrom.InitialCircle)
	}
	if parameterFrom.InitialRhombusGrid != nil {
		parameterTo.InitialRhombusGrid = CopyBranchRhombusGrid(mapOrigCopy, parameterFrom.InitialRhombusGrid)
	}
	if parameterFrom.InitialCircleGrid != nil {
		parameterTo.InitialCircleGrid = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.InitialCircleGrid)
	}
	if parameterFrom.InitialAxis != nil {
		parameterTo.InitialAxis = CopyBranchAxis(mapOrigCopy, parameterFrom.InitialAxis)
	}
	if parameterFrom.RotatedAxis != nil {
		parameterTo.RotatedAxis = CopyBranchAxis(mapOrigCopy, parameterFrom.RotatedAxis)
	}
	if parameterFrom.RotatedRhombus != nil {
		parameterTo.RotatedRhombus = CopyBranchRhombus(mapOrigCopy, parameterFrom.RotatedRhombus)
	}
	if parameterFrom.RotatedRhombusGrid != nil {
		parameterTo.RotatedRhombusGrid = CopyBranchRhombusGrid(mapOrigCopy, parameterFrom.RotatedRhombusGrid)
	}
	if parameterFrom.RotatedCircleGrid != nil {
		parameterTo.RotatedCircleGrid = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.RotatedCircleGrid)
	}
	if parameterFrom.NextRhombus != nil {
		parameterTo.NextRhombus = CopyBranchRhombus(mapOrigCopy, parameterFrom.NextRhombus)
	}
	if parameterFrom.NextCircle != nil {
		parameterTo.NextCircle = CopyBranchCircle(mapOrigCopy, parameterFrom.NextCircle)
	}
	if parameterFrom.GrowingRhombusGridSeed != nil {
		parameterTo.GrowingRhombusGridSeed = CopyBranchRhombus(mapOrigCopy, parameterFrom.GrowingRhombusGridSeed)
	}
	if parameterFrom.GrowingRhombusGrid != nil {
		parameterTo.GrowingRhombusGrid = CopyBranchRhombusGrid(mapOrigCopy, parameterFrom.GrowingRhombusGrid)
	}
	if parameterFrom.GrowingCircleGridSeed != nil {
		parameterTo.GrowingCircleGridSeed = CopyBranchCircle(mapOrigCopy, parameterFrom.GrowingCircleGridSeed)
	}
	if parameterFrom.GrowingCircleGrid != nil {
		parameterTo.GrowingCircleGrid = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.GrowingCircleGrid)
	}
	if parameterFrom.GrowingCircleGridLeftSeed != nil {
		parameterTo.GrowingCircleGridLeftSeed = CopyBranchCircle(mapOrigCopy, parameterFrom.GrowingCircleGridLeftSeed)
	}
	if parameterFrom.GrowingCircleGridLeft != nil {
		parameterTo.GrowingCircleGridLeft = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.GrowingCircleGridLeft)
	}
	if parameterFrom.ConstructionAxis != nil {
		parameterTo.ConstructionAxis = CopyBranchAxis(mapOrigCopy, parameterFrom.ConstructionAxis)
	}
	if parameterFrom.ConstructionAxisGrid != nil {
		parameterTo.ConstructionAxisGrid = CopyBranchAxisGrid(mapOrigCopy, parameterFrom.ConstructionAxisGrid)
	}
	if parameterFrom.ConstructionCircle != nil {
		parameterTo.ConstructionCircle = CopyBranchCircle(mapOrigCopy, parameterFrom.ConstructionCircle)
	}
	if parameterFrom.ConstructionCircleGrid != nil {
		parameterTo.ConstructionCircleGrid = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.ConstructionCircleGrid)
	}
	if parameterFrom.GrowthCurveSeed != nil {
		parameterTo.GrowthCurveSeed = CopyBranchBezier(mapOrigCopy, parameterFrom.GrowthCurveSeed)
	}
	if parameterFrom.GrowthCurve != nil {
		parameterTo.GrowthCurve = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.GrowthCurve)
	}
	if parameterFrom.GrowthCurveShiftedRightSeed != nil {
		parameterTo.GrowthCurveShiftedRightSeed = CopyBranchBezier(mapOrigCopy, parameterFrom.GrowthCurveShiftedRightSeed)
	}
	if parameterFrom.GrowthCurveShiftedRight != nil {
		parameterTo.GrowthCurveShiftedRight = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.GrowthCurveShiftedRight)
	}
	if parameterFrom.GrowthCurveNextSeed != nil {
		parameterTo.GrowthCurveNextSeed = CopyBranchBezier(mapOrigCopy, parameterFrom.GrowthCurveNextSeed)
	}
	if parameterFrom.GrowthCurveNext != nil {
		parameterTo.GrowthCurveNext = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.GrowthCurveNext)
	}
	if parameterFrom.GrowthCurveNextShiftedRightSeed != nil {
		parameterTo.GrowthCurveNextShiftedRightSeed = CopyBranchBezier(mapOrigCopy, parameterFrom.GrowthCurveNextShiftedRightSeed)
	}
	if parameterFrom.GrowthCurveNextShiftedRight != nil {
		parameterTo.GrowthCurveNextShiftedRight = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.GrowthCurveNextShiftedRight)
	}
	if parameterFrom.GrowthCurveStack != nil {
		parameterTo.GrowthCurveStack = CopyBranchBezierGridStack(mapOrigCopy, parameterFrom.GrowthCurveStack)
	}
	if parameterFrom.SpiralRhombusGridSeed != nil {
		parameterTo.SpiralRhombusGridSeed = CopyBranchSpiralRhombus(mapOrigCopy, parameterFrom.SpiralRhombusGridSeed)
	}
	if parameterFrom.SpiralRhombusGrid != nil {
		parameterTo.SpiralRhombusGrid = CopyBranchSpiralRhombusGrid(mapOrigCopy, parameterFrom.SpiralRhombusGrid)
	}
	if parameterFrom.SpiralCircleSeed != nil {
		parameterTo.SpiralCircleSeed = CopyBranchSpiralCircle(mapOrigCopy, parameterFrom.SpiralCircleSeed)
	}
	if parameterFrom.SpiralCircleGrid != nil {
		parameterTo.SpiralCircleGrid = CopyBranchSpiralCircleGrid(mapOrigCopy, parameterFrom.SpiralCircleGrid)
	}
	if parameterFrom.SpiralCircleFullGrid != nil {
		parameterTo.SpiralCircleFullGrid = CopyBranchSpiralCircleGrid(mapOrigCopy, parameterFrom.SpiralCircleFullGrid)
	}
	if parameterFrom.SpiralConstructionOuterLineSeed != nil {
		parameterTo.SpiralConstructionOuterLineSeed = CopyBranchSpiralLine(mapOrigCopy, parameterFrom.SpiralConstructionOuterLineSeed)
	}
	if parameterFrom.SpiralConstructionInnerLineSeed != nil {
		parameterTo.SpiralConstructionInnerLineSeed = CopyBranchSpiralLine(mapOrigCopy, parameterFrom.SpiralConstructionInnerLineSeed)
	}
	if parameterFrom.SpiralConstructionOuterLineGrid != nil {
		parameterTo.SpiralConstructionOuterLineGrid = CopyBranchSpiralLineGrid(mapOrigCopy, parameterFrom.SpiralConstructionOuterLineGrid)
	}
	if parameterFrom.SpiralConstructionInnerLineGrid != nil {
		parameterTo.SpiralConstructionInnerLineGrid = CopyBranchSpiralLineGrid(mapOrigCopy, parameterFrom.SpiralConstructionInnerLineGrid)
	}
	if parameterFrom.SpiralConstructionCircleGrid != nil {
		parameterTo.SpiralConstructionCircleGrid = CopyBranchSpiralCircleGrid(mapOrigCopy, parameterFrom.SpiralConstructionCircleGrid)
	}
	if parameterFrom.SpiralConstructionOuterLineFullGrid != nil {
		parameterTo.SpiralConstructionOuterLineFullGrid = CopyBranchSpiralLineGrid(mapOrigCopy, parameterFrom.SpiralConstructionOuterLineFullGrid)
	}
	if parameterFrom.SpiralBezierSeed != nil {
		parameterTo.SpiralBezierSeed = CopyBranchSpiralBezier(mapOrigCopy, parameterFrom.SpiralBezierSeed)
	}
	if parameterFrom.SpiralBezierGrid != nil {
		parameterTo.SpiralBezierGrid = CopyBranchSpiralBezierGrid(mapOrigCopy, parameterFrom.SpiralBezierGrid)
	}
	if parameterFrom.SpiralBezierFullGrid != nil {
		parameterTo.SpiralBezierFullGrid = CopyBranchSpiralBezierGrid(mapOrigCopy, parameterFrom.SpiralBezierFullGrid)
	}
	if parameterFrom.FrontCurveStack != nil {
		parameterTo.FrontCurveStack = CopyBranchFrontCurveStack(mapOrigCopy, parameterFrom.FrontCurveStack)
	}
	if parameterFrom.Fkey != nil {
		parameterTo.Fkey = CopyBranchKey(mapOrigCopy, parameterFrom.Fkey)
	}
	if parameterFrom.PitchLines != nil {
		parameterTo.PitchLines = CopyBranchAxisGrid(mapOrigCopy, parameterFrom.PitchLines)
	}
	if parameterFrom.BeatLines != nil {
		parameterTo.BeatLines = CopyBranchAxisGrid(mapOrigCopy, parameterFrom.BeatLines)
	}
	if parameterFrom.FirstVoice != nil {
		parameterTo.FirstVoice = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.FirstVoice)
	}
	if parameterFrom.FirstVoiceShiftedRigth != nil {
		parameterTo.FirstVoiceShiftedRigth = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.FirstVoiceShiftedRigth)
	}
	if parameterFrom.SecondVoice != nil {
		parameterTo.SecondVoice = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.SecondVoice)
	}
	if parameterFrom.SecondVoiceShiftedRight != nil {
		parameterTo.SecondVoiceShiftedRight = CopyBranchBezierGrid(mapOrigCopy, parameterFrom.SecondVoiceShiftedRight)
	}
	if parameterFrom.FirstVoiceNotes != nil {
		parameterTo.FirstVoiceNotes = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.FirstVoiceNotes)
	}
	if parameterFrom.FirstVoiceNotesShiftedRight != nil {
		parameterTo.FirstVoiceNotesShiftedRight = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.FirstVoiceNotesShiftedRight)
	}
	if parameterFrom.SecondVoiceNotes != nil {
		parameterTo.SecondVoiceNotes = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.SecondVoiceNotes)
	}
	if parameterFrom.SecondVoiceNotesShiftedRight != nil {
		parameterTo.SecondVoiceNotesShiftedRight = CopyBranchCircleGrid(mapOrigCopy, parameterFrom.SecondVoiceNotesShiftedRight)
	}
	if parameterFrom.HorizontalAxis != nil {
		parameterTo.HorizontalAxis = CopyBranchHorizontalAxis(mapOrigCopy, parameterFrom.HorizontalAxis)
	}
	if parameterFrom.VerticalAxis != nil {
		parameterTo.VerticalAxis = CopyBranchVerticalAxis(mapOrigCopy, parameterFrom.VerticalAxis)
	}
	if parameterFrom.SpiralOrigin != nil {
		parameterTo.SpiralOrigin = CopyBranchSpiralOrigin(mapOrigCopy, parameterFrom.SpiralOrigin)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRhombus(mapOrigCopy map[any]any, rhombusFrom *Rhombus) (rhombusTo *Rhombus) {

	// rhombusFrom has already been copied
	if _rhombusTo, ok := mapOrigCopy[rhombusFrom]; ok {
		rhombusTo = _rhombusTo.(*Rhombus)
		return
	}

	rhombusTo = new(Rhombus)
	mapOrigCopy[rhombusFrom] = rhombusTo
	rhombusFrom.CopyBasicFields(rhombusTo)

	//insertion point for the staging of instances referenced by pointers
	if rhombusFrom.ShapeCategory != nil {
		rhombusTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, rhombusFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRhombusGrid(mapOrigCopy map[any]any, rhombusgridFrom *RhombusGrid) (rhombusgridTo *RhombusGrid) {

	// rhombusgridFrom has already been copied
	if _rhombusgridTo, ok := mapOrigCopy[rhombusgridFrom]; ok {
		rhombusgridTo = _rhombusgridTo.(*RhombusGrid)
		return
	}

	rhombusgridTo = new(RhombusGrid)
	mapOrigCopy[rhombusgridFrom] = rhombusgridTo
	rhombusgridFrom.CopyBasicFields(rhombusgridTo)

	//insertion point for the staging of instances referenced by pointers
	if rhombusgridFrom.Reference != nil {
		rhombusgridTo.Reference = CopyBranchRhombus(mapOrigCopy, rhombusgridFrom.Reference)
	}
	if rhombusgridFrom.ShapeCategory != nil {
		rhombusgridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, rhombusgridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgridFrom.Rhombuses {
		rhombusgridTo.Rhombuses = append(rhombusgridTo.Rhombuses, CopyBranchRhombus(mapOrigCopy, _rhombus))
	}

	return
}

func CopyBranchShapeCategory(mapOrigCopy map[any]any, shapecategoryFrom *ShapeCategory) (shapecategoryTo *ShapeCategory) {

	// shapecategoryFrom has already been copied
	if _shapecategoryTo, ok := mapOrigCopy[shapecategoryFrom]; ok {
		shapecategoryTo = _shapecategoryTo.(*ShapeCategory)
		return
	}

	shapecategoryTo = new(ShapeCategory)
	mapOrigCopy[shapecategoryFrom] = shapecategoryTo
	shapecategoryFrom.CopyBasicFields(shapecategoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralBezier(mapOrigCopy map[any]any, spiralbezierFrom *SpiralBezier) (spiralbezierTo *SpiralBezier) {

	// spiralbezierFrom has already been copied
	if _spiralbezierTo, ok := mapOrigCopy[spiralbezierFrom]; ok {
		spiralbezierTo = _spiralbezierTo.(*SpiralBezier)
		return
	}

	spiralbezierTo = new(SpiralBezier)
	mapOrigCopy[spiralbezierFrom] = spiralbezierTo
	spiralbezierFrom.CopyBasicFields(spiralbezierTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralbezierFrom.ShapeCategory != nil {
		spiralbezierTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralbezierFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralBezierGrid(mapOrigCopy map[any]any, spiralbeziergridFrom *SpiralBezierGrid) (spiralbeziergridTo *SpiralBezierGrid) {

	// spiralbeziergridFrom has already been copied
	if _spiralbeziergridTo, ok := mapOrigCopy[spiralbeziergridFrom]; ok {
		spiralbeziergridTo = _spiralbeziergridTo.(*SpiralBezierGrid)
		return
	}

	spiralbeziergridTo = new(SpiralBezierGrid)
	mapOrigCopy[spiralbeziergridFrom] = spiralbeziergridTo
	spiralbeziergridFrom.CopyBasicFields(spiralbeziergridTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralbeziergridFrom.ShapeCategory != nil {
		spiralbeziergridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralbeziergridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralbezier := range spiralbeziergridFrom.SpiralBeziers {
		spiralbeziergridTo.SpiralBeziers = append(spiralbeziergridTo.SpiralBeziers, CopyBranchSpiralBezier(mapOrigCopy, _spiralbezier))
	}

	return
}

func CopyBranchSpiralCircle(mapOrigCopy map[any]any, spiralcircleFrom *SpiralCircle) (spiralcircleTo *SpiralCircle) {

	// spiralcircleFrom has already been copied
	if _spiralcircleTo, ok := mapOrigCopy[spiralcircleFrom]; ok {
		spiralcircleTo = _spiralcircleTo.(*SpiralCircle)
		return
	}

	spiralcircleTo = new(SpiralCircle)
	mapOrigCopy[spiralcircleFrom] = spiralcircleTo
	spiralcircleFrom.CopyBasicFields(spiralcircleTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralcircleFrom.ShapeCategory != nil {
		spiralcircleTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralcircleFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralCircleGrid(mapOrigCopy map[any]any, spiralcirclegridFrom *SpiralCircleGrid) (spiralcirclegridTo *SpiralCircleGrid) {

	// spiralcirclegridFrom has already been copied
	if _spiralcirclegridTo, ok := mapOrigCopy[spiralcirclegridFrom]; ok {
		spiralcirclegridTo = _spiralcirclegridTo.(*SpiralCircleGrid)
		return
	}

	spiralcirclegridTo = new(SpiralCircleGrid)
	mapOrigCopy[spiralcirclegridFrom] = spiralcirclegridTo
	spiralcirclegridFrom.CopyBasicFields(spiralcirclegridTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralcirclegridFrom.ShapeCategory != nil {
		spiralcirclegridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralcirclegridFrom.ShapeCategory)
	}
	if spiralcirclegridFrom.SpiralRhombusGrid != nil {
		spiralcirclegridTo.SpiralRhombusGrid = CopyBranchSpiralRhombusGrid(mapOrigCopy, spiralcirclegridFrom.SpiralRhombusGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralcircle := range spiralcirclegridFrom.SpiralCircles {
		spiralcirclegridTo.SpiralCircles = append(spiralcirclegridTo.SpiralCircles, CopyBranchSpiralCircle(mapOrigCopy, _spiralcircle))
	}

	return
}

func CopyBranchSpiralLine(mapOrigCopy map[any]any, spirallineFrom *SpiralLine) (spirallineTo *SpiralLine) {

	// spirallineFrom has already been copied
	if _spirallineTo, ok := mapOrigCopy[spirallineFrom]; ok {
		spirallineTo = _spirallineTo.(*SpiralLine)
		return
	}

	spirallineTo = new(SpiralLine)
	mapOrigCopy[spirallineFrom] = spirallineTo
	spirallineFrom.CopyBasicFields(spirallineTo)

	//insertion point for the staging of instances referenced by pointers
	if spirallineFrom.ShapeCategory != nil {
		spirallineTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spirallineFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralLineGrid(mapOrigCopy map[any]any, spirallinegridFrom *SpiralLineGrid) (spirallinegridTo *SpiralLineGrid) {

	// spirallinegridFrom has already been copied
	if _spirallinegridTo, ok := mapOrigCopy[spirallinegridFrom]; ok {
		spirallinegridTo = _spirallinegridTo.(*SpiralLineGrid)
		return
	}

	spirallinegridTo = new(SpiralLineGrid)
	mapOrigCopy[spirallinegridFrom] = spirallinegridTo
	spirallinegridFrom.CopyBasicFields(spirallinegridTo)

	//insertion point for the staging of instances referenced by pointers
	if spirallinegridFrom.ShapeCategory != nil {
		spirallinegridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spirallinegridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralline := range spirallinegridFrom.SpiralLines {
		spirallinegridTo.SpiralLines = append(spirallinegridTo.SpiralLines, CopyBranchSpiralLine(mapOrigCopy, _spiralline))
	}

	return
}

func CopyBranchSpiralOrigin(mapOrigCopy map[any]any, spiraloriginFrom *SpiralOrigin) (spiraloriginTo *SpiralOrigin) {

	// spiraloriginFrom has already been copied
	if _spiraloriginTo, ok := mapOrigCopy[spiraloriginFrom]; ok {
		spiraloriginTo = _spiraloriginTo.(*SpiralOrigin)
		return
	}

	spiraloriginTo = new(SpiralOrigin)
	mapOrigCopy[spiraloriginFrom] = spiraloriginTo
	spiraloriginFrom.CopyBasicFields(spiraloriginTo)

	//insertion point for the staging of instances referenced by pointers
	if spiraloriginFrom.ShapeCategory != nil {
		spiraloriginTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiraloriginFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralRhombus(mapOrigCopy map[any]any, spiralrhombusFrom *SpiralRhombus) (spiralrhombusTo *SpiralRhombus) {

	// spiralrhombusFrom has already been copied
	if _spiralrhombusTo, ok := mapOrigCopy[spiralrhombusFrom]; ok {
		spiralrhombusTo = _spiralrhombusTo.(*SpiralRhombus)
		return
	}

	spiralrhombusTo = new(SpiralRhombus)
	mapOrigCopy[spiralrhombusFrom] = spiralrhombusTo
	spiralrhombusFrom.CopyBasicFields(spiralrhombusTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombusFrom.ShapeCategory != nil {
		spiralrhombusTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralrhombusFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSpiralRhombusGrid(mapOrigCopy map[any]any, spiralrhombusgridFrom *SpiralRhombusGrid) (spiralrhombusgridTo *SpiralRhombusGrid) {

	// spiralrhombusgridFrom has already been copied
	if _spiralrhombusgridTo, ok := mapOrigCopy[spiralrhombusgridFrom]; ok {
		spiralrhombusgridTo = _spiralrhombusgridTo.(*SpiralRhombusGrid)
		return
	}

	spiralrhombusgridTo = new(SpiralRhombusGrid)
	mapOrigCopy[spiralrhombusgridFrom] = spiralrhombusgridTo
	spiralrhombusgridFrom.CopyBasicFields(spiralrhombusgridTo)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombusgridFrom.ShapeCategory != nil {
		spiralrhombusgridTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, spiralrhombusgridFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralrhombus := range spiralrhombusgridFrom.SpiralRhombuses {
		spiralrhombusgridTo.SpiralRhombuses = append(spiralrhombusgridTo.SpiralRhombuses, CopyBranchSpiralRhombus(mapOrigCopy, _spiralrhombus))
	}

	return
}

func CopyBranchVerticalAxis(mapOrigCopy map[any]any, verticalaxisFrom *VerticalAxis) (verticalaxisTo *VerticalAxis) {

	// verticalaxisFrom has already been copied
	if _verticalaxisTo, ok := mapOrigCopy[verticalaxisFrom]; ok {
		verticalaxisTo = _verticalaxisTo.(*VerticalAxis)
		return
	}

	verticalaxisTo = new(VerticalAxis)
	mapOrigCopy[verticalaxisFrom] = verticalaxisTo
	verticalaxisFrom.CopyBasicFields(verticalaxisTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalaxisFrom.ShapeCategory != nil {
		verticalaxisTo.ShapeCategory = CopyBranchShapeCategory(mapOrigCopy, verticalaxisFrom.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Axis:
		stage.UnstageBranchAxis(target)

	case *AxisGrid:
		stage.UnstageBranchAxisGrid(target)

	case *Bezier:
		stage.UnstageBranchBezier(target)

	case *BezierGrid:
		stage.UnstageBranchBezierGrid(target)

	case *BezierGridStack:
		stage.UnstageBranchBezierGridStack(target)

	case *Chapter:
		stage.UnstageBranchChapter(target)

	case *Circle:
		stage.UnstageBranchCircle(target)

	case *CircleGrid:
		stage.UnstageBranchCircleGrid(target)

	case *Content:
		stage.UnstageBranchContent(target)

	case *ExportToMusicxml:
		stage.UnstageBranchExportToMusicxml(target)

	case *FrontCurve:
		stage.UnstageBranchFrontCurve(target)

	case *FrontCurveStack:
		stage.UnstageBranchFrontCurveStack(target)

	case *HorizontalAxis:
		stage.UnstageBranchHorizontalAxis(target)

	case *Key:
		stage.UnstageBranchKey(target)

	case *Parameter:
		stage.UnstageBranchParameter(target)

	case *Rhombus:
		stage.UnstageBranchRhombus(target)

	case *RhombusGrid:
		stage.UnstageBranchRhombusGrid(target)

	case *ShapeCategory:
		stage.UnstageBranchShapeCategory(target)

	case *SpiralBezier:
		stage.UnstageBranchSpiralBezier(target)

	case *SpiralBezierGrid:
		stage.UnstageBranchSpiralBezierGrid(target)

	case *SpiralCircle:
		stage.UnstageBranchSpiralCircle(target)

	case *SpiralCircleGrid:
		stage.UnstageBranchSpiralCircleGrid(target)

	case *SpiralLine:
		stage.UnstageBranchSpiralLine(target)

	case *SpiralLineGrid:
		stage.UnstageBranchSpiralLineGrid(target)

	case *SpiralOrigin:
		stage.UnstageBranchSpiralOrigin(target)

	case *SpiralRhombus:
		stage.UnstageBranchSpiralRhombus(target)

	case *SpiralRhombusGrid:
		stage.UnstageBranchSpiralRhombusGrid(target)

	case *VerticalAxis:
		stage.UnstageBranchVerticalAxis(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAxis(axis *Axis) {

	// check if instance is already staged
	if !IsStaged(stage, axis) {
		return
	}

	axis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if axis.ShapeCategory != nil {
		UnstageBranch(stage, axis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchAxisGrid(axisgrid *AxisGrid) {

	// check if instance is already staged
	if !IsStaged(stage, axisgrid) {
		return
	}

	axisgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if axisgrid.Reference != nil {
		UnstageBranch(stage, axisgrid.Reference)
	}
	if axisgrid.ShapeCategory != nil {
		UnstageBranch(stage, axisgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _axis := range axisgrid.Axiss {
		UnstageBranch(stage, _axis)
	}

}

func (stage *Stage) UnstageBranchBezier(bezier *Bezier) {

	// check if instance is already staged
	if !IsStaged(stage, bezier) {
		return
	}

	bezier.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bezier.ShapeCategory != nil {
		UnstageBranch(stage, bezier.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBezierGrid(beziergrid *BezierGrid) {

	// check if instance is already staged
	if !IsStaged(stage, beziergrid) {
		return
	}

	beziergrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziergrid.Reference != nil {
		UnstageBranch(stage, beziergrid.Reference)
	}
	if beziergrid.ShapeCategory != nil {
		UnstageBranch(stage, beziergrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bezier := range beziergrid.Beziers {
		UnstageBranch(stage, _bezier)
	}

}

func (stage *Stage) UnstageBranchBezierGridStack(beziergridstack *BezierGridStack) {

	// check if instance is already staged
	if !IsStaged(stage, beziergridstack) {
		return
	}

	beziergridstack.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziergridstack.ShapeCategory != nil {
		UnstageBranch(stage, beziergridstack.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziergrid := range beziergridstack.BezierGrids {
		UnstageBranch(stage, _beziergrid)
	}

}

func (stage *Stage) UnstageBranchChapter(chapter *Chapter) {

	// check if instance is already staged
	if !IsStaged(stage, chapter) {
		return
	}

	chapter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if !IsStaged(stage, circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circle.ShapeCategory != nil {
		UnstageBranch(stage, circle.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCircleGrid(circlegrid *CircleGrid) {

	// check if instance is already staged
	if !IsStaged(stage, circlegrid) {
		return
	}

	circlegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circlegrid.Reference != nil {
		UnstageBranch(stage, circlegrid.Reference)
	}
	if circlegrid.ShapeCategory != nil {
		UnstageBranch(stage, circlegrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _circle := range circlegrid.Circles {
		UnstageBranch(stage, _circle)
	}

}

func (stage *Stage) UnstageBranchContent(content *Content) {

	// check if instance is already staged
	if !IsStaged(stage, content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		UnstageBranch(stage, _chapter)
	}

}

func (stage *Stage) UnstageBranchExportToMusicxml(exporttomusicxml *ExportToMusicxml) {

	// check if instance is already staged
	if !IsStaged(stage, exporttomusicxml) {
		return
	}

	exporttomusicxml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if exporttomusicxml.Parameter != nil {
		UnstageBranch(stage, exporttomusicxml.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFrontCurve(frontcurve *FrontCurve) {

	// check if instance is already staged
	if !IsStaged(stage, frontcurve) {
		return
	}

	frontcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFrontCurveStack(frontcurvestack *FrontCurveStack) {

	// check if instance is already staged
	if !IsStaged(stage, frontcurvestack) {
		return
	}

	frontcurvestack.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frontcurvestack.ShapeCategory != nil {
		UnstageBranch(stage, frontcurvestack.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frontcurve := range frontcurvestack.FrontCurves {
		UnstageBranch(stage, _frontcurve)
	}
	for _, _spiralcircle := range frontcurvestack.SpiralCircles {
		UnstageBranch(stage, _spiralcircle)
	}

}

func (stage *Stage) UnstageBranchHorizontalAxis(horizontalaxis *HorizontalAxis) {

	// check if instance is already staged
	if !IsStaged(stage, horizontalaxis) {
		return
	}

	horizontalaxis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if horizontalaxis.ShapeCategory != nil {
		UnstageBranch(stage, horizontalaxis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchKey(key *Key) {

	// check if instance is already staged
	if !IsStaged(stage, key) {
		return
	}

	key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.ShapeCategory != nil {
		UnstageBranch(stage, key.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParameter(parameter *Parameter) {

	// check if instance is already staged
	if !IsStaged(stage, parameter) {
		return
	}

	parameter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parameter.InitialRhombus != nil {
		UnstageBranch(stage, parameter.InitialRhombus)
	}
	if parameter.InitialCircle != nil {
		UnstageBranch(stage, parameter.InitialCircle)
	}
	if parameter.InitialRhombusGrid != nil {
		UnstageBranch(stage, parameter.InitialRhombusGrid)
	}
	if parameter.InitialCircleGrid != nil {
		UnstageBranch(stage, parameter.InitialCircleGrid)
	}
	if parameter.InitialAxis != nil {
		UnstageBranch(stage, parameter.InitialAxis)
	}
	if parameter.RotatedAxis != nil {
		UnstageBranch(stage, parameter.RotatedAxis)
	}
	if parameter.RotatedRhombus != nil {
		UnstageBranch(stage, parameter.RotatedRhombus)
	}
	if parameter.RotatedRhombusGrid != nil {
		UnstageBranch(stage, parameter.RotatedRhombusGrid)
	}
	if parameter.RotatedCircleGrid != nil {
		UnstageBranch(stage, parameter.RotatedCircleGrid)
	}
	if parameter.NextRhombus != nil {
		UnstageBranch(stage, parameter.NextRhombus)
	}
	if parameter.NextCircle != nil {
		UnstageBranch(stage, parameter.NextCircle)
	}
	if parameter.GrowingRhombusGridSeed != nil {
		UnstageBranch(stage, parameter.GrowingRhombusGridSeed)
	}
	if parameter.GrowingRhombusGrid != nil {
		UnstageBranch(stage, parameter.GrowingRhombusGrid)
	}
	if parameter.GrowingCircleGridSeed != nil {
		UnstageBranch(stage, parameter.GrowingCircleGridSeed)
	}
	if parameter.GrowingCircleGrid != nil {
		UnstageBranch(stage, parameter.GrowingCircleGrid)
	}
	if parameter.GrowingCircleGridLeftSeed != nil {
		UnstageBranch(stage, parameter.GrowingCircleGridLeftSeed)
	}
	if parameter.GrowingCircleGridLeft != nil {
		UnstageBranch(stage, parameter.GrowingCircleGridLeft)
	}
	if parameter.ConstructionAxis != nil {
		UnstageBranch(stage, parameter.ConstructionAxis)
	}
	if parameter.ConstructionAxisGrid != nil {
		UnstageBranch(stage, parameter.ConstructionAxisGrid)
	}
	if parameter.ConstructionCircle != nil {
		UnstageBranch(stage, parameter.ConstructionCircle)
	}
	if parameter.ConstructionCircleGrid != nil {
		UnstageBranch(stage, parameter.ConstructionCircleGrid)
	}
	if parameter.GrowthCurveSeed != nil {
		UnstageBranch(stage, parameter.GrowthCurveSeed)
	}
	if parameter.GrowthCurve != nil {
		UnstageBranch(stage, parameter.GrowthCurve)
	}
	if parameter.GrowthCurveShiftedRightSeed != nil {
		UnstageBranch(stage, parameter.GrowthCurveShiftedRightSeed)
	}
	if parameter.GrowthCurveShiftedRight != nil {
		UnstageBranch(stage, parameter.GrowthCurveShiftedRight)
	}
	if parameter.GrowthCurveNextSeed != nil {
		UnstageBranch(stage, parameter.GrowthCurveNextSeed)
	}
	if parameter.GrowthCurveNext != nil {
		UnstageBranch(stage, parameter.GrowthCurveNext)
	}
	if parameter.GrowthCurveNextShiftedRightSeed != nil {
		UnstageBranch(stage, parameter.GrowthCurveNextShiftedRightSeed)
	}
	if parameter.GrowthCurveNextShiftedRight != nil {
		UnstageBranch(stage, parameter.GrowthCurveNextShiftedRight)
	}
	if parameter.GrowthCurveStack != nil {
		UnstageBranch(stage, parameter.GrowthCurveStack)
	}
	if parameter.SpiralRhombusGridSeed != nil {
		UnstageBranch(stage, parameter.SpiralRhombusGridSeed)
	}
	if parameter.SpiralRhombusGrid != nil {
		UnstageBranch(stage, parameter.SpiralRhombusGrid)
	}
	if parameter.SpiralCircleSeed != nil {
		UnstageBranch(stage, parameter.SpiralCircleSeed)
	}
	if parameter.SpiralCircleGrid != nil {
		UnstageBranch(stage, parameter.SpiralCircleGrid)
	}
	if parameter.SpiralCircleFullGrid != nil {
		UnstageBranch(stage, parameter.SpiralCircleFullGrid)
	}
	if parameter.SpiralConstructionOuterLineSeed != nil {
		UnstageBranch(stage, parameter.SpiralConstructionOuterLineSeed)
	}
	if parameter.SpiralConstructionInnerLineSeed != nil {
		UnstageBranch(stage, parameter.SpiralConstructionInnerLineSeed)
	}
	if parameter.SpiralConstructionOuterLineGrid != nil {
		UnstageBranch(stage, parameter.SpiralConstructionOuterLineGrid)
	}
	if parameter.SpiralConstructionInnerLineGrid != nil {
		UnstageBranch(stage, parameter.SpiralConstructionInnerLineGrid)
	}
	if parameter.SpiralConstructionCircleGrid != nil {
		UnstageBranch(stage, parameter.SpiralConstructionCircleGrid)
	}
	if parameter.SpiralConstructionOuterLineFullGrid != nil {
		UnstageBranch(stage, parameter.SpiralConstructionOuterLineFullGrid)
	}
	if parameter.SpiralBezierSeed != nil {
		UnstageBranch(stage, parameter.SpiralBezierSeed)
	}
	if parameter.SpiralBezierGrid != nil {
		UnstageBranch(stage, parameter.SpiralBezierGrid)
	}
	if parameter.SpiralBezierFullGrid != nil {
		UnstageBranch(stage, parameter.SpiralBezierFullGrid)
	}
	if parameter.FrontCurveStack != nil {
		UnstageBranch(stage, parameter.FrontCurveStack)
	}
	if parameter.Fkey != nil {
		UnstageBranch(stage, parameter.Fkey)
	}
	if parameter.PitchLines != nil {
		UnstageBranch(stage, parameter.PitchLines)
	}
	if parameter.BeatLines != nil {
		UnstageBranch(stage, parameter.BeatLines)
	}
	if parameter.FirstVoice != nil {
		UnstageBranch(stage, parameter.FirstVoice)
	}
	if parameter.FirstVoiceShiftedRigth != nil {
		UnstageBranch(stage, parameter.FirstVoiceShiftedRigth)
	}
	if parameter.SecondVoice != nil {
		UnstageBranch(stage, parameter.SecondVoice)
	}
	if parameter.SecondVoiceShiftedRight != nil {
		UnstageBranch(stage, parameter.SecondVoiceShiftedRight)
	}
	if parameter.FirstVoiceNotes != nil {
		UnstageBranch(stage, parameter.FirstVoiceNotes)
	}
	if parameter.FirstVoiceNotesShiftedRight != nil {
		UnstageBranch(stage, parameter.FirstVoiceNotesShiftedRight)
	}
	if parameter.SecondVoiceNotes != nil {
		UnstageBranch(stage, parameter.SecondVoiceNotes)
	}
	if parameter.SecondVoiceNotesShiftedRight != nil {
		UnstageBranch(stage, parameter.SecondVoiceNotesShiftedRight)
	}
	if parameter.HorizontalAxis != nil {
		UnstageBranch(stage, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis != nil {
		UnstageBranch(stage, parameter.VerticalAxis)
	}
	if parameter.SpiralOrigin != nil {
		UnstageBranch(stage, parameter.SpiralOrigin)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRhombus(rhombus *Rhombus) {

	// check if instance is already staged
	if !IsStaged(stage, rhombus) {
		return
	}

	rhombus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rhombus.ShapeCategory != nil {
		UnstageBranch(stage, rhombus.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRhombusGrid(rhombusgrid *RhombusGrid) {

	// check if instance is already staged
	if !IsStaged(stage, rhombusgrid) {
		return
	}

	rhombusgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rhombusgrid.Reference != nil {
		UnstageBranch(stage, rhombusgrid.Reference)
	}
	if rhombusgrid.ShapeCategory != nil {
		UnstageBranch(stage, rhombusgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgrid.Rhombuses {
		UnstageBranch(stage, _rhombus)
	}

}

func (stage *Stage) UnstageBranchShapeCategory(shapecategory *ShapeCategory) {

	// check if instance is already staged
	if !IsStaged(stage, shapecategory) {
		return
	}

	shapecategory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralBezier(spiralbezier *SpiralBezier) {

	// check if instance is already staged
	if !IsStaged(stage, spiralbezier) {
		return
	}

	spiralbezier.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralbezier.ShapeCategory != nil {
		UnstageBranch(stage, spiralbezier.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralBezierGrid(spiralbeziergrid *SpiralBezierGrid) {

	// check if instance is already staged
	if !IsStaged(stage, spiralbeziergrid) {
		return
	}

	spiralbeziergrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralbeziergrid.ShapeCategory != nil {
		UnstageBranch(stage, spiralbeziergrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
		UnstageBranch(stage, _spiralbezier)
	}

}

func (stage *Stage) UnstageBranchSpiralCircle(spiralcircle *SpiralCircle) {

	// check if instance is already staged
	if !IsStaged(stage, spiralcircle) {
		return
	}

	spiralcircle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralcircle.ShapeCategory != nil {
		UnstageBranch(stage, spiralcircle.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralCircleGrid(spiralcirclegrid *SpiralCircleGrid) {

	// check if instance is already staged
	if !IsStaged(stage, spiralcirclegrid) {
		return
	}

	spiralcirclegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralcirclegrid.ShapeCategory != nil {
		UnstageBranch(stage, spiralcirclegrid.ShapeCategory)
	}
	if spiralcirclegrid.SpiralRhombusGrid != nil {
		UnstageBranch(stage, spiralcirclegrid.SpiralRhombusGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
		UnstageBranch(stage, _spiralcircle)
	}

}

func (stage *Stage) UnstageBranchSpiralLine(spiralline *SpiralLine) {

	// check if instance is already staged
	if !IsStaged(stage, spiralline) {
		return
	}

	spiralline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralline.ShapeCategory != nil {
		UnstageBranch(stage, spiralline.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralLineGrid(spirallinegrid *SpiralLineGrid) {

	// check if instance is already staged
	if !IsStaged(stage, spirallinegrid) {
		return
	}

	spirallinegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spirallinegrid.ShapeCategory != nil {
		UnstageBranch(stage, spirallinegrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralline := range spirallinegrid.SpiralLines {
		UnstageBranch(stage, _spiralline)
	}

}

func (stage *Stage) UnstageBranchSpiralOrigin(spiralorigin *SpiralOrigin) {

	// check if instance is already staged
	if !IsStaged(stage, spiralorigin) {
		return
	}

	spiralorigin.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralorigin.ShapeCategory != nil {
		UnstageBranch(stage, spiralorigin.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralRhombus(spiralrhombus *SpiralRhombus) {

	// check if instance is already staged
	if !IsStaged(stage, spiralrhombus) {
		return
	}

	spiralrhombus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombus.ShapeCategory != nil {
		UnstageBranch(stage, spiralrhombus.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSpiralRhombusGrid(spiralrhombusgrid *SpiralRhombusGrid) {

	// check if instance is already staged
	if !IsStaged(stage, spiralrhombusgrid) {
		return
	}

	spiralrhombusgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spiralrhombusgrid.ShapeCategory != nil {
		UnstageBranch(stage, spiralrhombusgrid.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
		UnstageBranch(stage, _spiralrhombus)
	}

}

func (stage *Stage) UnstageBranchVerticalAxis(verticalaxis *VerticalAxis) {

	// check if instance is already staged
	if !IsStaged(stage, verticalaxis) {
		return
	}

	verticalaxis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalaxis.ShapeCategory != nil {
		UnstageBranch(stage, verticalaxis.ShapeCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (axis *Axis) GongDiff(stage *Stage, axisOther *Axis) (diffs []string) {
	// insertion point for field diffs
	if axis.Name != axisOther.Name {
		diffs = append(diffs, axis.GongMarshallField(stage, "Name"))
	}
	if axis.IsDisplayed != axisOther.IsDisplayed {
		diffs = append(diffs, axis.GongMarshallField(stage, "IsDisplayed"))
	}
	if (axis.ShapeCategory == nil) != (axisOther.ShapeCategory == nil) {
		diffs = append(diffs, axis.GongMarshallField(stage, "ShapeCategory"))
	} else if axis.ShapeCategory != nil && axisOther.ShapeCategory != nil {
		if axis.ShapeCategory != axisOther.ShapeCategory {
			diffs = append(diffs, axis.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if axis.AngleDegree != axisOther.AngleDegree {
		diffs = append(diffs, axis.GongMarshallField(stage, "AngleDegree"))
	}
	if axis.Length != axisOther.Length {
		diffs = append(diffs, axis.GongMarshallField(stage, "Length"))
	}
	if axis.CenterX != axisOther.CenterX {
		diffs = append(diffs, axis.GongMarshallField(stage, "CenterX"))
	}
	if axis.CenterY != axisOther.CenterY {
		diffs = append(diffs, axis.GongMarshallField(stage, "CenterY"))
	}
	if axis.EndX != axisOther.EndX {
		diffs = append(diffs, axis.GongMarshallField(stage, "EndX"))
	}
	if axis.EndY != axisOther.EndY {
		diffs = append(diffs, axis.GongMarshallField(stage, "EndY"))
	}
	if axis.Color != axisOther.Color {
		diffs = append(diffs, axis.GongMarshallField(stage, "Color"))
	}
	if axis.FillOpacity != axisOther.FillOpacity {
		diffs = append(diffs, axis.GongMarshallField(stage, "FillOpacity"))
	}
	if axis.Stroke != axisOther.Stroke {
		diffs = append(diffs, axis.GongMarshallField(stage, "Stroke"))
	}
	if axis.StrokeOpacity != axisOther.StrokeOpacity {
		diffs = append(diffs, axis.GongMarshallField(stage, "StrokeOpacity"))
	}
	if axis.StrokeWidth != axisOther.StrokeWidth {
		diffs = append(diffs, axis.GongMarshallField(stage, "StrokeWidth"))
	}
	if axis.StrokeDashArray != axisOther.StrokeDashArray {
		diffs = append(diffs, axis.GongMarshallField(stage, "StrokeDashArray"))
	}
	if axis.StrokeDashArrayWhenSelected != axisOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, axis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if axis.Transform != axisOther.Transform {
		diffs = append(diffs, axis.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (axisgrid *AxisGrid) GongDiff(stage *Stage, axisgridOther *AxisGrid) (diffs []string) {
	// insertion point for field diffs
	if axisgrid.Name != axisgridOther.Name {
		diffs = append(diffs, axisgrid.GongMarshallField(stage, "Name"))
	}
	if (axisgrid.Reference == nil) != (axisgridOther.Reference == nil) {
		diffs = append(diffs, axisgrid.GongMarshallField(stage, "Reference"))
	} else if axisgrid.Reference != nil && axisgridOther.Reference != nil {
		if axisgrid.Reference != axisgridOther.Reference {
			diffs = append(diffs, axisgrid.GongMarshallField(stage, "Reference"))
		}
	}
	if axisgrid.IsDisplayed != axisgridOther.IsDisplayed {
		diffs = append(diffs, axisgrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (axisgrid.ShapeCategory == nil) != (axisgridOther.ShapeCategory == nil) {
		diffs = append(diffs, axisgrid.GongMarshallField(stage, "ShapeCategory"))
	} else if axisgrid.ShapeCategory != nil && axisgridOther.ShapeCategory != nil {
		if axisgrid.ShapeCategory != axisgridOther.ShapeCategory {
			diffs = append(diffs, axisgrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	AxissDifferent := false
	if len(axisgrid.Axiss) != len(axisgridOther.Axiss) {
		AxissDifferent = true
	} else {
		for i := range axisgrid.Axiss {
			if (axisgrid.Axiss[i] == nil) != (axisgridOther.Axiss[i] == nil) {
				AxissDifferent = true
				break
			} else if axisgrid.Axiss[i] != nil && axisgridOther.Axiss[i] != nil {
				// this is a pointer comparaison
				if axisgrid.Axiss[i] != axisgridOther.Axiss[i] {
					AxissDifferent = true
					break
				}
			}
		}
	}
	if AxissDifferent {
		ops := Diff(stage, axisgrid, axisgridOther, "Axiss", axisgridOther.Axiss, axisgrid.Axiss)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bezier *Bezier) GongDiff(stage *Stage, bezierOther *Bezier) (diffs []string) {
	// insertion point for field diffs
	if bezier.Name != bezierOther.Name {
		diffs = append(diffs, bezier.GongMarshallField(stage, "Name"))
	}
	if bezier.IsDisplayed != bezierOther.IsDisplayed {
		diffs = append(diffs, bezier.GongMarshallField(stage, "IsDisplayed"))
	}
	if (bezier.ShapeCategory == nil) != (bezierOther.ShapeCategory == nil) {
		diffs = append(diffs, bezier.GongMarshallField(stage, "ShapeCategory"))
	} else if bezier.ShapeCategory != nil && bezierOther.ShapeCategory != nil {
		if bezier.ShapeCategory != bezierOther.ShapeCategory {
			diffs = append(diffs, bezier.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if bezier.StartX != bezierOther.StartX {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StartX"))
	}
	if bezier.StartY != bezierOther.StartY {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StartY"))
	}
	if bezier.ControlPointStartX != bezierOther.ControlPointStartX {
		diffs = append(diffs, bezier.GongMarshallField(stage, "ControlPointStartX"))
	}
	if bezier.ControlPointStartY != bezierOther.ControlPointStartY {
		diffs = append(diffs, bezier.GongMarshallField(stage, "ControlPointStartY"))
	}
	if bezier.EndX != bezierOther.EndX {
		diffs = append(diffs, bezier.GongMarshallField(stage, "EndX"))
	}
	if bezier.EndY != bezierOther.EndY {
		diffs = append(diffs, bezier.GongMarshallField(stage, "EndY"))
	}
	if bezier.ControlPointEndX != bezierOther.ControlPointEndX {
		diffs = append(diffs, bezier.GongMarshallField(stage, "ControlPointEndX"))
	}
	if bezier.ControlPointEndY != bezierOther.ControlPointEndY {
		diffs = append(diffs, bezier.GongMarshallField(stage, "ControlPointEndY"))
	}
	if bezier.Color != bezierOther.Color {
		diffs = append(diffs, bezier.GongMarshallField(stage, "Color"))
	}
	if bezier.FillOpacity != bezierOther.FillOpacity {
		diffs = append(diffs, bezier.GongMarshallField(stage, "FillOpacity"))
	}
	if bezier.Stroke != bezierOther.Stroke {
		diffs = append(diffs, bezier.GongMarshallField(stage, "Stroke"))
	}
	if bezier.StrokeOpacity != bezierOther.StrokeOpacity {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StrokeOpacity"))
	}
	if bezier.StrokeWidth != bezierOther.StrokeWidth {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StrokeWidth"))
	}
	if bezier.StrokeDashArray != bezierOther.StrokeDashArray {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StrokeDashArray"))
	}
	if bezier.StrokeDashArrayWhenSelected != bezierOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, bezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if bezier.Transform != bezierOther.Transform {
		diffs = append(diffs, bezier.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beziergrid *BezierGrid) GongDiff(stage *Stage, beziergridOther *BezierGrid) (diffs []string) {
	// insertion point for field diffs
	if beziergrid.Name != beziergridOther.Name {
		diffs = append(diffs, beziergrid.GongMarshallField(stage, "Name"))
	}
	if (beziergrid.Reference == nil) != (beziergridOther.Reference == nil) {
		diffs = append(diffs, beziergrid.GongMarshallField(stage, "Reference"))
	} else if beziergrid.Reference != nil && beziergridOther.Reference != nil {
		if beziergrid.Reference != beziergridOther.Reference {
			diffs = append(diffs, beziergrid.GongMarshallField(stage, "Reference"))
		}
	}
	if beziergrid.IsDisplayed != beziergridOther.IsDisplayed {
		diffs = append(diffs, beziergrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (beziergrid.ShapeCategory == nil) != (beziergridOther.ShapeCategory == nil) {
		diffs = append(diffs, beziergrid.GongMarshallField(stage, "ShapeCategory"))
	} else if beziergrid.ShapeCategory != nil && beziergridOther.ShapeCategory != nil {
		if beziergrid.ShapeCategory != beziergridOther.ShapeCategory {
			diffs = append(diffs, beziergrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	BeziersDifferent := false
	if len(beziergrid.Beziers) != len(beziergridOther.Beziers) {
		BeziersDifferent = true
	} else {
		for i := range beziergrid.Beziers {
			if (beziergrid.Beziers[i] == nil) != (beziergridOther.Beziers[i] == nil) {
				BeziersDifferent = true
				break
			} else if beziergrid.Beziers[i] != nil && beziergridOther.Beziers[i] != nil {
				// this is a pointer comparaison
				if beziergrid.Beziers[i] != beziergridOther.Beziers[i] {
					BeziersDifferent = true
					break
				}
			}
		}
	}
	if BeziersDifferent {
		ops := Diff(stage, beziergrid, beziergridOther, "Beziers", beziergridOther.Beziers, beziergrid.Beziers)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beziergridstack *BezierGridStack) GongDiff(stage *Stage, beziergridstackOther *BezierGridStack) (diffs []string) {
	// insertion point for field diffs
	if beziergridstack.Name != beziergridstackOther.Name {
		diffs = append(diffs, beziergridstack.GongMarshallField(stage, "Name"))
	}
	if beziergridstack.IsDisplayed != beziergridstackOther.IsDisplayed {
		diffs = append(diffs, beziergridstack.GongMarshallField(stage, "IsDisplayed"))
	}
	if (beziergridstack.ShapeCategory == nil) != (beziergridstackOther.ShapeCategory == nil) {
		diffs = append(diffs, beziergridstack.GongMarshallField(stage, "ShapeCategory"))
	} else if beziergridstack.ShapeCategory != nil && beziergridstackOther.ShapeCategory != nil {
		if beziergridstack.ShapeCategory != beziergridstackOther.ShapeCategory {
			diffs = append(diffs, beziergridstack.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	BezierGridsDifferent := false
	if len(beziergridstack.BezierGrids) != len(beziergridstackOther.BezierGrids) {
		BezierGridsDifferent = true
	} else {
		for i := range beziergridstack.BezierGrids {
			if (beziergridstack.BezierGrids[i] == nil) != (beziergridstackOther.BezierGrids[i] == nil) {
				BezierGridsDifferent = true
				break
			} else if beziergridstack.BezierGrids[i] != nil && beziergridstackOther.BezierGrids[i] != nil {
				// this is a pointer comparaison
				if beziergridstack.BezierGrids[i] != beziergridstackOther.BezierGrids[i] {
					BezierGridsDifferent = true
					break
				}
			}
		}
	}
	if BezierGridsDifferent {
		ops := Diff(stage, beziergridstack, beziergridstackOther, "BezierGrids", beziergridstackOther.BezierGrids, beziergridstack.BezierGrids)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (chapter *Chapter) GongDiff(stage *Stage, chapterOther *Chapter) (diffs []string) {
	// insertion point for field diffs
	if chapter.Name != chapterOther.Name {
		diffs = append(diffs, chapter.GongMarshallField(stage, "Name"))
	}
	if chapter.MardownContent != chapterOther.MardownContent {
		diffs = append(diffs, chapter.GongMarshallField(stage, "MardownContent"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circle *Circle) GongDiff(stage *Stage, circleOther *Circle) (diffs []string) {
	// insertion point for field diffs
	if circle.Name != circleOther.Name {
		diffs = append(diffs, circle.GongMarshallField(stage, "Name"))
	}
	if circle.IsDisplayed != circleOther.IsDisplayed {
		diffs = append(diffs, circle.GongMarshallField(stage, "IsDisplayed"))
	}
	if (circle.ShapeCategory == nil) != (circleOther.ShapeCategory == nil) {
		diffs = append(diffs, circle.GongMarshallField(stage, "ShapeCategory"))
	} else if circle.ShapeCategory != nil && circleOther.ShapeCategory != nil {
		if circle.ShapeCategory != circleOther.ShapeCategory {
			diffs = append(diffs, circle.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if circle.CenterX != circleOther.CenterX {
		diffs = append(diffs, circle.GongMarshallField(stage, "CenterX"))
	}
	if circle.CenterY != circleOther.CenterY {
		diffs = append(diffs, circle.GongMarshallField(stage, "CenterY"))
	}
	if circle.HasBespokeRadius != circleOther.HasBespokeRadius {
		diffs = append(diffs, circle.GongMarshallField(stage, "HasBespokeRadius"))
	}
	if circle.BespopkeRadius != circleOther.BespopkeRadius {
		diffs = append(diffs, circle.GongMarshallField(stage, "BespopkeRadius"))
	}
	if circle.Color != circleOther.Color {
		diffs = append(diffs, circle.GongMarshallField(stage, "Color"))
	}
	if circle.FillOpacity != circleOther.FillOpacity {
		diffs = append(diffs, circle.GongMarshallField(stage, "FillOpacity"))
	}
	if circle.Stroke != circleOther.Stroke {
		diffs = append(diffs, circle.GongMarshallField(stage, "Stroke"))
	}
	if circle.StrokeOpacity != circleOther.StrokeOpacity {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeOpacity"))
	}
	if circle.StrokeWidth != circleOther.StrokeWidth {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeWidth"))
	}
	if circle.StrokeDashArray != circleOther.StrokeDashArray {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeDashArray"))
	}
	if circle.StrokeDashArrayWhenSelected != circleOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if circle.Transform != circleOther.Transform {
		diffs = append(diffs, circle.GongMarshallField(stage, "Transform"))
	}
	if circle.Pitch != circleOther.Pitch {
		diffs = append(diffs, circle.GongMarshallField(stage, "Pitch"))
	}
	if circle.ShowName != circleOther.ShowName {
		diffs = append(diffs, circle.GongMarshallField(stage, "ShowName"))
	}
	if circle.BeatNb != circleOther.BeatNb {
		diffs = append(diffs, circle.GongMarshallField(stage, "BeatNb"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circlegrid *CircleGrid) GongDiff(stage *Stage, circlegridOther *CircleGrid) (diffs []string) {
	// insertion point for field diffs
	if circlegrid.Name != circlegridOther.Name {
		diffs = append(diffs, circlegrid.GongMarshallField(stage, "Name"))
	}
	if (circlegrid.Reference == nil) != (circlegridOther.Reference == nil) {
		diffs = append(diffs, circlegrid.GongMarshallField(stage, "Reference"))
	} else if circlegrid.Reference != nil && circlegridOther.Reference != nil {
		if circlegrid.Reference != circlegridOther.Reference {
			diffs = append(diffs, circlegrid.GongMarshallField(stage, "Reference"))
		}
	}
	if circlegrid.IsDisplayed != circlegridOther.IsDisplayed {
		diffs = append(diffs, circlegrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (circlegrid.ShapeCategory == nil) != (circlegridOther.ShapeCategory == nil) {
		diffs = append(diffs, circlegrid.GongMarshallField(stage, "ShapeCategory"))
	} else if circlegrid.ShapeCategory != nil && circlegridOther.ShapeCategory != nil {
		if circlegrid.ShapeCategory != circlegridOther.ShapeCategory {
			diffs = append(diffs, circlegrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	CirclesDifferent := false
	if len(circlegrid.Circles) != len(circlegridOther.Circles) {
		CirclesDifferent = true
	} else {
		for i := range circlegrid.Circles {
			if (circlegrid.Circles[i] == nil) != (circlegridOther.Circles[i] == nil) {
				CirclesDifferent = true
				break
			} else if circlegrid.Circles[i] != nil && circlegridOther.Circles[i] != nil {
				// this is a pointer comparaison
				if circlegrid.Circles[i] != circlegridOther.Circles[i] {
					CirclesDifferent = true
					break
				}
			}
		}
	}
	if CirclesDifferent {
		ops := Diff(stage, circlegrid, circlegridOther, "Circles", circlegridOther.Circles, circlegrid.Circles)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (content *Content) GongDiff(stage *Stage, contentOther *Content) (diffs []string) {
	// insertion point for field diffs
	if content.Name != contentOther.Name {
		diffs = append(diffs, content.GongMarshallField(stage, "Name"))
	}
	if content.MardownContent != contentOther.MardownContent {
		diffs = append(diffs, content.GongMarshallField(stage, "MardownContent"))
	}
	if content.ContentPath != contentOther.ContentPath {
		diffs = append(diffs, content.GongMarshallField(stage, "ContentPath"))
	}
	if content.OutputPath != contentOther.OutputPath {
		diffs = append(diffs, content.GongMarshallField(stage, "OutputPath"))
	}
	if content.LayoutPath != contentOther.LayoutPath {
		diffs = append(diffs, content.GongMarshallField(stage, "LayoutPath"))
	}
	if content.StaticPath != contentOther.StaticPath {
		diffs = append(diffs, content.GongMarshallField(stage, "StaticPath"))
	}
	if content.Target != contentOther.Target {
		diffs = append(diffs, content.GongMarshallField(stage, "Target"))
	}
	ChaptersDifferent := false
	if len(content.Chapters) != len(contentOther.Chapters) {
		ChaptersDifferent = true
	} else {
		for i := range content.Chapters {
			if (content.Chapters[i] == nil) != (contentOther.Chapters[i] == nil) {
				ChaptersDifferent = true
				break
			} else if content.Chapters[i] != nil && contentOther.Chapters[i] != nil {
				// this is a pointer comparaison
				if content.Chapters[i] != contentOther.Chapters[i] {
					ChaptersDifferent = true
					break
				}
			}
		}
	}
	if ChaptersDifferent {
		ops := Diff(stage, content, contentOther, "Chapters", contentOther.Chapters, content.Chapters)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (exporttomusicxml *ExportToMusicxml) GongDiff(stage *Stage, exporttomusicxmlOther *ExportToMusicxml) (diffs []string) {
	// insertion point for field diffs
	if exporttomusicxml.Name != exporttomusicxmlOther.Name {
		diffs = append(diffs, exporttomusicxml.GongMarshallField(stage, "Name"))
	}
	if (exporttomusicxml.Parameter == nil) != (exporttomusicxmlOther.Parameter == nil) {
		diffs = append(diffs, exporttomusicxml.GongMarshallField(stage, "Parameter"))
	} else if exporttomusicxml.Parameter != nil && exporttomusicxmlOther.Parameter != nil {
		if exporttomusicxml.Parameter != exporttomusicxmlOther.Parameter {
			diffs = append(diffs, exporttomusicxml.GongMarshallField(stage, "Parameter"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (frontcurve *FrontCurve) GongDiff(stage *Stage, frontcurveOther *FrontCurve) (diffs []string) {
	// insertion point for field diffs
	if frontcurve.Name != frontcurveOther.Name {
		diffs = append(diffs, frontcurve.GongMarshallField(stage, "Name"))
	}
	if frontcurve.Path != frontcurveOther.Path {
		diffs = append(diffs, frontcurve.GongMarshallField(stage, "Path"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (frontcurvestack *FrontCurveStack) GongDiff(stage *Stage, frontcurvestackOther *FrontCurveStack) (diffs []string) {
	// insertion point for field diffs
	if frontcurvestack.Name != frontcurvestackOther.Name {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "Name"))
	}
	if frontcurvestack.IsDisplayed != frontcurvestackOther.IsDisplayed {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "IsDisplayed"))
	}
	if (frontcurvestack.ShapeCategory == nil) != (frontcurvestackOther.ShapeCategory == nil) {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "ShapeCategory"))
	} else if frontcurvestack.ShapeCategory != nil && frontcurvestackOther.ShapeCategory != nil {
		if frontcurvestack.ShapeCategory != frontcurvestackOther.ShapeCategory {
			diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	FrontCurvesDifferent := false
	if len(frontcurvestack.FrontCurves) != len(frontcurvestackOther.FrontCurves) {
		FrontCurvesDifferent = true
	} else {
		for i := range frontcurvestack.FrontCurves {
			if (frontcurvestack.FrontCurves[i] == nil) != (frontcurvestackOther.FrontCurves[i] == nil) {
				FrontCurvesDifferent = true
				break
			} else if frontcurvestack.FrontCurves[i] != nil && frontcurvestackOther.FrontCurves[i] != nil {
				// this is a pointer comparaison
				if frontcurvestack.FrontCurves[i] != frontcurvestackOther.FrontCurves[i] {
					FrontCurvesDifferent = true
					break
				}
			}
		}
	}
	if FrontCurvesDifferent {
		ops := Diff(stage, frontcurvestack, frontcurvestackOther, "FrontCurves", frontcurvestackOther.FrontCurves, frontcurvestack.FrontCurves)
		diffs = append(diffs, ops)
	}
	SpiralCirclesDifferent := false
	if len(frontcurvestack.SpiralCircles) != len(frontcurvestackOther.SpiralCircles) {
		SpiralCirclesDifferent = true
	} else {
		for i := range frontcurvestack.SpiralCircles {
			if (frontcurvestack.SpiralCircles[i] == nil) != (frontcurvestackOther.SpiralCircles[i] == nil) {
				SpiralCirclesDifferent = true
				break
			} else if frontcurvestack.SpiralCircles[i] != nil && frontcurvestackOther.SpiralCircles[i] != nil {
				// this is a pointer comparaison
				if frontcurvestack.SpiralCircles[i] != frontcurvestackOther.SpiralCircles[i] {
					SpiralCirclesDifferent = true
					break
				}
			}
		}
	}
	if SpiralCirclesDifferent {
		ops := Diff(stage, frontcurvestack, frontcurvestackOther, "SpiralCircles", frontcurvestackOther.SpiralCircles, frontcurvestack.SpiralCircles)
		diffs = append(diffs, ops)
	}
	if frontcurvestack.Color != frontcurvestackOther.Color {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "Color"))
	}
	if frontcurvestack.FillOpacity != frontcurvestackOther.FillOpacity {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "FillOpacity"))
	}
	if frontcurvestack.Stroke != frontcurvestackOther.Stroke {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "Stroke"))
	}
	if frontcurvestack.StrokeOpacity != frontcurvestackOther.StrokeOpacity {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "StrokeOpacity"))
	}
	if frontcurvestack.StrokeWidth != frontcurvestackOther.StrokeWidth {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "StrokeWidth"))
	}
	if frontcurvestack.StrokeDashArray != frontcurvestackOther.StrokeDashArray {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "StrokeDashArray"))
	}
	if frontcurvestack.StrokeDashArrayWhenSelected != frontcurvestackOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if frontcurvestack.Transform != frontcurvestackOther.Transform {
		diffs = append(diffs, frontcurvestack.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (horizontalaxis *HorizontalAxis) GongDiff(stage *Stage, horizontalaxisOther *HorizontalAxis) (diffs []string) {
	// insertion point for field diffs
	if horizontalaxis.Name != horizontalaxisOther.Name {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "Name"))
	}
	if horizontalaxis.IsDisplayed != horizontalaxisOther.IsDisplayed {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "IsDisplayed"))
	}
	if (horizontalaxis.ShapeCategory == nil) != (horizontalaxisOther.ShapeCategory == nil) {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "ShapeCategory"))
	} else if horizontalaxis.ShapeCategory != nil && horizontalaxisOther.ShapeCategory != nil {
		if horizontalaxis.ShapeCategory != horizontalaxisOther.ShapeCategory {
			diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if horizontalaxis.AxisHandleBorderLength != horizontalaxisOther.AxisHandleBorderLength {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
	}
	if horizontalaxis.Axis_Length != horizontalaxisOther.Axis_Length {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "Axis_Length"))
	}
	if horizontalaxis.Color != horizontalaxisOther.Color {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "Color"))
	}
	if horizontalaxis.FillOpacity != horizontalaxisOther.FillOpacity {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "FillOpacity"))
	}
	if horizontalaxis.Stroke != horizontalaxisOther.Stroke {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "Stroke"))
	}
	if horizontalaxis.StrokeOpacity != horizontalaxisOther.StrokeOpacity {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "StrokeOpacity"))
	}
	if horizontalaxis.StrokeWidth != horizontalaxisOther.StrokeWidth {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "StrokeWidth"))
	}
	if horizontalaxis.StrokeDashArray != horizontalaxisOther.StrokeDashArray {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "StrokeDashArray"))
	}
	if horizontalaxis.StrokeDashArrayWhenSelected != horizontalaxisOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if horizontalaxis.Transform != horizontalaxisOther.Transform {
		diffs = append(diffs, horizontalaxis.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key *Key) GongDiff(stage *Stage, keyOther *Key) (diffs []string) {
	// insertion point for field diffs
	if key.Name != keyOther.Name {
		diffs = append(diffs, key.GongMarshallField(stage, "Name"))
	}
	if key.IsDisplayed != keyOther.IsDisplayed {
		diffs = append(diffs, key.GongMarshallField(stage, "IsDisplayed"))
	}
	if (key.ShapeCategory == nil) != (keyOther.ShapeCategory == nil) {
		diffs = append(diffs, key.GongMarshallField(stage, "ShapeCategory"))
	} else if key.ShapeCategory != nil && keyOther.ShapeCategory != nil {
		if key.ShapeCategory != keyOther.ShapeCategory {
			diffs = append(diffs, key.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if key.Path != keyOther.Path {
		diffs = append(diffs, key.GongMarshallField(stage, "Path"))
	}
	if key.Color != keyOther.Color {
		diffs = append(diffs, key.GongMarshallField(stage, "Color"))
	}
	if key.FillOpacity != keyOther.FillOpacity {
		diffs = append(diffs, key.GongMarshallField(stage, "FillOpacity"))
	}
	if key.Stroke != keyOther.Stroke {
		diffs = append(diffs, key.GongMarshallField(stage, "Stroke"))
	}
	if key.StrokeOpacity != keyOther.StrokeOpacity {
		diffs = append(diffs, key.GongMarshallField(stage, "StrokeOpacity"))
	}
	if key.StrokeWidth != keyOther.StrokeWidth {
		diffs = append(diffs, key.GongMarshallField(stage, "StrokeWidth"))
	}
	if key.StrokeDashArray != keyOther.StrokeDashArray {
		diffs = append(diffs, key.GongMarshallField(stage, "StrokeDashArray"))
	}
	if key.StrokeDashArrayWhenSelected != keyOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, key.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if key.Transform != keyOther.Transform {
		diffs = append(diffs, key.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parameter *Parameter) GongDiff(stage *Stage, parameterOther *Parameter) (diffs []string) {
	// insertion point for field diffs
	if parameter.Name != parameterOther.Name {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Name"))
	}
	if parameter.BackendColor != parameterOther.BackendColor {
		diffs = append(diffs, parameter.GongMarshallField(stage, "BackendColor"))
	}
	if parameter.MinuteColor != parameterOther.MinuteColor {
		diffs = append(diffs, parameter.GongMarshallField(stage, "MinuteColor"))
	}
	if parameter.HourColor != parameterOther.HourColor {
		diffs = append(diffs, parameter.GongMarshallField(stage, "HourColor"))
	}
	if parameter.N != parameterOther.N {
		diffs = append(diffs, parameter.GongMarshallField(stage, "N"))
	}
	if parameter.M != parameterOther.M {
		diffs = append(diffs, parameter.GongMarshallField(stage, "M"))
	}
	if parameter.Z != parameterOther.Z {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Z"))
	}
	if parameter.InsideAngle != parameterOther.InsideAngle {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InsideAngle"))
	}
	if parameter.ShiftToNearestCircle != parameterOther.ShiftToNearestCircle {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ShiftToNearestCircle"))
	}
	if parameter.SideLength != parameterOther.SideLength {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SideLength"))
	}
	if (parameter.InitialRhombus == nil) != (parameterOther.InitialRhombus == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InitialRhombus"))
	} else if parameter.InitialRhombus != nil && parameterOther.InitialRhombus != nil {
		if parameter.InitialRhombus != parameterOther.InitialRhombus {
			diffs = append(diffs, parameter.GongMarshallField(stage, "InitialRhombus"))
		}
	}
	if (parameter.InitialCircle == nil) != (parameterOther.InitialCircle == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InitialCircle"))
	} else if parameter.InitialCircle != nil && parameterOther.InitialCircle != nil {
		if parameter.InitialCircle != parameterOther.InitialCircle {
			diffs = append(diffs, parameter.GongMarshallField(stage, "InitialCircle"))
		}
	}
	if (parameter.InitialRhombusGrid == nil) != (parameterOther.InitialRhombusGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InitialRhombusGrid"))
	} else if parameter.InitialRhombusGrid != nil && parameterOther.InitialRhombusGrid != nil {
		if parameter.InitialRhombusGrid != parameterOther.InitialRhombusGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "InitialRhombusGrid"))
		}
	}
	if (parameter.InitialCircleGrid == nil) != (parameterOther.InitialCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InitialCircleGrid"))
	} else if parameter.InitialCircleGrid != nil && parameterOther.InitialCircleGrid != nil {
		if parameter.InitialCircleGrid != parameterOther.InitialCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "InitialCircleGrid"))
		}
	}
	if (parameter.InitialAxis == nil) != (parameterOther.InitialAxis == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "InitialAxis"))
	} else if parameter.InitialAxis != nil && parameterOther.InitialAxis != nil {
		if parameter.InitialAxis != parameterOther.InitialAxis {
			diffs = append(diffs, parameter.GongMarshallField(stage, "InitialAxis"))
		}
	}
	if (parameter.RotatedAxis == nil) != (parameterOther.RotatedAxis == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedAxis"))
	} else if parameter.RotatedAxis != nil && parameterOther.RotatedAxis != nil {
		if parameter.RotatedAxis != parameterOther.RotatedAxis {
			diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedAxis"))
		}
	}
	if (parameter.RotatedRhombus == nil) != (parameterOther.RotatedRhombus == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedRhombus"))
	} else if parameter.RotatedRhombus != nil && parameterOther.RotatedRhombus != nil {
		if parameter.RotatedRhombus != parameterOther.RotatedRhombus {
			diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedRhombus"))
		}
	}
	if (parameter.RotatedRhombusGrid == nil) != (parameterOther.RotatedRhombusGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedRhombusGrid"))
	} else if parameter.RotatedRhombusGrid != nil && parameterOther.RotatedRhombusGrid != nil {
		if parameter.RotatedRhombusGrid != parameterOther.RotatedRhombusGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedRhombusGrid"))
		}
	}
	if (parameter.RotatedCircleGrid == nil) != (parameterOther.RotatedCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedCircleGrid"))
	} else if parameter.RotatedCircleGrid != nil && parameterOther.RotatedCircleGrid != nil {
		if parameter.RotatedCircleGrid != parameterOther.RotatedCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "RotatedCircleGrid"))
		}
	}
	if (parameter.NextRhombus == nil) != (parameterOther.NextRhombus == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NextRhombus"))
	} else if parameter.NextRhombus != nil && parameterOther.NextRhombus != nil {
		if parameter.NextRhombus != parameterOther.NextRhombus {
			diffs = append(diffs, parameter.GongMarshallField(stage, "NextRhombus"))
		}
	}
	if (parameter.NextCircle == nil) != (parameterOther.NextCircle == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NextCircle"))
	} else if parameter.NextCircle != nil && parameterOther.NextCircle != nil {
		if parameter.NextCircle != parameterOther.NextCircle {
			diffs = append(diffs, parameter.GongMarshallField(stage, "NextCircle"))
		}
	}
	if (parameter.GrowingRhombusGridSeed == nil) != (parameterOther.GrowingRhombusGridSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingRhombusGridSeed"))
	} else if parameter.GrowingRhombusGridSeed != nil && parameterOther.GrowingRhombusGridSeed != nil {
		if parameter.GrowingRhombusGridSeed != parameterOther.GrowingRhombusGridSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingRhombusGridSeed"))
		}
	}
	if (parameter.GrowingRhombusGrid == nil) != (parameterOther.GrowingRhombusGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingRhombusGrid"))
	} else if parameter.GrowingRhombusGrid != nil && parameterOther.GrowingRhombusGrid != nil {
		if parameter.GrowingRhombusGrid != parameterOther.GrowingRhombusGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingRhombusGrid"))
		}
	}
	if (parameter.GrowingCircleGridSeed == nil) != (parameterOther.GrowingCircleGridSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridSeed"))
	} else if parameter.GrowingCircleGridSeed != nil && parameterOther.GrowingCircleGridSeed != nil {
		if parameter.GrowingCircleGridSeed != parameterOther.GrowingCircleGridSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridSeed"))
		}
	}
	if (parameter.GrowingCircleGrid == nil) != (parameterOther.GrowingCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGrid"))
	} else if parameter.GrowingCircleGrid != nil && parameterOther.GrowingCircleGrid != nil {
		if parameter.GrowingCircleGrid != parameterOther.GrowingCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGrid"))
		}
	}
	if (parameter.GrowingCircleGridLeftSeed == nil) != (parameterOther.GrowingCircleGridLeftSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed"))
	} else if parameter.GrowingCircleGridLeftSeed != nil && parameterOther.GrowingCircleGridLeftSeed != nil {
		if parameter.GrowingCircleGridLeftSeed != parameterOther.GrowingCircleGridLeftSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed"))
		}
	}
	if (parameter.GrowingCircleGridLeft == nil) != (parameterOther.GrowingCircleGridLeft == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridLeft"))
	} else if parameter.GrowingCircleGridLeft != nil && parameterOther.GrowingCircleGridLeft != nil {
		if parameter.GrowingCircleGridLeft != parameterOther.GrowingCircleGridLeft {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowingCircleGridLeft"))
		}
	}
	if (parameter.ConstructionAxis == nil) != (parameterOther.ConstructionAxis == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionAxis"))
	} else if parameter.ConstructionAxis != nil && parameterOther.ConstructionAxis != nil {
		if parameter.ConstructionAxis != parameterOther.ConstructionAxis {
			diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionAxis"))
		}
	}
	if (parameter.ConstructionAxisGrid == nil) != (parameterOther.ConstructionAxisGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionAxisGrid"))
	} else if parameter.ConstructionAxisGrid != nil && parameterOther.ConstructionAxisGrid != nil {
		if parameter.ConstructionAxisGrid != parameterOther.ConstructionAxisGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionAxisGrid"))
		}
	}
	if (parameter.ConstructionCircle == nil) != (parameterOther.ConstructionCircle == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionCircle"))
	} else if parameter.ConstructionCircle != nil && parameterOther.ConstructionCircle != nil {
		if parameter.ConstructionCircle != parameterOther.ConstructionCircle {
			diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionCircle"))
		}
	}
	if (parameter.ConstructionCircleGrid == nil) != (parameterOther.ConstructionCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionCircleGrid"))
	} else if parameter.ConstructionCircleGrid != nil && parameterOther.ConstructionCircleGrid != nil {
		if parameter.ConstructionCircleGrid != parameterOther.ConstructionCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "ConstructionCircleGrid"))
		}
	}
	if (parameter.GrowthCurveSeed == nil) != (parameterOther.GrowthCurveSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveSeed"))
	} else if parameter.GrowthCurveSeed != nil && parameterOther.GrowthCurveSeed != nil {
		if parameter.GrowthCurveSeed != parameterOther.GrowthCurveSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveSeed"))
		}
	}
	if (parameter.GrowthCurve == nil) != (parameterOther.GrowthCurve == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurve"))
	} else if parameter.GrowthCurve != nil && parameterOther.GrowthCurve != nil {
		if parameter.GrowthCurve != parameterOther.GrowthCurve {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurve"))
		}
	}
	if (parameter.GrowthCurveShiftedRightSeed == nil) != (parameterOther.GrowthCurveShiftedRightSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed"))
	} else if parameter.GrowthCurveShiftedRightSeed != nil && parameterOther.GrowthCurveShiftedRightSeed != nil {
		if parameter.GrowthCurveShiftedRightSeed != parameterOther.GrowthCurveShiftedRightSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed"))
		}
	}
	if (parameter.GrowthCurveShiftedRight == nil) != (parameterOther.GrowthCurveShiftedRight == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveShiftedRight"))
	} else if parameter.GrowthCurveShiftedRight != nil && parameterOther.GrowthCurveShiftedRight != nil {
		if parameter.GrowthCurveShiftedRight != parameterOther.GrowthCurveShiftedRight {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveShiftedRight"))
		}
	}
	if (parameter.GrowthCurveNextSeed == nil) != (parameterOther.GrowthCurveNextSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextSeed"))
	} else if parameter.GrowthCurveNextSeed != nil && parameterOther.GrowthCurveNextSeed != nil {
		if parameter.GrowthCurveNextSeed != parameterOther.GrowthCurveNextSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextSeed"))
		}
	}
	if (parameter.GrowthCurveNext == nil) != (parameterOther.GrowthCurveNext == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNext"))
	} else if parameter.GrowthCurveNext != nil && parameterOther.GrowthCurveNext != nil {
		if parameter.GrowthCurveNext != parameterOther.GrowthCurveNext {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNext"))
		}
	}
	if (parameter.GrowthCurveNextShiftedRightSeed == nil) != (parameterOther.GrowthCurveNextShiftedRightSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed"))
	} else if parameter.GrowthCurveNextShiftedRightSeed != nil && parameterOther.GrowthCurveNextShiftedRightSeed != nil {
		if parameter.GrowthCurveNextShiftedRightSeed != parameterOther.GrowthCurveNextShiftedRightSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed"))
		}
	}
	if (parameter.GrowthCurveNextShiftedRight == nil) != (parameterOther.GrowthCurveNextShiftedRight == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight"))
	} else if parameter.GrowthCurveNextShiftedRight != nil && parameterOther.GrowthCurveNextShiftedRight != nil {
		if parameter.GrowthCurveNextShiftedRight != parameterOther.GrowthCurveNextShiftedRight {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight"))
		}
	}
	if (parameter.GrowthCurveStack == nil) != (parameterOther.GrowthCurveStack == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveStack"))
	} else if parameter.GrowthCurveStack != nil && parameterOther.GrowthCurveStack != nil {
		if parameter.GrowthCurveStack != parameterOther.GrowthCurveStack {
			diffs = append(diffs, parameter.GongMarshallField(stage, "GrowthCurveStack"))
		}
	}
	if parameter.StackWidth != parameterOther.StackWidth {
		diffs = append(diffs, parameter.GongMarshallField(stage, "StackWidth"))
	}
	if parameter.NbShitRight != parameterOther.NbShitRight {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NbShitRight"))
	}
	if parameter.StackHeight != parameterOther.StackHeight {
		diffs = append(diffs, parameter.GongMarshallField(stage, "StackHeight"))
	}
	if parameter.BezierControlLengthRatio != parameterOther.BezierControlLengthRatio {
		diffs = append(diffs, parameter.GongMarshallField(stage, "BezierControlLengthRatio"))
	}
	if (parameter.SpiralRhombusGridSeed == nil) != (parameterOther.SpiralRhombusGridSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralRhombusGridSeed"))
	} else if parameter.SpiralRhombusGridSeed != nil && parameterOther.SpiralRhombusGridSeed != nil {
		if parameter.SpiralRhombusGridSeed != parameterOther.SpiralRhombusGridSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralRhombusGridSeed"))
		}
	}
	if (parameter.SpiralRhombusGrid == nil) != (parameterOther.SpiralRhombusGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralRhombusGrid"))
	} else if parameter.SpiralRhombusGrid != nil && parameterOther.SpiralRhombusGrid != nil {
		if parameter.SpiralRhombusGrid != parameterOther.SpiralRhombusGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralRhombusGrid"))
		}
	}
	if (parameter.SpiralCircleSeed == nil) != (parameterOther.SpiralCircleSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleSeed"))
	} else if parameter.SpiralCircleSeed != nil && parameterOther.SpiralCircleSeed != nil {
		if parameter.SpiralCircleSeed != parameterOther.SpiralCircleSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleSeed"))
		}
	}
	if (parameter.SpiralCircleGrid == nil) != (parameterOther.SpiralCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleGrid"))
	} else if parameter.SpiralCircleGrid != nil && parameterOther.SpiralCircleGrid != nil {
		if parameter.SpiralCircleGrid != parameterOther.SpiralCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleGrid"))
		}
	}
	if (parameter.SpiralCircleFullGrid == nil) != (parameterOther.SpiralCircleFullGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleFullGrid"))
	} else if parameter.SpiralCircleFullGrid != nil && parameterOther.SpiralCircleFullGrid != nil {
		if parameter.SpiralCircleFullGrid != parameterOther.SpiralCircleFullGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralCircleFullGrid"))
		}
	}
	if (parameter.SpiralConstructionOuterLineSeed == nil) != (parameterOther.SpiralConstructionOuterLineSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed"))
	} else if parameter.SpiralConstructionOuterLineSeed != nil && parameterOther.SpiralConstructionOuterLineSeed != nil {
		if parameter.SpiralConstructionOuterLineSeed != parameterOther.SpiralConstructionOuterLineSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed"))
		}
	}
	if (parameter.SpiralConstructionInnerLineSeed == nil) != (parameterOther.SpiralConstructionInnerLineSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed"))
	} else if parameter.SpiralConstructionInnerLineSeed != nil && parameterOther.SpiralConstructionInnerLineSeed != nil {
		if parameter.SpiralConstructionInnerLineSeed != parameterOther.SpiralConstructionInnerLineSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed"))
		}
	}
	if (parameter.SpiralConstructionOuterLineGrid == nil) != (parameterOther.SpiralConstructionOuterLineGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid"))
	} else if parameter.SpiralConstructionOuterLineGrid != nil && parameterOther.SpiralConstructionOuterLineGrid != nil {
		if parameter.SpiralConstructionOuterLineGrid != parameterOther.SpiralConstructionOuterLineGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid"))
		}
	}
	if (parameter.SpiralConstructionInnerLineGrid == nil) != (parameterOther.SpiralConstructionInnerLineGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid"))
	} else if parameter.SpiralConstructionInnerLineGrid != nil && parameterOther.SpiralConstructionInnerLineGrid != nil {
		if parameter.SpiralConstructionInnerLineGrid != parameterOther.SpiralConstructionInnerLineGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid"))
		}
	}
	if (parameter.SpiralConstructionCircleGrid == nil) != (parameterOther.SpiralConstructionCircleGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid"))
	} else if parameter.SpiralConstructionCircleGrid != nil && parameterOther.SpiralConstructionCircleGrid != nil {
		if parameter.SpiralConstructionCircleGrid != parameterOther.SpiralConstructionCircleGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid"))
		}
	}
	if (parameter.SpiralConstructionOuterLineFullGrid == nil) != (parameterOther.SpiralConstructionOuterLineFullGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid"))
	} else if parameter.SpiralConstructionOuterLineFullGrid != nil && parameterOther.SpiralConstructionOuterLineFullGrid != nil {
		if parameter.SpiralConstructionOuterLineFullGrid != parameterOther.SpiralConstructionOuterLineFullGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid"))
		}
	}
	if (parameter.SpiralBezierSeed == nil) != (parameterOther.SpiralBezierSeed == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierSeed"))
	} else if parameter.SpiralBezierSeed != nil && parameterOther.SpiralBezierSeed != nil {
		if parameter.SpiralBezierSeed != parameterOther.SpiralBezierSeed {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierSeed"))
		}
	}
	if (parameter.SpiralBezierGrid == nil) != (parameterOther.SpiralBezierGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierGrid"))
	} else if parameter.SpiralBezierGrid != nil && parameterOther.SpiralBezierGrid != nil {
		if parameter.SpiralBezierGrid != parameterOther.SpiralBezierGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierGrid"))
		}
	}
	if (parameter.SpiralBezierFullGrid == nil) != (parameterOther.SpiralBezierFullGrid == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierFullGrid"))
	} else if parameter.SpiralBezierFullGrid != nil && parameterOther.SpiralBezierFullGrid != nil {
		if parameter.SpiralBezierFullGrid != parameterOther.SpiralBezierFullGrid {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierFullGrid"))
		}
	}
	if parameter.SpiralBezierStrength != parameterOther.SpiralBezierStrength {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralBezierStrength"))
	}
	if (parameter.FrontCurveStack == nil) != (parameterOther.FrontCurveStack == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FrontCurveStack"))
	} else if parameter.FrontCurveStack != nil && parameterOther.FrontCurveStack != nil {
		if parameter.FrontCurveStack != parameterOther.FrontCurveStack {
			diffs = append(diffs, parameter.GongMarshallField(stage, "FrontCurveStack"))
		}
	}
	if parameter.NbInterpolationPoints != parameterOther.NbInterpolationPoints {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NbInterpolationPoints"))
	}
	if (parameter.Fkey == nil) != (parameterOther.Fkey == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Fkey"))
	} else if parameter.Fkey != nil && parameterOther.Fkey != nil {
		if parameter.Fkey != parameterOther.Fkey {
			diffs = append(diffs, parameter.GongMarshallField(stage, "Fkey"))
		}
	}
	if parameter.FkeySizeRatio != parameterOther.FkeySizeRatio {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FkeySizeRatio"))
	}
	if parameter.FkeyOriginRelativeX != parameterOther.FkeyOriginRelativeX {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FkeyOriginRelativeX"))
	}
	if parameter.FkeyOriginRelativeY != parameterOther.FkeyOriginRelativeY {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FkeyOriginRelativeY"))
	}
	if (parameter.PitchLines == nil) != (parameterOther.PitchLines == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PitchLines"))
	} else if parameter.PitchLines != nil && parameterOther.PitchLines != nil {
		if parameter.PitchLines != parameterOther.PitchLines {
			diffs = append(diffs, parameter.GongMarshallField(stage, "PitchLines"))
		}
	}
	if parameter.PitchHeight != parameterOther.PitchHeight {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PitchHeight"))
	}
	if parameter.NbPitchLines != parameterOther.NbPitchLines {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NbPitchLines"))
	}
	if (parameter.BeatLines == nil) != (parameterOther.BeatLines == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "BeatLines"))
	} else if parameter.BeatLines != nil && parameterOther.BeatLines != nil {
		if parameter.BeatLines != parameterOther.BeatLines {
			diffs = append(diffs, parameter.GongMarshallField(stage, "BeatLines"))
		}
	}
	if parameter.BeatLinesHeightRatio != parameterOther.BeatLinesHeightRatio {
		diffs = append(diffs, parameter.GongMarshallField(stage, "BeatLinesHeightRatio"))
	}
	if parameter.NbBeatLines != parameterOther.NbBeatLines {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NbBeatLines"))
	}
	if parameter.NbOfBeatsInTheme != parameterOther.NbOfBeatsInTheme {
		diffs = append(diffs, parameter.GongMarshallField(stage, "NbOfBeatsInTheme"))
	}
	if (parameter.FirstVoice == nil) != (parameterOther.FirstVoice == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoice"))
	} else if parameter.FirstVoice != nil && parameterOther.FirstVoice != nil {
		if parameter.FirstVoice != parameterOther.FirstVoice {
			diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoice"))
		}
	}
	if (parameter.FirstVoiceShiftedRigth == nil) != (parameterOther.FirstVoiceShiftedRigth == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth"))
	} else if parameter.FirstVoiceShiftedRigth != nil && parameterOther.FirstVoiceShiftedRigth != nil {
		if parameter.FirstVoiceShiftedRigth != parameterOther.FirstVoiceShiftedRigth {
			diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth"))
		}
	}
	if parameter.FirstVoiceShiftX != parameterOther.FirstVoiceShiftX {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceShiftX"))
	}
	if parameter.FirstVoiceShiftY != parameterOther.FirstVoiceShiftY {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceShiftY"))
	}
	if (parameter.SecondVoice == nil) != (parameterOther.SecondVoice == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoice"))
	} else if parameter.SecondVoice != nil && parameterOther.SecondVoice != nil {
		if parameter.SecondVoice != parameterOther.SecondVoice {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoice"))
		}
	}
	if (parameter.SecondVoiceShiftedRight == nil) != (parameterOther.SecondVoiceShiftedRight == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceShiftedRight"))
	} else if parameter.SecondVoiceShiftedRight != nil && parameterOther.SecondVoiceShiftedRight != nil {
		if parameter.SecondVoiceShiftedRight != parameterOther.SecondVoiceShiftedRight {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceShiftedRight"))
		}
	}
	if parameter.PitchDifference != parameterOther.PitchDifference {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PitchDifference"))
	}
	if parameter.BeatsPerSecond != parameterOther.BeatsPerSecond {
		diffs = append(diffs, parameter.GongMarshallField(stage, "BeatsPerSecond"))
	}
	if parameter.Level != parameterOther.Level {
		diffs = append(diffs, parameter.GongMarshallField(stage, "Level"))
	}
	if (parameter.FirstVoiceNotes == nil) != (parameterOther.FirstVoiceNotes == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceNotes"))
	} else if parameter.FirstVoiceNotes != nil && parameterOther.FirstVoiceNotes != nil {
		if parameter.FirstVoiceNotes != parameterOther.FirstVoiceNotes {
			diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceNotes"))
		}
	}
	if (parameter.FirstVoiceNotesShiftedRight == nil) != (parameterOther.FirstVoiceNotesShiftedRight == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight"))
	} else if parameter.FirstVoiceNotesShiftedRight != nil && parameterOther.FirstVoiceNotesShiftedRight != nil {
		if parameter.FirstVoiceNotesShiftedRight != parameterOther.FirstVoiceNotesShiftedRight {
			diffs = append(diffs, parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight"))
		}
	}
	if (parameter.SecondVoiceNotes == nil) != (parameterOther.SecondVoiceNotes == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceNotes"))
	} else if parameter.SecondVoiceNotes != nil && parameterOther.SecondVoiceNotes != nil {
		if parameter.SecondVoiceNotes != parameterOther.SecondVoiceNotes {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceNotes"))
		}
	}
	if (parameter.SecondVoiceNotesShiftedRight == nil) != (parameterOther.SecondVoiceNotesShiftedRight == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight"))
	} else if parameter.SecondVoiceNotesShiftedRight != nil && parameterOther.SecondVoiceNotesShiftedRight != nil {
		if parameter.SecondVoiceNotesShiftedRight != parameterOther.SecondVoiceNotesShiftedRight {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight"))
		}
	}
	if parameter.IsMinor != parameterOther.IsMinor {
		diffs = append(diffs, parameter.GongMarshallField(stage, "IsMinor"))
	}
	if parameter.ThemeBinaryEncoding != parameterOther.ThemeBinaryEncoding {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ThemeBinaryEncoding"))
	}
	if parameter.OriginX != parameterOther.OriginX {
		diffs = append(diffs, parameter.GongMarshallField(stage, "OriginX"))
	}
	if parameter.OriginY != parameterOther.OriginY {
		diffs = append(diffs, parameter.GongMarshallField(stage, "OriginY"))
	}
	if (parameter.HorizontalAxis == nil) != (parameterOther.HorizontalAxis == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "HorizontalAxis"))
	} else if parameter.HorizontalAxis != nil && parameterOther.HorizontalAxis != nil {
		if parameter.HorizontalAxis != parameterOther.HorizontalAxis {
			diffs = append(diffs, parameter.GongMarshallField(stage, "HorizontalAxis"))
		}
	}
	if (parameter.VerticalAxis == nil) != (parameterOther.VerticalAxis == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "VerticalAxis"))
	} else if parameter.VerticalAxis != nil && parameterOther.VerticalAxis != nil {
		if parameter.VerticalAxis != parameterOther.VerticalAxis {
			diffs = append(diffs, parameter.GongMarshallField(stage, "VerticalAxis"))
		}
	}
	if (parameter.SpiralOrigin == nil) != (parameterOther.SpiralOrigin == nil) {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralOrigin"))
	} else if parameter.SpiralOrigin != nil && parameterOther.SpiralOrigin != nil {
		if parameter.SpiralOrigin != parameterOther.SpiralOrigin {
			diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralOrigin"))
		}
	}
	if parameter.SpiralOriginX != parameterOther.SpiralOriginX {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralOriginX"))
	}
	if parameter.SpiralOriginY != parameterOther.SpiralOriginY {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralOriginY"))
	}
	if parameter.OriginCrossWidth != parameterOther.OriginCrossWidth {
		diffs = append(diffs, parameter.GongMarshallField(stage, "OriginCrossWidth"))
	}
	if parameter.SpiralRadiusRatio != parameterOther.SpiralRadiusRatio {
		diffs = append(diffs, parameter.GongMarshallField(stage, "SpiralRadiusRatio"))
	}
	if parameter.ShowSpiralBezierConstruct != parameterOther.ShowSpiralBezierConstruct {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ShowSpiralBezierConstruct"))
	}
	if parameter.ShowInterpolationPoints != parameterOther.ShowInterpolationPoints {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ShowInterpolationPoints"))
	}
	if parameter.ActualBeatsTemporalShift != parameterOther.ActualBeatsTemporalShift {
		diffs = append(diffs, parameter.GongMarshallField(stage, "ActualBeatsTemporalShift"))
	}
	if parameter.PathToStaticFiles != parameterOther.PathToStaticFiles {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PathToStaticFiles"))
	}
	if parameter.PathToGeneratedSVG != parameterOther.PathToGeneratedSVG {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PathToGeneratedSVG"))
	}
	if parameter.PathToGeneratedScore != parameterOther.PathToGeneratedScore {
		diffs = append(diffs, parameter.GongMarshallField(stage, "PathToGeneratedScore"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rhombus *Rhombus) GongDiff(stage *Stage, rhombusOther *Rhombus) (diffs []string) {
	// insertion point for field diffs
	if rhombus.Name != rhombusOther.Name {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "Name"))
	}
	if rhombus.IsDisplayed != rhombusOther.IsDisplayed {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "IsDisplayed"))
	}
	if (rhombus.ShapeCategory == nil) != (rhombusOther.ShapeCategory == nil) {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "ShapeCategory"))
	} else if rhombus.ShapeCategory != nil && rhombusOther.ShapeCategory != nil {
		if rhombus.ShapeCategory != rhombusOther.ShapeCategory {
			diffs = append(diffs, rhombus.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if rhombus.CenterX != rhombusOther.CenterX {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "CenterX"))
	}
	if rhombus.CenterY != rhombusOther.CenterY {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "CenterY"))
	}
	if rhombus.SideLength != rhombusOther.SideLength {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "SideLength"))
	}
	if rhombus.AngleDegree != rhombusOther.AngleDegree {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "AngleDegree"))
	}
	if rhombus.InsideAngle != rhombusOther.InsideAngle {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "InsideAngle"))
	}
	if rhombus.Color != rhombusOther.Color {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "Color"))
	}
	if rhombus.FillOpacity != rhombusOther.FillOpacity {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "FillOpacity"))
	}
	if rhombus.Stroke != rhombusOther.Stroke {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "Stroke"))
	}
	if rhombus.StrokeOpacity != rhombusOther.StrokeOpacity {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rhombus.StrokeWidth != rhombusOther.StrokeWidth {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "StrokeWidth"))
	}
	if rhombus.StrokeDashArray != rhombusOther.StrokeDashArray {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rhombus.StrokeDashArrayWhenSelected != rhombusOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rhombus.Transform != rhombusOther.Transform {
		diffs = append(diffs, rhombus.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rhombusgrid *RhombusGrid) GongDiff(stage *Stage, rhombusgridOther *RhombusGrid) (diffs []string) {
	// insertion point for field diffs
	if rhombusgrid.Name != rhombusgridOther.Name {
		diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "Name"))
	}
	if (rhombusgrid.Reference == nil) != (rhombusgridOther.Reference == nil) {
		diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "Reference"))
	} else if rhombusgrid.Reference != nil && rhombusgridOther.Reference != nil {
		if rhombusgrid.Reference != rhombusgridOther.Reference {
			diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "Reference"))
		}
	}
	if rhombusgrid.IsDisplayed != rhombusgridOther.IsDisplayed {
		diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (rhombusgrid.ShapeCategory == nil) != (rhombusgridOther.ShapeCategory == nil) {
		diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "ShapeCategory"))
	} else if rhombusgrid.ShapeCategory != nil && rhombusgridOther.ShapeCategory != nil {
		if rhombusgrid.ShapeCategory != rhombusgridOther.ShapeCategory {
			diffs = append(diffs, rhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	RhombusesDifferent := false
	if len(rhombusgrid.Rhombuses) != len(rhombusgridOther.Rhombuses) {
		RhombusesDifferent = true
	} else {
		for i := range rhombusgrid.Rhombuses {
			if (rhombusgrid.Rhombuses[i] == nil) != (rhombusgridOther.Rhombuses[i] == nil) {
				RhombusesDifferent = true
				break
			} else if rhombusgrid.Rhombuses[i] != nil && rhombusgridOther.Rhombuses[i] != nil {
				// this is a pointer comparaison
				if rhombusgrid.Rhombuses[i] != rhombusgridOther.Rhombuses[i] {
					RhombusesDifferent = true
					break
				}
			}
		}
	}
	if RhombusesDifferent {
		ops := Diff(stage, rhombusgrid, rhombusgridOther, "Rhombuses", rhombusgridOther.Rhombuses, rhombusgrid.Rhombuses)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shapecategory *ShapeCategory) GongDiff(stage *Stage, shapecategoryOther *ShapeCategory) (diffs []string) {
	// insertion point for field diffs
	if shapecategory.Name != shapecategoryOther.Name {
		diffs = append(diffs, shapecategory.GongMarshallField(stage, "Name"))
	}
	if shapecategory.IsExpanded != shapecategoryOther.IsExpanded {
		diffs = append(diffs, shapecategory.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralbezier *SpiralBezier) GongDiff(stage *Stage, spiralbezierOther *SpiralBezier) (diffs []string) {
	// insertion point for field diffs
	if spiralbezier.Name != spiralbezierOther.Name {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "Name"))
	}
	if spiralbezier.IsDisplayed != spiralbezierOther.IsDisplayed {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralbezier.ShapeCategory == nil) != (spiralbezierOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralbezier.ShapeCategory != nil && spiralbezierOther.ShapeCategory != nil {
		if spiralbezier.ShapeCategory != spiralbezierOther.ShapeCategory {
			diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if spiralbezier.StartX != spiralbezierOther.StartX {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StartX"))
	}
	if spiralbezier.StartY != spiralbezierOther.StartY {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StartY"))
	}
	if spiralbezier.ControlPointStartX != spiralbezierOther.ControlPointStartX {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ControlPointStartX"))
	}
	if spiralbezier.ControlPointStartY != spiralbezierOther.ControlPointStartY {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ControlPointStartY"))
	}
	if spiralbezier.EndX != spiralbezierOther.EndX {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "EndX"))
	}
	if spiralbezier.EndY != spiralbezierOther.EndY {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "EndY"))
	}
	if spiralbezier.ControlPointEndX != spiralbezierOther.ControlPointEndX {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ControlPointEndX"))
	}
	if spiralbezier.ControlPointEndY != spiralbezierOther.ControlPointEndY {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "ControlPointEndY"))
	}
	if spiralbezier.Color != spiralbezierOther.Color {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "Color"))
	}
	if spiralbezier.FillOpacity != spiralbezierOther.FillOpacity {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "FillOpacity"))
	}
	if spiralbezier.Stroke != spiralbezierOther.Stroke {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "Stroke"))
	}
	if spiralbezier.StrokeOpacity != spiralbezierOther.StrokeOpacity {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StrokeOpacity"))
	}
	if spiralbezier.StrokeWidth != spiralbezierOther.StrokeWidth {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StrokeWidth"))
	}
	if spiralbezier.StrokeDashArray != spiralbezierOther.StrokeDashArray {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StrokeDashArray"))
	}
	if spiralbezier.StrokeDashArrayWhenSelected != spiralbezierOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if spiralbezier.Transform != spiralbezierOther.Transform {
		diffs = append(diffs, spiralbezier.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralbeziergrid *SpiralBezierGrid) GongDiff(stage *Stage, spiralbeziergridOther *SpiralBezierGrid) (diffs []string) {
	// insertion point for field diffs
	if spiralbeziergrid.Name != spiralbeziergridOther.Name {
		diffs = append(diffs, spiralbeziergrid.GongMarshallField(stage, "Name"))
	}
	if spiralbeziergrid.IsDisplayed != spiralbeziergridOther.IsDisplayed {
		diffs = append(diffs, spiralbeziergrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralbeziergrid.ShapeCategory == nil) != (spiralbeziergridOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralbeziergrid.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralbeziergrid.ShapeCategory != nil && spiralbeziergridOther.ShapeCategory != nil {
		if spiralbeziergrid.ShapeCategory != spiralbeziergridOther.ShapeCategory {
			diffs = append(diffs, spiralbeziergrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	SpiralBeziersDifferent := false
	if len(spiralbeziergrid.SpiralBeziers) != len(spiralbeziergridOther.SpiralBeziers) {
		SpiralBeziersDifferent = true
	} else {
		for i := range spiralbeziergrid.SpiralBeziers {
			if (spiralbeziergrid.SpiralBeziers[i] == nil) != (spiralbeziergridOther.SpiralBeziers[i] == nil) {
				SpiralBeziersDifferent = true
				break
			} else if spiralbeziergrid.SpiralBeziers[i] != nil && spiralbeziergridOther.SpiralBeziers[i] != nil {
				// this is a pointer comparaison
				if spiralbeziergrid.SpiralBeziers[i] != spiralbeziergridOther.SpiralBeziers[i] {
					SpiralBeziersDifferent = true
					break
				}
			}
		}
	}
	if SpiralBeziersDifferent {
		ops := Diff(stage, spiralbeziergrid, spiralbeziergridOther, "SpiralBeziers", spiralbeziergridOther.SpiralBeziers, spiralbeziergrid.SpiralBeziers)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralcircle *SpiralCircle) GongDiff(stage *Stage, spiralcircleOther *SpiralCircle) (diffs []string) {
	// insertion point for field diffs
	if spiralcircle.Name != spiralcircleOther.Name {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Name"))
	}
	if spiralcircle.IsDisplayed != spiralcircleOther.IsDisplayed {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralcircle.ShapeCategory == nil) != (spiralcircleOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralcircle.ShapeCategory != nil && spiralcircleOther.ShapeCategory != nil {
		if spiralcircle.ShapeCategory != spiralcircleOther.ShapeCategory {
			diffs = append(diffs, spiralcircle.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if spiralcircle.CenterX != spiralcircleOther.CenterX {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "CenterX"))
	}
	if spiralcircle.CenterY != spiralcircleOther.CenterY {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "CenterY"))
	}
	if spiralcircle.HasBespokeRadius != spiralcircleOther.HasBespokeRadius {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "HasBespokeRadius"))
	}
	if spiralcircle.BespopkeRadius != spiralcircleOther.BespopkeRadius {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "BespopkeRadius"))
	}
	if spiralcircle.Color != spiralcircleOther.Color {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Color"))
	}
	if spiralcircle.FillOpacity != spiralcircleOther.FillOpacity {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "FillOpacity"))
	}
	if spiralcircle.Stroke != spiralcircleOther.Stroke {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Stroke"))
	}
	if spiralcircle.StrokeOpacity != spiralcircleOther.StrokeOpacity {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "StrokeOpacity"))
	}
	if spiralcircle.StrokeWidth != spiralcircleOther.StrokeWidth {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "StrokeWidth"))
	}
	if spiralcircle.StrokeDashArray != spiralcircleOther.StrokeDashArray {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "StrokeDashArray"))
	}
	if spiralcircle.StrokeDashArrayWhenSelected != spiralcircleOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if spiralcircle.Transform != spiralcircleOther.Transform {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Transform"))
	}
	if spiralcircle.Pitch != spiralcircleOther.Pitch {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Pitch"))
	}
	if spiralcircle.ShowName != spiralcircleOther.ShowName {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "ShowName"))
	}
	if spiralcircle.BeatNb != spiralcircleOther.BeatNb {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "BeatNb"))
	}
	if spiralcircle.Path != spiralcircleOther.Path {
		diffs = append(diffs, spiralcircle.GongMarshallField(stage, "Path"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralcirclegrid *SpiralCircleGrid) GongDiff(stage *Stage, spiralcirclegridOther *SpiralCircleGrid) (diffs []string) {
	// insertion point for field diffs
	if spiralcirclegrid.Name != spiralcirclegridOther.Name {
		diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "Name"))
	}
	if spiralcirclegrid.IsDisplayed != spiralcirclegridOther.IsDisplayed {
		diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralcirclegrid.ShapeCategory == nil) != (spiralcirclegridOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralcirclegrid.ShapeCategory != nil && spiralcirclegridOther.ShapeCategory != nil {
		if spiralcirclegrid.ShapeCategory != spiralcirclegridOther.ShapeCategory {
			diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if (spiralcirclegrid.SpiralRhombusGrid == nil) != (spiralcirclegridOther.SpiralRhombusGrid == nil) {
		diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid"))
	} else if spiralcirclegrid.SpiralRhombusGrid != nil && spiralcirclegridOther.SpiralRhombusGrid != nil {
		if spiralcirclegrid.SpiralRhombusGrid != spiralcirclegridOther.SpiralRhombusGrid {
			diffs = append(diffs, spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid"))
		}
	}
	SpiralCirclesDifferent := false
	if len(spiralcirclegrid.SpiralCircles) != len(spiralcirclegridOther.SpiralCircles) {
		SpiralCirclesDifferent = true
	} else {
		for i := range spiralcirclegrid.SpiralCircles {
			if (spiralcirclegrid.SpiralCircles[i] == nil) != (spiralcirclegridOther.SpiralCircles[i] == nil) {
				SpiralCirclesDifferent = true
				break
			} else if spiralcirclegrid.SpiralCircles[i] != nil && spiralcirclegridOther.SpiralCircles[i] != nil {
				// this is a pointer comparaison
				if spiralcirclegrid.SpiralCircles[i] != spiralcirclegridOther.SpiralCircles[i] {
					SpiralCirclesDifferent = true
					break
				}
			}
		}
	}
	if SpiralCirclesDifferent {
		ops := Diff(stage, spiralcirclegrid, spiralcirclegridOther, "SpiralCircles", spiralcirclegridOther.SpiralCircles, spiralcirclegrid.SpiralCircles)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralline *SpiralLine) GongDiff(stage *Stage, spirallineOther *SpiralLine) (diffs []string) {
	// insertion point for field diffs
	if spiralline.Name != spirallineOther.Name {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "Name"))
	}
	if spiralline.IsDisplayed != spirallineOther.IsDisplayed {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralline.ShapeCategory == nil) != (spirallineOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralline.ShapeCategory != nil && spirallineOther.ShapeCategory != nil {
		if spiralline.ShapeCategory != spirallineOther.ShapeCategory {
			diffs = append(diffs, spiralline.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if spiralline.StartX != spirallineOther.StartX {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StartX"))
	}
	if spiralline.EndX != spirallineOther.EndX {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "EndX"))
	}
	if spiralline.StartY != spirallineOther.StartY {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StartY"))
	}
	if spiralline.EndY != spirallineOther.EndY {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "EndY"))
	}
	if spiralline.Color != spirallineOther.Color {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "Color"))
	}
	if spiralline.FillOpacity != spirallineOther.FillOpacity {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "FillOpacity"))
	}
	if spiralline.Stroke != spirallineOther.Stroke {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "Stroke"))
	}
	if spiralline.StrokeOpacity != spirallineOther.StrokeOpacity {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StrokeOpacity"))
	}
	if spiralline.StrokeWidth != spirallineOther.StrokeWidth {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StrokeWidth"))
	}
	if spiralline.StrokeDashArray != spirallineOther.StrokeDashArray {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StrokeDashArray"))
	}
	if spiralline.StrokeDashArrayWhenSelected != spirallineOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if spiralline.Transform != spirallineOther.Transform {
		diffs = append(diffs, spiralline.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spirallinegrid *SpiralLineGrid) GongDiff(stage *Stage, spirallinegridOther *SpiralLineGrid) (diffs []string) {
	// insertion point for field diffs
	if spirallinegrid.Name != spirallinegridOther.Name {
		diffs = append(diffs, spirallinegrid.GongMarshallField(stage, "Name"))
	}
	if spirallinegrid.IsDisplayed != spirallinegridOther.IsDisplayed {
		diffs = append(diffs, spirallinegrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spirallinegrid.ShapeCategory == nil) != (spirallinegridOther.ShapeCategory == nil) {
		diffs = append(diffs, spirallinegrid.GongMarshallField(stage, "ShapeCategory"))
	} else if spirallinegrid.ShapeCategory != nil && spirallinegridOther.ShapeCategory != nil {
		if spirallinegrid.ShapeCategory != spirallinegridOther.ShapeCategory {
			diffs = append(diffs, spirallinegrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	SpiralLinesDifferent := false
	if len(spirallinegrid.SpiralLines) != len(spirallinegridOther.SpiralLines) {
		SpiralLinesDifferent = true
	} else {
		for i := range spirallinegrid.SpiralLines {
			if (spirallinegrid.SpiralLines[i] == nil) != (spirallinegridOther.SpiralLines[i] == nil) {
				SpiralLinesDifferent = true
				break
			} else if spirallinegrid.SpiralLines[i] != nil && spirallinegridOther.SpiralLines[i] != nil {
				// this is a pointer comparaison
				if spirallinegrid.SpiralLines[i] != spirallinegridOther.SpiralLines[i] {
					SpiralLinesDifferent = true
					break
				}
			}
		}
	}
	if SpiralLinesDifferent {
		ops := Diff(stage, spirallinegrid, spirallinegridOther, "SpiralLines", spirallinegridOther.SpiralLines, spirallinegrid.SpiralLines)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralorigin *SpiralOrigin) GongDiff(stage *Stage, spiraloriginOther *SpiralOrigin) (diffs []string) {
	// insertion point for field diffs
	if spiralorigin.Name != spiraloriginOther.Name {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "Name"))
	}
	if spiralorigin.IsDisplayed != spiraloriginOther.IsDisplayed {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralorigin.ShapeCategory == nil) != (spiraloriginOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralorigin.ShapeCategory != nil && spiraloriginOther.ShapeCategory != nil {
		if spiralorigin.ShapeCategory != spiraloriginOther.ShapeCategory {
			diffs = append(diffs, spiralorigin.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if spiralorigin.Color != spiraloriginOther.Color {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "Color"))
	}
	if spiralorigin.FillOpacity != spiraloriginOther.FillOpacity {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "FillOpacity"))
	}
	if spiralorigin.Stroke != spiraloriginOther.Stroke {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "Stroke"))
	}
	if spiralorigin.StrokeOpacity != spiraloriginOther.StrokeOpacity {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "StrokeOpacity"))
	}
	if spiralorigin.StrokeWidth != spiraloriginOther.StrokeWidth {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "StrokeWidth"))
	}
	if spiralorigin.StrokeDashArray != spiraloriginOther.StrokeDashArray {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "StrokeDashArray"))
	}
	if spiralorigin.StrokeDashArrayWhenSelected != spiraloriginOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if spiralorigin.Transform != spiraloriginOther.Transform {
		diffs = append(diffs, spiralorigin.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralrhombus *SpiralRhombus) GongDiff(stage *Stage, spiralrhombusOther *SpiralRhombus) (diffs []string) {
	// insertion point for field diffs
	if spiralrhombus.Name != spiralrhombusOther.Name {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Name"))
	}
	if spiralrhombus.IsDisplayed != spiralrhombusOther.IsDisplayed {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralrhombus.ShapeCategory == nil) != (spiralrhombusOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralrhombus.ShapeCategory != nil && spiralrhombusOther.ShapeCategory != nil {
		if spiralrhombus.ShapeCategory != spiralrhombusOther.ShapeCategory {
			diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if spiralrhombus.X_r0 != spiralrhombusOther.X_r0 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "X_r0"))
	}
	if spiralrhombus.Y_r0 != spiralrhombusOther.Y_r0 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Y_r0"))
	}
	if spiralrhombus.X_r1 != spiralrhombusOther.X_r1 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "X_r1"))
	}
	if spiralrhombus.Y_r1 != spiralrhombusOther.Y_r1 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Y_r1"))
	}
	if spiralrhombus.X_r2 != spiralrhombusOther.X_r2 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "X_r2"))
	}
	if spiralrhombus.Y_r2 != spiralrhombusOther.Y_r2 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Y_r2"))
	}
	if spiralrhombus.X_r3 != spiralrhombusOther.X_r3 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "X_r3"))
	}
	if spiralrhombus.Y_r3 != spiralrhombusOther.Y_r3 {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Y_r3"))
	}
	if spiralrhombus.Color != spiralrhombusOther.Color {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Color"))
	}
	if spiralrhombus.FillOpacity != spiralrhombusOther.FillOpacity {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "FillOpacity"))
	}
	if spiralrhombus.Stroke != spiralrhombusOther.Stroke {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Stroke"))
	}
	if spiralrhombus.StrokeOpacity != spiralrhombusOther.StrokeOpacity {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "StrokeOpacity"))
	}
	if spiralrhombus.StrokeWidth != spiralrhombusOther.StrokeWidth {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "StrokeWidth"))
	}
	if spiralrhombus.StrokeDashArray != spiralrhombusOther.StrokeDashArray {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "StrokeDashArray"))
	}
	if spiralrhombus.StrokeDashArrayWhenSelected != spiralrhombusOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if spiralrhombus.Transform != spiralrhombusOther.Transform {
		diffs = append(diffs, spiralrhombus.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spiralrhombusgrid *SpiralRhombusGrid) GongDiff(stage *Stage, spiralrhombusgridOther *SpiralRhombusGrid) (diffs []string) {
	// insertion point for field diffs
	if spiralrhombusgrid.Name != spiralrhombusgridOther.Name {
		diffs = append(diffs, spiralrhombusgrid.GongMarshallField(stage, "Name"))
	}
	if spiralrhombusgrid.IsDisplayed != spiralrhombusgridOther.IsDisplayed {
		diffs = append(diffs, spiralrhombusgrid.GongMarshallField(stage, "IsDisplayed"))
	}
	if (spiralrhombusgrid.ShapeCategory == nil) != (spiralrhombusgridOther.ShapeCategory == nil) {
		diffs = append(diffs, spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory"))
	} else if spiralrhombusgrid.ShapeCategory != nil && spiralrhombusgridOther.ShapeCategory != nil {
		if spiralrhombusgrid.ShapeCategory != spiralrhombusgridOther.ShapeCategory {
			diffs = append(diffs, spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	SpiralRhombusesDifferent := false
	if len(spiralrhombusgrid.SpiralRhombuses) != len(spiralrhombusgridOther.SpiralRhombuses) {
		SpiralRhombusesDifferent = true
	} else {
		for i := range spiralrhombusgrid.SpiralRhombuses {
			if (spiralrhombusgrid.SpiralRhombuses[i] == nil) != (spiralrhombusgridOther.SpiralRhombuses[i] == nil) {
				SpiralRhombusesDifferent = true
				break
			} else if spiralrhombusgrid.SpiralRhombuses[i] != nil && spiralrhombusgridOther.SpiralRhombuses[i] != nil {
				// this is a pointer comparaison
				if spiralrhombusgrid.SpiralRhombuses[i] != spiralrhombusgridOther.SpiralRhombuses[i] {
					SpiralRhombusesDifferent = true
					break
				}
			}
		}
	}
	if SpiralRhombusesDifferent {
		ops := Diff(stage, spiralrhombusgrid, spiralrhombusgridOther, "SpiralRhombuses", spiralrhombusgridOther.SpiralRhombuses, spiralrhombusgrid.SpiralRhombuses)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalaxis *VerticalAxis) GongDiff(stage *Stage, verticalaxisOther *VerticalAxis) (diffs []string) {
	// insertion point for field diffs
	if verticalaxis.Name != verticalaxisOther.Name {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "Name"))
	}
	if verticalaxis.IsDisplayed != verticalaxisOther.IsDisplayed {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "IsDisplayed"))
	}
	if (verticalaxis.ShapeCategory == nil) != (verticalaxisOther.ShapeCategory == nil) {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "ShapeCategory"))
	} else if verticalaxis.ShapeCategory != nil && verticalaxisOther.ShapeCategory != nil {
		if verticalaxis.ShapeCategory != verticalaxisOther.ShapeCategory {
			diffs = append(diffs, verticalaxis.GongMarshallField(stage, "ShapeCategory"))
		}
	}
	if verticalaxis.AxisHandleBorderLength != verticalaxisOther.AxisHandleBorderLength {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
	}
	if verticalaxis.Axis_Length != verticalaxisOther.Axis_Length {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "Axis_Length"))
	}
	if verticalaxis.Color != verticalaxisOther.Color {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "Color"))
	}
	if verticalaxis.FillOpacity != verticalaxisOther.FillOpacity {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "FillOpacity"))
	}
	if verticalaxis.Stroke != verticalaxisOther.Stroke {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "Stroke"))
	}
	if verticalaxis.StrokeOpacity != verticalaxisOther.StrokeOpacity {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "StrokeOpacity"))
	}
	if verticalaxis.StrokeWidth != verticalaxisOther.StrokeWidth {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "StrokeWidth"))
	}
	if verticalaxis.StrokeDashArray != verticalaxisOther.StrokeDashArray {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "StrokeDashArray"))
	}
	if verticalaxis.StrokeDashArrayWhenSelected != verticalaxisOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if verticalaxis.Transform != verticalaxisOther.Transform {
		diffs = append(diffs, verticalaxis.GongMarshallField(stage, "Transform"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
