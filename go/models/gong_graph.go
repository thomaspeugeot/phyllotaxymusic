// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

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

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *CircleGrid:
		ok = stage.IsStagedCircleGrid(target)

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

	case *VerticalAxis:
		ok = stage.IsStagedVerticalAxis(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAxis(axis *Axis) (ok bool) {

	_, ok = stage.Axiss[axis]

	return
}

func (stage *StageStruct) IsStagedAxisGrid(axisgrid *AxisGrid) (ok bool) {

	_, ok = stage.AxisGrids[axisgrid]

	return
}

func (stage *StageStruct) IsStagedBezier(bezier *Bezier) (ok bool) {

	_, ok = stage.Beziers[bezier]

	return
}

func (stage *StageStruct) IsStagedBezierGrid(beziergrid *BezierGrid) (ok bool) {

	_, ok = stage.BezierGrids[beziergrid]

	return
}

func (stage *StageStruct) IsStagedBezierGridStack(beziergridstack *BezierGridStack) (ok bool) {

	_, ok = stage.BezierGridStacks[beziergridstack]

	return
}

func (stage *StageStruct) IsStagedCircle(circle *Circle) (ok bool) {

	_, ok = stage.Circles[circle]

	return
}

func (stage *StageStruct) IsStagedCircleGrid(circlegrid *CircleGrid) (ok bool) {

	_, ok = stage.CircleGrids[circlegrid]

	return
}

func (stage *StageStruct) IsStagedHorizontalAxis(horizontalaxis *HorizontalAxis) (ok bool) {

	_, ok = stage.HorizontalAxiss[horizontalaxis]

	return
}

func (stage *StageStruct) IsStagedKey(key *Key) (ok bool) {

	_, ok = stage.Keys[key]

	return
}

func (stage *StageStruct) IsStagedParameter(parameter *Parameter) (ok bool) {

	_, ok = stage.Parameters[parameter]

	return
}

func (stage *StageStruct) IsStagedRhombus(rhombus *Rhombus) (ok bool) {

	_, ok = stage.Rhombuss[rhombus]

	return
}

func (stage *StageStruct) IsStagedRhombusGrid(rhombusgrid *RhombusGrid) (ok bool) {

	_, ok = stage.RhombusGrids[rhombusgrid]

	return
}

func (stage *StageStruct) IsStagedShapeCategory(shapecategory *ShapeCategory) (ok bool) {

	_, ok = stage.ShapeCategorys[shapecategory]

	return
}

func (stage *StageStruct) IsStagedVerticalAxis(verticalaxis *VerticalAxis) (ok bool) {

	_, ok = stage.VerticalAxiss[verticalaxis]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Circle:
		stage.StageBranchCircle(target)

	case *CircleGrid:
		stage.StageBranchCircleGrid(target)

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

	case *VerticalAxis:
		stage.StageBranchVerticalAxis(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAxis(axis *Axis) {

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

func (stage *StageStruct) StageBranchAxisGrid(axisgrid *AxisGrid) {

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

func (stage *StageStruct) StageBranchBezier(bezier *Bezier) {

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

func (stage *StageStruct) StageBranchBezierGrid(beziergrid *BezierGrid) {

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

func (stage *StageStruct) StageBranchBezierGridStack(beziergridstack *BezierGridStack) {

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

func (stage *StageStruct) StageBranchCircle(circle *Circle) {

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

func (stage *StageStruct) StageBranchCircleGrid(circlegrid *CircleGrid) {

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

func (stage *StageStruct) StageBranchHorizontalAxis(horizontalaxis *HorizontalAxis) {

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

func (stage *StageStruct) StageBranchKey(key *Key) {

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

func (stage *StageStruct) StageBranchParameter(parameter *Parameter) {

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
	if parameter.GrowthCurveSegment != nil {
		StageBranch(stage, parameter.GrowthCurveSegment)
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
	if parameter.Fkey != nil {
		StageBranch(stage, parameter.Fkey)
	}
	if parameter.HorizontalAxis != nil {
		StageBranch(stage, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis != nil {
		StageBranch(stage, parameter.VerticalAxis)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRhombus(rhombus *Rhombus) {

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

func (stage *StageStruct) StageBranchRhombusGrid(rhombusgrid *RhombusGrid) {

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

func (stage *StageStruct) StageBranchShapeCategory(shapecategory *ShapeCategory) {

	// check if instance is already staged
	if IsStaged(stage, shapecategory) {
		return
	}

	shapecategory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchVerticalAxis(verticalaxis *VerticalAxis) {

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

	case *Circle:
		toT := CopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGrid:
		toT := CopyBranchCircleGrid(mapOrigCopy, fromT)
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
	if parameterFrom.GrowthCurveSegment != nil {
		parameterTo.GrowthCurveSegment = CopyBranchBezier(mapOrigCopy, parameterFrom.GrowthCurveSegment)
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
	if parameterFrom.Fkey != nil {
		parameterTo.Fkey = CopyBranchKey(mapOrigCopy, parameterFrom.Fkey)
	}
	if parameterFrom.HorizontalAxis != nil {
		parameterTo.HorizontalAxis = CopyBranchHorizontalAxis(mapOrigCopy, parameterFrom.HorizontalAxis)
	}
	if parameterFrom.VerticalAxis != nil {
		parameterTo.VerticalAxis = CopyBranchVerticalAxis(mapOrigCopy, parameterFrom.VerticalAxis)
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
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Circle:
		stage.UnstageBranchCircle(target)

	case *CircleGrid:
		stage.UnstageBranchCircleGrid(target)

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

	case *VerticalAxis:
		stage.UnstageBranchVerticalAxis(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAxis(axis *Axis) {

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

func (stage *StageStruct) UnstageBranchAxisGrid(axisgrid *AxisGrid) {

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

func (stage *StageStruct) UnstageBranchBezier(bezier *Bezier) {

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

func (stage *StageStruct) UnstageBranchBezierGrid(beziergrid *BezierGrid) {

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

func (stage *StageStruct) UnstageBranchBezierGridStack(beziergridstack *BezierGridStack) {

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

func (stage *StageStruct) UnstageBranchCircle(circle *Circle) {

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

func (stage *StageStruct) UnstageBranchCircleGrid(circlegrid *CircleGrid) {

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

func (stage *StageStruct) UnstageBranchHorizontalAxis(horizontalaxis *HorizontalAxis) {

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

func (stage *StageStruct) UnstageBranchKey(key *Key) {

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

func (stage *StageStruct) UnstageBranchParameter(parameter *Parameter) {

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
	if parameter.GrowthCurveSegment != nil {
		UnstageBranch(stage, parameter.GrowthCurveSegment)
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
	if parameter.Fkey != nil {
		UnstageBranch(stage, parameter.Fkey)
	}
	if parameter.HorizontalAxis != nil {
		UnstageBranch(stage, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis != nil {
		UnstageBranch(stage, parameter.VerticalAxis)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRhombus(rhombus *Rhombus) {

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

func (stage *StageStruct) UnstageBranchRhombusGrid(rhombusgrid *RhombusGrid) {

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

func (stage *StageStruct) UnstageBranchShapeCategory(shapecategory *ShapeCategory) {

	// check if instance is already staged
	if !IsStaged(stage, shapecategory) {
		return
	}

	shapecategory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchVerticalAxis(verticalaxis *VerticalAxis) {

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
