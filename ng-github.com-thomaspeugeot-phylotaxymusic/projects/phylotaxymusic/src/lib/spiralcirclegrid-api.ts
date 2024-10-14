// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { SpiralCircleAPI } from './spiralcircle-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralCircleGridAPI {

	static GONGSTRUCT_NAME = "SpiralCircleGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	SpiralCircleGridPointersEncoding: SpiralCircleGridPointersEncoding = new SpiralCircleGridPointersEncoding
}

export class SpiralCircleGridPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	SpiralRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, SpiralRhombusGrid.ID = 0

	SpiralCircles: number[] = []
}
