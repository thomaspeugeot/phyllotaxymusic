// generated code - do not edit

import { SpiralLineGridAPI } from './spirallinegrid-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'
import { SpiralLine } from './spiralline'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralLineGrid {

	static GONGSTRUCT_NAME = "SpiralLineGrid"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

	SpiralLines: Array<SpiralLine> = []
}

export function CopySpiralLineGridToSpiralLineGridAPI(spirallinegrid: SpiralLineGrid, spirallinegridAPI: SpiralLineGridAPI) {

	spirallinegridAPI.CreatedAt = spirallinegrid.CreatedAt
	spirallinegridAPI.DeletedAt = spirallinegrid.DeletedAt
	spirallinegridAPI.ID = spirallinegrid.ID

	// insertion point for basic fields copy operations
	spirallinegridAPI.Name = spirallinegrid.Name
	spirallinegridAPI.IsDisplayed = spirallinegrid.IsDisplayed

	// insertion point for pointer fields encoding
	spirallinegridAPI.SpiralLineGridPointersEncoding.ShapeCategoryID.Valid = true
	if (spirallinegrid.ShapeCategory != undefined) {
		spirallinegridAPI.SpiralLineGridPointersEncoding.ShapeCategoryID.Int64 = spirallinegrid.ShapeCategory.ID  
	} else {
		spirallinegridAPI.SpiralLineGridPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	spirallinegridAPI.SpiralLineGridPointersEncoding.SpiralLines = []
	for (let _spiralline of spirallinegrid.SpiralLines) {
		spirallinegridAPI.SpiralLineGridPointersEncoding.SpiralLines.push(_spiralline.ID)
	}

}

// CopySpiralLineGridAPIToSpiralLineGrid update basic, pointers and slice of pointers fields of spirallinegrid
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spirallinegridAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralLineGridAPIToSpiralLineGrid(spirallinegridAPI: SpiralLineGridAPI, spirallinegrid: SpiralLineGrid, frontRepo: FrontRepo) {

	spirallinegrid.CreatedAt = spirallinegridAPI.CreatedAt
	spirallinegrid.DeletedAt = spirallinegridAPI.DeletedAt
	spirallinegrid.ID = spirallinegridAPI.ID

	// insertion point for basic fields copy operations
	spirallinegrid.Name = spirallinegridAPI.Name
	spirallinegrid.IsDisplayed = spirallinegridAPI.IsDisplayed

	// insertion point for pointer fields encoding
	spirallinegrid.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spirallinegridAPI.SpiralLineGridPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(spirallinegridAPI.SpiralLineGridPointersEncoding.SpiralLines)) {
		console.error('Rects is not an array:', spirallinegridAPI.SpiralLineGridPointersEncoding.SpiralLines);
		return;
	}

	spirallinegrid.SpiralLines = new Array<SpiralLine>()
	for (let _id of spirallinegridAPI.SpiralLineGridPointersEncoding.SpiralLines) {
		let _spiralline = frontRepo.map_ID_SpiralLine.get(_id)
		if (_spiralline != undefined) {
			spirallinegrid.SpiralLines.push(_spiralline!)
		}
	}
}
