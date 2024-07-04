// generated code - do not edit

import { LineAPI } from './line-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Line {

	static GONGSTRUCT_NAME = "Line"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X1: number = 0
	Y1: number = 0
	X2: number = 0
	Y2: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyLineToLineAPI(line: Line, lineAPI: LineAPI) {

	lineAPI.CreatedAt = line.CreatedAt
	lineAPI.DeletedAt = line.DeletedAt
	lineAPI.ID = line.ID

	// insertion point for basic fields copy operations
	lineAPI.Name = line.Name
	lineAPI.X1 = line.X1
	lineAPI.Y1 = line.Y1
	lineAPI.X2 = line.X2
	lineAPI.Y2 = line.Y2

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLineAPIToLine update basic, pointers and slice of pointers fields of line
// from respectively the basic fields and encoded fields of pointers and slices of pointers of lineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLineAPIToLine(lineAPI: LineAPI, line: Line, frontRepo: FrontRepo) {

	line.CreatedAt = lineAPI.CreatedAt
	line.DeletedAt = lineAPI.DeletedAt
	line.ID = lineAPI.ID

	// insertion point for basic fields copy operations
	line.Name = lineAPI.Name
	line.X1 = lineAPI.X1
	line.Y1 = lineAPI.Y1
	line.X2 = lineAPI.X2
	line.Y2 = lineAPI.Y2

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
