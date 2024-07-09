// generated code - do not edit

import { CircleGridAPI } from './circlegrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Circle } from './circle'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CircleGrid {

	static GONGSTRUCT_NAME = "CircleGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Reference?: Circle

	Circles: Array<Circle> = []
}

export function CopyCircleGridToCircleGridAPI(circlegrid: CircleGrid, circlegridAPI: CircleGridAPI) {

	circlegridAPI.CreatedAt = circlegrid.CreatedAt
	circlegridAPI.DeletedAt = circlegrid.DeletedAt
	circlegridAPI.ID = circlegrid.ID

	// insertion point for basic fields copy operations
	circlegridAPI.Name = circlegrid.Name
	circlegridAPI.IsDisplayed = circlegrid.IsDisplayed

	// insertion point for pointer fields encoding
	circlegridAPI.CircleGridPointersEncoding.ReferenceID.Valid = true
	if (circlegrid.Reference != undefined) {
		circlegridAPI.CircleGridPointersEncoding.ReferenceID.Int64 = circlegrid.Reference.ID  
	} else {
		circlegridAPI.CircleGridPointersEncoding.ReferenceID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	circlegridAPI.CircleGridPointersEncoding.Circles = []
	for (let _circle of circlegrid.Circles) {
		circlegridAPI.CircleGridPointersEncoding.Circles.push(_circle.ID)
	}

}

// CopyCircleGridAPIToCircleGrid update basic, pointers and slice of pointers fields of circlegrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circlegridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleGridAPIToCircleGrid(circlegridAPI: CircleGridAPI, circlegrid: CircleGrid, frontRepo: FrontRepo) {

	circlegrid.CreatedAt = circlegridAPI.CreatedAt
	circlegrid.DeletedAt = circlegridAPI.DeletedAt
	circlegrid.ID = circlegridAPI.ID

	// insertion point for basic fields copy operations
	circlegrid.Name = circlegridAPI.Name
	circlegrid.IsDisplayed = circlegridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	circlegrid.Reference = frontRepo.map_ID_Circle.get(circlegridAPI.CircleGridPointersEncoding.ReferenceID.Int64)

	// insertion point for slice of pointers fields encoding
	circlegrid.Circles = new Array<Circle>()
	for (let _id of circlegridAPI.CircleGridPointersEncoding.Circles) {
		let _circle = frontRepo.map_ID_Circle.get(_id)
		if (_circle != undefined) {
			circlegrid.Circles.push(_circle!)
		}
	}
}
