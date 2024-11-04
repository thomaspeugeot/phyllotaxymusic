// generated code - do not edit

import { SpiralCircleGridAPI } from './spiralcirclegrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { SpiralRhombusGrid } from './spiralrhombusgrid'
import { SpiralCircle } from './spiralcircle'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralCircleGrid {

	static GONGSTRUCT_NAME = "SpiralCircleGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	SpiralRhombusGrid?: SpiralRhombusGrid

	SpiralCircles: Array<SpiralCircle> = []
}

export function CopySpiralCircleGridToSpiralCircleGridAPI(spiralcirclegrid: SpiralCircleGrid, spiralcirclegridAPI: SpiralCircleGridAPI) {

	spiralcirclegridAPI.CreatedAt = spiralcirclegrid.CreatedAt
	spiralcirclegridAPI.DeletedAt = spiralcirclegrid.DeletedAt
	spiralcirclegridAPI.ID = spiralcirclegrid.ID

	// insertion point for basic fields copy operations
	spiralcirclegridAPI.Name = spiralcirclegrid.Name
	spiralcirclegridAPI.IsDisplayed = spiralcirclegrid.IsDisplayed

	// insertion point for pointer fields encoding
	spiralcirclegridAPI.SpiralCircleGridPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralcirclegrid.ShapeCategory != undefined) {
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding.ShapeCategoryID.Int64 = spiralcirclegrid.ShapeCategory.ID  
	} else {
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}

	spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralRhombusGridID.Valid = true
	if (spiralcirclegrid.SpiralRhombusGrid != undefined) {
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralRhombusGridID.Int64 = spiralcirclegrid.SpiralRhombusGrid.ID  
	} else {
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralRhombusGridID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralCircles = []
	for (let _spiralcircle of spiralcirclegrid.SpiralCircles) {
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralCircles.push(_spiralcircle.ID)
	}

}

// CopySpiralCircleGridAPIToSpiralCircleGrid update basic, pointers and slice of pointers fields of spiralcirclegrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralcirclegridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralCircleGridAPIToSpiralCircleGrid(spiralcirclegridAPI: SpiralCircleGridAPI, spiralcirclegrid: SpiralCircleGrid, frontRepo: FrontRepo) {

	spiralcirclegrid.CreatedAt = spiralcirclegridAPI.CreatedAt
	spiralcirclegrid.DeletedAt = spiralcirclegridAPI.DeletedAt
	spiralcirclegrid.ID = spiralcirclegridAPI.ID

	// insertion point for basic fields copy operations
	spiralcirclegrid.Name = spiralcirclegridAPI.Name
	spiralcirclegrid.IsDisplayed = spiralcirclegridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spiralcirclegrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralcirclegridAPI.SpiralCircleGridPointersEncoding.ShapeCategoryID.Int64)
	spiralcirclegrid.SpiralRhombusGrid = frontRepo.map_ID_SpiralRhombusGrid.get(spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralRhombusGridID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralCircles)) {
		console.error('Rects is not an array:', spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralCircles);
		return;
	}

	spiralcirclegrid.SpiralCircles = new Array<SpiralCircle>()
	for (let _id of spiralcirclegridAPI.SpiralCircleGridPointersEncoding.SpiralCircles) {
		let _spiralcircle = frontRepo.map_ID_SpiralCircle.get(_id)
		if (_spiralcircle != undefined) {
			spiralcirclegrid.SpiralCircles.push(_spiralcircle!)
		}
	}
}
