// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FrontCurveAPI {

	static GONGSTRUCT_NAME = "FrontCurve"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Path: string = ""

	// insertion point for other decls

	FrontCurvePointersEncoding: FrontCurvePointersEncoding = new FrontCurvePointersEncoding
}

export class FrontCurvePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
