// generated code - do not edit

import { SpiralCircleAPI } from './spiralcircle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralCircle {

	static GONGSTRUCT_NAME = "SpiralCircle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	CenterX: number = 0
	CenterY: number = 0
	HasBespokeRadius: boolean = false
	BespopkeRadius: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	Pitch: number = 0
	ShowName: boolean = false
	BeatNb: number = 0
	Path: string = ""

	// insertion point for pointers and slices of pointers declarations
	ShapeCategory?: ShapeCategory

}

export function CopySpiralCircleToSpiralCircleAPI(spiralcircle: SpiralCircle, spiralcircleAPI: SpiralCircleAPI) {

	spiralcircleAPI.CreatedAt = spiralcircle.CreatedAt
	spiralcircleAPI.DeletedAt = spiralcircle.DeletedAt
	spiralcircleAPI.ID = spiralcircle.ID

	// insertion point for basic fields copy operations
	spiralcircleAPI.Name = spiralcircle.Name
	spiralcircleAPI.IsDisplayed = spiralcircle.IsDisplayed
	spiralcircleAPI.CenterX = spiralcircle.CenterX
	spiralcircleAPI.CenterY = spiralcircle.CenterY
	spiralcircleAPI.HasBespokeRadius = spiralcircle.HasBespokeRadius
	spiralcircleAPI.BespopkeRadius = spiralcircle.BespopkeRadius
	spiralcircleAPI.Color = spiralcircle.Color
	spiralcircleAPI.FillOpacity = spiralcircle.FillOpacity
	spiralcircleAPI.Stroke = spiralcircle.Stroke
	spiralcircleAPI.StrokeOpacity = spiralcircle.StrokeOpacity
	spiralcircleAPI.StrokeWidth = spiralcircle.StrokeWidth
	spiralcircleAPI.StrokeDashArray = spiralcircle.StrokeDashArray
	spiralcircleAPI.StrokeDashArrayWhenSelected = spiralcircle.StrokeDashArrayWhenSelected
	spiralcircleAPI.Transform = spiralcircle.Transform
	spiralcircleAPI.Pitch = spiralcircle.Pitch
	spiralcircleAPI.ShowName = spiralcircle.ShowName
	spiralcircleAPI.BeatNb = spiralcircle.BeatNb
	spiralcircleAPI.Path = spiralcircle.Path

	// insertion point for pointer fields encoding
	spiralcircleAPI.SpiralCirclePointersEncoding.ShapeCategoryID.Valid = true
	if (spiralcircle.ShapeCategory != undefined) {
		spiralcircleAPI.SpiralCirclePointersEncoding.ShapeCategoryID.Int64 = spiralcircle.ShapeCategory.ID  
	} else {
		spiralcircleAPI.SpiralCirclePointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySpiralCircleAPIToSpiralCircle update basic, pointers and slice of pointers fields of spiralcircle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spiralcircleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySpiralCircleAPIToSpiralCircle(spiralcircleAPI: SpiralCircleAPI, spiralcircle: SpiralCircle, frontRepo: FrontRepo) {

	spiralcircle.CreatedAt = spiralcircleAPI.CreatedAt
	spiralcircle.DeletedAt = spiralcircleAPI.DeletedAt
	spiralcircle.ID = spiralcircleAPI.ID

	// insertion point for basic fields copy operations
	spiralcircle.Name = spiralcircleAPI.Name
	spiralcircle.IsDisplayed = spiralcircleAPI.IsDisplayed
	spiralcircle.CenterX = spiralcircleAPI.CenterX
	spiralcircle.CenterY = spiralcircleAPI.CenterY
	spiralcircle.HasBespokeRadius = spiralcircleAPI.HasBespokeRadius
	spiralcircle.BespopkeRadius = spiralcircleAPI.BespopkeRadius
	spiralcircle.Color = spiralcircleAPI.Color
	spiralcircle.FillOpacity = spiralcircleAPI.FillOpacity
	spiralcircle.Stroke = spiralcircleAPI.Stroke
	spiralcircle.StrokeOpacity = spiralcircleAPI.StrokeOpacity
	spiralcircle.StrokeWidth = spiralcircleAPI.StrokeWidth
	spiralcircle.StrokeDashArray = spiralcircleAPI.StrokeDashArray
	spiralcircle.StrokeDashArrayWhenSelected = spiralcircleAPI.StrokeDashArrayWhenSelected
	spiralcircle.Transform = spiralcircleAPI.Transform
	spiralcircle.Pitch = spiralcircleAPI.Pitch
	spiralcircle.ShowName = spiralcircleAPI.ShowName
	spiralcircle.BeatNb = spiralcircleAPI.BeatNb
	spiralcircle.Path = spiralcircleAPI.Path

	// insertion point for pointer fields encoding
	spiralcircle.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(spiralcircleAPI.SpiralCirclePointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
