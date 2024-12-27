// generated code - do not edit

import { SpiralLineAPI } from './spiralline-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralLine {

	static GONGSTRUCT_NAME = "SpiralLine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	StartX: number = 0
	EndX: number = 0
	StartY: number = 0
	EndY: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

}

export function CopySpiralLineToSpiralLineAPI(spiralline: SpiralLine, spirallineAPI: SpiralLineAPI) {

	spirallineAPI.CreatedAt = spiralline.CreatedAt
	spirallineAPI.DeletedAt = spiralline.DeletedAt
	spirallineAPI.ID = spiralline.ID

	// insertion point for basic fields copy operations
	spirallineAPI.Name = spiralline.Name
	spirallineAPI.IsDisplayed = spiralline.IsDisplayed
	spirallineAPI.StartX = spiralline.StartX
	spirallineAPI.EndX = spiralline.EndX
	spirallineAPI.StartY = spiralline.StartY
	spirallineAPI.EndY = spiralline.EndY
	spirallineAPI.Color = spiralline.Color
	spirallineAPI.FillOpacity = spiralline.FillOpacity
	spirallineAPI.Stroke = spiralline.Stroke
	spirallineAPI.StrokeOpacity = spiralline.StrokeOpacity
	spirallineAPI.StrokeWidth = spiralline.StrokeWidth
	spirallineAPI.StrokeDashArray = spiralline.StrokeDashArray
	spirallineAPI.StrokeDashArrayWhenSelected = spiralline.StrokeDashArrayWhenSelected
	spirallineAPI.Transform = spiralline.Transform

	// insertion point for pointer fields encoding
	spirallineAPI.SpiralLinePointersEncoding.ShapeCategoryID.Valid = true
	if (spiralline.ShapeCategory != undefined) {
		spirallineAPI.SpiralLinePointersEncoding.ShapeCategoryID.Int64 = spiralline.ShapeCategory.ID  
	} else {
		spirallineAPI.SpiralLinePointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralLineAPIToSpiralLine update basic, pointers and slice of pointers fields of spiralline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spirallineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralLineAPIToSpiralLine(spirallineAPI: SpiralLineAPI, spiralline: SpiralLine, frontRepo: FrontRepo) {

	spiralline.CreatedAt = spirallineAPI.CreatedAt
	spiralline.DeletedAt = spirallineAPI.DeletedAt
	spiralline.ID = spirallineAPI.ID

	// insertion point for basic fields copy operations
	spiralline.Name = spirallineAPI.Name
	spiralline.IsDisplayed = spirallineAPI.IsDisplayed
	spiralline.StartX = spirallineAPI.StartX
	spiralline.EndX = spirallineAPI.EndX
	spiralline.StartY = spirallineAPI.StartY
	spiralline.EndY = spirallineAPI.EndY
	spiralline.Color = spirallineAPI.Color
	spiralline.FillOpacity = spirallineAPI.FillOpacity
	spiralline.Stroke = spirallineAPI.Stroke
	spiralline.StrokeOpacity = spirallineAPI.StrokeOpacity
	spiralline.StrokeWidth = spirallineAPI.StrokeWidth
	spiralline.StrokeDashArray = spirallineAPI.StrokeDashArray
	spiralline.StrokeDashArrayWhenSelected = spirallineAPI.StrokeDashArrayWhenSelected
	spiralline.Transform = spirallineAPI.Transform

	// insertion point for pointer fields encoding
	spiralline.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spirallineAPI.SpiralLinePointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
