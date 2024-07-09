// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Axis:
		if stage.OnAfterAxisCreateCallback != nil {
			stage.OnAfterAxisCreateCallback.OnAfterCreate(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleCreateCallback != nil {
			stage.OnAfterCircleCreateCallback.OnAfterCreate(stage, target)
		}
	case *CircleGrid:
		if stage.OnAfterCircleGridCreateCallback != nil {
			stage.OnAfterCircleGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *HorizontalAxis:
		if stage.OnAfterHorizontalAxisCreateCallback != nil {
			stage.OnAfterHorizontalAxisCreateCallback.OnAfterCreate(stage, target)
		}
	case *Line:
		if stage.OnAfterLineCreateCallback != nil {
			stage.OnAfterLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterCreateCallback != nil {
			stage.OnAfterParameterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rhombus:
		if stage.OnAfterRhombusCreateCallback != nil {
			stage.OnAfterRhombusCreateCallback.OnAfterCreate(stage, target)
		}
	case *RhombusGrid:
		if stage.OnAfterRhombusGridCreateCallback != nil {
			stage.OnAfterRhombusGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *VerticalAxis:
		if stage.OnAfterVerticalAxisCreateCallback != nil {
			stage.OnAfterVerticalAxisCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Axis:
		newTarget := any(new).(*Axis)
		if stage.OnAfterAxisUpdateCallback != nil {
			stage.OnAfterAxisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Circle:
		newTarget := any(new).(*Circle)
		if stage.OnAfterCircleUpdateCallback != nil {
			stage.OnAfterCircleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CircleGrid:
		newTarget := any(new).(*CircleGrid)
		if stage.OnAfterCircleGridUpdateCallback != nil {
			stage.OnAfterCircleGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *HorizontalAxis:
		newTarget := any(new).(*HorizontalAxis)
		if stage.OnAfterHorizontalAxisUpdateCallback != nil {
			stage.OnAfterHorizontalAxisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Line:
		newTarget := any(new).(*Line)
		if stage.OnAfterLineUpdateCallback != nil {
			stage.OnAfterLineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Parameter:
		newTarget := any(new).(*Parameter)
		if stage.OnAfterParameterUpdateCallback != nil {
			stage.OnAfterParameterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rhombus:
		newTarget := any(new).(*Rhombus)
		if stage.OnAfterRhombusUpdateCallback != nil {
			stage.OnAfterRhombusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RhombusGrid:
		newTarget := any(new).(*RhombusGrid)
		if stage.OnAfterRhombusGridUpdateCallback != nil {
			stage.OnAfterRhombusGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VerticalAxis:
		newTarget := any(new).(*VerticalAxis)
		if stage.OnAfterVerticalAxisUpdateCallback != nil {
			stage.OnAfterVerticalAxisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Axis:
		if stage.OnAfterAxisDeleteCallback != nil {
			staged := any(staged).(*Axis)
			stage.OnAfterAxisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Circle:
		if stage.OnAfterCircleDeleteCallback != nil {
			staged := any(staged).(*Circle)
			stage.OnAfterCircleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CircleGrid:
		if stage.OnAfterCircleGridDeleteCallback != nil {
			staged := any(staged).(*CircleGrid)
			stage.OnAfterCircleGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *HorizontalAxis:
		if stage.OnAfterHorizontalAxisDeleteCallback != nil {
			staged := any(staged).(*HorizontalAxis)
			stage.OnAfterHorizontalAxisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Line:
		if stage.OnAfterLineDeleteCallback != nil {
			staged := any(staged).(*Line)
			stage.OnAfterLineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Parameter:
		if stage.OnAfterParameterDeleteCallback != nil {
			staged := any(staged).(*Parameter)
			stage.OnAfterParameterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rhombus:
		if stage.OnAfterRhombusDeleteCallback != nil {
			staged := any(staged).(*Rhombus)
			stage.OnAfterRhombusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RhombusGrid:
		if stage.OnAfterRhombusGridDeleteCallback != nil {
			staged := any(staged).(*RhombusGrid)
			stage.OnAfterRhombusGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VerticalAxis:
		if stage.OnAfterVerticalAxisDeleteCallback != nil {
			staged := any(staged).(*VerticalAxis)
			stage.OnAfterVerticalAxisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Axis:
		if stage.OnAfterAxisReadCallback != nil {
			stage.OnAfterAxisReadCallback.OnAfterRead(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleReadCallback != nil {
			stage.OnAfterCircleReadCallback.OnAfterRead(stage, target)
		}
	case *CircleGrid:
		if stage.OnAfterCircleGridReadCallback != nil {
			stage.OnAfterCircleGridReadCallback.OnAfterRead(stage, target)
		}
	case *HorizontalAxis:
		if stage.OnAfterHorizontalAxisReadCallback != nil {
			stage.OnAfterHorizontalAxisReadCallback.OnAfterRead(stage, target)
		}
	case *Line:
		if stage.OnAfterLineReadCallback != nil {
			stage.OnAfterLineReadCallback.OnAfterRead(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterReadCallback != nil {
			stage.OnAfterParameterReadCallback.OnAfterRead(stage, target)
		}
	case *Rhombus:
		if stage.OnAfterRhombusReadCallback != nil {
			stage.OnAfterRhombusReadCallback.OnAfterRead(stage, target)
		}
	case *RhombusGrid:
		if stage.OnAfterRhombusGridReadCallback != nil {
			stage.OnAfterRhombusGridReadCallback.OnAfterRead(stage, target)
		}
	case *VerticalAxis:
		if stage.OnAfterVerticalAxisReadCallback != nil {
			stage.OnAfterVerticalAxisReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Axis:
		stage.OnAfterAxisUpdateCallback = any(callback).(OnAfterUpdateInterface[Axis])
	
	case *Circle:
		stage.OnAfterCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridUpdateCallback = any(callback).(OnAfterUpdateInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisUpdateCallback = any(callback).(OnAfterUpdateInterface[HorizontalAxis])
	
	case *Line:
		stage.OnAfterLineUpdateCallback = any(callback).(OnAfterUpdateInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterUpdateCallback = any(callback).(OnAfterUpdateInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusUpdateCallback = any(callback).(OnAfterUpdateInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusGrid])
	
	case *VerticalAxis:
		stage.OnAfterVerticalAxisUpdateCallback = any(callback).(OnAfterUpdateInterface[VerticalAxis])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Axis:
		stage.OnAfterAxisCreateCallback = any(callback).(OnAfterCreateInterface[Axis])
	
	case *Circle:
		stage.OnAfterCircleCreateCallback = any(callback).(OnAfterCreateInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridCreateCallback = any(callback).(OnAfterCreateInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisCreateCallback = any(callback).(OnAfterCreateInterface[HorizontalAxis])
	
	case *Line:
		stage.OnAfterLineCreateCallback = any(callback).(OnAfterCreateInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterCreateCallback = any(callback).(OnAfterCreateInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusCreateCallback = any(callback).(OnAfterCreateInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridCreateCallback = any(callback).(OnAfterCreateInterface[RhombusGrid])
	
	case *VerticalAxis:
		stage.OnAfterVerticalAxisCreateCallback = any(callback).(OnAfterCreateInterface[VerticalAxis])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Axis:
		stage.OnAfterAxisDeleteCallback = any(callback).(OnAfterDeleteInterface[Axis])
	
	case *Circle:
		stage.OnAfterCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridDeleteCallback = any(callback).(OnAfterDeleteInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisDeleteCallback = any(callback).(OnAfterDeleteInterface[HorizontalAxis])
	
	case *Line:
		stage.OnAfterLineDeleteCallback = any(callback).(OnAfterDeleteInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterDeleteCallback = any(callback).(OnAfterDeleteInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusDeleteCallback = any(callback).(OnAfterDeleteInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusGrid])
	
	case *VerticalAxis:
		stage.OnAfterVerticalAxisDeleteCallback = any(callback).(OnAfterDeleteInterface[VerticalAxis])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Axis:
		stage.OnAfterAxisReadCallback = any(callback).(OnAfterReadInterface[Axis])
	
	case *Circle:
		stage.OnAfterCircleReadCallback = any(callback).(OnAfterReadInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridReadCallback = any(callback).(OnAfterReadInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisReadCallback = any(callback).(OnAfterReadInterface[HorizontalAxis])
	
	case *Line:
		stage.OnAfterLineReadCallback = any(callback).(OnAfterReadInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterReadCallback = any(callback).(OnAfterReadInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusReadCallback = any(callback).(OnAfterReadInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridReadCallback = any(callback).(OnAfterReadInterface[RhombusGrid])
	
	case *VerticalAxis:
		stage.OnAfterVerticalAxisReadCallback = any(callback).(OnAfterReadInterface[VerticalAxis])
	
	}
}
