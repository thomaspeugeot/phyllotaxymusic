// generated code - do not edit

import { CursorAPI } from './cursor-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Cursor {

	static GONGSTRUCT_NAME = "Cursor"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
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
}

export function CopyCursorToCursorAPI(cursor: Cursor, cursorAPI: CursorAPI) {

	cursorAPI.CreatedAt = cursor.CreatedAt
	cursorAPI.DeletedAt = cursor.DeletedAt
	cursorAPI.ID = cursor.ID

	// insertion point for basic fields copy operations
	cursorAPI.Name = cursor.Name
	cursorAPI.AngleDegree = cursor.AngleDegree
	cursorAPI.Length = cursor.Length
	cursorAPI.CenterX = cursor.CenterX
	cursorAPI.CenterY = cursor.CenterY
	cursorAPI.StartX = cursor.StartX
	cursorAPI.EndX = cursor.EndX
	cursorAPI.DurationSeconds = cursor.DurationSeconds
	cursorAPI.IsMoving = cursor.IsMoving
	cursorAPI.Color = cursor.Color
	cursorAPI.FillOpacity = cursor.FillOpacity
	cursorAPI.Stroke = cursor.Stroke
	cursorAPI.StrokeOpacity = cursor.StrokeOpacity
	cursorAPI.StrokeWidth = cursor.StrokeWidth
	cursorAPI.StrokeDashArray = cursor.StrokeDashArray
	cursorAPI.StrokeDashArrayWhenSelected = cursor.StrokeDashArrayWhenSelected
	cursorAPI.Transform = cursor.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCursorAPIToCursor update basic, pointers and slice of pointers fields of cursor
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cursorAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCursorAPIToCursor(cursorAPI: CursorAPI, cursor: Cursor, frontRepo: FrontRepo) {

	cursor.CreatedAt = cursorAPI.CreatedAt
	cursor.DeletedAt = cursorAPI.DeletedAt
	cursor.ID = cursorAPI.ID

	// insertion point for basic fields copy operations
	cursor.Name = cursorAPI.Name
	cursor.AngleDegree = cursorAPI.AngleDegree
	cursor.Length = cursorAPI.Length
	cursor.CenterX = cursorAPI.CenterX
	cursor.CenterY = cursorAPI.CenterY
	cursor.StartX = cursorAPI.StartX
	cursor.EndX = cursorAPI.EndX
	cursor.DurationSeconds = cursorAPI.DurationSeconds
	cursor.IsMoving = cursorAPI.IsMoving
	cursor.Color = cursorAPI.Color
	cursor.FillOpacity = cursorAPI.FillOpacity
	cursor.Stroke = cursorAPI.Stroke
	cursor.StrokeOpacity = cursorAPI.StrokeOpacity
	cursor.StrokeWidth = cursorAPI.StrokeWidth
	cursor.StrokeDashArray = cursorAPI.StrokeDashArray
	cursor.StrokeDashArrayWhenSelected = cursorAPI.StrokeDashArrayWhenSelected
	cursor.Transform = cursorAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
