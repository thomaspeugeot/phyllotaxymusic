// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { RhombusGridAPI } from './rhombusgrid-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralRhombusGridAPI {

	static GONGSTRUCT_NAME = "SpiralRhombusGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	SpiralRhombusGridPointersEncoding: SpiralRhombusGridPointersEncoding = new SpiralRhombusGridPointersEncoding
}

export class SpiralRhombusGridPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	RhombusGridID: NullInt64 = new NullInt64 // if pointer is null, RhombusGrid.ID = 0

}
