// generated code - do not edit

import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { SpiralRhombus } from './spiralrhombus'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralRhombusGrid {

	static GONGSTRUCT_NAME = "SpiralRhombusGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	SpiralRhombuses: Array<SpiralRhombus> = []
}

export function CopySpiralRhombusGridToSpiralRhombusGridAPI(spiralrhombusgrid: SpiralRhombusGrid, spiralrhombusgridAPI: SpiralRhombusGridAPI) {

	spiralrhombusgridAPI.CreatedAt = spiralrhombusgrid.CreatedAt
	spiralrhombusgridAPI.DeletedAt = spiralrhombusgrid.DeletedAt
	spiralrhombusgridAPI.ID = spiralrhombusgrid.ID

	// insertion point for basic fields copy operations
	spiralrhombusgridAPI.Name = spiralrhombusgrid.Name
	spiralrhombusgridAPI.IsDisplayed = spiralrhombusgrid.IsDisplayed

	// insertion point for pointer fields encoding
	spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralrhombusgrid.ShapeCategory != undefined) {
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.ShapeCategoryID.Int64 = spiralrhombusgrid.ShapeCategory.ID  
	} else {
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.SpiralRhombuses = []
	for (let _spiralrhombus of spiralrhombusgrid.SpiralRhombuses) {
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.SpiralRhombuses.push(_spiralrhombus.ID)
	}

}

// CopySpiralRhombusGridAPIToSpiralRhombusGrid update basic, pointers and slice of pointers fields of spiralrhombusgrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralrhombusgridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralRhombusGridAPIToSpiralRhombusGrid(spiralrhombusgridAPI: SpiralRhombusGridAPI, spiralrhombusgrid: SpiralRhombusGrid, frontRepo: FrontRepo) {

	spiralrhombusgrid.CreatedAt = spiralrhombusgridAPI.CreatedAt
	spiralrhombusgrid.DeletedAt = spiralrhombusgridAPI.DeletedAt
	spiralrhombusgrid.ID = spiralrhombusgridAPI.ID

	// insertion point for basic fields copy operations
	spiralrhombusgrid.Name = spiralrhombusgridAPI.Name
	spiralrhombusgrid.IsDisplayed = spiralrhombusgridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spiralrhombusgrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	spiralrhombusgrid.SpiralRhombuses = new Array<SpiralRhombus>()
	for (let _id of spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.SpiralRhombuses) {
		let _spiralrhombus = frontRepo.map_ID_SpiralRhombus.get(_id)
		if (_spiralrhombus != undefined) {
			spiralrhombusgrid.SpiralRhombuses.push(_spiralrhombus!)
		}
	}
}
