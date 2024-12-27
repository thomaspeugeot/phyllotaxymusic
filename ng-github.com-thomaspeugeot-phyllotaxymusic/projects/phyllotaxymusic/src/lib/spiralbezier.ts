// generated code - do not edit

import { SpiralBezierAPI } from './spiralbezier-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralBezier {

	static GONGSTRUCT_NAME = "SpiralBezier"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	StartX: number = 0
	StartY: number = 0
	ControlPointStartX: number = 0
	ControlPointStartY: number = 0
	EndX: number = 0
	EndY: number = 0
	ControlPointEndX: number = 0
	ControlPointEndY: number = 0
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

export function CopySpiralBezierToSpiralBezierAPI(spiralbezier: SpiralBezier, spiralbezierAPI: SpiralBezierAPI) {

	spiralbezierAPI.CreatedAt = spiralbezier.CreatedAt
	spiralbezierAPI.DeletedAt = spiralbezier.DeletedAt
	spiralbezierAPI.ID = spiralbezier.ID

	// insertion point for basic fields copy operations
	spiralbezierAPI.Name = spiralbezier.Name
	spiralbezierAPI.IsDisplayed = spiralbezier.IsDisplayed
	spiralbezierAPI.StartX = spiralbezier.StartX
	spiralbezierAPI.StartY = spiralbezier.StartY
	spiralbezierAPI.ControlPointStartX = spiralbezier.ControlPointStartX
	spiralbezierAPI.ControlPointStartY = spiralbezier.ControlPointStartY
	spiralbezierAPI.EndX = spiralbezier.EndX
	spiralbezierAPI.EndY = spiralbezier.EndY
	spiralbezierAPI.ControlPointEndX = spiralbezier.ControlPointEndX
	spiralbezierAPI.ControlPointEndY = spiralbezier.ControlPointEndY
	spiralbezierAPI.Color = spiralbezier.Color
	spiralbezierAPI.FillOpacity = spiralbezier.FillOpacity
	spiralbezierAPI.Stroke = spiralbezier.Stroke
	spiralbezierAPI.StrokeOpacity = spiralbezier.StrokeOpacity
	spiralbezierAPI.StrokeWidth = spiralbezier.StrokeWidth
	spiralbezierAPI.StrokeDashArray = spiralbezier.StrokeDashArray
	spiralbezierAPI.StrokeDashArrayWhenSelected = spiralbezier.StrokeDashArrayWhenSelected
	spiralbezierAPI.Transform = spiralbezier.Transform

	// insertion point for pointer fields encoding
	spiralbezierAPI.SpiralBezierPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralbezier.ShapeCategory != undefined) {
		spiralbezierAPI.SpiralBezierPointersEncoding.ShapeCategoryID.Int64 = spiralbezier.ShapeCategory.ID  
	} else {
		spiralbezierAPI.SpiralBezierPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralBezierAPIToSpiralBezier update basic, pointers and slice of pointers fields of spiralbezier
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralbezierAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralBezierAPIToSpiralBezier(spiralbezierAPI: SpiralBezierAPI, spiralbezier: SpiralBezier, frontRepo: FrontRepo) {

	spiralbezier.CreatedAt = spiralbezierAPI.CreatedAt
	spiralbezier.DeletedAt = spiralbezierAPI.DeletedAt
	spiralbezier.ID = spiralbezierAPI.ID

	// insertion point for basic fields copy operations
	spiralbezier.Name = spiralbezierAPI.Name
	spiralbezier.IsDisplayed = spiralbezierAPI.IsDisplayed
	spiralbezier.StartX = spiralbezierAPI.StartX
	spiralbezier.StartY = spiralbezierAPI.StartY
	spiralbezier.ControlPointStartX = spiralbezierAPI.ControlPointStartX
	spiralbezier.ControlPointStartY = spiralbezierAPI.ControlPointStartY
	spiralbezier.EndX = spiralbezierAPI.EndX
	spiralbezier.EndY = spiralbezierAPI.EndY
	spiralbezier.ControlPointEndX = spiralbezierAPI.ControlPointEndX
	spiralbezier.ControlPointEndY = spiralbezierAPI.ControlPointEndY
	spiralbezier.Color = spiralbezierAPI.Color
	spiralbezier.FillOpacity = spiralbezierAPI.FillOpacity
	spiralbezier.Stroke = spiralbezierAPI.Stroke
	spiralbezier.StrokeOpacity = spiralbezierAPI.StrokeOpacity
	spiralbezier.StrokeWidth = spiralbezierAPI.StrokeWidth
	spiralbezier.StrokeDashArray = spiralbezierAPI.StrokeDashArray
	spiralbezier.StrokeDashArrayWhenSelected = spiralbezierAPI.StrokeDashArrayWhenSelected
	spiralbezier.Transform = spiralbezierAPI.Transform

	// insertion point for pointer fields encoding
	spiralbezier.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralbezierAPI.SpiralBezierPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
