// generated code - do not edit

import { BezierGridAPI } from './beziergrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Bezier } from './bezier'
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierGrid {

	static GONGSTRUCT_NAME = "BezierGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Reference?: Bezier

	ShapeCategory?: ShapeCategory

	Beziers: Array<Bezier> = []
}

export function CopyBezierGridToBezierGridAPI(beziergrid: BezierGrid, beziergridAPI: BezierGridAPI) {

	beziergridAPI.CreatedAt = beziergrid.CreatedAt
	beziergridAPI.DeletedAt = beziergrid.DeletedAt
	beziergridAPI.ID = beziergrid.ID

	// insertion point for basic fields copy operations
	beziergridAPI.Name = beziergrid.Name
	beziergridAPI.IsDisplayed = beziergrid.IsDisplayed

	// insertion point for pointer fields encoding
	beziergridAPI.BezierGridPointersEncoding.ReferenceID.Valid = true
	if (beziergrid.Reference != undefined) {
		beziergridAPI.BezierGridPointersEncoding.ReferenceID.Int64 = beziergrid.Reference.ID  
	} else {
		beziergridAPI.BezierGridPointersEncoding.ReferenceID.Int64 = 0 		
	}

	beziergridAPI.BezierGridPointersEncoding.ShapeCategoryID.Valid = true
	if (beziergrid.ShapeCategory != undefined) {
		beziergridAPI.BezierGridPointersEncoding.ShapeCategoryID.Int64 = beziergrid.ShapeCategory.ID  
	} else {
		beziergridAPI.BezierGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	beziergridAPI.BezierGridPointersEncoding.Beziers = []
	for (let _bezier of beziergrid.Beziers) {
		beziergridAPI.BezierGridPointersEncoding.Beziers.push(_bezier.ID)
	}

}

// CopyBezierGridAPIToBezierGrid update basic, pointers and slice of pointers fields of beziergrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of beziergridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBezierGridAPIToBezierGrid(beziergridAPI: BezierGridAPI, beziergrid: BezierGrid, frontRepo: FrontRepo) {

	beziergrid.CreatedAt = beziergridAPI.CreatedAt
	beziergrid.DeletedAt = beziergridAPI.DeletedAt
	beziergrid.ID = beziergridAPI.ID

	// insertion point for basic fields copy operations
	beziergrid.Name = beziergridAPI.Name
	beziergrid.IsDisplayed = beziergridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	beziergrid.Reference = frontRepo.map_ID_Bezier.get(beziergridAPI.BezierGridPointersEncoding.ReferenceID.Int64)
	beziergrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(beziergridAPI.BezierGridPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(beziergridAPI.BezierGridPointersEncoding.Beziers)) {
		console.error('Rects is not an array:', beziergridAPI.BezierGridPointersEncoding.Beziers);
		return;
	}

	beziergrid.Beziers = new Array<Bezier>()
	for (let _id of beziergridAPI.BezierGridPointersEncoding.Beziers) {
		let _bezier = frontRepo.map_ID_Bezier.get(_id)
		if (_bezier != undefined) {
			beziergrid.Beziers.push(_bezier!)
		}
	}
}
