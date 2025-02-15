// insertion point for imports
import { ParameterAPI } from './parameter-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ExportToMusicxmlAPI {

	static GONGSTRUCT_NAME = "ExportToMusicxml"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	ExportToMusicxmlPointersEncoding: ExportToMusicxmlPointersEncoding = new ExportToMusicxmlPointersEncoding
}

export class ExportToMusicxmlPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ParameterID: NullInt64 = new NullInt64 // if pointer is null, Parameter.ID = 0

}
