// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { RhombusAPI } from './rhombus-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralRhombusAPI {

	static GONGSTRUCT_NAME = "SpiralRhombus"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	SpiralRhombusPointersEncoding: SpiralRhombusPointersEncoding = new SpiralRhombusPointersEncoding
}

export class SpiralRhombusPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	RhombusID: NullInt64 = new NullInt64 // if pointer is null, Rhombus.ID = 0

}
