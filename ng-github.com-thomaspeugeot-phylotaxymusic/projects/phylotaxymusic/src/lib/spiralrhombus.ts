// generated code - do not edit

import { SpiralRhombusAPI } from './spiralrhombus-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { Rhombus } from './rhombus'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralRhombus {

	static GONGSTRUCT_NAME = "SpiralRhombus"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	Rhombus?: Rhombus

}

export function CopySpiralRhombusToSpiralRhombusAPI(spiralrhombus: SpiralRhombus, spiralrhombusAPI: SpiralRhombusAPI) {

	spiralrhombusAPI.CreatedAt = spiralrhombus.CreatedAt
	spiralrhombusAPI.DeletedAt = spiralrhombus.DeletedAt
	spiralrhombusAPI.ID = spiralrhombus.ID

	// insertion point for basic fields copy operations
	spiralrhombusAPI.Name = spiralrhombus.Name
	spiralrhombusAPI.IsDisplayed = spiralrhombus.IsDisplayed

	// insertion point for pointer fields encoding
	spiralrhombusAPI.SpiralRhombusPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralrhombus.ShapeCategory != undefined) {
		spiralrhombusAPI.SpiralRhombusPointersEncoding.ShapeCategoryID.Int64 = spiralrhombus.ShapeCategory.ID  
	} else {
		spiralrhombusAPI.SpiralRhombusPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}

	spiralrhombusAPI.SpiralRhombusPointersEncoding.RhombusID.Valid = true
	if (spiralrhombus.Rhombus != undefined) {
		spiralrhombusAPI.SpiralRhombusPointersEncoding.RhombusID.Int64 = spiralrhombus.Rhombus.ID  
	} else {
		spiralrhombusAPI.SpiralRhombusPointersEncoding.RhombusID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralRhombusAPIToSpiralRhombus update basic, pointers and slice of pointers fields of spiralrhombus
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralrhombusAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralRhombusAPIToSpiralRhombus(spiralrhombusAPI: SpiralRhombusAPI, spiralrhombus: SpiralRhombus, frontRepo: FrontRepo) {

	spiralrhombus.CreatedAt = spiralrhombusAPI.CreatedAt
	spiralrhombus.DeletedAt = spiralrhombusAPI.DeletedAt
	spiralrhombus.ID = spiralrhombusAPI.ID

	// insertion point for basic fields copy operations
	spiralrhombus.Name = spiralrhombusAPI.Name
	spiralrhombus.IsDisplayed = spiralrhombusAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spiralrhombus.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralrhombusAPI.SpiralRhombusPointersEncoding.ShapeCategoryID.Int64)
	spiralrhombus.Rhombus = frontRepo.map_ID_Rhombus.get(spiralrhombusAPI.SpiralRhombusPointersEncoding.RhombusID.Int64)

	// insertion point for slice of pointers fields encoding
}
