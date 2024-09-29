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
	case *AxisGrid:
		if stage.OnAfterAxisGridCreateCallback != nil {
			stage.OnAfterAxisGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bezier:
		if stage.OnAfterBezierCreateCallback != nil {
			stage.OnAfterBezierCreateCallback.OnAfterCreate(stage, target)
		}
	case *BezierGrid:
		if stage.OnAfterBezierGridCreateCallback != nil {
			stage.OnAfterBezierGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *BezierGridStack:
		if stage.OnAfterBezierGridStackCreateCallback != nil {
			stage.OnAfterBezierGridStackCreateCallback.OnAfterCreate(stage, target)
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
	case *Key:
		if stage.OnAfterKeyCreateCallback != nil {
			stage.OnAfterKeyCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteInfo:
		if stage.OnAfterNoteInfoCreateCallback != nil {
			stage.OnAfterNoteInfoCreateCallback.OnAfterCreate(stage, target)
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
	case *ShapeCategory:
		if stage.OnAfterShapeCategoryCreateCallback != nil {
			stage.OnAfterShapeCategoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *SpiralAxis:
		if stage.OnAfterSpiralAxisCreateCallback != nil {
			stage.OnAfterSpiralAxisCreateCallback.OnAfterCreate(stage, target)
		}
	case *SpiralCircle:
		if stage.OnAfterSpiralCircleCreateCallback != nil {
			stage.OnAfterSpiralCircleCreateCallback.OnAfterCreate(stage, target)
		}
	case *SpiralCircleGrid:
		if stage.OnAfterSpiralCircleGridCreateCallback != nil {
			stage.OnAfterSpiralCircleGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *SpiralRhombus:
		if stage.OnAfterSpiralRhombusCreateCallback != nil {
			stage.OnAfterSpiralRhombusCreateCallback.OnAfterCreate(stage, target)
		}
	case *SpiralRhombusGrid:
		if stage.OnAfterSpiralRhombusGridCreateCallback != nil {
			stage.OnAfterSpiralRhombusGridCreateCallback.OnAfterCreate(stage, target)
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
	case *AxisGrid:
		newTarget := any(new).(*AxisGrid)
		if stage.OnAfterAxisGridUpdateCallback != nil {
			stage.OnAfterAxisGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bezier:
		newTarget := any(new).(*Bezier)
		if stage.OnAfterBezierUpdateCallback != nil {
			stage.OnAfterBezierUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BezierGrid:
		newTarget := any(new).(*BezierGrid)
		if stage.OnAfterBezierGridUpdateCallback != nil {
			stage.OnAfterBezierGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BezierGridStack:
		newTarget := any(new).(*BezierGridStack)
		if stage.OnAfterBezierGridStackUpdateCallback != nil {
			stage.OnAfterBezierGridStackUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Key:
		newTarget := any(new).(*Key)
		if stage.OnAfterKeyUpdateCallback != nil {
			stage.OnAfterKeyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteInfo:
		newTarget := any(new).(*NoteInfo)
		if stage.OnAfterNoteInfoUpdateCallback != nil {
			stage.OnAfterNoteInfoUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *ShapeCategory:
		newTarget := any(new).(*ShapeCategory)
		if stage.OnAfterShapeCategoryUpdateCallback != nil {
			stage.OnAfterShapeCategoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SpiralAxis:
		newTarget := any(new).(*SpiralAxis)
		if stage.OnAfterSpiralAxisUpdateCallback != nil {
			stage.OnAfterSpiralAxisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SpiralCircle:
		newTarget := any(new).(*SpiralCircle)
		if stage.OnAfterSpiralCircleUpdateCallback != nil {
			stage.OnAfterSpiralCircleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SpiralCircleGrid:
		newTarget := any(new).(*SpiralCircleGrid)
		if stage.OnAfterSpiralCircleGridUpdateCallback != nil {
			stage.OnAfterSpiralCircleGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SpiralRhombus:
		newTarget := any(new).(*SpiralRhombus)
		if stage.OnAfterSpiralRhombusUpdateCallback != nil {
			stage.OnAfterSpiralRhombusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SpiralRhombusGrid:
		newTarget := any(new).(*SpiralRhombusGrid)
		if stage.OnAfterSpiralRhombusGridUpdateCallback != nil {
			stage.OnAfterSpiralRhombusGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *AxisGrid:
		if stage.OnAfterAxisGridDeleteCallback != nil {
			staged := any(staged).(*AxisGrid)
			stage.OnAfterAxisGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bezier:
		if stage.OnAfterBezierDeleteCallback != nil {
			staged := any(staged).(*Bezier)
			stage.OnAfterBezierDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BezierGrid:
		if stage.OnAfterBezierGridDeleteCallback != nil {
			staged := any(staged).(*BezierGrid)
			stage.OnAfterBezierGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BezierGridStack:
		if stage.OnAfterBezierGridStackDeleteCallback != nil {
			staged := any(staged).(*BezierGridStack)
			stage.OnAfterBezierGridStackDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Key:
		if stage.OnAfterKeyDeleteCallback != nil {
			staged := any(staged).(*Key)
			stage.OnAfterKeyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteInfo:
		if stage.OnAfterNoteInfoDeleteCallback != nil {
			staged := any(staged).(*NoteInfo)
			stage.OnAfterNoteInfoDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *ShapeCategory:
		if stage.OnAfterShapeCategoryDeleteCallback != nil {
			staged := any(staged).(*ShapeCategory)
			stage.OnAfterShapeCategoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SpiralAxis:
		if stage.OnAfterSpiralAxisDeleteCallback != nil {
			staged := any(staged).(*SpiralAxis)
			stage.OnAfterSpiralAxisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SpiralCircle:
		if stage.OnAfterSpiralCircleDeleteCallback != nil {
			staged := any(staged).(*SpiralCircle)
			stage.OnAfterSpiralCircleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SpiralCircleGrid:
		if stage.OnAfterSpiralCircleGridDeleteCallback != nil {
			staged := any(staged).(*SpiralCircleGrid)
			stage.OnAfterSpiralCircleGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SpiralRhombus:
		if stage.OnAfterSpiralRhombusDeleteCallback != nil {
			staged := any(staged).(*SpiralRhombus)
			stage.OnAfterSpiralRhombusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SpiralRhombusGrid:
		if stage.OnAfterSpiralRhombusGridDeleteCallback != nil {
			staged := any(staged).(*SpiralRhombusGrid)
			stage.OnAfterSpiralRhombusGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *AxisGrid:
		if stage.OnAfterAxisGridReadCallback != nil {
			stage.OnAfterAxisGridReadCallback.OnAfterRead(stage, target)
		}
	case *Bezier:
		if stage.OnAfterBezierReadCallback != nil {
			stage.OnAfterBezierReadCallback.OnAfterRead(stage, target)
		}
	case *BezierGrid:
		if stage.OnAfterBezierGridReadCallback != nil {
			stage.OnAfterBezierGridReadCallback.OnAfterRead(stage, target)
		}
	case *BezierGridStack:
		if stage.OnAfterBezierGridStackReadCallback != nil {
			stage.OnAfterBezierGridStackReadCallback.OnAfterRead(stage, target)
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
	case *Key:
		if stage.OnAfterKeyReadCallback != nil {
			stage.OnAfterKeyReadCallback.OnAfterRead(stage, target)
		}
	case *NoteInfo:
		if stage.OnAfterNoteInfoReadCallback != nil {
			stage.OnAfterNoteInfoReadCallback.OnAfterRead(stage, target)
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
	case *ShapeCategory:
		if stage.OnAfterShapeCategoryReadCallback != nil {
			stage.OnAfterShapeCategoryReadCallback.OnAfterRead(stage, target)
		}
	case *SpiralAxis:
		if stage.OnAfterSpiralAxisReadCallback != nil {
			stage.OnAfterSpiralAxisReadCallback.OnAfterRead(stage, target)
		}
	case *SpiralCircle:
		if stage.OnAfterSpiralCircleReadCallback != nil {
			stage.OnAfterSpiralCircleReadCallback.OnAfterRead(stage, target)
		}
	case *SpiralCircleGrid:
		if stage.OnAfterSpiralCircleGridReadCallback != nil {
			stage.OnAfterSpiralCircleGridReadCallback.OnAfterRead(stage, target)
		}
	case *SpiralRhombus:
		if stage.OnAfterSpiralRhombusReadCallback != nil {
			stage.OnAfterSpiralRhombusReadCallback.OnAfterRead(stage, target)
		}
	case *SpiralRhombusGrid:
		if stage.OnAfterSpiralRhombusGridReadCallback != nil {
			stage.OnAfterSpiralRhombusGridReadCallback.OnAfterRead(stage, target)
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
	
	case *AxisGrid:
		stage.OnAfterAxisGridUpdateCallback = any(callback).(OnAfterUpdateInterface[AxisGrid])
	
	case *Bezier:
		stage.OnAfterBezierUpdateCallback = any(callback).(OnAfterUpdateInterface[Bezier])
	
	case *BezierGrid:
		stage.OnAfterBezierGridUpdateCallback = any(callback).(OnAfterUpdateInterface[BezierGrid])
	
	case *BezierGridStack:
		stage.OnAfterBezierGridStackUpdateCallback = any(callback).(OnAfterUpdateInterface[BezierGridStack])
	
	case *Circle:
		stage.OnAfterCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridUpdateCallback = any(callback).(OnAfterUpdateInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisUpdateCallback = any(callback).(OnAfterUpdateInterface[HorizontalAxis])
	
	case *Key:
		stage.OnAfterKeyUpdateCallback = any(callback).(OnAfterUpdateInterface[Key])
	
	case *NoteInfo:
		stage.OnAfterNoteInfoUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteInfo])
	
	case *Parameter:
		stage.OnAfterParameterUpdateCallback = any(callback).(OnAfterUpdateInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusUpdateCallback = any(callback).(OnAfterUpdateInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusGrid])
	
	case *ShapeCategory:
		stage.OnAfterShapeCategoryUpdateCallback = any(callback).(OnAfterUpdateInterface[ShapeCategory])
	
	case *SpiralAxis:
		stage.OnAfterSpiralAxisUpdateCallback = any(callback).(OnAfterUpdateInterface[SpiralAxis])
	
	case *SpiralCircle:
		stage.OnAfterSpiralCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[SpiralCircle])
	
	case *SpiralCircleGrid:
		stage.OnAfterSpiralCircleGridUpdateCallback = any(callback).(OnAfterUpdateInterface[SpiralCircleGrid])
	
	case *SpiralRhombus:
		stage.OnAfterSpiralRhombusUpdateCallback = any(callback).(OnAfterUpdateInterface[SpiralRhombus])
	
	case *SpiralRhombusGrid:
		stage.OnAfterSpiralRhombusGridUpdateCallback = any(callback).(OnAfterUpdateInterface[SpiralRhombusGrid])
	
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
	
	case *AxisGrid:
		stage.OnAfterAxisGridCreateCallback = any(callback).(OnAfterCreateInterface[AxisGrid])
	
	case *Bezier:
		stage.OnAfterBezierCreateCallback = any(callback).(OnAfterCreateInterface[Bezier])
	
	case *BezierGrid:
		stage.OnAfterBezierGridCreateCallback = any(callback).(OnAfterCreateInterface[BezierGrid])
	
	case *BezierGridStack:
		stage.OnAfterBezierGridStackCreateCallback = any(callback).(OnAfterCreateInterface[BezierGridStack])
	
	case *Circle:
		stage.OnAfterCircleCreateCallback = any(callback).(OnAfterCreateInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridCreateCallback = any(callback).(OnAfterCreateInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisCreateCallback = any(callback).(OnAfterCreateInterface[HorizontalAxis])
	
	case *Key:
		stage.OnAfterKeyCreateCallback = any(callback).(OnAfterCreateInterface[Key])
	
	case *NoteInfo:
		stage.OnAfterNoteInfoCreateCallback = any(callback).(OnAfterCreateInterface[NoteInfo])
	
	case *Parameter:
		stage.OnAfterParameterCreateCallback = any(callback).(OnAfterCreateInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusCreateCallback = any(callback).(OnAfterCreateInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridCreateCallback = any(callback).(OnAfterCreateInterface[RhombusGrid])
	
	case *ShapeCategory:
		stage.OnAfterShapeCategoryCreateCallback = any(callback).(OnAfterCreateInterface[ShapeCategory])
	
	case *SpiralAxis:
		stage.OnAfterSpiralAxisCreateCallback = any(callback).(OnAfterCreateInterface[SpiralAxis])
	
	case *SpiralCircle:
		stage.OnAfterSpiralCircleCreateCallback = any(callback).(OnAfterCreateInterface[SpiralCircle])
	
	case *SpiralCircleGrid:
		stage.OnAfterSpiralCircleGridCreateCallback = any(callback).(OnAfterCreateInterface[SpiralCircleGrid])
	
	case *SpiralRhombus:
		stage.OnAfterSpiralRhombusCreateCallback = any(callback).(OnAfterCreateInterface[SpiralRhombus])
	
	case *SpiralRhombusGrid:
		stage.OnAfterSpiralRhombusGridCreateCallback = any(callback).(OnAfterCreateInterface[SpiralRhombusGrid])
	
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
	
	case *AxisGrid:
		stage.OnAfterAxisGridDeleteCallback = any(callback).(OnAfterDeleteInterface[AxisGrid])
	
	case *Bezier:
		stage.OnAfterBezierDeleteCallback = any(callback).(OnAfterDeleteInterface[Bezier])
	
	case *BezierGrid:
		stage.OnAfterBezierGridDeleteCallback = any(callback).(OnAfterDeleteInterface[BezierGrid])
	
	case *BezierGridStack:
		stage.OnAfterBezierGridStackDeleteCallback = any(callback).(OnAfterDeleteInterface[BezierGridStack])
	
	case *Circle:
		stage.OnAfterCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridDeleteCallback = any(callback).(OnAfterDeleteInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisDeleteCallback = any(callback).(OnAfterDeleteInterface[HorizontalAxis])
	
	case *Key:
		stage.OnAfterKeyDeleteCallback = any(callback).(OnAfterDeleteInterface[Key])
	
	case *NoteInfo:
		stage.OnAfterNoteInfoDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteInfo])
	
	case *Parameter:
		stage.OnAfterParameterDeleteCallback = any(callback).(OnAfterDeleteInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusDeleteCallback = any(callback).(OnAfterDeleteInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusGrid])
	
	case *ShapeCategory:
		stage.OnAfterShapeCategoryDeleteCallback = any(callback).(OnAfterDeleteInterface[ShapeCategory])
	
	case *SpiralAxis:
		stage.OnAfterSpiralAxisDeleteCallback = any(callback).(OnAfterDeleteInterface[SpiralAxis])
	
	case *SpiralCircle:
		stage.OnAfterSpiralCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[SpiralCircle])
	
	case *SpiralCircleGrid:
		stage.OnAfterSpiralCircleGridDeleteCallback = any(callback).(OnAfterDeleteInterface[SpiralCircleGrid])
	
	case *SpiralRhombus:
		stage.OnAfterSpiralRhombusDeleteCallback = any(callback).(OnAfterDeleteInterface[SpiralRhombus])
	
	case *SpiralRhombusGrid:
		stage.OnAfterSpiralRhombusGridDeleteCallback = any(callback).(OnAfterDeleteInterface[SpiralRhombusGrid])
	
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
	
	case *AxisGrid:
		stage.OnAfterAxisGridReadCallback = any(callback).(OnAfterReadInterface[AxisGrid])
	
	case *Bezier:
		stage.OnAfterBezierReadCallback = any(callback).(OnAfterReadInterface[Bezier])
	
	case *BezierGrid:
		stage.OnAfterBezierGridReadCallback = any(callback).(OnAfterReadInterface[BezierGrid])
	
	case *BezierGridStack:
		stage.OnAfterBezierGridStackReadCallback = any(callback).(OnAfterReadInterface[BezierGridStack])
	
	case *Circle:
		stage.OnAfterCircleReadCallback = any(callback).(OnAfterReadInterface[Circle])
	
	case *CircleGrid:
		stage.OnAfterCircleGridReadCallback = any(callback).(OnAfterReadInterface[CircleGrid])
	
	case *HorizontalAxis:
		stage.OnAfterHorizontalAxisReadCallback = any(callback).(OnAfterReadInterface[HorizontalAxis])
	
	case *Key:
		stage.OnAfterKeyReadCallback = any(callback).(OnAfterReadInterface[Key])
	
	case *NoteInfo:
		stage.OnAfterNoteInfoReadCallback = any(callback).(OnAfterReadInterface[NoteInfo])
	
	case *Parameter:
		stage.OnAfterParameterReadCallback = any(callback).(OnAfterReadInterface[Parameter])
	
	case *Rhombus:
		stage.OnAfterRhombusReadCallback = any(callback).(OnAfterReadInterface[Rhombus])
	
	case *RhombusGrid:
		stage.OnAfterRhombusGridReadCallback = any(callback).(OnAfterReadInterface[RhombusGrid])
	
	case *ShapeCategory:
		stage.OnAfterShapeCategoryReadCallback = any(callback).(OnAfterReadInterface[ShapeCategory])
	
	case *SpiralAxis:
		stage.OnAfterSpiralAxisReadCallback = any(callback).(OnAfterReadInterface[SpiralAxis])
	
	case *SpiralCircle:
		stage.OnAfterSpiralCircleReadCallback = any(callback).(OnAfterReadInterface[SpiralCircle])
	
	case *SpiralCircleGrid:
		stage.OnAfterSpiralCircleGridReadCallback = any(callback).(OnAfterReadInterface[SpiralCircleGrid])
	
	case *SpiralRhombus:
		stage.OnAfterSpiralRhombusReadCallback = any(callback).(OnAfterReadInterface[SpiralRhombus])
	
	case *SpiralRhombusGrid:
		stage.OnAfterSpiralRhombusGridReadCallback = any(callback).(OnAfterReadInterface[SpiralRhombusGrid])
	
	case *VerticalAxis:
		stage.OnAfterVerticalAxisReadCallback = any(callback).(OnAfterReadInterface[VerticalAxis])
	
	}
}
