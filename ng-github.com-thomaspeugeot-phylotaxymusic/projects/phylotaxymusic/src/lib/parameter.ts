// generated code - do not edit

import { ParameterAPI } from './parameter-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rhombus } from './rhombus'
import { Circle } from './circle'
import { RhombusGrid } from './rhombusgrid'
import { CircleGrid } from './circlegrid'
import { Axis } from './axis'
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
	InsideAngle: number = 0
	SideLength: number = 0
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
	parameterAPI.InsideAngle = parameter.InsideAngle
	parameterAPI.SideLength = parameter.SideLength
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
	parameter.InsideAngle = parameterAPI.InsideAngle
	parameter.SideLength = parameterAPI.SideLength
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
	parameter.HorizontalAxis = frontRepo.map_ID_HorizontalAxis.get(parameterAPI.ParameterPointersEncoding.HorizontalAxisID.Int64)
	parameter.VerticalAxis = frontRepo.map_ID_VerticalAxis.get(parameterAPI.ParameterPointersEncoding.VerticalAxisID.Int64)

	// insertion point for slice of pointers fields encoding
}
