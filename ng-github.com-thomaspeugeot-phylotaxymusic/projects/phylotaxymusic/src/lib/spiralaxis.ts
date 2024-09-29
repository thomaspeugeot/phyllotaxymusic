// generated code - do not edit

import { SpiralAxisAPI } from './spiralaxis-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralAxis {

	static GONGSTRUCT_NAME = "SpiralAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	Angle: number = 0
	Length: number = 0
	CenterX: number = 0
	CenterY: number = 0
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

export function CopySpiralAxisToSpiralAxisAPI(spiralaxis: SpiralAxis, spiralaxisAPI: SpiralAxisAPI) {

	spiralaxisAPI.CreatedAt = spiralaxis.CreatedAt
	spiralaxisAPI.DeletedAt = spiralaxis.DeletedAt
	spiralaxisAPI.ID = spiralaxis.ID

	// insertion point for basic fields copy operations
	spiralaxisAPI.Name = spiralaxis.Name
	spiralaxisAPI.IsDisplayed = spiralaxis.IsDisplayed
	spiralaxisAPI.Angle = spiralaxis.Angle
	spiralaxisAPI.Length = spiralaxis.Length
	spiralaxisAPI.CenterX = spiralaxis.CenterX
	spiralaxisAPI.CenterY = spiralaxis.CenterY
	spiralaxisAPI.Color = spiralaxis.Color
	spiralaxisAPI.FillOpacity = spiralaxis.FillOpacity
	spiralaxisAPI.Stroke = spiralaxis.Stroke
	spiralaxisAPI.StrokeOpacity = spiralaxis.StrokeOpacity
	spiralaxisAPI.StrokeWidth = spiralaxis.StrokeWidth
	spiralaxisAPI.StrokeDashArray = spiralaxis.StrokeDashArray
	spiralaxisAPI.StrokeDashArrayWhenSelected = spiralaxis.StrokeDashArrayWhenSelected
	spiralaxisAPI.Transform = spiralaxis.Transform

	// insertion point for pointer fields encoding
	spiralaxisAPI.SpiralAxisPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralaxis.ShapeCategory != undefined) {
		spiralaxisAPI.SpiralAxisPointersEncoding.ShapeCategoryID.Int64 = spiralaxis.ShapeCategory.ID  
	} else {
		spiralaxisAPI.SpiralAxisPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralAxisAPIToSpiralAxis update basic, pointers and slice of pointers fields of spiralaxis
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralaxisAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralAxisAPIToSpiralAxis(spiralaxisAPI: SpiralAxisAPI, spiralaxis: SpiralAxis, frontRepo: FrontRepo) {

	spiralaxis.CreatedAt = spiralaxisAPI.CreatedAt
	spiralaxis.DeletedAt = spiralaxisAPI.DeletedAt
	spiralaxis.ID = spiralaxisAPI.ID

	// insertion point for basic fields copy operations
	spiralaxis.Name = spiralaxisAPI.Name
	spiralaxis.IsDisplayed = spiralaxisAPI.IsDisplayed
	spiralaxis.Angle = spiralaxisAPI.Angle
	spiralaxis.Length = spiralaxisAPI.Length
	spiralaxis.CenterX = spiralaxisAPI.CenterX
	spiralaxis.CenterY = spiralaxisAPI.CenterY
	spiralaxis.Color = spiralaxisAPI.Color
	spiralaxis.FillOpacity = spiralaxisAPI.FillOpacity
	spiralaxis.Stroke = spiralaxisAPI.Stroke
	spiralaxis.StrokeOpacity = spiralaxisAPI.StrokeOpacity
	spiralaxis.StrokeWidth = spiralaxisAPI.StrokeWidth
	spiralaxis.StrokeDashArray = spiralaxisAPI.StrokeDashArray
	spiralaxis.StrokeDashArrayWhenSelected = spiralaxisAPI.StrokeDashArrayWhenSelected
	spiralaxis.Transform = spiralaxisAPI.Transform

	// insertion point for pointer fields encoding
	spiralaxis.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralaxisAPI.SpiralAxisPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
