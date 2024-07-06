// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *HorizontalAxis:
		ok = stage.IsStagedHorizontalAxis(target)

	case *Line:
		ok = stage.IsStagedLine(target)

	case *Parameter:
		ok = stage.IsStagedParameter(target)

	case *Rhombus:
		ok = stage.IsStagedRhombus(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
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

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *HorizontalAxis:
		stage.StageBranchHorizontalAxis(target)

	case *Line:
		stage.StageBranchLine(target)

	case *Parameter:
		stage.StageBranchParameter(target)

	case *Rhombus:
		stage.StageBranchRhombus(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	if parameter.HorizontalAxis != nil {
		StageBranch(stage, parameter.HorizontalAxis)
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

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
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

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	if parameterFrom.HorizontalAxis != nil {
		parameterTo.HorizontalAxis = CopyBranchHorizontalAxis(mapOrigCopy, parameterFrom.HorizontalAxis)
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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *HorizontalAxis:
		stage.UnstageBranchHorizontalAxis(target)

	case *Line:
		stage.UnstageBranchLine(target)

	case *Parameter:
		stage.UnstageBranchParameter(target)

	case *Rhombus:
		stage.UnstageBranchRhombus(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	if parameter.HorizontalAxis != nil {
		UnstageBranch(stage, parameter.HorizontalAxis)
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
