// generated code - do not edit

import { CircleAPI } from './circle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Circle {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	CenterX: number = 0
	CenterY: number = 0
	StrokeWidth: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCircleToCircleAPI(circle: Circle, circleAPI: CircleAPI) {

	circleAPI.CreatedAt = circle.CreatedAt
	circleAPI.DeletedAt = circle.DeletedAt
	circleAPI.ID = circle.ID

	// insertion point for basic fields copy operations
	circleAPI.Name = circle.Name
	circleAPI.IsDisplayed = circle.IsDisplayed
	circleAPI.CenterX = circle.CenterX
	circleAPI.CenterY = circle.CenterY
	circleAPI.StrokeWidth = circle.StrokeWidth

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCircleAPIToCircle update basic, pointers and slice of pointers fields of circle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleAPIToCircle(circleAPI: CircleAPI, circle: Circle, frontRepo: FrontRepo) {

	circle.CreatedAt = circleAPI.CreatedAt
	circle.DeletedAt = circleAPI.DeletedAt
	circle.ID = circleAPI.ID

	// insertion point for basic fields copy operations
	circle.Name = circleAPI.Name
	circle.IsDisplayed = circleAPI.IsDisplayed
	circle.CenterX = circleAPI.CenterX
	circle.CenterY = circleAPI.CenterY
	circle.StrokeWidth = circleAPI.StrokeWidth

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
