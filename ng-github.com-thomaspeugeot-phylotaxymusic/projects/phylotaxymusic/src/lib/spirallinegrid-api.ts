// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { SpiralLineAPI } from './spiralline-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralLineGridAPI {

	static GONGSTRUCT_NAME = "SpiralLineGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	SpiralLineGridPointersEncoding: SpiralLineGridPointersEncoding = new SpiralLineGridPointersEncoding
}

export class SpiralLineGridPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	SpiralLines: number[] = []
}
