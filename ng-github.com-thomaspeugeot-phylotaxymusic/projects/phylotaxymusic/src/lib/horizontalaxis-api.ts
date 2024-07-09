// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class HorizontalAxisAPI {

	static GONGSTRUCT_NAME = "HorizontalAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	AxisHandleBorderLength: number = 0
	Axis_Length: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	HorizontalAxisPointersEncoding: HorizontalAxisPointersEncoding = new HorizontalAxisPointersEncoding
}

export class HorizontalAxisPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
