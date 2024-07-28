// generated code - do not edit

import { VerticalAxisAPI } from './verticalaxis-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VerticalAxis {

	static GONGSTRUCT_NAME = "VerticalAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	AxisHandleBorderLength: number = 0
	Axis_Length: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

}

export function CopyVerticalAxisToVerticalAxisAPI(verticalaxis: VerticalAxis, verticalaxisAPI: VerticalAxisAPI) {

	verticalaxisAPI.CreatedAt = verticalaxis.CreatedAt
	verticalaxisAPI.DeletedAt = verticalaxis.DeletedAt
	verticalaxisAPI.ID = verticalaxis.ID

	// insertion point for basic fields copy operations
	verticalaxisAPI.Name = verticalaxis.Name
	verticalaxisAPI.IsDisplayed = verticalaxis.IsDisplayed
	verticalaxisAPI.AxisHandleBorderLength = verticalaxis.AxisHandleBorderLength
	verticalaxisAPI.Axis_Length = verticalaxis.Axis_Length
	verticalaxisAPI.Color = verticalaxis.Color
	verticalaxisAPI.FillOpacity = verticalaxis.FillOpacity
	verticalaxisAPI.Stroke = verticalaxis.Stroke
	verticalaxisAPI.StrokeOpacity = verticalaxis.StrokeOpacity
	verticalaxisAPI.StrokeWidth = verticalaxis.StrokeWidth
	verticalaxisAPI.StrokeDashArray = verticalaxis.StrokeDashArray
	verticalaxisAPI.StrokeDashArrayWhenSelected = verticalaxis.StrokeDashArrayWhenSelected
	verticalaxisAPI.Transform = verticalaxis.Transform

	// insertion point for pointer fields encoding
	verticalaxisAPI.VerticalAxisPointersEncoding.ShapeCategoryID.Valid = true
	if (verticalaxis.ShapeCategory != undefined) {
		verticalaxisAPI.VerticalAxisPointersEncoding.ShapeCategoryID.Int64 = verticalaxis.ShapeCategory.ID  
	} else {
		verticalaxisAPI.VerticalAxisPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyVerticalAxisAPIToVerticalAxis update basic, pointers and slice of pointers fields of verticalaxis
// from respectively the basic fields and encoded fields of pointers and slices of pointers of verticalaxisAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVerticalAxisAPIToVerticalAxis(verticalaxisAPI: VerticalAxisAPI, verticalaxis: VerticalAxis, frontRepo: FrontRepo) {

	verticalaxis.CreatedAt = verticalaxisAPI.CreatedAt
	verticalaxis.DeletedAt = verticalaxisAPI.DeletedAt
	verticalaxis.ID = verticalaxisAPI.ID

	// insertion point for basic fields copy operations
	verticalaxis.Name = verticalaxisAPI.Name
	verticalaxis.IsDisplayed = verticalaxisAPI.IsDisplayed
	verticalaxis.AxisHandleBorderLength = verticalaxisAPI.AxisHandleBorderLength
	verticalaxis.Axis_Length = verticalaxisAPI.Axis_Length
	verticalaxis.Color = verticalaxisAPI.Color
	verticalaxis.FillOpacity = verticalaxisAPI.FillOpacity
	verticalaxis.Stroke = verticalaxisAPI.Stroke
	verticalaxis.StrokeOpacity = verticalaxisAPI.StrokeOpacity
	verticalaxis.StrokeWidth = verticalaxisAPI.StrokeWidth
	verticalaxis.StrokeDashArray = verticalaxisAPI.StrokeDashArray
	verticalaxis.StrokeDashArrayWhenSelected = verticalaxisAPI.StrokeDashArrayWhenSelected
	verticalaxis.Transform = verticalaxisAPI.Transform

	// insertion point for pointer fields encoding
	verticalaxis.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(verticalaxisAPI.VerticalAxisPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
