// generated code - do not edit

import { ShapeCategoryAPI } from './shapecategory-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ShapeCategory {

	static GONGSTRUCT_NAME = "ShapeCategory"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyShapeCategoryToShapeCategoryAPI(shapecategory: ShapeCategory, shapecategoryAPI: ShapeCategoryAPI) {

	shapecategoryAPI.CreatedAt = shapecategory.CreatedAt
	shapecategoryAPI.DeletedAt = shapecategory.DeletedAt
	shapecategoryAPI.ID = shapecategory.ID

	// insertion point for basic fields copy operations
	shapecategoryAPI.Name = shapecategory.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyShapeCategoryAPIToShapeCategory update basic, pointers and slice of pointers fields of shapecategory
// from respectively the basic fields and encoded fields of pointers and slices of pointers of shapecategoryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyShapeCategoryAPIToShapeCategory(shapecategoryAPI: ShapeCategoryAPI, shapecategory: ShapeCategory, frontRepo: FrontRepo) {

	shapecategory.CreatedAt = shapecategoryAPI.CreatedAt
	shapecategory.DeletedAt = shapecategoryAPI.DeletedAt
	shapecategory.ID = shapecategoryAPI.ID

	// insertion point for basic fields copy operations
	shapecategory.Name = shapecategoryAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
