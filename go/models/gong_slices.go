// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Axis:
		// insertion point per field

	case *AxisGrid:
		// insertion point per field
		if fieldName == "Axiss" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*AxisGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*AxisGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Axiss).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Axiss = _inferedTypeInstance.Axiss[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Axiss =
								append(_inferedTypeInstance.Axiss, any(fieldInstance).(*Axis))
						}
					}
				}
			}
		}

	case *Bezier:
		// insertion point per field

	case *BezierGrid:
		// insertion point per field
		if fieldName == "Beziers" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*BezierGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*BezierGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Beziers).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Beziers = _inferedTypeInstance.Beziers[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Beziers =
								append(_inferedTypeInstance.Beziers, any(fieldInstance).(*Bezier))
						}
					}
				}
			}
		}

	case *BezierGridStack:
		// insertion point per field
		if fieldName == "BezierGrids" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*BezierGridStack) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*BezierGridStack)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.BezierGrids).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.BezierGrids = _inferedTypeInstance.BezierGrids[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.BezierGrids =
								append(_inferedTypeInstance.BezierGrids, any(fieldInstance).(*BezierGrid))
						}
					}
				}
			}
		}

	case *Circle:
		// insertion point per field

	case *CircleGrid:
		// insertion point per field
		if fieldName == "Circles" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*CircleGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*CircleGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Circles).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Circles = _inferedTypeInstance.Circles[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Circles =
								append(_inferedTypeInstance.Circles, any(fieldInstance).(*Circle))
						}
					}
				}
			}
		}

	case *HorizontalAxis:
		// insertion point per field

	case *Key:
		// insertion point per field

	case *NoteInfo:
		// insertion point per field

	case *Parameter:
		// insertion point per field
		if fieldName == "NoteInfos" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Parameter) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Parameter)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.NoteInfos).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.NoteInfos = _inferedTypeInstance.NoteInfos[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.NoteInfos =
								append(_inferedTypeInstance.NoteInfos, any(fieldInstance).(*NoteInfo))
						}
					}
				}
			}
		}

	case *Rhombus:
		// insertion point per field

	case *RhombusGrid:
		// insertion point per field
		if fieldName == "Rhombuses" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RhombusGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RhombusGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rhombuses).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rhombuses = _inferedTypeInstance.Rhombuses[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rhombuses =
								append(_inferedTypeInstance.Rhombuses, any(fieldInstance).(*Rhombus))
						}
					}
				}
			}
		}

	case *ShapeCategory:
		// insertion point per field

	case *SpiralBezier:
		// insertion point per field

	case *SpiralBezierGrid:
		// insertion point per field
		if fieldName == "SpiralBeziers" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SpiralBezierGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SpiralBezierGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SpiralBeziers).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SpiralBeziers = _inferedTypeInstance.SpiralBeziers[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SpiralBeziers =
								append(_inferedTypeInstance.SpiralBeziers, any(fieldInstance).(*SpiralBezier))
						}
					}
				}
			}
		}

	case *SpiralCircle:
		// insertion point per field

	case *SpiralCircleGrid:
		// insertion point per field
		if fieldName == "SpiralCircles" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SpiralCircleGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SpiralCircleGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SpiralCircles).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SpiralCircles = _inferedTypeInstance.SpiralCircles[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SpiralCircles =
								append(_inferedTypeInstance.SpiralCircles, any(fieldInstance).(*SpiralCircle))
						}
					}
				}
			}
		}

	case *SpiralLine:
		// insertion point per field

	case *SpiralLineGrid:
		// insertion point per field
		if fieldName == "SpiralLines" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SpiralLineGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SpiralLineGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SpiralLines).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SpiralLines = _inferedTypeInstance.SpiralLines[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SpiralLines =
								append(_inferedTypeInstance.SpiralLines, any(fieldInstance).(*SpiralLine))
						}
					}
				}
			}
		}

	case *SpiralOrigin:
		// insertion point per field

	case *SpiralRhombus:
		// insertion point per field

	case *SpiralRhombusGrid:
		// insertion point per field
		if fieldName == "SpiralRhombuses" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SpiralRhombusGrid) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SpiralRhombusGrid)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SpiralRhombuses).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SpiralRhombuses = _inferedTypeInstance.SpiralRhombuses[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SpiralRhombuses =
								append(_inferedTypeInstance.SpiralRhombuses, any(fieldInstance).(*SpiralRhombus))
						}
					}
				}
			}
		}

	case *VerticalAxis:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Axis
	// insertion point per field

	// Compute reverse map for named struct AxisGrid
	// insertion point per field
	clear(stage.AxisGrid_Axiss_reverseMap)
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
	clear(stage.BezierGrid_Beziers_reverseMap)
	stage.BezierGrid_Beziers_reverseMap = make(map[*Bezier]*BezierGrid)
	for beziergrid := range stage.BezierGrids {
		_ = beziergrid
		for _, _bezier := range beziergrid.Beziers {
			stage.BezierGrid_Beziers_reverseMap[_bezier] = beziergrid
		}
	}

	// Compute reverse map for named struct BezierGridStack
	// insertion point per field
	clear(stage.BezierGridStack_BezierGrids_reverseMap)
	stage.BezierGridStack_BezierGrids_reverseMap = make(map[*BezierGrid]*BezierGridStack)
	for beziergridstack := range stage.BezierGridStacks {
		_ = beziergridstack
		for _, _beziergrid := range beziergridstack.BezierGrids {
			stage.BezierGridStack_BezierGrids_reverseMap[_beziergrid] = beziergridstack
		}
	}

	// Compute reverse map for named struct Circle
	// insertion point per field

	// Compute reverse map for named struct CircleGrid
	// insertion point per field
	clear(stage.CircleGrid_Circles_reverseMap)
	stage.CircleGrid_Circles_reverseMap = make(map[*Circle]*CircleGrid)
	for circlegrid := range stage.CircleGrids {
		_ = circlegrid
		for _, _circle := range circlegrid.Circles {
			stage.CircleGrid_Circles_reverseMap[_circle] = circlegrid
		}
	}

	// Compute reverse map for named struct HorizontalAxis
	// insertion point per field

	// Compute reverse map for named struct Key
	// insertion point per field

	// Compute reverse map for named struct NoteInfo
	// insertion point per field

	// Compute reverse map for named struct Parameter
	// insertion point per field
	clear(stage.Parameter_NoteInfos_reverseMap)
	stage.Parameter_NoteInfos_reverseMap = make(map[*NoteInfo]*Parameter)
	for parameter := range stage.Parameters {
		_ = parameter
		for _, _noteinfo := range parameter.NoteInfos {
			stage.Parameter_NoteInfos_reverseMap[_noteinfo] = parameter
		}
	}

	// Compute reverse map for named struct Rhombus
	// insertion point per field

	// Compute reverse map for named struct RhombusGrid
	// insertion point per field
	clear(stage.RhombusGrid_Rhombuses_reverseMap)
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
	clear(stage.SpiralBezierGrid_SpiralBeziers_reverseMap)
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
	clear(stage.SpiralCircleGrid_SpiralCircles_reverseMap)
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
	clear(stage.SpiralLineGrid_SpiralLines_reverseMap)
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
	clear(stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap)
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
