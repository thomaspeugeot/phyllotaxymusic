// generated code - do not edit

import { FrontCurveStackAPI } from './frontcurvestack-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { FrontCurve } from './frontcurve'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FrontCurveStack {

	static GONGSTRUCT_NAME = "FrontCurveStack"

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

	FrontCurves: Array<FrontCurve> = []
}

export function CopyFrontCurveStackToFrontCurveStackAPI(frontcurvestack: FrontCurveStack, frontcurvestackAPI: FrontCurveStackAPI) {

	frontcurvestackAPI.CreatedAt = frontcurvestack.CreatedAt
	frontcurvestackAPI.DeletedAt = frontcurvestack.DeletedAt
	frontcurvestackAPI.ID = frontcurvestack.ID

	// insertion point for basic fields copy operations
	frontcurvestackAPI.Name = frontcurvestack.Name
	frontcurvestackAPI.IsDisplayed = frontcurvestack.IsDisplayed
	frontcurvestackAPI.Color = frontcurvestack.Color
	frontcurvestackAPI.FillOpacity = frontcurvestack.FillOpacity
	frontcurvestackAPI.Stroke = frontcurvestack.Stroke
	frontcurvestackAPI.StrokeOpacity = frontcurvestack.StrokeOpacity
	frontcurvestackAPI.StrokeWidth = frontcurvestack.StrokeWidth
	frontcurvestackAPI.StrokeDashArray = frontcurvestack.StrokeDashArray
	frontcurvestackAPI.StrokeDashArrayWhenSelected = frontcurvestack.StrokeDashArrayWhenSelected
	frontcurvestackAPI.Transform = frontcurvestack.Transform

	// insertion point for pointer fields encoding
	frontcurvestackAPI.FrontCurveStackPointersEncoding.ShapeCategoryID.Valid = true
	if (frontcurvestack.ShapeCategory != undefined) {
		frontcurvestackAPI.FrontCurveStackPointersEncoding.ShapeCategoryID.Int64 = frontcurvestack.ShapeCategory.ID  
	} else {
		frontcurvestackAPI.FrontCurveStackPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	frontcurvestackAPI.FrontCurveStackPointersEncoding.FrontCurves = []
	for (let _frontcurve of frontcurvestack.FrontCurves) {
		frontcurvestackAPI.FrontCurveStackPointersEncoding.FrontCurves.push(_frontcurve.ID)
	}

}

// CopyFrontCurveStackAPIToFrontCurveStack update basic, pointers and slice of pointers fields of frontcurvestack
// from respectively the basic fields and encoded fields of pointers and slices of pointers of frontcurvestackAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFrontCurveStackAPIToFrontCurveStack(frontcurvestackAPI: FrontCurveStackAPI, frontcurvestack: FrontCurveStack, frontRepo: FrontRepo) {

	frontcurvestack.CreatedAt = frontcurvestackAPI.CreatedAt
	frontcurvestack.DeletedAt = frontcurvestackAPI.DeletedAt
	frontcurvestack.ID = frontcurvestackAPI.ID

	// insertion point for basic fields copy operations
	frontcurvestack.Name = frontcurvestackAPI.Name
	frontcurvestack.IsDisplayed = frontcurvestackAPI.IsDisplayed
	frontcurvestack.Color = frontcurvestackAPI.Color
	frontcurvestack.FillOpacity = frontcurvestackAPI.FillOpacity
	frontcurvestack.Stroke = frontcurvestackAPI.Stroke
	frontcurvestack.StrokeOpacity = frontcurvestackAPI.StrokeOpacity
	frontcurvestack.StrokeWidth = frontcurvestackAPI.StrokeWidth
	frontcurvestack.StrokeDashArray = frontcurvestackAPI.StrokeDashArray
	frontcurvestack.StrokeDashArrayWhenSelected = frontcurvestackAPI.StrokeDashArrayWhenSelected
	frontcurvestack.Transform = frontcurvestackAPI.Transform

	// insertion point for pointer fields encoding
	frontcurvestack.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(frontcurvestackAPI.FrontCurveStackPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	frontcurvestack.FrontCurves = new Array<FrontCurve>()
	for (let _id of frontcurvestackAPI.FrontCurveStackPointersEncoding.FrontCurves) {
		let _frontcurve = frontRepo.map_ID_FrontCurve.get(_id)
		if (_frontcurve != undefined) {
			frontcurvestack.FrontCurves.push(_frontcurve!)
		}
	}
}
