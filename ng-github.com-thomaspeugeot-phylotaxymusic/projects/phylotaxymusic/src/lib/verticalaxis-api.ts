// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VerticalAxisAPI {

	static GONGSTRUCT_NAME = "VerticalAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	AxisHandleBorderLength: number = 0
	Axis_Length: number = 0
	StrokeWidth: number = 0

	// insertion point for other decls

	VerticalAxisPointersEncoding: VerticalAxisPointersEncoding = new VerticalAxisPointersEncoding
}

export class VerticalAxisPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
