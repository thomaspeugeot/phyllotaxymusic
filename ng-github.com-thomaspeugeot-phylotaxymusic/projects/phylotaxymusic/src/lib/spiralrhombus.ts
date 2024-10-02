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
	X_r0: number = 0
	Y_r0: number = 0
	X_r1: number = 0
	Y_r1: number = 0
	X_r2: number = 0
	Y_r2: number = 0
	X_r3: number = 0
	Y_r3: number = 0

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
	spiralrhombusAPI.X_r0 = spiralrhombus.X_r0
	spiralrhombusAPI.Y_r0 = spiralrhombus.Y_r0
	spiralrhombusAPI.X_r1 = spiralrhombus.X_r1
	spiralrhombusAPI.Y_r1 = spiralrhombus.Y_r1
	spiralrhombusAPI.X_r2 = spiralrhombus.X_r2
	spiralrhombusAPI.Y_r2 = spiralrhombus.Y_r2
	spiralrhombusAPI.X_r3 = spiralrhombus.X_r3
	spiralrhombusAPI.Y_r3 = spiralrhombus.Y_r3

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
	spiralrhombus.X_r0 = spiralrhombusAPI.X_r0
	spiralrhombus.Y_r0 = spiralrhombusAPI.Y_r0
	spiralrhombus.X_r1 = spiralrhombusAPI.X_r1
	spiralrhombus.Y_r1 = spiralrhombusAPI.Y_r1
	spiralrhombus.X_r2 = spiralrhombusAPI.X_r2
	spiralrhombus.Y_r2 = spiralrhombusAPI.Y_r2
	spiralrhombus.X_r3 = spiralrhombusAPI.X_r3
	spiralrhombus.Y_r3 = spiralrhombusAPI.Y_r3

	// insertion point for pointer fields encoding
	spiralrhombus.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralrhombusAPI.SpiralRhombusPointersEncoding.ShapeCategoryID.Int64)
	spiralrhombus.Rhombus = frontRepo.map_ID_Rhombus.get(spiralrhombusAPI.SpiralRhombusPointersEncoding.RhombusID.Int64)

	// insertion point for slice of pointers fields encoding
}
