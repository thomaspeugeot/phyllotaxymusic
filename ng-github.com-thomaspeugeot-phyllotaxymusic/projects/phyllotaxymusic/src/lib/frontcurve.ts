// generated code - do not edit

import { FrontCurveAPI } from './frontcurve-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FrontCurve {

	static GONGSTRUCT_NAME = "FrontCurve"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Path: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFrontCurveToFrontCurveAPI(frontcurve: FrontCurve, frontcurveAPI: FrontCurveAPI) {

	frontcurveAPI.CreatedAt = frontcurve.CreatedAt
	frontcurveAPI.DeletedAt = frontcurve.DeletedAt
	frontcurveAPI.ID = frontcurve.ID

	// insertion point for basic fields copy operations
	frontcurveAPI.Name = frontcurve.Name
	frontcurveAPI.Path = frontcurve.Path

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFrontCurveAPIToFrontCurve update basic, pointers and slice of pointers fields of frontcurve
// from respectively the basic fields and encoded fields of pointers and slices of pointers of frontcurveAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFrontCurveAPIToFrontCurve(frontcurveAPI: FrontCurveAPI, frontcurve: FrontCurve, frontRepo: FrontRepo) {

	frontcurve.CreatedAt = frontcurveAPI.CreatedAt
	frontcurve.DeletedAt = frontcurveAPI.DeletedAt
	frontcurve.ID = frontcurveAPI.ID

	// insertion point for basic fields copy operations
	frontcurve.Name = frontcurveAPI.Name
	frontcurve.Path = frontcurveAPI.Path

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
