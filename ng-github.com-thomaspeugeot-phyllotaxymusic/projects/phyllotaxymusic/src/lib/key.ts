// generated code - do not edit

import { KeyAPI } from './key-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ShapeCategory } from './shapecategory'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Key {

	static GONGSTRUCT_NAME = "Key"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	Path: string = ""
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

export function CopyKeyToKeyAPI(key: Key, keyAPI: KeyAPI) {

	keyAPI.CreatedAt = key.CreatedAt
	keyAPI.DeletedAt = key.DeletedAt
	keyAPI.ID = key.ID

	// insertion point for basic fields copy operations
	keyAPI.Name = key.Name
	keyAPI.IsDisplayed = key.IsDisplayed
	keyAPI.Path = key.Path
	keyAPI.Color = key.Color
	keyAPI.FillOpacity = key.FillOpacity
	keyAPI.Stroke = key.Stroke
	keyAPI.StrokeOpacity = key.StrokeOpacity
	keyAPI.StrokeWidth = key.StrokeWidth
	keyAPI.StrokeDashArray = key.StrokeDashArray
	keyAPI.StrokeDashArrayWhenSelected = key.StrokeDashArrayWhenSelected
	keyAPI.Transform = key.Transform

	// insertion point for pointer fields encoding
	keyAPI.KeyPointersEncoding.ShapeCategoryID.Valid = true
	if (key.ShapeCategory != undefined) {
		keyAPI.KeyPointersEncoding.ShapeCategoryID.Int64 = key.ShapeCategory.ID  
	} else {
		keyAPI.KeyPointersEncoding.ShapeCategoryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyKeyAPIToKey update basic, pointers and slice of pointers fields of key
// from respectively the basic fields and encoded fields of pointers and slices of pointers of keyAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyKeyAPIToKey(keyAPI: KeyAPI, key: Key, frontRepo: FrontRepo) {

	key.CreatedAt = keyAPI.CreatedAt
	key.DeletedAt = keyAPI.DeletedAt
	key.ID = keyAPI.ID

	// insertion point for basic fields copy operations
	key.Name = keyAPI.Name
	key.IsDisplayed = keyAPI.IsDisplayed
	key.Path = keyAPI.Path
	key.Color = keyAPI.Color
	key.FillOpacity = keyAPI.FillOpacity
	key.Stroke = keyAPI.Stroke
	key.StrokeOpacity = keyAPI.StrokeOpacity
	key.StrokeWidth = keyAPI.StrokeWidth
	key.StrokeDashArray = keyAPI.StrokeDashArray
	key.StrokeDashArrayWhenSelected = keyAPI.StrokeDashArrayWhenSelected
	key.Transform = keyAPI.Transform

	// insertion point for pointer fields encoding
	key.ShapeCategory = frontRepo.map_ID_ShapeCategory.get(keyAPI.KeyPointersEncoding.ShapeCategoryID.Int64)

	// insertion point for slice of pointers fields encoding
}
