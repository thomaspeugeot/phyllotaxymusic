// generated code - do not edit

import { BezierAPI } from './bezier-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Bezier {

	static GONGSTRUCT_NAME = "Bezier"

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

export function CopyBezierToBezierAPI(bezier: Bezier, bezierAPI: BezierAPI) {

	bezierAPI.CreatedAt = bezier.CreatedAt
	bezierAPI.DeletedAt = bezier.DeletedAt
	bezierAPI.ID = bezier.ID

	// insertion point for basic fields copy operations
	bezierAPI.Name = bezier.Name
	bezierAPI.IsDisplayed = bezier.IsDisplayed
	bezierAPI.StartX = bezier.StartX
	bezierAPI.StartY = bezier.StartY
	bezierAPI.ControlPointStartX = bezier.ControlPointStartX
	bezierAPI.ControlPointStartY = bezier.ControlPointStartY
	bezierAPI.EndX = bezier.EndX
	bezierAPI.EndY = bezier.EndY
	bezierAPI.ControlPointEndX = bezier.ControlPointEndX
	bezierAPI.ControlPointEndY = bezier.ControlPointEndY
	bezierAPI.Color = bezier.Color
	bezierAPI.FillOpacity = bezier.FillOpacity
	bezierAPI.Stroke = bezier.Stroke
	bezierAPI.StrokeOpacity = bezier.StrokeOpacity
	bezierAPI.StrokeWidth = bezier.StrokeWidth
	bezierAPI.StrokeDashArray = bezier.StrokeDashArray
	bezierAPI.StrokeDashArrayWhenSelected = bezier.StrokeDashArrayWhenSelected
	bezierAPI.Transform = bezier.Transform

	// insertion point for pointer fields encoding
	bezierAPI.BezierPointersEncoding.ShapeCategoryID.Valid = true
	if (bezier.ShapeCategory != undefined) {
		bezierAPI.BezierPointersEncoding.ShapeCategoryID.Int64 = bezier.ShapeCategory.ID  
	} else {
		bezierAPI.BezierPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyBezierAPIToBezier update basic, pointers and slice of pointers fields of bezier
// from respectively the basic fields and encoded fields of pointers and slices of pointers of bezierAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBezierAPIToBezier(bezierAPI: BezierAPI, bezier: Bezier, frontRepo: FrontRepo) {

	bezier.CreatedAt = bezierAPI.CreatedAt
	bezier.DeletedAt = bezierAPI.DeletedAt
	bezier.ID = bezierAPI.ID

	// insertion point for basic fields copy operations
	bezier.Name = bezierAPI.Name
	bezier.IsDisplayed = bezierAPI.IsDisplayed
	bezier.StartX = bezierAPI.StartX
	bezier.StartY = bezierAPI.StartY
	bezier.ControlPointStartX = bezierAPI.ControlPointStartX
	bezier.ControlPointStartY = bezierAPI.ControlPointStartY
	bezier.EndX = bezierAPI.EndX
	bezier.EndY = bezierAPI.EndY
	bezier.ControlPointEndX = bezierAPI.ControlPointEndX
	bezier.ControlPointEndY = bezierAPI.ControlPointEndY
	bezier.Color = bezierAPI.Color
	bezier.FillOpacity = bezierAPI.FillOpacity
	bezier.Stroke = bezierAPI.Stroke
	bezier.StrokeOpacity = bezierAPI.StrokeOpacity
	bezier.StrokeWidth = bezierAPI.StrokeWidth
	bezier.StrokeDashArray = bezierAPI.StrokeDashArray
	bezier.StrokeDashArrayWhenSelected = bezierAPI.StrokeDashArrayWhenSelected
	bezier.Transform = bezierAPI.Transform

	// insertion point for pointer fields encoding
	bezier.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(bezierAPI.BezierPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
