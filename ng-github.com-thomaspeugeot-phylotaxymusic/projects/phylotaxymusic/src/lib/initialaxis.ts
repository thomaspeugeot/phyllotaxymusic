// generated code - do not edit

import { InitialAxisAPI } from './initialaxis-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class InitialAxis {

	static GONGSTRUCT_NAME = "InitialAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	Angle: number = 0
	Length: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyInitialAxisToInitialAxisAPI(initialaxis: InitialAxis, initialaxisAPI: InitialAxisAPI) {

	initialaxisAPI.CreatedAt = initialaxis.CreatedAt
	initialaxisAPI.DeletedAt = initialaxis.DeletedAt
	initialaxisAPI.ID = initialaxis.ID

	// insertion point for basic fields copy operations
	initialaxisAPI.Name = initialaxis.Name
	initialaxisAPI.IsDisplayed = initialaxis.IsDisplayed
	initialaxisAPI.Angle = initialaxis.Angle
	initialaxisAPI.Length = initialaxis.Length
	initialaxisAPI.Color = initialaxis.Color
	initialaxisAPI.FillOpacity = initialaxis.FillOpacity
	initialaxisAPI.Stroke = initialaxis.Stroke
	initialaxisAPI.StrokeOpacity = initialaxis.StrokeOpacity
	initialaxisAPI.StrokeWidth = initialaxis.StrokeWidth
	initialaxisAPI.StrokeDashArray = initialaxis.StrokeDashArray
	initialaxisAPI.StrokeDashArrayWhenSelected = initialaxis.StrokeDashArrayWhenSelected
	initialaxisAPI.Transform = initialaxis.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyInitialAxisAPIToInitialAxis update basic, pointers and slice of pointers fields of initialaxis
// from respectively the basic fields and encoded fields of pointers and slices of pointers of initialaxisAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyInitialAxisAPIToInitialAxis(initialaxisAPI: InitialAxisAPI, initialaxis: InitialAxis, frontRepo: FrontRepo) {

	initialaxis.CreatedAt = initialaxisAPI.CreatedAt
	initialaxis.DeletedAt = initialaxisAPI.DeletedAt
	initialaxis.ID = initialaxisAPI.ID

	// insertion point for basic fields copy operations
	initialaxis.Name = initialaxisAPI.Name
	initialaxis.IsDisplayed = initialaxisAPI.IsDisplayed
	initialaxis.Angle = initialaxisAPI.Angle
	initialaxis.Length = initialaxisAPI.Length
	initialaxis.Color = initialaxisAPI.Color
	initialaxis.FillOpacity = initialaxisAPI.FillOpacity
	initialaxis.Stroke = initialaxisAPI.Stroke
	initialaxis.StrokeOpacity = initialaxisAPI.StrokeOpacity
	initialaxis.StrokeWidth = initialaxisAPI.StrokeWidth
	initialaxis.StrokeDashArray = initialaxisAPI.StrokeDashArray
	initialaxis.StrokeDashArrayWhenSelected = initialaxisAPI.StrokeDashArrayWhenSelected
	initialaxis.Transform = initialaxisAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
