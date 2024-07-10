// insertion point for imports
import { RhombusAPI } from './rhombus-api'
import { CircleAPI } from './circle-api'
import { RhombusGridAPI } from './rhombusgrid-api'
import { CircleGridAPI } from './circlegrid-api'
import { AxisAPI } from './axis-api'
import { HorizontalAxisAPI } from './horizontalaxis-api'
import { VerticalAxisAPI } from './verticalaxis-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ParameterAPI {

	static GONGSTRUCT_NAME = "Parameter"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	N: number = 0
	M: number = 0
	InsideAngle: number = 0
	SideLength: number = 0
	OriginX: number = 0
	OriginY: number = 0

	// insertion point for other decls

	ParameterPointersEncoding: ParameterPointersEncoding = new ParameterPointersEncoding
}

export class ParameterPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	InitialRhombusID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombus.ID = 0

	InitialCircleID: NullInt64 = new NullInt64 // if pointer is null, InitialCircle.ID = 0

	InitialRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombusGrid.ID = 0

	InitialCircleGridID: NullInt64 = new NullInt64 // if pointer is null, InitialCircleGrid.ID = 0

	InitialAxisID: NullInt64 = new NullInt64 // if pointer is null, InitialAxis.ID = 0

	RotatedAxisID: NullInt64 = new NullInt64 // if pointer is null, RotatedAxis.ID = 0

	RotatedRhombusID: NullInt64 = new NullInt64 // if pointer is null, RotatedRhombus.ID = 0

	RotatedRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, RotatedRhombusGrid.ID = 0

	RotatedCircleGridID: NullInt64 = new NullInt64 // if pointer is null, RotatedCircleGrid.ID = 0

	NextRhombusID: NullInt64 = new NullInt64 // if pointer is null, NextRhombus.ID = 0

	NextCircleID: NullInt64 = new NullInt64 // if pointer is null, NextCircle.ID = 0

	HorizontalAxisID: NullInt64 = new NullInt64 // if pointer is null, HorizontalAxis.ID = 0

	VerticalAxisID: NullInt64 = new NullInt64 // if pointer is null, VerticalAxis.ID = 0

}
