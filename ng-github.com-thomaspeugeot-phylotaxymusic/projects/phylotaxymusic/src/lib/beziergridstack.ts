// generated code - do not edit

import { BezierGridStackAPI } from './beziergridstack-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { BezierGrid } from './beziergrid'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierGridStack {

	static GONGSTRUCT_NAME = "BezierGridStack"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	BezierGrids: Array<BezierGrid> = []
}

export function CopyBezierGridStackToBezierGridStackAPI(beziergridstack: BezierGridStack, beziergridstackAPI: BezierGridStackAPI) {

	beziergridstackAPI.CreatedAt = beziergridstack.CreatedAt
	beziergridstackAPI.DeletedAt = beziergridstack.DeletedAt
	beziergridstackAPI.ID = beziergridstack.ID

	// insertion point for basic fields copy operations
	beziergridstackAPI.Name = beziergridstack.Name
	beziergridstackAPI.IsDisplayed = beziergridstack.IsDisplayed

	// insertion point for pointer fields encoding
	beziergridstackAPI.BezierGridStackPointersEncoding.ShapeCategoryID.Valid = true
	if (beziergridstack.ShapeCategory != undefined) {
		beziergridstackAPI.BezierGridStackPointersEncoding.ShapeCategoryID.Int64 = beziergridstack.ShapeCategory.ID  
	} else {
		beziergridstackAPI.BezierGridStackPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	beziergridstackAPI.BezierGridStackPointersEncoding.BezierGrids = []
	for (let _beziergrid of beziergridstack.BezierGrids) {
		beziergridstackAPI.BezierGridStackPointersEncoding.BezierGrids.push(_beziergrid.ID)
	}

}

// CopyBezierGridStackAPIToBezierGridStack update basic, pointers and slice of pointers fields of beziergridstack
// from respectively the basic fields and encoded fields of pointers and slices of pointers of beziergridstackAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBezierGridStackAPIToBezierGridStack(beziergridstackAPI: BezierGridStackAPI, beziergridstack: BezierGridStack, frontRepo: FrontRepo) {

	beziergridstack.CreatedAt = beziergridstackAPI.CreatedAt
	beziergridstack.DeletedAt = beziergridstackAPI.DeletedAt
	beziergridstack.ID = beziergridstackAPI.ID

	// insertion point for basic fields copy operations
	beziergridstack.Name = beziergridstackAPI.Name
	beziergridstack.IsDisplayed = beziergridstackAPI.IsDisplayed

	// insertion point for pointer fields encoding
	beziergridstack.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(beziergridstackAPI.BezierGridStackPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	beziergridstack.BezierGrids = new Array<BezierGrid>()
	for (let _id of beziergridstackAPI.BezierGridStackPointersEncoding.BezierGrids) {
		let _beziergrid = frontRepo.map_ID_BezierGrid.get(_id)
		if (_beziergrid != undefined) {
			beziergridstack.BezierGrids.push(_beziergrid!)
		}
	}
}
