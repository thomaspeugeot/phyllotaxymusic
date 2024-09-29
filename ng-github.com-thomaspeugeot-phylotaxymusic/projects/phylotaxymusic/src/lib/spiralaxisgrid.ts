// generated code - do not edit

import { SpiralAxisGridAPI } from './spiralaxisgrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { SpiralAxis } from './spiralaxis'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralAxisGrid {

	static GONGSTRUCT_NAME = "SpiralAxisGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	SpiralAxises: Array<SpiralAxis> = []
}

export function CopySpiralAxisGridToSpiralAxisGridAPI(spiralaxisgrid: SpiralAxisGrid, spiralaxisgridAPI: SpiralAxisGridAPI) {

	spiralaxisgridAPI.CreatedAt = spiralaxisgrid.CreatedAt
	spiralaxisgridAPI.DeletedAt = spiralaxisgrid.DeletedAt
	spiralaxisgridAPI.ID = spiralaxisgrid.ID

	// insertion point for basic fields copy operations
	spiralaxisgridAPI.Name = spiralaxisgrid.Name
	spiralaxisgridAPI.IsDisplayed = spiralaxisgrid.IsDisplayed

	// insertion point for pointer fields encoding
	spiralaxisgridAPI.SpiralAxisGridPointersEncoding.ShapeCategoryID.Valid = true
	if (spiralaxisgrid.ShapeCategory != undefined) {
		spiralaxisgridAPI.SpiralAxisGridPointersEncoding.ShapeCategoryID.Int64 = spiralaxisgrid.ShapeCategory.ID  
	} else {
		spiralaxisgridAPI.SpiralAxisGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	spiralaxisgridAPI.SpiralAxisGridPointersEncoding.SpiralAxises = []
	for (let _spiralaxis of spiralaxisgrid.SpiralAxises) {
		spiralaxisgridAPI.SpiralAxisGridPointersEncoding.SpiralAxises.push(_spiralaxis.ID)
	}

}

// CopySpiralAxisGridAPIToSpiralAxisGrid update basic, pointers and slice of pointers fields of spiralaxisgrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralaxisgridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralAxisGridAPIToSpiralAxisGrid(spiralaxisgridAPI: SpiralAxisGridAPI, spiralaxisgrid: SpiralAxisGrid, frontRepo: FrontRepo) {

	spiralaxisgrid.CreatedAt = spiralaxisgridAPI.CreatedAt
	spiralaxisgrid.DeletedAt = spiralaxisgridAPI.DeletedAt
	spiralaxisgrid.ID = spiralaxisgridAPI.ID

	// insertion point for basic fields copy operations
	spiralaxisgrid.Name = spiralaxisgridAPI.Name
	spiralaxisgrid.IsDisplayed = spiralaxisgridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spiralaxisgrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralaxisgridAPI.SpiralAxisGridPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	spiralaxisgrid.SpiralAxises = new Array<SpiralAxis>()
	for (let _id of spiralaxisgridAPI.SpiralAxisGridPointersEncoding.SpiralAxises) {
		let _spiralaxis = frontRepo.map_ID_SpiralAxis.get(_id)
		if (_spiralaxis != undefined) {
			spiralaxisgrid.SpiralAxises.push(_spiralaxis!)
		}
	}
}
