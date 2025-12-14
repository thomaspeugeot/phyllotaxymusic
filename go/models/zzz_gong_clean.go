// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Axis
func (axis *Axis) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	axis.ShapeCategory = GongCleanPointer(stage, axis.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by AxisGrid
func (axisgrid *AxisGrid) GongClean(stage *Stage) {
	// insertion point per field
	axisgrid.Axiss = GongCleanSlice(stage, axisgrid.Axiss)
	// insertion point per field
	axisgrid.Reference = GongCleanPointer(stage, axisgrid.Reference)
	axisgrid.ShapeCategory = GongCleanPointer(stage, axisgrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by Bezier
func (bezier *Bezier) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	bezier.ShapeCategory = GongCleanPointer(stage, bezier.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by BezierGrid
func (beziergrid *BezierGrid) GongClean(stage *Stage) {
	// insertion point per field
	beziergrid.Beziers = GongCleanSlice(stage, beziergrid.Beziers)
	// insertion point per field
	beziergrid.Reference = GongCleanPointer(stage, beziergrid.Reference)
	beziergrid.ShapeCategory = GongCleanPointer(stage, beziergrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by BezierGridStack
func (beziergridstack *BezierGridStack) GongClean(stage *Stage) {
	// insertion point per field
	beziergridstack.BezierGrids = GongCleanSlice(stage, beziergridstack.BezierGrids)
	// insertion point per field
	beziergridstack.ShapeCategory = GongCleanPointer(stage, beziergridstack.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by Chapter
func (chapter *Chapter) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Circle
func (circle *Circle) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	circle.ShapeCategory = GongCleanPointer(stage, circle.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by CircleGrid
func (circlegrid *CircleGrid) GongClean(stage *Stage) {
	// insertion point per field
	circlegrid.Circles = GongCleanSlice(stage, circlegrid.Circles)
	// insertion point per field
	circlegrid.Reference = GongCleanPointer(stage, circlegrid.Reference)
	circlegrid.ShapeCategory = GongCleanPointer(stage, circlegrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by Content
func (content *Content) GongClean(stage *Stage) {
	// insertion point per field
	content.Chapters = GongCleanSlice(stage, content.Chapters)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ExportToMusicxml
func (exporttomusicxml *ExportToMusicxml) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	exporttomusicxml.Parameter = GongCleanPointer(stage, exporttomusicxml.Parameter)
}

// Clean garbage collect unstaged instances that are referenced by FrontCurve
func (frontcurve *FrontCurve) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FrontCurveStack
func (frontcurvestack *FrontCurveStack) GongClean(stage *Stage) {
	// insertion point per field
	frontcurvestack.FrontCurves = GongCleanSlice(stage, frontcurvestack.FrontCurves)
	frontcurvestack.SpiralCircles = GongCleanSlice(stage, frontcurvestack.SpiralCircles)
	// insertion point per field
	frontcurvestack.ShapeCategory = GongCleanPointer(stage, frontcurvestack.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by HorizontalAxis
func (horizontalaxis *HorizontalAxis) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	horizontalaxis.ShapeCategory = GongCleanPointer(stage, horizontalaxis.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by Key
func (key *Key) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	key.ShapeCategory = GongCleanPointer(stage, key.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by Parameter
func (parameter *Parameter) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	parameter.InitialRhombus = GongCleanPointer(stage, parameter.InitialRhombus)
	parameter.InitialCircle = GongCleanPointer(stage, parameter.InitialCircle)
	parameter.InitialRhombusGrid = GongCleanPointer(stage, parameter.InitialRhombusGrid)
	parameter.InitialCircleGrid = GongCleanPointer(stage, parameter.InitialCircleGrid)
	parameter.InitialAxis = GongCleanPointer(stage, parameter.InitialAxis)
	parameter.RotatedAxis = GongCleanPointer(stage, parameter.RotatedAxis)
	parameter.RotatedRhombus = GongCleanPointer(stage, parameter.RotatedRhombus)
	parameter.RotatedRhombusGrid = GongCleanPointer(stage, parameter.RotatedRhombusGrid)
	parameter.RotatedCircleGrid = GongCleanPointer(stage, parameter.RotatedCircleGrid)
	parameter.NextRhombus = GongCleanPointer(stage, parameter.NextRhombus)
	parameter.NextCircle = GongCleanPointer(stage, parameter.NextCircle)
	parameter.GrowingRhombusGridSeed = GongCleanPointer(stage, parameter.GrowingRhombusGridSeed)
	parameter.GrowingRhombusGrid = GongCleanPointer(stage, parameter.GrowingRhombusGrid)
	parameter.GrowingCircleGridSeed = GongCleanPointer(stage, parameter.GrowingCircleGridSeed)
	parameter.GrowingCircleGrid = GongCleanPointer(stage, parameter.GrowingCircleGrid)
	parameter.GrowingCircleGridLeftSeed = GongCleanPointer(stage, parameter.GrowingCircleGridLeftSeed)
	parameter.GrowingCircleGridLeft = GongCleanPointer(stage, parameter.GrowingCircleGridLeft)
	parameter.ConstructionAxis = GongCleanPointer(stage, parameter.ConstructionAxis)
	parameter.ConstructionAxisGrid = GongCleanPointer(stage, parameter.ConstructionAxisGrid)
	parameter.ConstructionCircle = GongCleanPointer(stage, parameter.ConstructionCircle)
	parameter.ConstructionCircleGrid = GongCleanPointer(stage, parameter.ConstructionCircleGrid)
	parameter.GrowthCurveSeed = GongCleanPointer(stage, parameter.GrowthCurveSeed)
	parameter.GrowthCurve = GongCleanPointer(stage, parameter.GrowthCurve)
	parameter.GrowthCurveShiftedRightSeed = GongCleanPointer(stage, parameter.GrowthCurveShiftedRightSeed)
	parameter.GrowthCurveShiftedRight = GongCleanPointer(stage, parameter.GrowthCurveShiftedRight)
	parameter.GrowthCurveNextSeed = GongCleanPointer(stage, parameter.GrowthCurveNextSeed)
	parameter.GrowthCurveNext = GongCleanPointer(stage, parameter.GrowthCurveNext)
	parameter.GrowthCurveNextShiftedRightSeed = GongCleanPointer(stage, parameter.GrowthCurveNextShiftedRightSeed)
	parameter.GrowthCurveNextShiftedRight = GongCleanPointer(stage, parameter.GrowthCurveNextShiftedRight)
	parameter.GrowthCurveStack = GongCleanPointer(stage, parameter.GrowthCurveStack)
	parameter.SpiralRhombusGridSeed = GongCleanPointer(stage, parameter.SpiralRhombusGridSeed)
	parameter.SpiralRhombusGrid = GongCleanPointer(stage, parameter.SpiralRhombusGrid)
	parameter.SpiralCircleSeed = GongCleanPointer(stage, parameter.SpiralCircleSeed)
	parameter.SpiralCircleGrid = GongCleanPointer(stage, parameter.SpiralCircleGrid)
	parameter.SpiralCircleFullGrid = GongCleanPointer(stage, parameter.SpiralCircleFullGrid)
	parameter.SpiralConstructionOuterLineSeed = GongCleanPointer(stage, parameter.SpiralConstructionOuterLineSeed)
	parameter.SpiralConstructionInnerLineSeed = GongCleanPointer(stage, parameter.SpiralConstructionInnerLineSeed)
	parameter.SpiralConstructionOuterLineGrid = GongCleanPointer(stage, parameter.SpiralConstructionOuterLineGrid)
	parameter.SpiralConstructionInnerLineGrid = GongCleanPointer(stage, parameter.SpiralConstructionInnerLineGrid)
	parameter.SpiralConstructionCircleGrid = GongCleanPointer(stage, parameter.SpiralConstructionCircleGrid)
	parameter.SpiralConstructionOuterLineFullGrid = GongCleanPointer(stage, parameter.SpiralConstructionOuterLineFullGrid)
	parameter.SpiralBezierSeed = GongCleanPointer(stage, parameter.SpiralBezierSeed)
	parameter.SpiralBezierGrid = GongCleanPointer(stage, parameter.SpiralBezierGrid)
	parameter.SpiralBezierFullGrid = GongCleanPointer(stage, parameter.SpiralBezierFullGrid)
	parameter.FrontCurveStack = GongCleanPointer(stage, parameter.FrontCurveStack)
	parameter.Fkey = GongCleanPointer(stage, parameter.Fkey)
	parameter.PitchLines = GongCleanPointer(stage, parameter.PitchLines)
	parameter.BeatLines = GongCleanPointer(stage, parameter.BeatLines)
	parameter.FirstVoice = GongCleanPointer(stage, parameter.FirstVoice)
	parameter.FirstVoiceShiftedRigth = GongCleanPointer(stage, parameter.FirstVoiceShiftedRigth)
	parameter.SecondVoice = GongCleanPointer(stage, parameter.SecondVoice)
	parameter.SecondVoiceShiftedRight = GongCleanPointer(stage, parameter.SecondVoiceShiftedRight)
	parameter.FirstVoiceNotes = GongCleanPointer(stage, parameter.FirstVoiceNotes)
	parameter.FirstVoiceNotesShiftedRight = GongCleanPointer(stage, parameter.FirstVoiceNotesShiftedRight)
	parameter.SecondVoiceNotes = GongCleanPointer(stage, parameter.SecondVoiceNotes)
	parameter.SecondVoiceNotesShiftedRight = GongCleanPointer(stage, parameter.SecondVoiceNotesShiftedRight)
	parameter.HorizontalAxis = GongCleanPointer(stage, parameter.HorizontalAxis)
	parameter.VerticalAxis = GongCleanPointer(stage, parameter.VerticalAxis)
	parameter.SpiralOrigin = GongCleanPointer(stage, parameter.SpiralOrigin)
}

// Clean garbage collect unstaged instances that are referenced by Rhombus
func (rhombus *Rhombus) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	rhombus.ShapeCategory = GongCleanPointer(stage, rhombus.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by RhombusGrid
func (rhombusgrid *RhombusGrid) GongClean(stage *Stage) {
	// insertion point per field
	rhombusgrid.Rhombuses = GongCleanSlice(stage, rhombusgrid.Rhombuses)
	// insertion point per field
	rhombusgrid.Reference = GongCleanPointer(stage, rhombusgrid.Reference)
	rhombusgrid.ShapeCategory = GongCleanPointer(stage, rhombusgrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by ShapeCategory
func (shapecategory *ShapeCategory) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by SpiralBezier
func (spiralbezier *SpiralBezier) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spiralbezier.ShapeCategory = GongCleanPointer(stage, spiralbezier.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralBezierGrid
func (spiralbeziergrid *SpiralBezierGrid) GongClean(stage *Stage) {
	// insertion point per field
	spiralbeziergrid.SpiralBeziers = GongCleanSlice(stage, spiralbeziergrid.SpiralBeziers)
	// insertion point per field
	spiralbeziergrid.ShapeCategory = GongCleanPointer(stage, spiralbeziergrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralCircle
func (spiralcircle *SpiralCircle) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spiralcircle.ShapeCategory = GongCleanPointer(stage, spiralcircle.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralCircleGrid
func (spiralcirclegrid *SpiralCircleGrid) GongClean(stage *Stage) {
	// insertion point per field
	spiralcirclegrid.SpiralCircles = GongCleanSlice(stage, spiralcirclegrid.SpiralCircles)
	// insertion point per field
	spiralcirclegrid.ShapeCategory = GongCleanPointer(stage, spiralcirclegrid.ShapeCategory)
	spiralcirclegrid.SpiralRhombusGrid = GongCleanPointer(stage, spiralcirclegrid.SpiralRhombusGrid)
}

// Clean garbage collect unstaged instances that are referenced by SpiralLine
func (spiralline *SpiralLine) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spiralline.ShapeCategory = GongCleanPointer(stage, spiralline.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralLineGrid
func (spirallinegrid *SpiralLineGrid) GongClean(stage *Stage) {
	// insertion point per field
	spirallinegrid.SpiralLines = GongCleanSlice(stage, spirallinegrid.SpiralLines)
	// insertion point per field
	spirallinegrid.ShapeCategory = GongCleanPointer(stage, spirallinegrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralOrigin
func (spiralorigin *SpiralOrigin) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spiralorigin.ShapeCategory = GongCleanPointer(stage, spiralorigin.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralRhombus
func (spiralrhombus *SpiralRhombus) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spiralrhombus.ShapeCategory = GongCleanPointer(stage, spiralrhombus.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by SpiralRhombusGrid
func (spiralrhombusgrid *SpiralRhombusGrid) GongClean(stage *Stage) {
	// insertion point per field
	spiralrhombusgrid.SpiralRhombuses = GongCleanSlice(stage, spiralrhombusgrid.SpiralRhombuses)
	// insertion point per field
	spiralrhombusgrid.ShapeCategory = GongCleanPointer(stage, spiralrhombusgrid.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by VerticalAxis
func (verticalaxis *VerticalAxis) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	verticalaxis.ShapeCategory = GongCleanPointer(stage, verticalaxis.ShapeCategory)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
