// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { BezierGridAPI } from './beziergrid-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierGridStackAPI {

	static GONGSTRUCT_NAME = "BezierGridStack"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	BezierGridStackPointersEncoding: BezierGridStackPointersEncoding = new BezierGridStackPointersEncoding
}

export class BezierGridStackPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	BezierGrids: number[] = []
}
