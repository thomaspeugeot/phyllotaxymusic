// insertion point for imports
import { RhombusAPI } from './rhombus-api'

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
	Angle: number = 0
	DiamondSideLenght: number = 0
	IsHorizontalAxisDisplayed: boolean = false
	AxisHandleBorderLength: number = 0
	OriginX: number = 0
	OriginY: number = 0
	HorizontalAxis_Length: number = 0
	VerticalAxis_Length: number = 0
	Axis_StrokeWidth: number = 0

	// insertion point for other decls

	ParameterPointersEncoding: ParameterPointersEncoding = new ParameterPointersEncoding
}

export class ParameterPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	InitialRhombusID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombus.ID = 0

}
