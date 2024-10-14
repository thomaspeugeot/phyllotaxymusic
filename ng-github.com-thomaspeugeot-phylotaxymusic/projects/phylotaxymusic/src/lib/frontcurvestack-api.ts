// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'
import { FrontCurveAPI } from './frontcurve-api'
import { SpiralCircleAPI } from './spiralcircle-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FrontCurveStackAPI {

	static GONGSTRUCT_NAME = "FrontCurveStack"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	FrontCurveStackPointersEncoding: FrontCurveStackPointersEncoding = new FrontCurveStackPointersEncoding
}

export class FrontCurveStackPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

	FrontCurves: number[] = []
	SpiralCircles: number[] = []
}
