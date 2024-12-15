// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MovingLineAPI {

	static GONGSTRUCT_NAME = "MovingLine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	AngleDegree: number = 0
	Length: number = 0
	CenterX: number = 0
	CenterY: number = 0
	SpeedX: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	MovingLinePointersEncoding: MovingLinePointersEncoding = new MovingLinePointersEncoding
}

export class MovingLinePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

}
