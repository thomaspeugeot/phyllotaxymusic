package models

import (
	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// ShapeInterface
type ShapeInterface interface {
	GetName() string
	GetIsDisplayed() bool
	GetIsDisplayedPointer() *bool
	GetShapeCategory() *ShapeCategory
	Draw(svgStage *svg_models.StageStruct,
		layer *svg_models.Layer,
		parameter *Parameter)
}
