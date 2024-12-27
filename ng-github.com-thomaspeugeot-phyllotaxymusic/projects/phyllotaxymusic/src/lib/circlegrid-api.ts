// insertion point for imports
import { CircleAPI } from './circle-api'
import { ShapeCategoryAPI } from './shapecategory-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CircleGridAPI {

	static GONGSTRUCT_NAME = "CircleGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	CircleGridPointersEncoding: CircleGridPointersEncoding = new CircleGridPointersEncoding
}

export class CircleGridPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ReferenceID: NullInt64 = new NullInt64 // if pointer is null, Reference.ID = 0

	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	Circles: number[] = []
}
