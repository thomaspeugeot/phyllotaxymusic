// generated code - do not edit

import { RhombusAPI } from './rhombus-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Rhombus {

	static GONGSTRUCT_NAME = "Rhombus"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	CenterX: number = 0
	CenterY: number = 0
	SideLength: number = 0
	Angle: number = 0
	InsideAngle: number = 0
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

export function CopyRhombusToRhombusAPI(rhombus: Rhombus, rhombusAPI: RhombusAPI) {

	rhombusAPI.CreatedAt = rhombus.CreatedAt
	rhombusAPI.DeletedAt = rhombus.DeletedAt
	rhombusAPI.ID = rhombus.ID

	// insertion point for basic fields copy operations
	rhombusAPI.Name = rhombus.Name
	rhombusAPI.IsDisplayed = rhombus.IsDisplayed
	rhombusAPI.CenterX = rhombus.CenterX
	rhombusAPI.CenterY = rhombus.CenterY
	rhombusAPI.SideLength = rhombus.SideLength
	rhombusAPI.Angle = rhombus.Angle
	rhombusAPI.InsideAngle = rhombus.InsideAngle
	rhombusAPI.Color = rhombus.Color
	rhombusAPI.FillOpacity = rhombus.FillOpacity
	rhombusAPI.Stroke = rhombus.Stroke
	rhombusAPI.StrokeOpacity = rhombus.StrokeOpacity
	rhombusAPI.StrokeWidth = rhombus.StrokeWidth
	rhombusAPI.StrokeDashArray = rhombus.StrokeDashArray
	rhombusAPI.StrokeDashArrayWhenSelected = rhombus.StrokeDashArrayWhenSelected
	rhombusAPI.Transform = rhombus.Transform

	// insertion point for pointer fields encoding
	rhombusAPI.RhombusPointersEncoding.ShapeCategoryID.Valid = true
	if (rhombus.ShapeCategory != undefined) {
		rhombusAPI.RhombusPointersEncoding.ShapeCategoryID.Int64 = rhombus.ShapeCategory.ID  
	} else {
		rhombusAPI.RhombusPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyRhombusAPIToRhombus update basic, pointers and slice of pointers fields of rhombus
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rhombusAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRhombusAPIToRhombus(rhombusAPI: RhombusAPI, rhombus: Rhombus, frontRepo: FrontRepo) {

	rhombus.CreatedAt = rhombusAPI.CreatedAt
	rhombus.DeletedAt = rhombusAPI.DeletedAt
	rhombus.ID = rhombusAPI.ID

	// insertion point for basic fields copy operations
	rhombus.Name = rhombusAPI.Name
	rhombus.IsDisplayed = rhombusAPI.IsDisplayed
	rhombus.CenterX = rhombusAPI.CenterX
	rhombus.CenterY = rhombusAPI.CenterY
	rhombus.SideLength = rhombusAPI.SideLength
	rhombus.Angle = rhombusAPI.Angle
	rhombus.InsideAngle = rhombusAPI.InsideAngle
	rhombus.Color = rhombusAPI.Color
	rhombus.FillOpacity = rhombusAPI.FillOpacity
	rhombus.Stroke = rhombusAPI.Stroke
	rhombus.StrokeOpacity = rhombusAPI.StrokeOpacity
	rhombus.StrokeWidth = rhombusAPI.StrokeWidth
	rhombus.StrokeDashArray = rhombusAPI.StrokeDashArray
	rhombus.StrokeDashArrayWhenSelected = rhombusAPI.StrokeDashArrayWhenSelected
	rhombus.Transform = rhombusAPI.Transform

	// insertion point for pointer fields encoding
	rhombus.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(rhombusAPI.RhombusPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
