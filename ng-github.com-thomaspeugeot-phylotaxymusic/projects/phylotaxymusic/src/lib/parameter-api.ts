// insertion point for imports
import { RhombusAPI } from './rhombus-api'
import { RhombusGridAPI } from './rhombusgrid-api'
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
	DiamondSideLenght: number = 0
	OriginX: number = 0
	OriginY: number = 0

	// insertion point for other decls

	ParameterPointersEncoding: ParameterPointersEncoding = new ParameterPointersEncoding
}

export class ParameterPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	InitialRhombusID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombus.ID = 0

	InitialRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombusGrid.ID = 0

	HorizontalAxisID: NullInt64 = new NullInt64 // if pointer is null, HorizontalAxis.ID = 0

	VerticalAxisID: NullInt64 = new NullInt64 // if pointer is null, VerticalAxis.ID = 0

}
