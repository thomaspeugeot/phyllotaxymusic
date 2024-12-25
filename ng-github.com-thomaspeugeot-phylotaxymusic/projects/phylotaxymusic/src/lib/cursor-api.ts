// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CursorAPI {

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

	// insertion point for other decls

	CursorPointersEncoding: CursorPointersEncoding = new CursorPointersEncoding
}

export class CursorPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
