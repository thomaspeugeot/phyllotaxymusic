// generated code - do not edit

import { RhombusGridAPI } from './rhombusgrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rhombus } from './rhombus'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RhombusGrid {

	static GONGSTRUCT_NAME = "RhombusGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Reference?: Rhombus

	Rhombuses: Array<Rhombus> = []
}

export function CopyRhombusGridToRhombusGridAPI(rhombusgrid: RhombusGrid, rhombusgridAPI: RhombusGridAPI) {

	rhombusgridAPI.CreatedAt = rhombusgrid.CreatedAt
	rhombusgridAPI.DeletedAt = rhombusgrid.DeletedAt
	rhombusgridAPI.ID = rhombusgrid.ID

	// insertion point for basic fields copy operations
	rhombusgridAPI.Name = rhombusgrid.Name
	rhombusgridAPI.IsDisplayed = rhombusgrid.IsDisplayed

	// insertion point for pointer fields encoding
	rhombusgridAPI.RhombusGridPointersEncoding.ReferenceID.Valid = true
	if (rhombusgrid.Reference != undefined) {
		rhombusgridAPI.RhombusGridPointersEncoding.ReferenceID.Int64 = rhombusgrid.Reference.ID  
	} else {
		rhombusgridAPI.RhombusGridPointersEncoding.ReferenceID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	rhombusgridAPI.RhombusGridPointersEncoding.Rhombuses = []
	for (let _rhombus of rhombusgrid.Rhombuses) {
		rhombusgridAPI.RhombusGridPointersEncoding.Rhombuses.push(_rhombus.ID)
	}

}

// CopyRhombusGridAPIToRhombusGrid update basic, pointers and slice of pointers fields of rhombusgrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rhombusgridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRhombusGridAPIToRhombusGrid(rhombusgridAPI: RhombusGridAPI, rhombusgrid: RhombusGrid, frontRepo: FrontRepo) {

	rhombusgrid.CreatedAt = rhombusgridAPI.CreatedAt
	rhombusgrid.DeletedAt = rhombusgridAPI.DeletedAt
	rhombusgrid.ID = rhombusgridAPI.ID

	// insertion point for basic fields copy operations
	rhombusgrid.Name = rhombusgridAPI.Name
	rhombusgrid.IsDisplayed = rhombusgridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	rhombusgrid.Reference = frontRepo.map_ID_Rhombus.get(rhombusgridAPI.RhombusGridPointersEncoding.ReferenceID.Int64)

	// insertion point for slice of pointers fields encoding
	rhombusgrid.Rhombuses = new Array<Rhombus>()
	for (let _id of rhombusgridAPI.RhombusGridPointersEncoding.Rhombuses) {
		let _rhombus = frontRepo.map_ID_Rhombus.get(_id)
		if (_rhombus != undefined) {
			rhombusgrid.Rhombuses.push(_rhombus!)
		}
	}
}
