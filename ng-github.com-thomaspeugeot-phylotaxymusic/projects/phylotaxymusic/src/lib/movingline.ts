// generated code - do not edit

import { MovingLineAPI } from './movingline-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MovingLine {

	static GONGSTRUCT_NAME = "MovingLine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	AngleDegree: number = 0
	Length: number = 0
	CenterX: number = 0
	CenterY: number = 0
	StartX: number = 0
	EndX: number = 0
	DurationSeconds: number = 0
	IsMoving: boolean = false
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

export function CopyMovingLineToMovingLineAPI(movingline: MovingLine, movinglineAPI: MovingLineAPI) {

	movinglineAPI.CreatedAt = movingline.CreatedAt
	movinglineAPI.DeletedAt = movingline.DeletedAt
	movinglineAPI.ID = movingline.ID

	// insertion point for basic fields copy operations
	movinglineAPI.Name = movingline.Name
	movinglineAPI.IsDisplayed = movingline.IsDisplayed
	movinglineAPI.AngleDegree = movingline.AngleDegree
	movinglineAPI.Length = movingline.Length
	movinglineAPI.CenterX = movingline.CenterX
	movinglineAPI.CenterY = movingline.CenterY
	movinglineAPI.StartX = movingline.StartX
	movinglineAPI.EndX = movingline.EndX
	movinglineAPI.DurationSeconds = movingline.DurationSeconds
	movinglineAPI.IsMoving = movingline.IsMoving
	movinglineAPI.Color = movingline.Color
	movinglineAPI.FillOpacity = movingline.FillOpacity
	movinglineAPI.Stroke = movingline.Stroke
	movinglineAPI.StrokeOpacity = movingline.StrokeOpacity
	movinglineAPI.StrokeWidth = movingline.StrokeWidth
	movinglineAPI.StrokeDashArray = movingline.StrokeDashArray
	movinglineAPI.StrokeDashArrayWhenSelected = movingline.StrokeDashArrayWhenSelected
	movinglineAPI.Transform = movingline.Transform

	// insertion point for pointer fields encoding
	movinglineAPI.MovingLinePointersEncoding.ShapeCategoryID.Valid = true
	if (movingline.ShapeCategory != undefined) {
		movinglineAPI.MovingLinePointersEncoding.ShapeCategoryID.Int64 = movingline.ShapeCategory.ID  
	} else {
		movinglineAPI.MovingLinePointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyMovingLineAPIToMovingLine update basic, pointers and slice of pointers fields of movingline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of movinglineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMovingLineAPIToMovingLine(movinglineAPI: MovingLineAPI, movingline: MovingLine, frontRepo: FrontRepo) {

	movingline.CreatedAt = movinglineAPI.CreatedAt
	movingline.DeletedAt = movinglineAPI.DeletedAt
	movingline.ID = movinglineAPI.ID

	// insertion point for basic fields copy operations
	movingline.Name = movinglineAPI.Name
	movingline.IsDisplayed = movinglineAPI.IsDisplayed
	movingline.AngleDegree = movinglineAPI.AngleDegree
	movingline.Length = movinglineAPI.Length
	movingline.CenterX = movinglineAPI.CenterX
	movingline.CenterY = movinglineAPI.CenterY
	movingline.StartX = movinglineAPI.StartX
	movingline.EndX = movinglineAPI.EndX
	movingline.DurationSeconds = movinglineAPI.DurationSeconds
	movingline.IsMoving = movinglineAPI.IsMoving
	movingline.Color = movinglineAPI.Color
	movingline.FillOpacity = movinglineAPI.FillOpacity
	movingline.Stroke = movinglineAPI.Stroke
	movingline.StrokeOpacity = movinglineAPI.StrokeOpacity
	movingline.StrokeWidth = movinglineAPI.StrokeWidth
	movingline.StrokeDashArray = movinglineAPI.StrokeDashArray
	movingline.StrokeDashArrayWhenSelected = movinglineAPI.StrokeDashArrayWhenSelected
	movingline.Transform = movinglineAPI.Transform

	// insertion point for pointer fields encoding
	movingline.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(movinglineAPI.MovingLinePointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
