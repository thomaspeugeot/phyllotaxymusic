// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { SpiralBezierAPI } from './spiralbezier-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralBezierGridAPI {

	static GONGSTRUCT_NAME = "SpiralBezierGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	SpiralBezierGridPointersEncoding: SpiralBezierGridPointersEncoding = new SpiralBezierGridPointersEncoding
}

export class SpiralBezierGridPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	SpiralBeziers: number[] = []
}
