// generated code - do not edit

import { ParameterAPI } from './parameter-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rhombus } from './rhombus'
import { Circle } from './circle'
import { RhombusGrid } from './rhombusgrid'
import { CircleGrid } from './circlegrid'
import { Axis } from './axis'
import { AxisGrid } from './axisgrid'
import { Bezier } from './bezier'
import { BezierGrid } from './beziergrid'
import { BezierGridStack } from './beziergridstack'
import { Key } from './key'
import { HorizontalAxis } from './horizontalaxis'
import { VerticalAxis } from './verticalaxis'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Parameter {

	static GONGSTRUCT_NAME = "Parameter"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	N: number = 0
	M: number = 0
	Z: number = 0
	InsideAngle: number = 0
	SideLength: number = 0
	StackWidth: number = 0
	NbShitRight: number = 0
	StackHeight: number = 0
	BezierControlLengthRatio: number = 0
	FkeySizeRatio: number = 0
	FkeyOriginRelativeX: number = 0
	FkeyOriginRelativeY: number = 0
	PitchHeight: number = 0
	NbPitchLines: number = 0
	MeasureLinesHeightRatio: number = 0
	NbMeasureLines: number = 0
	NbMeasureLinesPerCurve: number = 0
	FirstVoiceShiftX: number = 0
	FirstVoiceShiftY: number = 0
	PitchDifference: number = 0
	Speed: number = 0
	Level: number = 0
	OriginX: number = 0
	OriginY: number = 0

	// insertion point for pointers and slices of pointers declarations
	InitialRhombus?: Rhombus

	InitialCircle?: Circle

	InitialRhombusGrid?: RhombusGrid

	InitialCircleGrid?: CircleGrid

	InitialAxis?: Axis

	RotatedAxis?: Axis

	RotatedRhombus?: Rhombus

	RotatedRhombusGrid?: RhombusGrid

	RotatedCircleGrid?: CircleGrid

	NextRhombus?: Rhombus

	NextCircle?: Circle

	GrowingRhombusGridSeed?: Rhombus

	GrowingRhombusGrid?: RhombusGrid

	GrowingCircleGridSeed?: Circle

	GrowingCircleGrid?: CircleGrid

	GrowingCircleGridLeftSeed?: Circle

	GrowingCircleGridLeft?: CircleGrid

	ConstructionAxis?: Axis

	ConstructionAxisGrid?: AxisGrid

	ConstructionCircle?: Circle

	ConstructionCircleGrid?: CircleGrid

	GrowthCurveSegment?: Bezier

	GrowthCurve?: BezierGrid

	GrowthCurveShiftedRightSeed?: Bezier

	GrowthCurveShiftedRight?: BezierGrid

	GrowthCurveNextSeed?: Bezier

	GrowthCurveNext?: BezierGrid

	GrowthCurveNextShiftedRightSeed?: Bezier

	GrowthCurveNextShiftedRight?: BezierGrid

	GrowthCurveStack?: BezierGridStack

	Fkey?: Key

	PitchLines?: AxisGrid

	MeasureLines?: AxisGrid

	FirstVoice?: BezierGrid

	FirstVoiceShiftRigth?: BezierGrid

	SecondVoice?: BezierGrid

	SecondVoiceShiftedRight?: BezierGrid

	FirstVoiceNotes?: CircleGrid

	FirstVoiceNotesShiftedRight?: CircleGrid

	HorizontalAxis?: HorizontalAxis

	VerticalAxis?: VerticalAxis

}

export function CopyParameterToParameterAPI(parameter: Parameter, parameterAPI: ParameterAPI) {

	parameterAPI.CreatedAt = parameter.CreatedAt
	parameterAPI.DeletedAt = parameter.DeletedAt
	parameterAPI.ID = parameter.ID

	// insertion point for basic fields copy operations
	parameterAPI.Name = parameter.Name
	parameterAPI.N = parameter.N
	parameterAPI.M = parameter.M
	parameterAPI.Z = parameter.Z
	parameterAPI.InsideAngle = parameter.InsideAngle
	parameterAPI.SideLength = parameter.SideLength
	parameterAPI.StackWidth = parameter.StackWidth
	parameterAPI.NbShitRight = parameter.NbShitRight
	parameterAPI.StackHeight = parameter.StackHeight
	parameterAPI.BezierControlLengthRatio = parameter.BezierControlLengthRatio
	parameterAPI.FkeySizeRatio = parameter.FkeySizeRatio
	parameterAPI.FkeyOriginRelativeX = parameter.FkeyOriginRelativeX
	parameterAPI.FkeyOriginRelativeY = parameter.FkeyOriginRelativeY
	parameterAPI.PitchHeight = parameter.PitchHeight
	parameterAPI.NbPitchLines = parameter.NbPitchLines
	parameterAPI.MeasureLinesHeightRatio = parameter.MeasureLinesHeightRatio
	parameterAPI.NbMeasureLines = parameter.NbMeasureLines
	parameterAPI.NbMeasureLinesPerCurve = parameter.NbMeasureLinesPerCurve
	parameterAPI.FirstVoiceShiftX = parameter.FirstVoiceShiftX
	parameterAPI.FirstVoiceShiftY = parameter.FirstVoiceShiftY
	parameterAPI.PitchDifference = parameter.PitchDifference
	parameterAPI.Speed = parameter.Speed
	parameterAPI.Level = parameter.Level
	parameterAPI.OriginX = parameter.OriginX
	parameterAPI.OriginY = parameter.OriginY

	// insertion point for pointer fields encoding
	parameterAPI.ParameterPointersEncoding.InitialRhombusID.Valid = true
	if (parameter.InitialRhombus != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64 = parameter.InitialRhombus.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.InitialCircleID.Valid = true
	if (parameter.InitialCircle != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialCircleID.Int64 = parameter.InitialCircle.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialCircleID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.InitialRhombusGridID.Valid = true
	if (parameter.InitialRhombusGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialRhombusGridID.Int64 = parameter.InitialRhombusGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialRhombusGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.InitialCircleGridID.Valid = true
	if (parameter.InitialCircleGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialCircleGridID.Int64 = parameter.InitialCircleGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialCircleGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.InitialAxisID.Valid = true
	if (parameter.InitialAxis != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialAxisID.Int64 = parameter.InitialAxis.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialAxisID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.RotatedAxisID.Valid = true
	if (parameter.RotatedAxis != undefined) {
		parameterAPI.ParameterPointersEncoding.RotatedAxisID.Int64 = parameter.RotatedAxis.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.RotatedAxisID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.RotatedRhombusID.Valid = true
	if (parameter.RotatedRhombus != undefined) {
		parameterAPI.ParameterPointersEncoding.RotatedRhombusID.Int64 = parameter.RotatedRhombus.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.RotatedRhombusID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.RotatedRhombusGridID.Valid = true
	if (parameter.RotatedRhombusGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.RotatedRhombusGridID.Int64 = parameter.RotatedRhombusGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.RotatedRhombusGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.RotatedCircleGridID.Valid = true
	if (parameter.RotatedCircleGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.RotatedCircleGridID.Int64 = parameter.RotatedCircleGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.RotatedCircleGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.NextRhombusID.Valid = true
	if (parameter.NextRhombus != undefined) {
		parameterAPI.ParameterPointersEncoding.NextRhombusID.Int64 = parameter.NextRhombus.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.NextRhombusID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.NextCircleID.Valid = true
	if (parameter.NextCircle != undefined) {
		parameterAPI.ParameterPointersEncoding.NextCircleID.Int64 = parameter.NextCircle.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.NextCircleID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingRhombusGridSeedID.Valid = true
	if (parameter.GrowingRhombusGridSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingRhombusGridSeedID.Int64 = parameter.GrowingRhombusGridSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingRhombusGridSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingRhombusGridID.Valid = true
	if (parameter.GrowingRhombusGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingRhombusGridID.Int64 = parameter.GrowingRhombusGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingRhombusGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingCircleGridSeedID.Valid = true
	if (parameter.GrowingCircleGridSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridSeedID.Int64 = parameter.GrowingCircleGridSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingCircleGridID.Valid = true
	if (parameter.GrowingCircleGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridID.Int64 = parameter.GrowingCircleGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftSeedID.Valid = true
	if (parameter.GrowingCircleGridLeftSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftSeedID.Int64 = parameter.GrowingCircleGridLeftSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftID.Valid = true
	if (parameter.GrowingCircleGridLeft != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftID.Int64 = parameter.GrowingCircleGridLeft.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.ConstructionAxisID.Valid = true
	if (parameter.ConstructionAxis != undefined) {
		parameterAPI.ParameterPointersEncoding.ConstructionAxisID.Int64 = parameter.ConstructionAxis.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.ConstructionAxisID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.ConstructionAxisGridID.Valid = true
	if (parameter.ConstructionAxisGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.ConstructionAxisGridID.Int64 = parameter.ConstructionAxisGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.ConstructionAxisGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.ConstructionCircleID.Valid = true
	if (parameter.ConstructionCircle != undefined) {
		parameterAPI.ParameterPointersEncoding.ConstructionCircleID.Int64 = parameter.ConstructionCircle.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.ConstructionCircleID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.ConstructionCircleGridID.Valid = true
	if (parameter.ConstructionCircleGrid != undefined) {
		parameterAPI.ParameterPointersEncoding.ConstructionCircleGridID.Int64 = parameter.ConstructionCircleGrid.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.ConstructionCircleGridID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveSegmentID.Valid = true
	if (parameter.GrowthCurveSegment != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveSegmentID.Int64 = parameter.GrowthCurveSegment.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveSegmentID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveID.Valid = true
	if (parameter.GrowthCurve != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveID.Int64 = parameter.GrowthCurve.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightSeedID.Valid = true
	if (parameter.GrowthCurveShiftedRightSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightSeedID.Int64 = parameter.GrowthCurveShiftedRightSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightID.Valid = true
	if (parameter.GrowthCurveShiftedRight != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightID.Int64 = parameter.GrowthCurveShiftedRight.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveNextSeedID.Valid = true
	if (parameter.GrowthCurveNextSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextSeedID.Int64 = parameter.GrowthCurveNextSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveNextID.Valid = true
	if (parameter.GrowthCurveNext != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextID.Int64 = parameter.GrowthCurveNext.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightSeedID.Valid = true
	if (parameter.GrowthCurveNextShiftedRightSeed != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightSeedID.Int64 = parameter.GrowthCurveNextShiftedRightSeed.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightSeedID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightID.Valid = true
	if (parameter.GrowthCurveNextShiftedRight != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightID.Int64 = parameter.GrowthCurveNextShiftedRight.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.GrowthCurveStackID.Valid = true
	if (parameter.GrowthCurveStack != undefined) {
		parameterAPI.ParameterPointersEncoding.GrowthCurveStackID.Int64 = parameter.GrowthCurveStack.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.GrowthCurveStackID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.FkeyID.Valid = true
	if (parameter.Fkey != undefined) {
		parameterAPI.ParameterPointersEncoding.FkeyID.Int64 = parameter.Fkey.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.FkeyID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.PitchLinesID.Valid = true
	if (parameter.PitchLines != undefined) {
		parameterAPI.ParameterPointersEncoding.PitchLinesID.Int64 = parameter.PitchLines.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.PitchLinesID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.MeasureLinesID.Valid = true
	if (parameter.MeasureLines != undefined) {
		parameterAPI.ParameterPointersEncoding.MeasureLinesID.Int64 = parameter.MeasureLines.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.MeasureLinesID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.FirstVoiceID.Valid = true
	if (parameter.FirstVoice != undefined) {
		parameterAPI.ParameterPointersEncoding.FirstVoiceID.Int64 = parameter.FirstVoice.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.FirstVoiceID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.FirstVoiceShiftRigthID.Valid = true
	if (parameter.FirstVoiceShiftRigth != undefined) {
		parameterAPI.ParameterPointersEncoding.FirstVoiceShiftRigthID.Int64 = parameter.FirstVoiceShiftRigth.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.FirstVoiceShiftRigthID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.SecondVoiceID.Valid = true
	if (parameter.SecondVoice != undefined) {
		parameterAPI.ParameterPointersEncoding.SecondVoiceID.Int64 = parameter.SecondVoice.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.SecondVoiceID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.SecondVoiceShiftedRightID.Valid = true
	if (parameter.SecondVoiceShiftedRight != undefined) {
		parameterAPI.ParameterPointersEncoding.SecondVoiceShiftedRightID.Int64 = parameter.SecondVoiceShiftedRight.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.SecondVoiceShiftedRightID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.FirstVoiceNotesID.Valid = true
	if (parameter.FirstVoiceNotes != undefined) {
		parameterAPI.ParameterPointersEncoding.FirstVoiceNotesID.Int64 = parameter.FirstVoiceNotes.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.FirstVoiceNotesID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.FirstVoiceNotesShiftedRightID.Valid = true
	if (parameter.FirstVoiceNotesShiftedRight != undefined) {
		parameterAPI.ParameterPointersEncoding.FirstVoiceNotesShiftedRightID.Int64 = parameter.FirstVoiceNotesShiftedRight.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.FirstVoiceNotesShiftedRightID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.HorizontalAxisID.Valid = true
	if (parameter.HorizontalAxis != undefined) {
		parameterAPI.ParameterPointersEncoding.HorizontalAxisID.Int64 = parameter.HorizontalAxis.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.HorizontalAxisID.Int64 = 0 		
	}

	parameterAPI.ParameterPointersEncoding.VerticalAxisID.Valid = true
	if (parameter.VerticalAxis != undefined) {
		parameterAPI.ParameterPointersEncoding.VerticalAxisID.Int64 = parameter.VerticalAxis.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.VerticalAxisID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyParameterAPIToParameter update basic, pointers and slice of pointers fields of parameter
// from respectively the basic fields and encoded fields of pointers and slices of pointers of parameterAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyParameterAPIToParameter(parameterAPI: ParameterAPI, parameter: Parameter, frontRepo: FrontRepo) {

	parameter.CreatedAt = parameterAPI.CreatedAt
	parameter.DeletedAt = parameterAPI.DeletedAt
	parameter.ID = parameterAPI.ID

	// insertion point for basic fields copy operations
	parameter.Name = parameterAPI.Name
	parameter.N = parameterAPI.N
	parameter.M = parameterAPI.M
	parameter.Z = parameterAPI.Z
	parameter.InsideAngle = parameterAPI.InsideAngle
	parameter.SideLength = parameterAPI.SideLength
	parameter.StackWidth = parameterAPI.StackWidth
	parameter.NbShitRight = parameterAPI.NbShitRight
	parameter.StackHeight = parameterAPI.StackHeight
	parameter.BezierControlLengthRatio = parameterAPI.BezierControlLengthRatio
	parameter.FkeySizeRatio = parameterAPI.FkeySizeRatio
	parameter.FkeyOriginRelativeX = parameterAPI.FkeyOriginRelativeX
	parameter.FkeyOriginRelativeY = parameterAPI.FkeyOriginRelativeY
	parameter.PitchHeight = parameterAPI.PitchHeight
	parameter.NbPitchLines = parameterAPI.NbPitchLines
	parameter.MeasureLinesHeightRatio = parameterAPI.MeasureLinesHeightRatio
	parameter.NbMeasureLines = parameterAPI.NbMeasureLines
	parameter.NbMeasureLinesPerCurve = parameterAPI.NbMeasureLinesPerCurve
	parameter.FirstVoiceShiftX = parameterAPI.FirstVoiceShiftX
	parameter.FirstVoiceShiftY = parameterAPI.FirstVoiceShiftY
	parameter.PitchDifference = parameterAPI.PitchDifference
	parameter.Speed = parameterAPI.Speed
	parameter.Level = parameterAPI.Level
	parameter.OriginX = parameterAPI.OriginX
	parameter.OriginY = parameterAPI.OriginY

	// insertion point for pointer fields encoding
	parameter.InitialRhombus = frontRepo.map_ID_Rhombus.get(parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64)
	parameter.InitialCircle = frontRepo.map_ID_Circle.get(parameterAPI.ParameterPointersEncoding.InitialCircleID.Int64)
	parameter.InitialRhombusGrid = frontRepo.map_ID_RhombusGrid.get(parameterAPI.ParameterPointersEncoding.InitialRhombusGridID.Int64)
	parameter.InitialCircleGrid = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.InitialCircleGridID.Int64)
	parameter.InitialAxis = frontRepo.map_ID_Axis.get(parameterAPI.ParameterPointersEncoding.InitialAxisID.Int64)
	parameter.RotatedAxis = frontRepo.map_ID_Axis.get(parameterAPI.ParameterPointersEncoding.RotatedAxisID.Int64)
	parameter.RotatedRhombus = frontRepo.map_ID_Rhombus.get(parameterAPI.ParameterPointersEncoding.RotatedRhombusID.Int64)
	parameter.RotatedRhombusGrid = frontRepo.map_ID_RhombusGrid.get(parameterAPI.ParameterPointersEncoding.RotatedRhombusGridID.Int64)
	parameter.RotatedCircleGrid = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.RotatedCircleGridID.Int64)
	parameter.NextRhombus = frontRepo.map_ID_Rhombus.get(parameterAPI.ParameterPointersEncoding.NextRhombusID.Int64)
	parameter.NextCircle = frontRepo.map_ID_Circle.get(parameterAPI.ParameterPointersEncoding.NextCircleID.Int64)
	parameter.GrowingRhombusGridSeed = frontRepo.map_ID_Rhombus.get(parameterAPI.ParameterPointersEncoding.GrowingRhombusGridSeedID.Int64)
	parameter.GrowingRhombusGrid = frontRepo.map_ID_RhombusGrid.get(parameterAPI.ParameterPointersEncoding.GrowingRhombusGridID.Int64)
	parameter.GrowingCircleGridSeed = frontRepo.map_ID_Circle.get(parameterAPI.ParameterPointersEncoding.GrowingCircleGridSeedID.Int64)
	parameter.GrowingCircleGrid = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.GrowingCircleGridID.Int64)
	parameter.GrowingCircleGridLeftSeed = frontRepo.map_ID_Circle.get(parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftSeedID.Int64)
	parameter.GrowingCircleGridLeft = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.GrowingCircleGridLeftID.Int64)
	parameter.ConstructionAxis = frontRepo.map_ID_Axis.get(parameterAPI.ParameterPointersEncoding.ConstructionAxisID.Int64)
	parameter.ConstructionAxisGrid = frontRepo.map_ID_AxisGrid.get(parameterAPI.ParameterPointersEncoding.ConstructionAxisGridID.Int64)
	parameter.ConstructionCircle = frontRepo.map_ID_Circle.get(parameterAPI.ParameterPointersEncoding.ConstructionCircleID.Int64)
	parameter.ConstructionCircleGrid = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.ConstructionCircleGridID.Int64)
	parameter.GrowthCurveSegment = frontRepo.map_ID_Bezier.get(parameterAPI.ParameterPointersEncoding.GrowthCurveSegmentID.Int64)
	parameter.GrowthCurve = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.GrowthCurveID.Int64)
	parameter.GrowthCurveShiftedRightSeed = frontRepo.map_ID_Bezier.get(parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightSeedID.Int64)
	parameter.GrowthCurveShiftedRight = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.GrowthCurveShiftedRightID.Int64)
	parameter.GrowthCurveNextSeed = frontRepo.map_ID_Bezier.get(parameterAPI.ParameterPointersEncoding.GrowthCurveNextSeedID.Int64)
	parameter.GrowthCurveNext = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.GrowthCurveNextID.Int64)
	parameter.GrowthCurveNextShiftedRightSeed = frontRepo.map_ID_Bezier.get(parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightSeedID.Int64)
	parameter.GrowthCurveNextShiftedRight = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.GrowthCurveNextShiftedRightID.Int64)
	parameter.GrowthCurveStack = frontRepo.map_ID_BezierGridStack.get(parameterAPI.ParameterPointersEncoding.GrowthCurveStackID.Int64)
	parameter.Fkey = frontRepo.map_ID_Key.get(parameterAPI.ParameterPointersEncoding.FkeyID.Int64)
	parameter.PitchLines = frontRepo.map_ID_AxisGrid.get(parameterAPI.ParameterPointersEncoding.PitchLinesID.Int64)
	parameter.MeasureLines = frontRepo.map_ID_AxisGrid.get(parameterAPI.ParameterPointersEncoding.MeasureLinesID.Int64)
	parameter.FirstVoice = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.FirstVoiceID.Int64)
	parameter.FirstVoiceShiftRigth = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.FirstVoiceShiftRigthID.Int64)
	parameter.SecondVoice = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.SecondVoiceID.Int64)
	parameter.SecondVoiceShiftedRight = frontRepo.map_ID_BezierGrid.get(parameterAPI.ParameterPointersEncoding.SecondVoiceShiftedRightID.Int64)
	parameter.FirstVoiceNotes = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.FirstVoiceNotesID.Int64)
	parameter.FirstVoiceNotesShiftedRight = frontRepo.map_ID_CircleGrid.get(parameterAPI.ParameterPointersEncoding.FirstVoiceNotesShiftedRightID.Int64)
	parameter.HorizontalAxis = frontRepo.map_ID_HorizontalAxis.get(parameterAPI.ParameterPointersEncoding.HorizontalAxisID.Int64)
	parameter.VerticalAxis = frontRepo.map_ID_VerticalAxis.get(parameterAPI.ParameterPointersEncoding.VerticalAxisID.Int64)

	// insertion point for slice of pointers fields encoding
}
