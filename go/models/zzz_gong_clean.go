// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Axis
func (axis *Axis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &axis.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AxisGrid
func (axisgrid *AxisGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &axisgrid.Axiss) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &axisgrid.Reference) || modified
	modified = GongCleanPointer(stage, &axisgrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bezier
func (bezier *Bezier) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &bezier.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BezierGrid
func (beziergrid *BezierGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &beziergrid.Beziers) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &beziergrid.Reference) || modified
	modified = GongCleanPointer(stage, &beziergrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BezierGridStack
func (beziergridstack *BezierGridStack) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &beziergridstack.BezierGrids) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &beziergridstack.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Chapter
func (chapter *Chapter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Circle
func (circle *Circle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &circle.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleGrid
func (circlegrid *CircleGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &circlegrid.Circles) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &circlegrid.Reference) || modified
	modified = GongCleanPointer(stage, &circlegrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Content
func (content *Content) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &content.Chapters) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExportToMusicxml
func (exporttomusicxml *ExportToMusicxml) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &exporttomusicxml.Parameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FrontCurve
func (frontcurve *FrontCurve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FrontCurveStack
func (frontcurvestack *FrontCurveStack) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &frontcurvestack.FrontCurves) || modified
	modified = GongCleanSlice(stage, &frontcurvestack.SpiralCircles) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &frontcurvestack.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HorizontalAxis
func (horizontalaxis *HorizontalAxis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &horizontalaxis.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Key
func (key *Key) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &key.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Parameter
func (parameter *Parameter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &parameter.InitialRhombus) || modified
	modified = GongCleanPointer(stage, &parameter.InitialCircle) || modified
	modified = GongCleanPointer(stage, &parameter.InitialRhombusGrid) || modified
	modified = GongCleanPointer(stage, &parameter.InitialCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.InitialAxis) || modified
	modified = GongCleanPointer(stage, &parameter.RotatedAxis) || modified
	modified = GongCleanPointer(stage, &parameter.RotatedRhombus) || modified
	modified = GongCleanPointer(stage, &parameter.RotatedRhombusGrid) || modified
	modified = GongCleanPointer(stage, &parameter.RotatedCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.NextRhombus) || modified
	modified = GongCleanPointer(stage, &parameter.NextCircle) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingRhombusGridSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingRhombusGrid) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingCircleGridSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingCircleGridLeftSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowingCircleGridLeft) || modified
	modified = GongCleanPointer(stage, &parameter.ConstructionAxis) || modified
	modified = GongCleanPointer(stage, &parameter.ConstructionAxisGrid) || modified
	modified = GongCleanPointer(stage, &parameter.ConstructionCircle) || modified
	modified = GongCleanPointer(stage, &parameter.ConstructionCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurve) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveShiftedRightSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveShiftedRight) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveNextSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveNext) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveNextShiftedRightSeed) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveNextShiftedRight) || modified
	modified = GongCleanPointer(stage, &parameter.GrowthCurveStack) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralRhombusGridSeed) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralRhombusGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralCircleSeed) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralCircleFullGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionOuterLineSeed) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionInnerLineSeed) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionOuterLineGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionInnerLineGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionCircleGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralConstructionOuterLineFullGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralBezierSeed) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralBezierGrid) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralBezierFullGrid) || modified
	modified = GongCleanPointer(stage, &parameter.FrontCurveStack) || modified
	modified = GongCleanPointer(stage, &parameter.Fkey) || modified
	modified = GongCleanPointer(stage, &parameter.PitchLines) || modified
	modified = GongCleanPointer(stage, &parameter.BeatLines) || modified
	modified = GongCleanPointer(stage, &parameter.FirstVoice) || modified
	modified = GongCleanPointer(stage, &parameter.FirstVoiceShiftedRigth) || modified
	modified = GongCleanPointer(stage, &parameter.SecondVoice) || modified
	modified = GongCleanPointer(stage, &parameter.SecondVoiceShiftedRight) || modified
	modified = GongCleanPointer(stage, &parameter.FirstVoiceNotes) || modified
	modified = GongCleanPointer(stage, &parameter.FirstVoiceNotesShiftedRight) || modified
	modified = GongCleanPointer(stage, &parameter.SecondVoiceNotes) || modified
	modified = GongCleanPointer(stage, &parameter.SecondVoiceNotesShiftedRight) || modified
	modified = GongCleanPointer(stage, &parameter.HorizontalAxis) || modified
	modified = GongCleanPointer(stage, &parameter.VerticalAxis) || modified
	modified = GongCleanPointer(stage, &parameter.SpiralOrigin) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Rhombus
func (rhombus *Rhombus) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &rhombus.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusGrid
func (rhombusgrid *RhombusGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &rhombusgrid.Rhombuses) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &rhombusgrid.Reference) || modified
	modified = GongCleanPointer(stage, &rhombusgrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ShapeCategory
func (shapecategory *ShapeCategory) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralBezier
func (spiralbezier *SpiralBezier) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralbezier.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralBezierGrid
func (spiralbeziergrid *SpiralBezierGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &spiralbeziergrid.SpiralBeziers) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralbeziergrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralCircle
func (spiralcircle *SpiralCircle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralcircle.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralCircleGrid
func (spiralcirclegrid *SpiralCircleGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &spiralcirclegrid.SpiralCircles) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralcirclegrid.ShapeCategory) || modified
	modified = GongCleanPointer(stage, &spiralcirclegrid.SpiralRhombusGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralLine
func (spiralline *SpiralLine) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralline.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralLineGrid
func (spirallinegrid *SpiralLineGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &spirallinegrid.SpiralLines) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &spirallinegrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralOrigin
func (spiralorigin *SpiralOrigin) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralorigin.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralRhombus
func (spiralrhombus *SpiralRhombus) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralrhombus.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpiralRhombusGrid
func (spiralrhombusgrid *SpiralRhombusGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &spiralrhombusgrid.SpiralRhombuses) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &spiralrhombusgrid.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalAxis
func (verticalaxis *VerticalAxis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &verticalaxis.ShapeCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
