// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteInfoAPI {

	static GONGSTRUCT_NAME = "NoteInfo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsKept: boolean = false

	// insertion point for other decls

	NoteInfoPointersEncoding: NoteInfoPointersEncoding = new NoteInfoPointersEncoding
}

export class NoteInfoPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
