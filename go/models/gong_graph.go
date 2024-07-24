// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Axis:
		ok = stage.IsStagedAxis(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *CircleGrid:
		ok = stage.IsStagedCircleGrid(target)

	case *HorizontalAxis:
		ok = stage.IsStagedHorizontalAxis(target)

	case *Line:
		ok = stage.IsStagedLine(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *Rhombus:
		ok = stage.IsStagedRhombus(target)

	case *RhombusGrid:
		ok = stage.IsStagedRhombusGrid(target)

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

func (stage *StageStruct) IsStagedLine(line *Line) (ok bool) {

	_, ok = stage.Lines[line]

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

	case *Circle:
		stage.StageBranchCircle(target)

	case *CircleGrid:
		stage.StageBranchCircleGrid(target)

	case *HorizontalAxis:
		stage.StageBranchHorizontalAxis(target)

	case *Line:
		stage.StageBranchLine(target)

	case *Parameter:
		stage.StageBranchParameter(target)

	case *Rhombus:
		stage.StageBranchRhombus(target)

	case *RhombusGrid:
		stage.StageBranchRhombusGrid(target)

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if IsStaged(stage, circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLine(line *Line) {

	// check if instance is already staged
	if IsStaged(stage, line) {
		return
	}

	line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgrid.Rhombuses {
		StageBranch(stage, _rhombus)
	}

}

func (stage *StageStruct) StageBranchVerticalAxis(verticalaxis *VerticalAxis) {

	// check if instance is already staged
	if IsStaged(stage, verticalaxis) {
		return
	}

	verticalaxis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	case *Circle:
		toT := CopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGrid:
		toT := CopyBranchCircleGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HorizontalAxis:
		toT := CopyBranchHorizontalAxis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line:
		toT := CopyBranchLine(mapOrigCopy, fromT)
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

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLine(mapOrigCopy map[any]any, lineFrom *Line) (lineTo *Line) {

	// lineFrom has already been copied
	if _lineTo, ok := mapOrigCopy[lineFrom]; ok {
		lineTo = _lineTo.(*Line)
		return
	}

	lineTo = new(Line)
	mapOrigCopy[lineFrom] = lineTo
	lineFrom.CopyBasicFields(lineTo)

	//insertion point for the staging of instances referenced by pointers

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgridFrom.Rhombuses {
		rhombusgridTo.Rhombuses = append(rhombusgridTo.Rhombuses, CopyBranchRhombus(mapOrigCopy, _rhombus))
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

	case *Circle:
		stage.UnstageBranchCircle(target)

	case *CircleGrid:
		stage.UnstageBranchCircleGrid(target)

	case *HorizontalAxis:
		stage.UnstageBranchHorizontalAxis(target)

	case *Line:
		stage.UnstageBranchLine(target)

	case *Parameter:
		stage.UnstageBranchParameter(target)

	case *Rhombus:
		stage.UnstageBranchRhombus(target)

	case *RhombusGrid:
		stage.UnstageBranchRhombusGrid(target)

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if !IsStaged(stage, circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLine(line *Line) {

	// check if instance is already staged
	if !IsStaged(stage, line) {
		return
	}

	line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombus := range rhombusgrid.Rhombuses {
		UnstageBranch(stage, _rhombus)
	}

}

func (stage *StageStruct) UnstageBranchVerticalAxis(verticalaxis *VerticalAxis) {

	// check if instance is already staged
	if !IsStaged(stage, verticalaxis) {
		return
	}

	verticalaxis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
