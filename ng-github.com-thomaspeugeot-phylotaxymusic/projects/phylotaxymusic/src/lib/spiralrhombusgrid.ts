// generated code - do not edit

import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { RhombusGrid } from './rhombusgrid'

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

	RhombusGrid?: RhombusGrid

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

	spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.RhombusGridID.Valid = true
	if (spiralrhombusgrid.RhombusGrid != undefined) {
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.RhombusGridID.Int64 = spiralrhombusgrid.RhombusGrid.ID  
	} else {
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.RhombusGridID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
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
	spiralrhombusgrid.RhombusGrid = frontRepo.map_ID_RhombusGrid.get(spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding.RhombusGridID.Int64)

	// insertion point for slice of pointers fields encoding
}
