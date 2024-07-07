// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class InitialAxisAPI {

	static GONGSTRUCT_NAME = "InitialAxis"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	Angle: number = 0
	Length: number = 0
	StrokeWidth: number = 0

	// insertion point for other decls

	InitialAxisPointersEncoding: InitialAxisPointersEncoding = new InitialAxisPointersEncoding
}

export class InitialAxisPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
