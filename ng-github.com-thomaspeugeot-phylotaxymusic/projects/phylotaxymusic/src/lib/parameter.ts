// generated code - do not edit

import { ParameterAPI } from './parameter-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rhombus } from './rhombus'

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
	Angle: number = 0
	OriginX: number = 0
	OriginY: number = 0
	DiamondSideLenght: number = 0

	// insertion point for pointers and slices of pointers declarations
	InitialRhombus?: Rhombus

}

export function CopyParameterToParameterAPI(parameter: Parameter, parameterAPI: ParameterAPI) {

	parameterAPI.CreatedAt = parameter.CreatedAt
	parameterAPI.DeletedAt = parameter.DeletedAt
	parameterAPI.ID = parameter.ID

	// insertion point for basic fields copy operations
	parameterAPI.Name = parameter.Name
	parameterAPI.N = parameter.N
	parameterAPI.M = parameter.M
	parameterAPI.Angle = parameter.Angle
	parameterAPI.OriginX = parameter.OriginX
	parameterAPI.OriginY = parameter.OriginY
	parameterAPI.DiamondSideLenght = parameter.DiamondSideLenght

	// insertion point for pointer fields encoding
	parameterAPI.ParameterPointersEncoding.InitialRhombusID.Valid = true
	if (parameter.InitialRhombus != undefined) {
		parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64 = parameter.InitialRhombus.ID  
	} else {
		parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64 = 0 		
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
	parameter.Angle = parameterAPI.Angle
	parameter.OriginX = parameterAPI.OriginX
	parameter.OriginY = parameterAPI.OriginY
	parameter.DiamondSideLenght = parameterAPI.DiamondSideLenght

	// insertion point for pointer fields encoding
	parameter.InitialRhombus = frontRepo.map_ID_Rhombus.get(parameterAPI.ParameterPointersEncoding.InitialRhombusID.Int64)

	// insertion point for slice of pointers fields encoding
}
