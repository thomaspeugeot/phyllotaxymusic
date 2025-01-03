// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralCircleAPI {

	static GONGSTRUCT_NAME = "SpiralCircle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	CenterX: number = 0
	CenterY: number = 0
	HasBespokeRadius: boolean = false
	BespopkeRadius: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	Pitch: number = 0
	ShowName: boolean = false
	BeatNb: number = 0
	Path: string = ""

	// insertion point for other decls

	SpiralCirclePointersEncoding: SpiralCirclePointersEncoding = new SpiralCirclePointersEncoding
}

export class SpiralCirclePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

}
