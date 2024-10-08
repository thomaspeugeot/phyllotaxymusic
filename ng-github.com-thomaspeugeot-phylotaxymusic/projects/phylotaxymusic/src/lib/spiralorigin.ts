// generated code - do not edit

import { SpiralOriginAPI } from './spiralorigin-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralOrigin {

	static GONGSTRUCT_NAME = "SpiralOrigin"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
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

export function CopySpiralOriginToSpiralOriginAPI(spiralorigin: SpiralOrigin, spiraloriginAPI: SpiralOriginAPI) {

	spiraloriginAPI.CreatedAt = spiralorigin.CreatedAt
	spiraloriginAPI.DeletedAt = spiralorigin.DeletedAt
	spiraloriginAPI.ID = spiralorigin.ID

	// insertion point for basic fields copy operations
	spiraloriginAPI.Name = spiralorigin.Name
	spiraloriginAPI.IsDisplayed = spiralorigin.IsDisplayed
	spiraloriginAPI.Color = spiralorigin.Color
	spiraloriginAPI.FillOpacity = spiralorigin.FillOpacity
	spiraloriginAPI.Stroke = spiralorigin.Stroke
	spiraloriginAPI.StrokeOpacity = spiralorigin.StrokeOpacity
	spiraloriginAPI.StrokeWidth = spiralorigin.StrokeWidth
	spiraloriginAPI.StrokeDashArray = spiralorigin.StrokeDashArray
	spiraloriginAPI.StrokeDashArrayWhenSelected = spiralorigin.StrokeDashArrayWhenSelected
	spiraloriginAPI.Transform = spiralorigin.Transform

	// insertion point for pointer fields encoding
	spiraloriginAPI.SpiralOriginPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralorigin.ShapeCategory != undefined) {
		spiraloriginAPI.SpiralOriginPointersEncoding.ShapeCategoryID.Int64 = spiralorigin.ShapeCategory.ID  
	} else {
		spiraloriginAPI.SpiralOriginPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralOriginAPIToSpiralOrigin update basic, pointers and slice of pointers fields of spiralorigin
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiraloriginAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralOriginAPIToSpiralOrigin(spiraloriginAPI: SpiralOriginAPI, spiralorigin: SpiralOrigin, frontRepo: FrontRepo) {

	spiralorigin.CreatedAt = spiraloriginAPI.CreatedAt
	spiralorigin.DeletedAt = spiraloriginAPI.DeletedAt
	spiralorigin.ID = spiraloriginAPI.ID

	// insertion point for basic fields copy operations
	spiralorigin.Name = spiraloriginAPI.Name
	spiralorigin.IsDisplayed = spiraloriginAPI.IsDisplayed
	spiralorigin.Color = spiraloriginAPI.Color
	spiralorigin.FillOpacity = spiraloriginAPI.FillOpacity
	spiralorigin.Stroke = spiraloriginAPI.Stroke
	spiralorigin.StrokeOpacity = spiraloriginAPI.StrokeOpacity
	spiralorigin.StrokeWidth = spiraloriginAPI.StrokeWidth
	spiralorigin.StrokeDashArray = spiraloriginAPI.StrokeDashArray
	spiralorigin.StrokeDashArrayWhenSelected = spiraloriginAPI.StrokeDashArrayWhenSelected
	spiralorigin.Transform = spiraloriginAPI.Transform

	// insertion point for pointer fields encoding
	spiralorigin.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiraloriginAPI.SpiralOriginPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
