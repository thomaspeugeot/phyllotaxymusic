// generated code - do not edit

import { AxisAPI } from './axis-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Axis {

	static GONGSTRUCT_NAME = "Axis"

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

export function CopyAxisToAxisAPI(axis: Axis, axisAPI: AxisAPI) {

	axisAPI.CreatedAt = axis.CreatedAt
	axisAPI.DeletedAt = axis.DeletedAt
	axisAPI.ID = axis.ID

	// insertion point for basic fields copy operations
	axisAPI.Name = axis.Name
	axisAPI.IsDisplayed = axis.IsDisplayed
	axisAPI.Angle = axis.Angle
	axisAPI.Length = axis.Length
	axisAPI.Color = axis.Color
	axisAPI.FillOpacity = axis.FillOpacity
	axisAPI.Stroke = axis.Stroke
	axisAPI.StrokeOpacity = axis.StrokeOpacity
	axisAPI.StrokeWidth = axis.StrokeWidth
	axisAPI.StrokeDashArray = axis.StrokeDashArray
	axisAPI.StrokeDashArrayWhenSelected = axis.StrokeDashArrayWhenSelected
	axisAPI.Transform = axis.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAxisAPIToAxis update basic, pointers and slice of pointers fields of axis
// from respectively the basic fields and encoded fields of pointers and slices of pointers of axisAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAxisAPIToAxis(axisAPI: AxisAPI, axis: Axis, frontRepo: FrontRepo) {

	axis.CreatedAt = axisAPI.CreatedAt
	axis.DeletedAt = axisAPI.DeletedAt
	axis.ID = axisAPI.ID

	// insertion point for basic fields copy operations
	axis.Name = axisAPI.Name
	axis.IsDisplayed = axisAPI.IsDisplayed
	axis.Angle = axisAPI.Angle
	axis.Length = axisAPI.Length
	axis.Color = axisAPI.Color
	axis.FillOpacity = axisAPI.FillOpacity
	axis.Stroke = axisAPI.Stroke
	axis.StrokeOpacity = axisAPI.StrokeOpacity
	axis.StrokeWidth = axisAPI.StrokeWidth
	axis.StrokeDashArray = axisAPI.StrokeDashArray
	axis.StrokeDashArrayWhenSelected = axisAPI.StrokeDashArrayWhenSelected
	axis.Transform = axisAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}