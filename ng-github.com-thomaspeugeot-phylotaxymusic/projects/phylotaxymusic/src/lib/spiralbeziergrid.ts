// generated code - do not edit

import { SpiralBezierGridAPI } from './spiralbeziergrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { SpiralBezier } from './spiralbezier'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralBezierGrid {

	static GONGSTRUCT_NAME = "SpiralBezierGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	SpiralBeziers: Array<SpiralBezier> = []
}

export function CopySpiralBezierGridToSpiralBezierGridAPI(spiralbeziergrid: SpiralBezierGrid, spiralbeziergridAPI: SpiralBezierGridAPI) {

	spiralbeziergridAPI.CreatedAt = spiralbeziergrid.CreatedAt
	spiralbeziergridAPI.DeletedAt = spiralbeziergrid.DeletedAt
	spiralbeziergridAPI.ID = spiralbeziergrid.ID

	// insertion point for basic fields copy operations
	spiralbeziergridAPI.Name = spiralbeziergrid.Name
	spiralbeziergridAPI.IsDisplayed = spiralbeziergrid.IsDisplayed

	// insertion point for pointer fields encoding
	spiralbeziergridAPI.SpiralBezierGridPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralbeziergrid.ShapeCategory != undefined) {
		spiralbeziergridAPI.SpiralBezierGridPointersEncoding.ShapeCategoryID.Int64 = spiralbeziergrid.ShapeCategory.ID  
	} else {
		spiralbeziergridAPI.SpiralBezierGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	spiralbeziergridAPI.SpiralBezierGridPointersEncoding.SpiralBeziers = []
	for (let _spiralbezier of spiralbeziergrid.SpiralBeziers) {
		spiralbeziergridAPI.SpiralBezierGridPointersEncoding.SpiralBeziers.push(_spiralbezier.ID)
	}

}

// CopySpiralBezierGridAPIToSpiralBezierGrid update basic, pointers and slice of pointers fields of spiralbeziergrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralbeziergridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralBezierGridAPIToSpiralBezierGrid(spiralbeziergridAPI: SpiralBezierGridAPI, spiralbeziergrid: SpiralBezierGrid, frontRepo: FrontRepo) {

	spiralbeziergrid.CreatedAt = spiralbeziergridAPI.CreatedAt
	spiralbeziergrid.DeletedAt = spiralbeziergridAPI.DeletedAt
	spiralbeziergrid.ID = spiralbeziergridAPI.ID

	// insertion point for basic fields copy operations
	spiralbeziergrid.Name = spiralbeziergridAPI.Name
	spiralbeziergrid.IsDisplayed = spiralbeziergridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spiralbeziergrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralbeziergridAPI.SpiralBezierGridPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(spiralbeziergridAPI.SpiralBezierGridPointersEncoding.SpiralBeziers)) {
		console.error('Rects is not an array:', spiralbeziergridAPI.SpiralBezierGridPointersEncoding.SpiralBeziers);
		return;
	}

	spiralbeziergrid.SpiralBeziers = new Array<SpiralBezier>()
	for (let _id of spiralbeziergridAPI.SpiralBezierGridPointersEncoding.SpiralBeziers) {
		let _spiralbezier = frontRepo.map_ID_SpiralBezier.get(_id)
		if (_spiralbezier != undefined) {
			spiralbeziergrid.SpiralBeziers.push(_spiralbezier!)
		}
	}
}
