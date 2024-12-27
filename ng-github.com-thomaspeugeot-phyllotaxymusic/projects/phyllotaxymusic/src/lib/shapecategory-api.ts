// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ShapeCategoryAPI {

	static GONGSTRUCT_NAME = "ShapeCategory"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsExpanded: boolean = false

	// insertion point for other decls

	ShapeCategoryPointersEncoding: ShapeCategoryPointersEncoding = new ShapeCategoryPointersEncoding
}

export class ShapeCategoryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
