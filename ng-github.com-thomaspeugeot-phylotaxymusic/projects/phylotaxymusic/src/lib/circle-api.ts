// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CircleAPI {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	CenterX: number = 0
	CenterY: number = 0
	StrokeWidth: number = 0

	// insertion point for other decls

	CirclePointersEncoding: CirclePointersEncoding = new CirclePointersEncoding
}

export class CirclePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
