// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RhombusAPI {

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

	// insertion point for other decls

	RhombusPointersEncoding: RhombusPointersEncoding = new RhombusPointersEncoding
}

export class RhombusPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
