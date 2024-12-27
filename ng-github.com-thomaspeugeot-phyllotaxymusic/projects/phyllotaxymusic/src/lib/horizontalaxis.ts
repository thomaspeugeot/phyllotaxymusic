// generated code - do not edit

import { HorizontalAxisAPI } from './horizontalaxis-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class HorizontalAxis {

	static GONGSTRUCT_NAME = "HorizontalAxis"

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

export function CopyHorizontalAxisToHorizontalAxisAPI(horizontalaxis: HorizontalAxis, horizontalaxisAPI: HorizontalAxisAPI) {

	horizontalaxisAPI.CreatedAt = horizontalaxis.CreatedAt
	horizontalaxisAPI.DeletedAt = horizontalaxis.DeletedAt
	horizontalaxisAPI.ID = horizontalaxis.ID

	// insertion point for basic fields copy operations
	horizontalaxisAPI.Name = horizontalaxis.Name
	horizontalaxisAPI.IsDisplayed = horizontalaxis.IsDisplayed
	horizontalaxisAPI.AxisHandleBorderLength = horizontalaxis.AxisHandleBorderLength
	horizontalaxisAPI.Axis_Length = horizontalaxis.Axis_Length
	horizontalaxisAPI.Color = horizontalaxis.Color
	horizontalaxisAPI.FillOpacity = horizontalaxis.FillOpacity
	horizontalaxisAPI.Stroke = horizontalaxis.Stroke
	horizontalaxisAPI.StrokeOpacity = horizontalaxis.StrokeOpacity
	horizontalaxisAPI.StrokeWidth = horizontalaxis.StrokeWidth
	horizontalaxisAPI.StrokeDashArray = horizontalaxis.StrokeDashArray
	horizontalaxisAPI.StrokeDashArrayWhenSelected = horizontalaxis.StrokeDashArrayWhenSelected
	horizontalaxisAPI.Transform = horizontalaxis.Transform

	// insertion point for pointer fields encoding
	horizontalaxisAPI.HorizontalAxisPointersEncoding.ShapeCategoryID.Valid = true
	if (horizontalaxis.ShapeCategory != undefined) {
		horizontalaxisAPI.HorizontalAxisPointersEncoding.ShapeCategoryID.Int64 = horizontalaxis.ShapeCategory.ID  
	} else {
		horizontalaxisAPI.HorizontalAxisPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyHorizontalAxisAPIToHorizontalAxis update basic, pointers and slice of pointers fields of horizontalaxis
// from respectively the basic fields and encoded fields of pointers and slices of pointers of horizontalaxisAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyHorizontalAxisAPIToHorizontalAxis(horizontalaxisAPI: HorizontalAxisAPI, horizontalaxis: HorizontalAxis, frontRepo: FrontRepo) {

	horizontalaxis.CreatedAt = horizontalaxisAPI.CreatedAt
	horizontalaxis.DeletedAt = horizontalaxisAPI.DeletedAt
	horizontalaxis.ID = horizontalaxisAPI.ID

	// insertion point for basic fields copy operations
	horizontalaxis.Name = horizontalaxisAPI.Name
	horizontalaxis.IsDisplayed = horizontalaxisAPI.IsDisplayed
	horizontalaxis.AxisHandleBorderLength = horizontalaxisAPI.AxisHandleBorderLength
	horizontalaxis.Axis_Length = horizontalaxisAPI.Axis_Length
	horizontalaxis.Color = horizontalaxisAPI.Color
	horizontalaxis.FillOpacity = horizontalaxisAPI.FillOpacity
	horizontalaxis.Stroke = horizontalaxisAPI.Stroke
	horizontalaxis.StrokeOpacity = horizontalaxisAPI.StrokeOpacity
	horizontalaxis.StrokeWidth = horizontalaxisAPI.StrokeWidth
	horizontalaxis.StrokeDashArray = horizontalaxisAPI.StrokeDashArray
	horizontalaxis.StrokeDashArrayWhenSelected = horizontalaxisAPI.StrokeDashArrayWhenSelected
	horizontalaxis.Transform = horizontalaxisAPI.Transform

	// insertion point for pointer fields encoding
	horizontalaxis.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(horizontalaxisAPI.HorizontalAxisPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
