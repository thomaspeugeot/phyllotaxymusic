// generated code - do not edit

import { AxisGridAPI } from './axisgrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Axis } from './axis'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AxisGrid {

	static GONGSTRUCT_NAME = "AxisGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Reference?: Axis

	Axiss: Array<Axis> = []
}

export function CopyAxisGridToAxisGridAPI(axisgrid: AxisGrid, axisgridAPI: AxisGridAPI) {

	axisgridAPI.CreatedAt = axisgrid.CreatedAt
	axisgridAPI.DeletedAt = axisgrid.DeletedAt
	axisgridAPI.ID = axisgrid.ID

	// insertion point for basic fields copy operations
	axisgridAPI.Name = axisgrid.Name
	axisgridAPI.IsDisplayed = axisgrid.IsDisplayed

	// insertion point for pointer fields encoding
	axisgridAPI.AxisGridPointersEncoding.ReferenceID.Valid = true
	if (axisgrid.Reference != undefined) {
		axisgridAPI.AxisGridPointersEncoding.ReferenceID.Int64 = axisgrid.Reference.ID  
	} else {
		axisgridAPI.AxisGridPointersEncoding.ReferenceID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	axisgridAPI.AxisGridPointersEncoding.Axiss = []
	for (let _axis of axisgrid.Axiss) {
		axisgridAPI.AxisGridPointersEncoding.Axiss.push(_axis.ID)
	}

}

// CopyAxisGridAPIToAxisGrid update basic, pointers and slice of pointers fields of axisgrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of axisgridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAxisGridAPIToAxisGrid(axisgridAPI: AxisGridAPI, axisgrid: AxisGrid, frontRepo: FrontRepo) {

	axisgrid.CreatedAt = axisgridAPI.CreatedAt
	axisgrid.DeletedAt = axisgridAPI.DeletedAt
	axisgrid.ID = axisgridAPI.ID

	// insertion point for basic fields copy operations
	axisgrid.Name = axisgridAPI.Name
	axisgrid.IsDisplayed = axisgridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	axisgrid.Reference = frontRepo.map_ID_Axis.get(axisgridAPI.AxisGridPointersEncoding.ReferenceID.Int64)

	// insertion point for slice of pointers fields encoding
	axisgrid.Axiss = new Array<Axis>()
	for (let _id of axisgridAPI.AxisGridPointersEncoding.Axiss) {
		let _axis = frontRepo.map_ID_Axis.get(_id)
		if (_axis != undefined) {
			axisgrid.Axiss.push(_axis!)
		}
	}
}
