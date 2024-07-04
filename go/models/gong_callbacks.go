// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Line:
		if stage.OnAfterLineCreateCallback != nil {
			stage.OnAfterLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterCreateCallback != nil {
			stage.OnAfterParameterCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
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
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
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
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Line:
		if stage.OnAfterLineReadCallback != nil {
			stage.OnAfterLineReadCallback.OnAfterRead(stage, target)
		}
	case *Parameter:
		if stage.OnAfterParameterReadCallback != nil {
			stage.OnAfterParameterReadCallback.OnAfterRead(stage, target)
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
	case *Line:
		stage.OnAfterLineUpdateCallback = any(callback).(OnAfterUpdateInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterUpdateCallback = any(callback).(OnAfterUpdateInterface[Parameter])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Line:
		stage.OnAfterLineCreateCallback = any(callback).(OnAfterCreateInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterCreateCallback = any(callback).(OnAfterCreateInterface[Parameter])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Line:
		stage.OnAfterLineDeleteCallback = any(callback).(OnAfterDeleteInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterDeleteCallback = any(callback).(OnAfterDeleteInterface[Parameter])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Line:
		stage.OnAfterLineReadCallback = any(callback).(OnAfterReadInterface[Line])
	
	case *Parameter:
		stage.OnAfterParameterReadCallback = any(callback).(OnAfterReadInterface[Parameter])
	
	}
}
